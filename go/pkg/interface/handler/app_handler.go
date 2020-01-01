package handler

import (
	"errors"
	"net/http"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/interface/server/logger"
	"l-semi-chat/pkg/interface/server/middleware"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
)

type appHandler struct {
	AccountHandler AccountHandler
	AuthHandler    AuthHandler
}

// AppHandler ApplicationHandler
type AppHandler interface {
	// account
	ManageAccount() http.HandlerFunc

	// auth
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
}

// NewAppHandler create application handler
func NewAppHandler(sh repository.SQLHandler, ph interactor.PasswordHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(
			interactor.NewAccountInteractor(
				repository.NewAccountRepository(sh),
				ph,
			),
		),
		AuthHandler: NewAuthHandler(
			interactor.NewAuthInteractor(
				repository.NewAuthRepository(sh),
				ph,
			),
		),
	}
}

func (ah *appHandler) ManageAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.AccountHandler.GetAccount).ServeHTTP(w, r)
		case http.MethodPost:
			ah.AccountHandler.CreateAccount(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.AccountHandler.UpdateAccount).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.AccountHandler.DeleteAccount).ServeHTTP(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ah.AuthHandler.Login(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			ah.AuthHandler.Logout(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}
