package router

import (
	"l-semi-chat/pkg/interface/handler"
	"l-semi-chat/pkg/interface/server"
)

// SetupRouter urlのマッピングを行う
func SetupRouter(s server.Server, h handler.AppHandler) {

	// account
	s.Handle("/accounts", h.ManageAccount())
	s.Handle("/accounts/tags", h.ManageAccountTags())
	s.Handle("/accounts/tags/{tagID}", h.ManageAccountTag())

	// auth
	s.Handle("/login", h.Login())
	s.Handle("/logout", h.Logout())

	// tags
	s.Handle("/tags", h.ManageTags())
	s.Handle("/tags/{id}", h.ManageTag())

	// archive
	s.Handle("/threads/{threadID}/archives", h.ManageArchive())

}
