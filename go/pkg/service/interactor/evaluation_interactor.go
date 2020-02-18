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
	AddEvaluation(item string)(domain.Evaluation, error)
	UpdateEvaluation(evaluationID, item string)(domain.Evaluation, error)
	ShowEvaluations()(domain.Evaluation, error)
	DeleteEvaluation(evaluationID string)error

	CheckIsAdmin(userID string)(domain.Evaluation, error)
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

func (ei *evaluationRepository) UpdateEvaluation(evaluationID, item string)(evaluation domain.Evaluation, err error) {
	err := ei.EvaluationRepositoy.UpdateEvaluation(evaluationID, item)
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
		logger.Warn("thread checkAdmin: userのIDがない。 request userID=" userID)
		return false, err
	}
	if user.IsAdmin != 1 {
		logger.Warn("thread checkIsAdmin: 評価項目を操作する権限がない。 request userID=", userID)
		return false, domain.Unauthorized(errors.New("評価項目を操作する権限がありません"))
	} 
	return true, nil
}
