package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/internal/repo"
	"main/logger"
	"net/http"
	"strings"
)

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	name, found := mux.Vars(r)["name"]

	defer r.Body.Close()

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name = strings.Join(strings.Split(name, "-"), " ")

	group, found := repo.GetTeacher(name)

	if !found {
		w.WriteHeader(http.StatusNotFound)
		logger.Log.Warn().Timestamp().Msg("Not found teacher: " + name)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Log.Error().Caller().Err(err).Msg("error was occurred while encoding a structure")
		return
	}

	w.WriteHeader(http.StatusOK)
}
