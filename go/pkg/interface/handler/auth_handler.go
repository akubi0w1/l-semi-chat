package handler

import (
	"encoding/json"
	"io/ioutil"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
	"net/http"

	"l-semi-chat/pkg/interface/auth"
)

type authHandler struct {
	AuthInteractor interactor.AuthInteractor
}

// AuthHandler login, logout
type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// NewAuthHandler create authorized handler
func NewAuthHandler(sh repository.SQLHandler, ph interactor.PasswordHandler) AuthHandler {
	return &authHandler{
		AuthInteractor: interactor.NewAuthInteractor(
			repository.NewAuthRepository(sh),
			ph,
		),
	}
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	// bodyの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req loginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// 認証処理
	user, err := ah.AuthInteractor.Login(req.UserID, req.Password)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}

	// tokenの作成
	token, err := auth.CreateToken(user.ID, user.UserID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// cookieに載せる
	cookie := http.Cookie{
		Name:  "x-token",
		Value: token,
	}
	http.SetCookie(w, &cookie)

	// レスポンス
	response.Success(w, &loginResponse{
		Message: "login success",
	})

}

func (ah *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// check cookie
	cookie, err := r.Cookie("x-token")
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}

	// check token
	_, err = auth.VerifyToken(cookie.Value)
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}

	// cookieの無効化
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	response.NoContent(w)
}

type loginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type loginResponse struct {
	Message string `json:"message"`
}
