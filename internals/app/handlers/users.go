package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/alexeyval/test-task-makves-group/internals/app/processors"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	processor *processors.UsersProcessor
}

func NewUsersHandler(processor *processors.UsersProcessor) *UsersHandler {
	return &UsersHandler{
		processor: processor,
	}
}

func (handler *UsersHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))

		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)

		return
	}

	user, err := handler.processor.FindUser(id)
	if err != nil {
		WrapError(w, err)

		return
	}

	m := map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
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
