package handler

import (
	"net/http"

	"l-semi-chat/pkg/interface/server/middleware"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
)

type appHandler struct {
	AccountHandler AccountHandler
	AuthHandler    AuthHandler
}

type AppHandler interface {
	// account
	ManageAccount() http.HandlerFunc

	// auth
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
}

func NewAppHandler(sh repository.SQLHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(
			interactor.NewAccountInteractor(
				repository.NewAccountRepository(sh),
			),
		),
		AuthHandler: NewAuthHandler(
			interactor.NewAuthInteractor(
				repository.NewAuthRepository(sh),
			),
		),
	}
}

func (ah *appHandler) ManageAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var handler http.HandlerFunc
		if r.Method == http.MethodGet {
			handler = middleware.Authorized(ah.AccountHandler.GetAccount)
			handler(w, r)
			return

		} else if r.Method == http.MethodPost {
			ah.AccountHandler.CreateAccount(w, r)
			return

		} else if r.Method == http.MethodPut {
			handler = middleware.Authorized(ah.AccountHandler.UpdateAccount)
			handler(w, r)
			return

		} else if r.Method == http.MethodDelete {
			handler = middleware.Authorized(ah.AccountHandler.DeleteAccount)
			handler(w, r)
			return

		}
		response.MethodNotAllowed(w, "Method not allowed")
	}
}

func (ah *appHandler) Login() http.HandlerFunc {
	return ah.AuthHandler.Login
}

func (ah *appHandler) Logout() http.HandlerFunc {
	return ah.AuthHandler.Logout
}
