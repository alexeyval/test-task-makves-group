package processors

import (
	"errors"

	"github.com/alexeyval/test-task-makves-group/internals/app/models"
	"github.com/alexeyval/test-task-makves-group/internals/app/storage"
)

type UsersProcessor struct {
	storage *storage.UsersStorage
}

func NewUsersProcessor(storage *storage.UsersStorage) *UsersProcessor {
	return &UsersProcessor{
		storage: storage,
	}
}

func (p *UsersProcessor) FindUser(id int64) (models.User, error) {
	user, ok := p.storage.GetUserById(id)

	if !ok {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (p *UsersProcessor) ListUser(ids []int64) ([]models.User, error) {
	return p.storage.GetUsersList(ids), nil
}
