module github.com/gofred-io/gofred-website

go 1.25.0

//replace github.com/gofred-io/gofred => ../gofred

require (
	github.com/fsnotify/fsnotify v1.9.0
	github.com/gofred-io/gofred v0.0.0-20250903154902-9e68352543e4
	github.com/gorilla/websocket v1.5.3
)

require golang.org/x/sys v0.35.0 // indirect
