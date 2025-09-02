module github.com/gofred-io/gofred-website

go 1.25.0

replace github.com/gofred-io/gofred => ../gofred

require (
	github.com/fsnotify/fsnotify v1.9.0
	github.com/gofred-io/gofred v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.3
)

require (
	golang.org/x/exp/shiny v0.0.0-20250606033433-dcc06ee1d476 // indirect
	golang.org/x/image v0.28.0 // indirect
	golang.org/x/mobile v0.0.0-20250606033058-a2a15c67f36f // indirect
	golang.org/x/sys v0.35.0 // indirect
)
