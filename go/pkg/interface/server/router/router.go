package router

import (
	"l-semi-chat/pkg/interface/handler"
	"l-semi-chat/pkg/interface/server"
)

// SetupRouter routing
func SetupRouter(s server.Server, h handler.AppHandler) {

	// account
	s.Handle("/accounts", h.ManageAccount())

	// auth
	s.Handle("/login", h.Login())
	s.Handle("/logout", h.Logout())

	// tags
	s.Handle("/tags", h.ManageTag())
	s.Handle("/tags/{id}", h.ManageSpecificTag())

}
