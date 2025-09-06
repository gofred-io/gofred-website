module github.com/gofred-io/gofred-website

go 1.25.0

//replace github.com/gofred-io/gofred => ../gofred

require (
	github.com/fsnotify/fsnotify v1.9.0
	github.com/gofred-io/gofred v0.0.2
	github.com/gorilla/websocket v1.5.3
)

require golang.org/x/sys v0.35.0 // indirect
