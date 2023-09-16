package main

import (
	"fmt"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"github.com/gorilla/mux"
	"main/config"
	"main/http/handlers"
	"main/internal/repo"
	"net/http"
)

func main() {
	//init logger
	logger.Init()

	logger.Log.Info().Msg("Cache start initialization...")

	if err := config.Init(); err != nil {
		logger.Log.Error().Err(err).Msg("Couldn't init config")
		return
	}
	logger.Log.Info().Msg("Config successfully initialized!")

	if err := repo.Init(); err != nil {
		logger.Log.Error().Err(err).Msg("Couldn't init repository")
		return
	}
	logger.Log.Info().Msg("Repository successfully initialized!")

	router := mux.NewRouter()

	//groups
	router.HandleFunc("/api/group/{number}/day", handlers.GetGroupDay).Methods(http.MethodGet)
	router.HandleFunc("/api/group/{number}", handlers.GetGroup).Methods(http.MethodGet)

	//teachers
	router.HandleFunc("/api/teacher/{name}/day", handlers.GetTeacherDay).Methods(http.MethodGet)
	router.HandleFunc("/api/teacher/{name}", handlers.GetTeacher).Methods(http.MethodGet)

	//get all teachers names
	router.HandleFunc("/api/names", handlers.GetNames).Methods(http.MethodGet)

	//old group check rout
	router.HandleFunc("/api/check/{number}", handlers.CheckGroup).Methods(http.MethodGet)

	//check group existence
	router.HandleFunc("/api/check/group/{number}", handlers.CheckGroup).Methods(http.MethodGet)
	router.HandleFunc("/api/check/teacher/{name}", handlers.CheckTeacher).Methods(http.MethodGet)

	addr := fmt.Sprintf("%s:%s", config.Cfg.Server.Host, config.Cfg.Server.Port)

	logger.Log.Info().Msg("Cache server start listening on " + addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Log.Error().Err(err).Msg("Cache stopped with the error")
		return
	}

	logger.Log.Info().Msg("Cache stopped")
}
