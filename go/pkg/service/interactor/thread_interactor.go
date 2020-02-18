package interactor
import (
	"erroes"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/service/repository"
	"time"

	"github.com/google/uuid"
)
type threadInteractor struct {
	ThreadRepository repository.ThreadRepository
	AuthRepository repository.AuthRepository
}

type threadInteractor interface{
	AddThread(userID, name, description string, limitUsers, isPublic int)(domain.Thread,error)
	UpdateThread(threadID, Name, Description string, LimitUsers, IsPublic int)(domain.Thread, error)
	ShowThreads()(domain.Threads, error)
	ShowThreadByID(threadID, userID string)(domain.Thread, error)
	DeleteThread(threadID string) error
	ShowUserByThreadID(threadID string)(domain.users_thread, error)
	JoinParticipantsByThreadID(threadID, userID string)error
	RemoveParticipantsByUser(threadID,userID  string)error
	RemoveParticipantsByUserID(threadID, userID string)error

	ShowAccountByID(userID string)(domain.Thread, error)
	AddUsersThread(threadID, userID string)error
	CheckIsAdmin(threadID, userID string)(bool, error)
	ShowTagsByUserID(userID string)(domain.Tags, error)
	ShowEvaluationScoresByUserID(userID string)(domain.EvaluationScores, error)
	ShowUserIDByThreadID(threadID string)(userID string, err error)
}

func (ti *threadInteractor) AddThread(userID, name, description string, limitUsers, isPublic int) (thread domain.Thread, err error) {
	
	
	if name == "" {
		return thread, domain.BadRequest(errors.New("thread name is empty"))
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return thread, domain.InternalServerError(err)
	}
	createdAt := time.Now()
	updatedAt := createdAt

	err = ti.ThreadRepository.StoreThread(
		id.String(),
		name,
		description,
		limitUsers,
		userID,
		isPublic,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return thread, domain.InternalServerError(err)
	}
	thread.Id = id
	thread.Name = name
	thread.Description = description
	thread.LimitUsers = limitUsers
	thread.IsPublic = isPublic
	thread.CreatedAt = createdAt
	thread.UpdatedAt = updatedAt

	return thread, nil
}

func (ti *threadInteractor) ShowThreads()(domain.Threads, error) {
	return ti.ThreadRepository.FindThreads()
}

func (ti *threadInteractor) ShowThreadByID(threadID, userID string)(thread domain.Thread, err error) {
	thread, err := ti.ThreadRepository.FindThreadByID(threadID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	thread.user.Tags, err = ti.ThreadRepository.FindTagsByUserID(userID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	thread.user.Evaluations, err = ti.ThreadRepository.FindEvaluationsByUserID(userID)
	if err != nil {
		response.HttpError(w, domain.InternalServerError(err))
		return
	}
	return
}

func (ti *threadInteractor) UpdateThread(threadID, Name, Description string, LimitUsers, IsPublic int)(thread domain.Thread, err error) {
	err = ti.ThreadRepository.UpdateThread(threadID, Name, Description, LimitUsers, IsPublic)
	if err != nil {
		return thread, domain.InternalServerError(err)
	}
	
	thread, err = ti.ThreadRepository.FindThreadByID(threadID)
	if err != nil {
		return thread, domain.InternalServerError(err)
	}
	return
}


func (ti *threadInteractor) ShowAccountByID(userID string)(domain.User, error) {
	return ti.ThreadRepository.FindAccountByID(userID)
}

func (ti *threadInteractor) DeleteThread(threadID string) error {
	return ti.ThreadRepository.DeleteThread(threadID)
}

func (ti *threadInteractor) CheckIsAdmin(threadID, userID string) (bool error) {
	thread, err := ti.ThreadRepository.FindThreadByID(threadID)
	if err != nil {
		logger.Error(domain.InternalServerError(err))
		return false, err
	}
	if thread.Admin.UserID != userID {
		logger.Warn("thread checkIsAdmin: スレッドを削除する権限がない。 request userID=", userID)
		return false, domain.Unauthorized(errors.New("スレッドを削除する権限がありません"))
	}
	return true, nil
}

func (ti *threadInteractor) ShowUserByThreadID(threadID string)(domain.Users, error) {
	return ti.ThreadRepository.FindUsersByThreadID(threadID)
}

func (ti *threadInteractor) JoinParticipantsByThreadID(threadID, UserID string) error {
	ID, err := uuid.NewRandom()
	if err != nil {
		logger.Error(domain.InternalServerError(err))
		return domain.InternalServerError(err)
	}
	return ti.ThreadRepository.JoinParticipantsByThreadID(ID, threadID, UserID)
}

func (ti *threadInteractor) RemoveParticipantsByUserID(threadID, UserID string) error {
	return ti.ThreadRepository.RemoveParticipantsByUserID(threadID, UserID)
}

func (ti *threadInteractor) AddUsersThread(threadID, UserID string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		logger.Error(domain.InternalServerError(err))
		return domain.InternalServerError(err)
	}
	return ti.ThreadRepository.StoreUsersThread(id,threadID,UserID)
}

func (ti *threadInteractor) ShowTagsByUserID(userID string) (domain.Tags, error) {
	return ti.ThreadRepository.FindTagsByUserID(userID)
}

func (ti *threadInteractor) ShowEvaluationScoresByUserID(userID string) (domain.EvaluationScores, error) {
	return ti.ThreadRepository.FindEvaluationsByUserID(userID)
}

func (ti *threadInteractor) ShowUserIDByThreadID(threadID string)(string, error) {
	return ti.ThreadRepository.FindUserByThreadID(threadID)
}

