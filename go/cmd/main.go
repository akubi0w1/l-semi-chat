package main

import (
	"l-semi-chat/pkg/interface/auth"
	"l-semi-chat/pkg/interface/database"
	"l-semi-chat/pkg/interface/handler"
	"l-semi-chat/pkg/interface/server"
	"l-semi-chat/pkg/interface/server/router"
)

func main() {
	// TODO: load config

	// connect db
	sh := database.NewSQLHandler()

	// create handler
	ph := auth.NewPasswordHandler()
	appHandler := handler.NewAppHandler(sh, ph)

	// create server
	serv := server.NewServer("localhost", "8080")

	// setup router
	router.SetupRouter(serv, appHandler)

	// server running
	serv.Serve()

}