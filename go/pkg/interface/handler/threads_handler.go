package handler
import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
)

type threadHandler static{
	threadHandler interactor.threadHandler
}

type ThreadHandler interactor{
	CreateThread(w http.ResponseWriter, r *http.Request)
	GetThread(w http.ResponseWriter, r *http.Request)
	UpdateThread(w http.ResponseWriter, r *http.Request)
	DeleteThread(w http.ResponseWriter, r *http.Request)
}

func NewThreadHandler(ti interactor.ThreadInteractor) ThreadHandler{
	return &ThreadHandler{
		ThreadInteractor: ti,
	}
} 

func (th *threadHandler) CreateThread(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Bpdy)
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req CreateThreadRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, InternalServerError(err))
		return
	}

	thread, err := ThreadInteractor.AddThread(req.Name, req.Description, req.LimitUsers, req.IsPublic)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}

	response.Success(w, &CreateThreadResponse{
		ID:				thread.ID
		Name:			thread.Name
		Description:	thread.Description
		LimitUsers:		thread.LimitUsers
		IsPublic:		thread.IsPublic
		CreatedAt:		thread.CreatedAt
		UpdatedAt:		thread.UpdatedAt
	})
}

type CreateThreadRequest struct {
	userID			string `json:"user_id"`
	Name			string `json:"name"`
	Description		string `json:"description"`
	LimitUsers		string `json:"limitUsers"`
	IsPublic		int    `json:"isPublic"`
}

type CreateThreadResponse struct {
	Id				string `json:"id"`
	Name			string `json:"name"`
	Description		string `json:"description"`
	LimitUsers		string `json:"limitUsers"`
	IsPublic		int    `json:"isPublic"`
	CreatedAt		string `json:"createdAt"`
	UpdatedAt		string `json:"updatedAt"`
}