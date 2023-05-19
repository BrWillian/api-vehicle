package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	. "vehicle-search/repository"
)

var vehicleRepository VehiclesRepository

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

var GetPlaca = func(w http.ResponseWriter, r *http.Request) {
	placaParam := mux.Vars(r)
	placaString := placaParam["Placa"]
	vehicle, err := vehicleRepository.Find(placaString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Info(vehicle)
	respondWithJson(w, http.StatusOK, vehicle)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var GetVersion = func(w http.ResponseWriter, r *http.Request) {
	version := map[string]float32{"Version": 1.0}
	respondWithJson(w, http.StatusOK, version)
}
var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
