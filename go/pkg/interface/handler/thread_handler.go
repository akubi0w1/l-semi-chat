package handler
import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"

	"l-semi-chat/pkg/interface/dcontext"
	"github.com/gorilla/mux"
)

type threadHandler static{
	ThreadInteractor interactor.ThreadInteractor
}

type ThreadHandler interactor{
	CreateThread(w http.ResponseWriter, r *http.Request)
	GetThreads(w http.ResponseWriter, r *http.Request)
	GetThreadByThreadID(w http.ResponseWriter, r *http.Request)
	UpdateThread(w http.ResponseWriter, r *http.Request)
	DeleteThread(w http.ResponseWriter, r *http.Request)
	GetParticipants(w http.ResponseWriter, r *http.Request)
	JoinParticipants(w http.ResponseWriter, r *http.Request)
	RemoveParticipants(w http.ResponseWriter, r *http.Request)
	RemoveParticipantsByUser(w http.ResponseWriter, r *http.Request)
}


func NewThreadHandler(ti interactor.ThreadInteractor) ThreadHandler{
	return &ThreadHandler{
		ThreadInteractor: ti,
	}
} 

func (th *threadHandler) CreateThread(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()
	id, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		logger.Warn("create thread: %s",err.Error())
		response.HttpError(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Bpdy)
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req createThreadRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	thread, err := th.ThreadInteractor.AddThread(req.Name, req.Description, req.LimitUsers, req.IsPublic)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	err = th.ThreadInteractor.AddUsersThread(thread.ID, id)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	
	admin, err := th.ThreadInteractor.ShowAccountByID(id)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domani.InternalServerError(err))
		return
	}

	admin.user.Tags, err = th.ThreadInteractor.ShowTagsByUserID(id)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	admin.user.Evaluations, err = th.ThreadInteractor.ShowEvaluationScoresByUserID(id)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}


	response.Success(w, &createThreadResponse{
		ID:             thread.ID,
		Name:           thread.Name,
		Description:    thread.Description,
		LimitUsers:     thread.LimitUsers,
		Admin:          admin,
		IsPublic:       thread.IsPublic,
		CreatedAt:      thread.CreatedAt,
		UpdatedAt:      thread.UpdatedAt,
	})
}

type createThreadRequest struct {
	Name			string `json:"name"`
	Description		string `json:"description"`
	LimitUsers		int    `json:"limit_users"`
	IsPublic		int    `json:"is_public"`
}

type createThreadResponse struct {
	ID				string             `json:"id"`
	Name			string             `json:"name"`
	Description		string             `json:"description"`
	LimitUsers		int                `json:"limitu_sers"`
	Admin			getAccountResponse `json:"admin"`
	IsPublic		int                `json:"is_public"`
	CreatedAt		string             `json:"created_at"`
	UpdatedAt		string             `json:"updated_at"`
}

func (th *threadHandler) GetThreads(w http.ResponseWriter, r *http.Request) {
	threads, err :=th.ThreadInteractor.ShowThreads()
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	var res getThreadsResponse

	for _, thread := range threads {
		ctx := r.Context()
		id, err := dcontext.GetIDFromContext(ctx)
		if err != nil {
			response.HttpError(w, domain.BadRequest(err))
			return
		}

		thread, err := th.ThreadInteractor.ShowThreadByID(threadID)
		if err != nil {
			response.HttpError(w, err)
			return
		}
		thread.user.Tags, err = th.ThreadInteractor.ShowTagsByUserID(id)
		if err != nil {
			response.HttpError(w, domain.InternalServerError(err))
			return
		}
		thread.user.Evaluations, err = th.ThreadInteractor.ShowEvaluationScoresByUserID(id)
		if err != nil {
			response.HttpError(w, domain.InternalServerError(err))
			return
		}
			res.Threads = append(res.Threads, convertThreadToResponse(thread))
	}

	response.Success(w, res)
}

func convertThreadToResponse(thread domain.Thread) (res getThreadResponse) {
	res.ID = thread.ID
	res.Name = thread.Name
	res.Description = thread.Description
	res.LimitUsers = thread.LimitUsers
	res.Admin = convertAccountToResponse(thread.UserID)
	res.IsPublic = thread.IsPublic
	res.CreateAt = thread.CreateAt
	res.UpdateAt = thread.UpdateAt

}

func (th *threadHandler) GetThreadByThreadID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["id"]
	if threadID == "" {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
	}

	ctx := r.Context()
	id, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}

	thread, err := th.ThreadInteractor.ShowThreadByID(threadID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	thread.user.Tags, err = th.ThreadInteractor.ShowTagsByUserID(id)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	thread.user.Evaluations, err = th.ThreadInteractor.ShowEvaluationScoresByUserID(id)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.Success(w,convertThreadToResponse(thread))
}

type getThreadsResponse struct {
	Threads  []getThreadResponse `jsom:"threads"`
}

type getThreadResponse struct {
	ID				string `json:"id"`
	Name			string `json:"name"`
	Description		string `json:"description"`
	LimitUsers		int    `json:"limit_users"`
	Admin			getAccountResponse `json:"admin"`
	IsPublic		int    `json:"is_public"`
	CreatedAt		string `json:"created_at"`
	UpdatedAt		string `json:"updated_at"`
}

func (th *threadHandler) UpdateThread(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	threadID := vars["id"]
	if threadID == "" {
		logger.Error(err)
		response.HttpError(w, domain.BadRequest(err))
	}

	ctx := r.Context()
	id, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateThreadRequest
	err = json.Unmarshal(body, req)
	if err != nil {
		logger.Error(fmt.Sprintf("update thread: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}


	thread, err := th.ThreadInteractor.UpdateThread(threadID, req.Name, req.Description, req.LimitUsers, req.IsPublic)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	thread.user.Tags, err = th.ThreadInteractor.ShowTagsByUserID(id)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	thread.user.Evaluations, err = th.ThreadInteractor.ShowEvaluationScoresByUserID(id)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.Success(w, convertThreadToResponse(thread))
}

type updateThreadRequest struct {
	Name			string `json:"name"`
	Description		string `json:"description"`
	LimitUsers		int    `json:"limit_users"`
	IsPublic		int    `json:"is_public"`
}

type updateThreadResponse struct {
	ID				string `json:"id"`
	Name			string `json:"name"`
	Description		string `json:"description"`
	LimitUsers		int    `json:"limitu_sers"`
	Admin			getAccountResponse `json:"admin"`
	IsPublic		int    `json:"is_public"`
	CreatedAt		string `json:"created_at"`
	UpdatedAt		string `json:"updated_at"`
}

func (th *threadHandler) DeleteThread(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	vars := mux.Vars(r)
	threadID := vars["id"]
	if threadID == "" {
		logger.Error(err)
		response.HttpError(w, domain.BadRequest(err))
	}

	_, err = th.ThreadInteractor.CheckIsAdmin(threadID, userID)
	if err != nil {
		response.HttpError(w, domain.Unauthorized(err))
		return
	}

	err = th.ThreadInteractor.DeleteThread(threadID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.NoContent(w)
}

func (th *threadHandler) GetParticipants(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["id"]
	if threadID == "" {
		logger.Error(err)
		response.HttpError(w, domain.BadRequest(err))
	}

	users, err := th.ThreadInteractor.ShowUserByThreadID(threadID)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var res getAccountsResponse
	for _, user := range users{
		// id, err := th.ThreadInteractor.GetIDByUserID(user.userID)
		ctx := r.Context()
		id, err := dcontext.GetIDFromContext(ctx)
		if err != nil {
			response.HttpError(w, domain.BadRequest(err))
			return
		}
		if err != nil {
			response.HttpError(w, domain.InternalServerError(err))
			return
		}
		user.Tags, err = th.ThreadInteractor.ShowTagsByUserID(id)
		if err != nil {
			response.HttpError(w, domain.InternalServerError(err))
			return
		}
		user.Evaluations, err = th.ThreadInteractor.ShowEvaluationScoresByUserID(id)
		if err != nil {
			response.HttpError(w, domain.InternalServerError(err))
			return
		}

		res.Users = append(res.Users, convertAccountToResponse(user))
	} 
	response.Success(w, res)
}

type getAccountsResponse struct {
	accounts []getAccountResponse `json:"accounts"`
}

func (th *threadHandler) JoinParticipants(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["id"]
	if threadID == "" {
		logger.fmt.Sprintf("thread error: not get threadID")
		// response.HttpError(w, domain.BadRequest(err))
	}
	ctx := r.Context()
	userID, err := dcontext.GetrIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}

	err = th.ThreadInteractor.JoinParticipantsByThreadID(threadID,userID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.Success(w, &deleteAccountResponse{
		Message: "success join thread",
	})
}

func (th *threadHandler)RemoveParticipants(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["id"]
	if threadID == "" {
		logger.errors.New(fmt.Sprintf("thread error: not get threadID"))
		response.HttpError(w, domain.BadRequest(err))
	}
	userID, err := dcontext.GetrIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	err :=th.ThreadInteractor.RemoveParticipantsByUserID(threadID, userID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.NoContent(w)
}


func (th *threadHandler)RemoveParticipantsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["t_id"]
	if threadID == "" {
		logger.fmt.Sprintf("thread error: not get threadID")
		response.HttpError(w, domain.BadRequest(err))
	}
	userID := vars["m_id"]
	if userID == "" {
		logger.Errorfmt.Sprintf("thread error: not get userID")
		response.HttpError(w, domain.BadRequest(err))
	}




	err :=th.ThreadInteractor.RemoveParticipantsByUser(threadID, ID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	response.NoContent(w)
}