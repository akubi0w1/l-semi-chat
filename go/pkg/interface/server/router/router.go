package router

import (
	"l-semi-chat/pkg/interface/handler"
	"l-semi-chat/pkg/interface/server"
)

// SetupRouter urlのマッピングを行う
func SetupRouter(s server.Server, h handler.AppHandler) {

	// account
	s.Handle("/accounts", h.ManageAccount())

	// auth
	s.Handle("/login", h.Login())
	s.Handle("/logout", h.Logout())

	// archive
	s.Handle("/threads/{threadID}/archives", h.ManageArchive())

}
