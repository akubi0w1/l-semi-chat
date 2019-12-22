package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
)

type accountHandler struct {
	AccountInteractor interactor.AccountInteractor
}

type AccountHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
}

// NewAccountHandler
func NewAccountHandler(ai interactor.AccountInteractor) AccountHandler {
	return &accountHandler{
		AccountInteractor: ai,
	}
}

func (ah *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// requestの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	var req createAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	// 登録
	user, err := ah.AccountInteractor.AddAccount(req.UserID, req.Name, req.Mail, req.Image, req.Profile, req.Password)
	if err != nil {
		response.InternalServerError(w, err.Error())
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
