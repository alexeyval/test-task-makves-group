package storage

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/alexeyval/test-task-makves-group/internals/app/models"
)

type Storage interface {
	Upload(fileName string)
	GetUsersList(ids []int64) []models.User
}

type UsersStorage struct {
	users map[int64]*models.User
}

func NewStorage(batchSize int) Storage {
	return &UsersStorage{
		users: make(map[int64]*models.User, batchSize),
	}
}

func (storage *UsersStorage) Upload(fileName string) {
	log.Printf("Read file: %s", filepath.Base(fileName))

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

		newUser, err := models.NewUser(record)
		if err != nil {
			log.Fatalln(err)
		}

		storage.users[newUser.ID] = newUser
	}
}

func (storage *UsersStorage) GetUsersList(ids []int64) []models.User {
	items := make([]models.User, 0, len(ids))

	for _, id := range ids {
		item, ok := storage.users[id]
		if ok && item != nil {
			items = append(items, *item)
		}
	}

	return items
}
