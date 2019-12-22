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
	CreateAccount() http.HandlerFunc
	GetAccount() http.HandlerFunc
	DeleteAccount() http.HandlerFunc

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

func (ah *appHandler) CreateAccount() http.HandlerFunc {
	return ah.AccountHandler.CreateAccount
}

func (ah *appHandler) GetAccount() http.HandlerFunc {
	return ah.AccountHandler.GetAccount
}

func (ah *appHandler) DeleteAccount() http.HandlerFunc {
	return ah.AccountHandler.DeleteAccount
}

func (ah *appHandler) Login() http.HandlerFunc {
	return ah.AuthHandler.Login
}

func (ah *appHandler) Logout() http.HandlerFunc {
	return ah.AuthHandler.Logout
}
