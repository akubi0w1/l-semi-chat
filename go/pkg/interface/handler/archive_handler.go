package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/interface/dcontext"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type archiveHandler struct {
	ArchiveInteractor interactor.ArchiveInteractor
}

// ArchiveHandler archive handler
type ArchiveHandler interface {
	CreateArchive(w http.ResponseWriter, r *http.Request)
	GetArchive(w http.ResponseWriter, r *http.Request)
	UpdateArchive(w http.ResponseWriter, r *http.Request)
	DeleteArchive(w http.ResponseWriter, r *http.Request)
}

// NewArchiveHandler create archive handler
func NewArchiveHandler(sh repository.SQLHandler, ph interactor.PasswordHandler) ArchiveHandler {
	return &archiveHandler{
		ArchiveInteractor: interactor.NewArchiveInteractor(
			repository.NewArchiveRepository(sh),
			repository.NewAccountRepository(sh),
			ph,
		),
	}
}

func (ai *archiveHandler) GetArchive(w http.ResponseWriter, r *http.Request) {
	// queryの読み出し
	threadID := mux.Vars(r)["threadID"]

	// header からpassword取得
	password := r.Header.Get("_password")

	archive, err := ai.ArchiveInteractor.ShowArchive(threadID, password)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	response.Success(w, convertArchiveToResponse(archive))
}

func (ai *archiveHandler) CreateArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// get threadID
	threadID := mux.Vars(r)["threadID"]

	// requestの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateArchiveRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	isAdmin, err := ai.ArchiveInteractor.CheckIsAdmin(threadID, userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	if !isAdmin {
		response.HttpError(w, domain.Unauthorized(errors.New("archiveを作成する権限がありません")))
		return
	}

	if req.IsPublic == 0 && req.Password == "" {
		response.HttpError(w, domain.BadRequest(errors.New("非公開アーカイブを作成する場合、パスワードは必須です")))
		return
	}

	// threadの登録
	_, err = ai.ArchiveInteractor.AddArchive(req.Password, threadID, req.IsPublic)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	archive, err := ai.ArchiveInteractor.ShowArchive(threadID, req.Password)
	if err != nil {
		response.HttpError(w, err)
	}

	response.Success(w, convertArchiveToResponse(archive))

}

func (ai *archiveHandler) UpdateArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	threadID := mux.Vars(r)["threadID"]

	isAdmin, err := ai.ArchiveInteractor.CheckIsAdmin(threadID, userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	if !isAdmin {
		response.HttpError(w, domain.Unauthorized(errors.New("archiveを作成する権限がありません")))
		return
	}

	// requestの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateArchiveRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// threadの登録
	_, err = ai.ArchiveInteractor.UpdateArchive(threadID, req.Password, req.IsPublic)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	archive, err := ai.ArchiveInteractor.ShowArchive(threadID, req.Password)
	if err != nil {
		response.HttpError(w, err)
	}

	response.Success(w, convertArchiveToResponse(archive))
}

func (ai *archiveHandler) DeleteArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	threadID := mux.Vars(r)["threadID"]

	isAdmin, err := ai.ArchiveInteractor.CheckIsAdmin(threadID, userID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	if !isAdmin {
		response.HttpError(w, domain.Unauthorized(errors.New("archiveを作成する権限がありません")))
		return
	}

	err = ai.ArchiveInteractor.DeleteArchive(threadID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	response.NoContent(w)
}

// responseのコンバータ
func convertArchiveToResponse(archive domain.Archive) (res getArchiveResponse) {
	res.ID = archive.ID
	res.Path = archive.Path
	res.IsPublic = archive.IsPublic
	res.Password = archive.Password
	res.Thread.ID = archive.Thread.ID
	res.Thread.Name = archive.Thread.Name
	res.Thread.Description = archive.Thread.Description
	res.Thread.LimitUsers = archive.Thread.LimitUsers
	res.Thread.CreatedAt = archive.Thread.CreatedAt
	res.Thread.UpdatedAt = archive.Thread.UpdatedAt
	res.Thread.Admin.UserID = archive.Thread.Admin.UserID
	res.Thread.Admin.Name = archive.Thread.Admin.Name
	res.Thread.Admin.Mail = archive.Thread.Admin.Mail
	res.Thread.Admin.Image = archive.Thread.Admin.Image
	res.Thread.Admin.Profile = archive.Thread.Admin.Profile
	for _, tag := range archive.Thread.Admin.Tags {
		res.Thread.Admin.Tags = append(
			res.Thread.Admin.Tags,
			getTagResponse{
				ID:  tag.ID,
				Tag: tag.Tag,
				Category: getCategoryResponse{
					ID:       tag.Category.ID,
					Category: tag.Category.Category,
				},
			},
		)
	}
	for _, score := range archive.Thread.Admin.Evaluations {
		res.Thread.Admin.Evaluations = append(
			res.Thread.Admin.Evaluations,
			getEvaluationScoreResponse{
				ID:    score.ID,
				Item:  score.Item,
				Score: score.Score,
			},
		)
	}
	return
}

type getArchiveRequest struct {
	Password string `json:"password"`
}

type getArchiveResponse struct {
	ID       string            `json:"id"`
	Path     string            `json:"path"`
	IsPublic int               `json:"is_public"`
	Password string            `json:"password"`
	Thread   getThreadResponse `json:"thread"`
}

type updateArchiveRequest struct {
	Password string `json:"password"`
	IsPublic int    `json:"is_public"`
}

type updateArchiveResponse struct {
	ID       string            `json:"id"`
	Path     string            `json:"path"`
	IsPublic int               `json:"is_public"`
	Password string            `json:"password"`
	Thread   getThreadResponse `json:"thread"`
}

// TODO: けして
type getThreadResponse struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	LimitUsers  int                `json:"limit_users"`
	IsPublic    int                `json:"isPublic"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
	Admin       getAccountResponse `json:"admin"`
}
