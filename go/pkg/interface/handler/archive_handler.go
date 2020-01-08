package handler

import (
	"encoding/json"
	"io/ioutil"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/interface/dcontext"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"net/http"
	"strings"
)

type archiveHandler struct {
	ArchiveInteractor interactor.ArchiveInteractor
}

type ArchiveHandler interface {
	CreateArchive(w http.ResponseWriter, r *http.Request)
	GetArchive(w http.ResponseWriter, r *http.Request)
	UpdateArchive(w http.ResponseWriter, r *http.Request)
	DeleteArchive(w http.ResponseWriter, r *http.Request)
}

func NewArchiveHandler(ai interactor.ArchiveInteractor) ArchiveHandler {
	return &archiveHandler{
		ArchiveInteractor: ai,
	}
}

func (ai *archiveHandler) GetArchive(w http.ResponseWriter, r *http.Request) {
	// queryの読み出し
	threadID := strings.TrimPrefix(r.URL.Path, "/threads/")
	threadID = strings.TrimSuffix(threadID, "/archives")

	archive, err := ai.ArchiveInteractor.ShowArchive(threadID)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	response.Success(w, &getArchiveResponse{
		ID:       archive.ID,
		Path:     archive.Path,
		IsPublic: archive.IsPublic,
		Password: archive.Password,
		Thread:   archive.Thread,
	})

}

type getArchiveResponse struct {
	ID       string        `json:"id"`
	Path     string        `json:"path"`
	IsPublic int           `json:"is_public"`
	Password string        `json:"password"`
	Thread   domain.Thread `json:"thread"`
}

func (ai *archiveHandler) CreateArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	threadID := strings.TrimPrefix(r.URL.Path, "/threads/")
	threadID = strings.TrimSuffix(threadID, "/archives")

	// TODO: thread管理者であるか確認

	// requestの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req createArchiveRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// threadの登録
	archive, err := ai.ArchiveInteractor.AddArchive(req.Password, threadID, req.IsPublic)

	response.Success(w, &createArchiveResponse{
		ID:       archive.ID,
		Path:     archive.Path,
		IsPublic: archive.IsPublic,
		Password: archive.Password,
		Thread:   archive.Thread,
	})

}

type createArchiveRequest struct {
	Password string `json:"password"`
	IsPublic int    `json:"is_public"`
}

type createArchiveResponse struct {
	ID       string        `json:"id"`
	Path     string        `json:"path"`
	IsPublic int           `json:"is_public"`
	Password string        `json:"password"`
	Thread   domain.Thread `json:"thread"`
}

func (ai *archiveHandler) UpdateArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	threadID := strings.TrimPrefix(r.URL.Path, "/threads/")
	threadID = strings.TrimSuffix(threadID, "/archives")

	// TODO: thread管理者であるか確認

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
	archive, err := ai.ArchiveInteractor.UpdateArchive(threadID, req.Password, req.IsPublic)

	response.Success(w, &updateArchiveResponse{
		ID:       archive.ID,
		Path:     archive.Path,
		IsPublic: archive.IsPublic,
		Password: archive.Password,
		Thread:   archive.Thread,
	})
}

type updateArchiveRequest struct {
	Password string `json:"password"`
	IsPublic int    `json:"is_public"`
}

type updateArchiveResponse struct {
	ID       string        `json:"id"`
	Path     string        `json:"path"`
	IsPublic int           `json:"is_public"`
	Password string        `json:"password"`
	Thread   domain.Thread `json:"thread"`
}

func (ai *archiveHandler) DeleteArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := dcontext.GetUserIDFromContext(ctx)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	threadID := strings.TrimPrefix(r.URL.Path, "/threads/")
	threadID = strings.TrimSuffix(threadID, "/archives")

	// TODO: thread管理者であるか確認

	err = ai.ArchiveInteractor.DeleteArchive(threadID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	response.NoContent(w)
}
