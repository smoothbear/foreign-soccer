package handler

import (
	"net/http"
)

func (h _default) CreateNewUser(w http.ResponseWriter, r *http.Request) {

	accessor, err := h.accessManage.BeginTx()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Something went wrong."))
	}

	accessor.CreateUser()
}
