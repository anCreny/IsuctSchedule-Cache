package handlers

import (
	"encoding/json"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"github.com/gorilla/mux"
	"main/internal/repo"
	"net/http"
)

func GetTeacherCommonNames(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	name := mux.Vars(r)["name"]

	commonNames := repo.GetCommonTeachers(name)
	if err := json.NewEncoder(w).Encode(commonNames); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Log.Error().Caller().Err(err).Msg("error was occurred while encoding a structure")
		return
	}
	w.WriteHeader(200)
}
