package interactor

import (
	"errors"
	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/service/repository"

	"github.com/google/uuid"
)

type tagInteractor struct {
	TagRepository repository.TagRepository
}

type TagInteractor interface {
	AddTag(tagName, categoryID string) (domain.Tag, error)
	ShowTagByID(tagID string) (domain.Tag, error)
	ShowTags() (domain.Tags, error)
}

func NewTagInteractor(tr repository.TagRepository) TagInteractor {
	return &tagInteractor{
		TagRepository: tr,
	}
}

func (ti *tagInteractor) AddTag(tagName, categoryID string) (tag domain.Tag, err error) {
	// 入力のバリデーション
	if tagName == "" {
		return tag, domain.BadRequest(errors.New("tag is empty"))
	}
	if categoryID == "" {
		return tag, domain.BadRequest(errors.New("categoryID is empty"))
	}

	// uuidの生成
	id, err := uuid.NewRandom()
	if err != nil {
		return tag, domain.InternalServerError(err)
	}

	// TODO: カテゴリiDの正当性...

	err = ti.TagRepository.StoreTag(id.String(), tagName, categoryID)
	if err != nil {
		return
	}
	tag.ID = id.String()
	tag.Tag = tagName
	tag.Category = tag.Category
	// TODO: カテゴリの処理

	return
}

func (ti *tagInteractor) ShowTagByID(tagID string) (domain.Tag, error) {
	return ti.TagRepository.FindTagByTagID(tagID)
}

func (ti *tagInteractor) ShowTags() (domain.Tags, error) {
	return ti.TagRepository.FindTags()
}
