package main

import (
	"l-semi-chat/conf"
	"l-semi-chat/pkg/interface/auth"
	"l-semi-chat/pkg/interface/database"
	"l-semi-chat/pkg/interface/handler"
	"l-semi-chat/pkg/interface/server"
	"l-semi-chat/pkg/interface/server/router"
)

func main() {
	// load config
	servConf := conf.LoadServerConfig()

	// connect db
	sh := database.NewSQLHandler()

	// create handler
	ph := auth.NewPasswordHandler()
	appHandler := handler.NewAppHandler(sh, ph)

	// create server
	serv := server.NewServer(servConf["addr"], servConf["port"])

	// setup router
	router.SetupRouter(serv, appHandler)

	// server running
	serv.Serve()

}
