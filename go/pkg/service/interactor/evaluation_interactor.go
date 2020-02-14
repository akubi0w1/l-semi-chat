package interactor

import (
	"errors"
	"fmt"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/service/repository"
	"time"

	"github.com/google/uuid"
)

type evaluationInteractor struct {
	EvaluationRepository repository.EvaluationRepository
}

type EvaluationInteractor interface {
	AddEvaluation(string)(domain.Evaluation, error)
	UpdateEvaluation(string)(domain.Evaluation, error)
	ShowEvaluations()(domain.Evaluation, error)
	DeleteEvaluation(string)error

	CheckIsAdmin(string)(domain.Evaluation, error)
}

func (ei *evaluationRepository) AddEvaluation(item string)(evaluation domain.Evaluation, err error) {
	if item == "" {
		logger.Warn("create evaluation: evaluation is empty")
		return tag, domain.BadRequest(errors.New("evaluation is empty"))
	}

	id, err := uuid.NewRandom()
	if err != nil {
		logger.Error(fmt.Sprintf("create evaluation: %s", err.Error()))
		return evaluation, domain.InternalServerError(err)
	}

	err = ei.EvaluationRepository.StoreEvaluation(id.String(), item)
	if err != nil {
		return
	}

	evaluation.ID = id.String()
	evaluation.Item = item
	return

}

func (ei *evaluationRepository) ShowEvaluations()(domain.Evaluation, error) {
	return ei.EvaluationRepository.FindEvaluations()
}

func (ei *evaluationRepository) UpdateEvaluation(evaluationID, item string)(domain.Evaluation, error) {
	err = ei.EvaluationRepositoy.UpdateEvaluation(evaluationID, item)
	if err != nil {
		logger.Error(fmt.Sprintf("update evaluation: %s", err.Error()))
		return
	}
	
	evaluation, err = ei.EvaluationRepository.FindEvaluationByID(evaluationID)
	if err != nil {
		logger.Error(fmt.Sprintf("update evaluation: %s", err.Error()))
	}
	return
}

func (ei *evaluationRepository) DeleteEvaluation(evaluationID string) error {
	return ei.EvaluationRepository.DeleteEvaluation(evaluationID)
}


func (ei *evaluationRepository) CheckIsAdmin(userID string)(bool, error) {
	user, err := ei.EvaluationRepository.FindUserByID(userID)
	if err != nil {
		return false, err
	}
	if user.IsAdmin != 1 {
		return false, domain.Unauthorized(err)
	} 
	return true, nil
}
