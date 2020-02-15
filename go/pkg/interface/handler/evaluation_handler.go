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

	"l-semi-chat/pkg/interface/dcontext"
)

type evaluationHandler struct {
	EvaluationInteractor interactor.EvaluationInteractor
}

type EvaluationHandler interface {
	CreateEvaluation(w http.ResponseWriter, r *http.Request)
	GetEvaluations(w http.ResponseWriter, r *http.Request)
	UpdateEvaluation(w http.ResponseWriter, r *http.Request)
	DeleteEvaluation(w http.ResponseWriter, r *http.Request)
}

func NewEvaluationHandler(sh repository.SQLHandler) EvaluationHandler {
	return &evaluationHandler{
		EvaluationInteractor: interactor.NewEvaluationInteractor(
			repository.NewEvaluationRepository(sh),
		),
	}
}

func (eh *evaluationHandler) CreateEvaluation(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()
	userID, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("create evaluation: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	_, err :=eh.EvaluationInteractor.CheckIsAdmin(userID)
	if err != nil {
		response.HttpError(w, domain.Unauthorized(err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(fmt.Sprintf("create evaluation: %s", err.Error()))
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req createEvaluationRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("create evaluation: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	evaluation, err := eh.EvaluationInteractor.AddEvaluation(req.Item)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.Success(w, &createEvaluationResponse{
		ID:   evaluation.ID,
		Item: evaluation.Item,
	})
}

type createEvaluationRequest struct {
	Item string `json:"item"`
}

type createEvaluationResponse struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

func (eh *evaluationHandler) GetEvaluations(w http.ResponseWriter, r *http.Request) {
	evaluations, err := eh.EvaluationInteractor.ShowEvaluations()
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	var res GetEvaluationsResponse
	for _, evaluation := range evaluations {
		res.Evaluations = append(res.Evaluations, getEvaluationResponse{
			ID:   evaluation.ID,
			Item: evaluation.Item
		})
	} 
	response.Success(w, res)
}

type getEvaluationResponse struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

type getEvaluationsResponse struct {
	Evaluations []getEvaluationResponse `json:"evaluations"`
}

func (eh *evaluationHandler) UpdateEvaluation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("update evaluation: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	_, err :=eh.EvaluationInteractor.CheckIsAdmin(userID)
	if err != nil {
		response.HttpError(w, domain.Unauthorized(err))
		return
	}

	evaluationID := mux.Vars(r)["evaluationID"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn(fmt.Sprintf("update evaluation: %s", err.Error()))
		response.HttpError(w, domain.BadRequest(err))
		return
	}

	var req updateEvaluationRequest
	if err != nil {
		logger.Error(fmt.Sprintf("update evaluation: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	evaluation, err := eh.EvaluationInteractor.UpdateEvaluation(evaluationID ,erq.Item)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	response.Success(w, &updateEvaluationResponse{
		ID:   evaluation.ID,
		Item: evaluations.Item,
	})
}

type updateEvaluationRequest struct {
	Item string `json:"item"`
}

type updateEvaluationResponse struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

func (eh *evaluationHandler) DeleteEvaluation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := dcontext.GetIDFromContext(ctx)
	if err != nil {
		logger.Warn(fmt.Sprintf("delete evaluation: %s", err.Error()))
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	_, err :=eh.EvaluationInteractor.CheckIsAdmin(userID)
	if err != nil {
		response.HttpError(w, domain.Unauthorized(err))
		return
	}

	evaluationID := mux.Vars(r)["evaluationID"]

	err = eh.EvaluationHandler.DeleteEvaluation(evaluationID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	response.NoContent(w)
}