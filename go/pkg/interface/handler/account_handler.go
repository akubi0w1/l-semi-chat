package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"l-semi-chat/pkg/interface/auth"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"

	jwt "github.com/dgrijalva/jwt-go"
)

type accountHandler struct {
	AccountInteractor interactor.AccountInteractor
}

type AccountHandler interface {
	ManageAccount(w http.ResponseWriter, r *http.Request)

	createAccount(w http.ResponseWriter, r *http.Request)
	getAccount(w http.ResponseWriter, r *http.Request)
	updateAccount(w http.ResponseWriter, r *http.Request)
	deleteAccount(w http.ResponseWriter, r *http.Request)
}

// NewAccountHandler
func NewAccountHandler(ai interactor.AccountInteractor) AccountHandler {
	return &accountHandler{
		AccountInteractor: ai,
	}
}

func (ah *accountHandler) ManageAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.getAccount(w, r)
	} else if r.Method == http.MethodPost {
		ah.createAccount(w, r)
	} else if r.Method == http.MethodPut {
		ah.updateAccount(w, r)
	} else if r.Method == http.MethodDelete {
		ah.deleteAccount(w, r)
	}

	response.MethodNotAllowed(w, "method not allowed")
	return
}

func (ah *accountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
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

func (ah *accountHandler) getAccount(w http.ResponseWriter, r *http.Request) {
	// cookieからtokenを取得
	cookie, err := r.Cookie("x-token")
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	// check token
	token, err := auth.VerifyToken(cookie.Value)
	if err != nil {
		response.BadRequest(w, err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := claims["user_id"].(string)

	// getData
	user, err := ah.AccountInteractor.ShowAccount(userID)

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

func (ah *accountHandler) updateAccount(w http.ResponseWriter, r *http.Request) {
	// cookieからtokenを取得
	cookie, err := r.Cookie("x-token")
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	// check token
	token, err := auth.VerifyToken(cookie.Value)
	if err != nil {
		response.BadRequest(w, err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := claims["user_id"].(string)
	log.Println(userID)

	// bodyの取得

	// 更新用データの作成

	// update

	// response

}

func (ah *accountHandler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	// cookieからtokenを取得
	cookie, err := r.Cookie("x-token")
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	// check token
	token, err := auth.VerifyToken(cookie.Value)
	if err != nil {
		response.BadRequest(w, err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := claims["user_id"].(string)

	// delete
	err = ah.AccountInteractor.DeleteAccount(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
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
