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
	AddThread(string,string,string,string,int)(domain.Thread,error)
	UpdateThreads
	ShowThread()(domain.Thread, error)
	DeleteThreads
}

func (ti *threadInteractor) AddThread(userID, name, description, limitUsers, isPublic) (domain.Thread, error) {
	var thread domain.Thread

	user,err=ti.AuthRepository.FindUserByUserID(userID)
	if err != nil {
		return user, domain.Unauthorized(errors.New("Unauthorized"))
	}
	if name == "" {
		return thread, domain.BadRequest(errors.New("thread name is empty"))
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return thread, domain.InternalServerError(err)
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	err = ti.ThreadRepository.StoreThread(
		id.String(),
		name,
		description,
		limitUsers,
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

func (ti *threadInteractor) ShowThread()(domain.Thread, error) {
	return ti.ThreadRepository.FindThread()
}