package main

import (
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var (
	chReload = make(chan struct{})
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	localURL       string
	netURL         string
	colorSuccess   = 32
	colorCompiling = 33
)

type wsMsg struct {
	Command string `json:"cmd"`
}

func main() {
	go watch()
	go runWebsockerServer()

	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer addr.Close()

	port := addr.Addr().(*net.TCPAddr).Port
	localURL = fmt.Sprintf("http://localhost:%d", port)

	// find the IP address of the machine
	ip, err := net.LookupIP("localhost")
	if err != nil {
		log.Printf("error: %v", err)
	}
	netURL = fmt.Sprintf("http://%s:%d", ip[0].String(), port)

	printSuccessMessage()

	openURL(localURL)

	rootDir := "./server"
	//log.Fatal(http.Serve(addr, http.FileServer(http.Dir("./server"))))
	log.Fatal(http.Serve(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the file is not found, serve the index.html file
		if _, err := os.Stat(rootDir + r.URL.Path); os.IsNotExist(err) {
			http.ServeFile(w, r, rootDir+"/index.html")
			return
		}

		http.FileServer(http.Dir(rootDir)).ServeHTTP(w, r)
		w.Header().Set("Content-Security-Policy", "script-src 'self' 'unsafe-inline'; default-src 'self'")
	})))
}

func runWebsockerServer() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var (
			quit = make(chan struct{})
		)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error: %v", err)
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			return
		}

		defer conn.Close()
		defer close(quit)

		go func() {
			for {
				select {
				case <-quit:
					return
				case <-chReload:
					err := conn.WriteJSON(wsMsg{Command: "reload"})
					if err != nil {
						log.Printf("error writing to websocket: %v", err)
					}
				}
			}
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				//log.Printf("error: %v", err)
				return
			}
		}
	})

	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer addr.Close()

	port := addr.Addr().(*net.TCPAddr).Port
	writeEnvJS(port)

	http.Serve(addr, nil)
}

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("error: %v", err)
	}
	defer watcher.Close()

	dir := "."
	err = watcher.Add(dir)
	if err != nil {
		log.Printf("error: %v", err)
	}

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.HasPrefix(path, ".") && strings.HasSuffix(path, ".go") {
			watcher.Add(path)
		}

		return nil
	})
	if err != nil {
		log.Printf("error: %v", err)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	hasChanged := atomic.Bool{}

	for {
		select {
		case _, ok := <-watcher.Events:
			if !ok {
				return
			}
			hasChanged.Store(true)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("error: %v", err)
		case <-ticker.C:
			if hasChanged.CompareAndSwap(true, false) {
				removeLines(6)
				printCompilingMessage()
				build()
				removeLines(1)
				printSuccessMessage()
				chReload <- struct{}{}
			}
		}
	}
}

func build() {
	process := exec.Command("go", "build", "-o", "server/main.wasm", ".")
	process.Env = append(os.Environ(), "GOARCH=wasm", "GOOS=js")
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	err := process.Run()
	if err != nil {
		log.Printf("error: %v", err)
	}
}

// openURL opens the specified URL in the default browser of the user.
func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		// Check if running under WSL
		if isWSL() {
			// Use 'cmd.exe /c start' to open the URL in the default Windows browser
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			// Use xdg-open on native Linux environments
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

func printSuccessMessage() {
	fmt.Printf("\x1b[%dmCompiled successfully!\x1b[0m\n\n", colorSuccess)
	fmt.Printf("Open your browser and go to:\n\n")
	fmt.Printf("   Local: %s\n", localURL)
	fmt.Printf("   Network: %s\n", netURL)
}

func printCompilingMessage() {
	fmt.Printf("\x1b[%dmCompiling...\x1b[0m\n", colorCompiling)
}

func removeLines(n int) {
	fmt.Print(strings.Repeat("\033[1A\033[K", n))
}

func writeEnvJS(port int) {
	var (
		content = strings.Builder{}
	)

	content.WriteString("window.env = {\n")
	content.WriteString(fmt.Sprintf("\tLIVE_PORT: %d\n", port))
	content.WriteString("}\n")

	err := os.WriteFile("server/env.js", []byte(content.String()), 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
