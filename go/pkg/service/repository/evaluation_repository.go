package repository

import (
	"fmt"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
)

type evaluationRepository struct {
	SQLHandler SQLHandler
}

type EvaluationRepository interface {
	StoreEvaluation(string, string) error
	FindEvaluations()(evaluations domain.Evaluations, err error)
	UpdateEvaluation(string,string) error
	DeleteEvaluation(string)error

	FindEvaluationByID(string)(domain.Evaluation, error)
	FindUserByID(string)(domain.User, error)
}

func NewEvaluationRepository(sh SQLHandler) EvaluationRepository {
	return &evaluationRepository{
		SQLHandler: sh,
	}
}

func (er *evaluationRepository) StoreEvaluation(ID, item string) error {
	_, err := er.SQLHandler.Execute(
		"INSERT INTO evaluations(id, item) VALUES (?,?)",
		ID,
		item,
	)

	if err != nil {
		logger.Error(fmt.Sprintf("create evaluation: %s", err.Error()))
		return domain.InternalServerError(err)
	}
	return nil
}

func (er *evaluationRepository) FindEvaluations()(evaluations domain.Evaluations,err error) {
	rows, err := er.SQLHandler.Query(
		`SELECT id, item
		FROM evaluations`
	)
	if err != nil {
		logger.Error(fmt.Sprintf("find evaluations: %s", err.Error()))
		return evaluations, domain.InternalServerError(err)
	}

	for rows.Next(){
		var evaluation domain.Evaluation
		if err = rows.Scan(&evaluation.ID, &evaluation.Item); err !=nil {
			continue
		}
		evaluations = append(evaluations, evaluation)
	}
	return
}

func (er *evaluationRepository) UpdateEvaluation(evaluationID, item string)(evaluations domain.Evaluations,err error) {
	query := "UPDATE evaluations"
	var values := []interface{}
	if item != "" {
		query += "SET item=?"
		values = append(values, item)
	}
	query += "WHERE id=?;"
	values = append(values, query)

	_, err := er.SQLHandler.Execute(query, values...)

	return domain.InternalServerError(err)
}

func (er *evaluationRepository) DeleteEvaluation(evaluationID string)error {
	_, err = er.SQLHandler.Execute("DELETE FROM evaluations id=?", evaluationID)
	return domain.InternalServerError(err)
}


func (er *evaluationRepository) FindUserByID(userID string) (user domain.User, err error) {
	row := er.SQLHandler.QueryRow("SELECT id, user_id, name, mail, image, is_admin, profile FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.UserID, &user.Name, &user.Mail, &user.Image, &user.IsAdmin, &user.Profile); err != nil {
		return user, domain.InternalServerError(err)
	}
	return user, nil
}

func (er *evaluationRepository) FindEvaluationByID(evaluationID string)(evaluation domain.Evaluation,err error) {
	row := er.SQLHandler.QueryRow("SELECT id, item FROM evaluations WHERE id=?", evaluationID)
	if err = row.Scan(&evaluation.ID, &evaluation.Item); err != nil {
		return evaluation, domain.InternalServerError(err)
	}
	return evaluation, nil
}