package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"

	"l-semi-chat/pkg/interface/dcontext"
)

type accountHandler struct {
	AccountInteractor interactor.AccountInteractor
}

// AccountHandler CRUD of account
type AccountHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}

// NewAccountHandler create account handler
func NewAccountHandler(ai interactor.AccountInteractor) AccountHandler {
	return &accountHandler{
		AccountInteractor: ai,
	}
}

func (ah *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// requestの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req createAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// 登録
	user, err := ah.AccountInteractor.AddAccount(req.UserID, req.Name, req.Mail, req.Image, req.Profile, req.Password)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// レスポンスの作成
	response.Success(w, &createAccountResponse{
		UserID:  user.UserID,
		Name:    user.Name,
		Mail:    user.Mail,
		Image:   user.Image,
		Profile: user.Profile,
	})
}

type createAccountRequest struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Image    string `json:"image"`
	Profile  string `json:"profile"`
	Password string `json:"password"`
}

type createAccountResponse struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Image   string `json:"image"`
	Profile string `json:"profile"`
	// Tags
	// Evaluations
}

func (ah *accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	// contextからuserIDの読み出し
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	// getData
	user, err := ah.AccountInteractor.ShowAccount(userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// create response
	response.Success(w, &getAccountResponse{
		UserID:  user.UserID,
		Name:    user.Name,
		Mail:    user.Mail,
		Image:   user.Image,
		Profile: user.Profile,
	})

}

type getAccountResponse struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Image   string `json:"image"`
	Profile string `json:"profile"`
	// Tags
	// Evaluations
}

func (ah *accountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIDを取得
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	// bodyの取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// 更新用データの作成
	user, err := ah.AccountInteractor.UpdateAccount(userID, req.UserID, req.Name, req.Mail, req.Image, req.Profile, req.Password)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// response
	response.Success(w, &updateAccountResponse{
		UserID:  user.UserID,
		Name:    user.Name,
		Mail:    user.Mail,
		Image:   user.Image,
		Profile: user.Profile,
	})

}

type updateAccountRequest struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Image    string `json:"image"`
	Profile  string `json:"profile"`
	Password string `json:"password"`
}

type updateAccountResponse struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Image   string `json:"image"`
	Profile string `json:"profile"`
	// Tags
	// Evaluations
}

func (ah *accountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	// delete
	err := ah.AccountInteractor.DeleteAccount(userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// response
	response.Success(w, &deleteAccountResponse{
		Message: "success delete account",
	})
}

type deleteAccountResponse struct {
	Message string `json:"message"`
}
