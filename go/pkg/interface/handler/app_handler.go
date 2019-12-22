package handler

import (
	"net/http"

	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
)

type appHandler struct {
	AccountHandler AccountHandler
}

type AppHandler interface {
	// account
	CreateAccount() http.HandlerFunc
}

func NewAppHandler(sh repository.SQLHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(
			interactor.NewAccountInteractor(
				repository.NewAccountRepository(sh),
			),
		),
	}
}

func (ah *appHandler) CreateAccount() http.HandlerFunc {
	return ah.AccountHandler.CreateAccount
}
