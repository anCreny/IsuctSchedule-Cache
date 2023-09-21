package handlers

import (
	"encoding/json"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"github.com/gorilla/mux"
	"main/internal/repo"
	"net/http"
)

func GetGroup(w http.ResponseWriter, r *http.Request) {
	groupNumber, found := mux.Vars(r)["number"]

	defer r.Body.Close()

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	group, found := repo.GetGroup(groupNumber)

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(group)
	if err != nil {
		logger.Log.Error().Caller().Err(err).Msg("error was occurred while encoding a structure")
	}

	w.WriteHeader(http.StatusOK)
}
