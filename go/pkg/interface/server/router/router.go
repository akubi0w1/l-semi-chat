package router

import(
	"l-semi-chat/pkg/interface/server"
	"l-semi-chat/pkg/interface/handler"
)

func SetupRouter(s server.Server, h handler.AppHandler) {
	s.Post("/accounts", h.CreateAccount())
}
