package handlers

import (
	"encoding/json"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"main/internal/repo"
	"net/http"
)

func GetNames(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	names := repo.GetNames()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(names); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Log.Error().Caller().Err(err).Msg("error was occurred while encoding a structure")
		return
	}

	w.WriteHeader(http.StatusOK)
}
