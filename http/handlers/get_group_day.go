package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/internal/repo"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"net/http"
	"strconv"
)

func GetGroupDay(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	groupNumber := mux.Vars(r)["number"]
	offsetS := r.URL.Query().Get("offset")

	if offsetS == "" {
		offsetS = "0"
	}

	offset, err := strconv.Atoi(offsetS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	day, err := repo.GetGroupDay(groupNumber, offset)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(day); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Log.Error().Caller().Err(err).Msg("error was occurred while encoding a structure")
		return
	}

	w.WriteHeader(http.StatusOK)
}
