package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/alexeyval/test-task-makves-group/internals/app/processors"
)

type UsersHandler struct {
	processor processors.UserService
}

func NewUsersHandler(processor processors.UserService) *UsersHandler {
	return &UsersHandler{
		processor: processor,
	}
}

func (handler *UsersHandler) List(w http.ResponseWriter, r *http.Request) {
	queryIDss, ok := r.URL.Query()["id"]
	if !ok || len(queryIDss) == 0 {
		WrapError(w, errors.New("no id in parameters"))

		return
	}

	ids := make([]int64, 0, len(queryIDss))

	for _, queryIDs := range queryIDss {
		for _, queryID := range strings.Split(queryIDs, ",") {
			id, err := strconv.ParseInt(queryID, 10, 64)
			if err != nil {
				WrapError(w, err)

				return
			}
			ids = append(ids, id)
		}
	}

	list, err := handler.processor.ListUser(ids)
	if err != nil {
		WrapError(w, err)

		return
	}

	m := map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}
