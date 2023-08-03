package processors

import (
	"github.com/alexeyval/test-task-makves-group/internals/app/storage"

	"github.com/alexeyval/test-task-makves-group/internals/app/models"
)

type UserService interface {
	ListUser(ids []int64) ([]models.User, error)
}

type UsersProcessor struct {
	storage storage.Storage
}

func NewUsersProcessor(storage storage.Storage) UserService {
	return &UsersProcessor{
		storage: storage,
	}
}

func (p *UsersProcessor) ListUser(ids []int64) ([]models.User, error) {
	return p.storage.GetUsersList(ids), nil
}
