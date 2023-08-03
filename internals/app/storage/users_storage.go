package storage

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/alexeyval/test-task-makves-group/internals/app/models"
	"github.com/alexeyval/test-task-makves-group/internals/app/user"
)

type UsersStorage struct {
	users map[int64]*models.User
}

func NewStorage(batchSize int) *UsersStorage {
	return &UsersStorage{
		users: make(map[int64]*models.User, batchSize),
	}
}

func (storage *UsersStorage) Upload(fileName string) {
	log.Printf("read file: %s", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	csvReader := csv.NewReader(reader)
	csvReader.Comment = '#'

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		newUser, err := user.NewUser(record)
		if err != nil {
			log.Fatalln(err)
		}

		storage.users[newUser.ID] = newUser
	}
}

func (storage *UsersStorage) GetUserByID(id int64) (models.User, bool) {
	if item, ok := storage.users[id]; ok && item != nil {
		return *item, true
	}

	return models.User{}, false
}

func (storage *UsersStorage) GetUsersList(ids []int64) []models.User {
	items := make([]models.User, 0, len(ids))

	for _, id := range ids {
		item, ok := storage.users[id]
		if ok {
			items = append(items, *item)
		}
	}

	return items
}
