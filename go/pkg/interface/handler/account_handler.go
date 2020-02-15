package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"

	"github.com/gorilla/mux"
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

	SetTag(w http.ResponseWriter, r *http.Request)
	RemoveTag(w http.ResponseWriter, r *http.Request)
}

// NewAccountHandler create account handler
func NewAccountHandler(sh repository.SQLHandler, ph interactor.PasswordHandler) AccountHandler {
	return &accountHandler{
		AccountInteractor: interactor.NewAccountInteractor(
			repository.NewAccountRepository(sh),
			ph,
		),
	}
}

func (ah *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// requestの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(fmt.Sprintf("create account: %s", err.Error()))
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("create account: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// 登録
	user, err := ah.AccountInteractor.AddAccount(req.UserID, req.Name, req.Mail, req.Image, req.Profile, req.Password)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	var evaluationScores []getEvaluationScoreResponse
	for _, v := range user.Evaluations {
		evaluationScores = append(evaluationScores, getEvaluationScoreResponse{ID: v.ID, Item: v.Item, Score: v.Score})
	}

	// レスポンスの作成
	response.Success(w, &getAccountResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		Mail:        user.Mail,
		Image:       user.Image,
		Profile:     user.Profile,
		Evaluations: evaluationScores,
	})
}

func (ah *accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	// contextからuserIDの読み出し
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("get account: %s", err.Error()))
		response.HttpError(w, err)
		return
	}

	id, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("get account: %s", err.Error()))
		response.HttpError(w, err)
		return
	}

	// getData
	user, err := ah.AccountInteractor.ShowAccount(userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	user.Tags, err = ah.AccountInteractor.ShowTagsByUserID(id)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	user.Evaluations, err = ah.AccountInteractor.ShowEvaluationScoresByUserID(id)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// create response
	response.Success(w, convertAccountToResponse(user))
}

func (ah *accountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIDを取得
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	id, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("get account: %s", err.Error()))
		response.HttpError(w, err)
		return
	}

	// bodyの取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("update account: %s", err.Error()))
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("update account: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// 更新用データの作成
	user, err := ah.AccountInteractor.UpdateAccount(userID, req.UserID, req.Name, req.Mail, req.Image, req.Profile, req.Password)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	// tagの取得
	user.Tags, err = ah.AccountInteractor.ShowTagsByUserID(id)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	// evaluationの取得
	user.Evaluations, err = ah.AccountInteractor.ShowEvaluationScoresByUserID(id)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// response
	response.Success(w, convertAccountToResponse(user))
}

func (ah *accountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// delete
	err = ah.AccountInteractor.DeleteAccount(userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// response
	// TODO: ここnocontentでよくねえ...?
	response.Success(w, &deleteAccountResponse{
		Message: "success delete account",
	})
}

func (ah *accountHandler) SetTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("set tag(account): %s", err.Error()))
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req setAccountTagRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("set tag(account): %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	tag, err := ah.AccountInteractor.AddAccountTag(userID, req.Tag, req.CategoryID)
	if err != nil {
		logger.Error(fmt.Sprintf("set tag(account): %s", err.Error()))
		response.HttpError(w, err)
		return
	}

	response.Success(w, convertTagToResponse(tag))
}

func (ah *accountHandler) RemoveTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	vars := mux.Vars(r)
	tagID := vars["tagID"]

	err = ah.AccountInteractor.DeleteAccountTag(userID, tagID)
	if err != nil {
		logger.Error(fmt.Sprintf("remove account tag: %s", err.Error()))
		response.HttpError(w, err)
		return
	}

	response.NoContent(w)
}

type updateAccountRequest struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Image    string `json:"image"`
	Profile  string `json:"profile"`
	Password string `json:"password"`
}

type deleteAccountResponse struct {
	Message string `json:"message"`
}

type setAccountTagRequest struct {
	Tag        string `json:"tag"`
	CategoryID string `json:"category_id"`
}

type setAccountTagResponse struct {
	ID       string              `json:"id"`
	Tag      string              `json:"tag"`
	Category getCategoryResponse `json:"category"`
}

type getAccountResponse struct {
	UserID      string                       `json:"user_id"`
	Name        string                       `json:"name"`
	Mail        string                       `json:"mail"`
	Image       string                       `json:"image"`
	Profile     string                       `json:"profile"`
	Tags        []getTagResponse             `json:"tags"`
	Evaluations []getEvaluationScoreResponse `json:"evaluations"`
}

func convertAccountToResponse(user domain.User) (res getAccountResponse) {
	res.UserID = user.UserID
	res.Name = user.Name
	res.Mail = user.Mail
	res.Image = user.Image
	res.Profile = user.Profile

	for _, tag := range user.Tags {
		res.Tags = append(res.Tags, convertTagToResponse(tag))
	}
	for _, score := range user.Evaluations {
		res.Evaluations = append(
			res.Evaluations,
			getEvaluationScoreResponse{
				ID:    score.ID,
				Item:  score.Item,
				Score: score.Score,
			},
		)
	}

	return
}

// TODO: 消して?
type getEvaluationScoreResponse struct {
	ID    string `json:"id"`
	Item  string `json:"item"`
	Score int    `json:"score"`
}
