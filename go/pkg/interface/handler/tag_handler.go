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
	"l-semi-chat/pkg/service/repository"
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
func NewTagHandler(sh repository.SQLHandler) TagHandler {
	return &tagHandler{
		TagInteractor: interactor.NewTagInteractor(
			repository.NewTagRepository(sh),
		),
	}
}

func (th *tagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn("tag create: ", err)
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req updateTagRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error("tag create: ", err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	tag, err := th.TagInteractor.AddTag(req.Tag, req.CategoryID)
	if err != nil {
		response.HttpError(w, err)
		return
	}

	response.Success(w, convertTagToResponse(tag))
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

	response.Success(w, convertTagToResponse(tag))
}

func (th *tagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := th.TagInteractor.ShowTags()
	if err != nil {
		response.HttpError(w, err)
		return
	}
	var res getTagsResponse
	for _, tag := range tags {
		res.Tags = append(res.Tags, convertTagToResponse(tag))
	}

	response.Success(w, res)
}

type getTagResponse struct {
	ID       string              `json:"id"`
	Tag      string              `json:"tag"`
	Category getCategoryResponse `json:"category"`
}

type getTagsResponse struct {
	Tags []getTagResponse `json:"tags"`
}

type updateTagRequest struct {
	Tag        string `json:"tag"`
	CategoryID string `json:"category_id"`
}

type updateTagResponse struct {
	ID       string              `json:"id"`
	Tag      string              `json:"tag"`
	Category getCategoryResponse `json:"category"`
}

func convertTagToResponse(tag domain.Tag) (res getTagResponse) {
	res.ID = tag.ID
	res.Tag = tag.Tag
	res.Category = getCategoryResponse{
		ID:       tag.Category.ID,
		Category: tag.Category.Category,
	}
	return
}
