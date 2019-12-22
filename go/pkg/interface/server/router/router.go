package router

import (
	"l-semi-chat/pkg/interface/handler"
	"l-semi-chat/pkg/interface/server"
)

func SetupRouter(s server.Server, h handler.AppHandler) {
	s.Post("/accounts", h.CreateAccount())

	// auth
	s.Post("/login", h.Login())
	s.Delete("/logout", h.Logout())

}
