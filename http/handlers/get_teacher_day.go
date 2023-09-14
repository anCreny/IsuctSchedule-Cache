package handlers

import (
	"encoding/json"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"github.com/gorilla/mux"
	"main/internal/repo"
	"net/http"
	"strconv"
	"strings"
)

func GetTeacherDay(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	name := mux.Vars(r)["name"]
	offsetS := r.URL.Query().Get("offset")

	if offsetS == "" {
		offsetS = "0"
	}

	offset, err := strconv.Atoi(offsetS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Log.Error().Caller().Err(err)
		return
	}

	name = strings.Join(strings.Split(name, "-"), " ")

	day, err := repo.GetTeacherDay(name, offset)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		logger.Log.Error().Caller().Err(err)
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
