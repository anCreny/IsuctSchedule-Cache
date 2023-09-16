package handlers

import (
	"github.com/gorilla/mux"
	"main/internal/repo"
	"net/http"
)

func CheckTeacher(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	name, found := mux.Vars(r)["name"]

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found = repo.CheckTeacher(name)

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
