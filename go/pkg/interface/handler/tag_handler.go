package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"net/http"

	"github.com/gorilla/mux"
)

type tagHandler struct {
	TagInteractor interactor.TagInteractor
}

// TagHandler tag handler
type TagHandler interface {
	CreateTag(w http.ResponseWriter, r *http.Request)
	GetTagByTagID(w http.ResponseWriter, r *http.Request)
	GetTags(w http.ResponseWriter, r *http.Request)
}

// NewTagHandler create tagHandler
func NewTagHandler(ti interactor.TagInteractor) TagHandler {
	return &tagHandler{
		TagInteractor: ti,
	}
}

func (th *tagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req createTagRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	tag, err := th.TagInteractor.AddTag(req.Tag, req.CategoryID)
	if err != nil {
		response.HttpError(w, err)
		return
	}
	response.Success(w, &createTagResponse{
		ID:       tag.ID,
		Tag:      tag.Tag,
		Category: tag.Category,
	})
}

type createTagRequest struct {
	Tag        string `json:"tag"`
	CategoryID string `json:"category_id"`
}

type createTagResponse struct {
	ID       string          `json:"id"`
	Tag      string          `json:"tag"`
	Category domain.Category // TODO: ここ、handlerのカテゴリに変更して
}

func (th *tagHandler) GetTagByTagID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID := vars["id"]
	if tagID == "" {
		logger.Warn(fmt.Sprintf("tagID is empty. path = %s", r.URL.Path))
		response.HttpError(w, errors.New("tagID is empty"))
		return
	}

	tag, err := th.TagInteractor.ShowTagByID(tagID)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	// TODO: カテゴリ...
	response.Success(w, &getTagResponse{
		ID:       tag.ID,
		Tag:      tag.Tag,
		Category: tag.Category,
	})
}

type getTagResponse struct {
	ID       string          `json:"id"`
	Tag      string          `json:"tag"`
	Category domain.Category // TODO: ここ、handlerのカテゴリに変更して
}

func (th *tagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	// TODO: カテゴリへの対応...?
	tags, err := th.TagInteractor.ShowTags()
	if err != nil {
		response.HttpError(w, err)
		return
	}
	var res getTagsResponse
	for _, tag := range tags {
		res.Tags = append(res.Tags, getTagResponse{
			ID:       tag.ID,
			Tag:      tag.Tag,
			Category: tag.Category,
		})
	}

	response.Success(w, res)
}

type getTagsResponse struct {
	Tags []getTagResponse
}
