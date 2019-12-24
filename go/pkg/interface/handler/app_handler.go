package handler

import (
	"net/http"

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
	return ah.AccountHandler.ManageAccount
}

func (ah *appHandler) Login() http.HandlerFunc {
	return ah.AuthHandler.Login
}

func (ah *appHandler) Logout() http.HandlerFunc {
	return ah.AuthHandler.Logout
}
