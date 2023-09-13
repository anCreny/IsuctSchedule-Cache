package handlers

import (
	"github.com/gorilla/mux"
	"main/internal/repo"
	"net/http"
)

func CheckGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	groupNumber, found := mux.Vars(r)["number"]

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found = repo.CheckGroup(groupNumber)

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
