package server

import (
	"flag"
	"os"
)

// define important variables
var (
	addr = flag.String("addr",":", os.Getenv("PORT"))
	cert = flag.String("cert","","")
	key = flag.String("key","","")
)

// define the function run
func Run() error {
	// parse the flag
	flag.Parse()

	// assign default port number if port is blank
	if *addr == ":" {
		*addr = ":8080"
	}

	// assign function on GET endpoints
	app.GET("/", handlers.Welcome)
	app.GET("/room/create", handlers.RoomCreate)
	app.GET("/room/:uuid", handlers.Room)
	app.GET("/room/:uuid/websocket",)
	app.GET("/room/:uuid/chat", handlers.RoomChat)
	app.GET("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.GET("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.GET("/stream/:ssuid", handlers.Stream)
	app.GET("/stream/:ssuid/websocket", handlers.)
	app.GET("/stream/:ssuid/chat/websocket", handlers.)
	app.GET("/stream/:ssuid/viewer/websocket", handlers.)

}