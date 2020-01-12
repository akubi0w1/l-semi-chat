package handler

import (
	"errors"
	"net/http"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/middleware"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
)

type appHandler struct {
	AccountHandler AccountHandler
	AuthHandler    AuthHandler
	ArchiveHandler ArchiveHandler
}

// AppHandler ApplicationHandler
type AppHandler interface {
	// account
	ManageAccount() http.HandlerFunc

	// auth
	Login() http.HandlerFunc
	Logout() http.HandlerFunc

	// archive
	ManageArchive() http.HandlerFunc
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
		ArchiveHandler: NewArchiveHandler(
			interactor.NewArchiveInteractor(
				repository.NewArchiveRepository(sh),
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

func (ah *appHandler) ManageArchive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.ArchiveHandler.GetArchive).ServeHTTP(w, r)
		case http.MethodPost:
			middleware.Authorized(ah.ArchiveHandler.CreateArchive).ServeHTTP(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.ArchiveHandler.UpdateArchive).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.ArchiveHandler.DeleteArchive).ServeHTTP(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}
