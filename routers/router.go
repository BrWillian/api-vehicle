package routers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"vehicle-search/controllers"
)

func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/version", controllers.GetVersion).Methods("GET")
	s.HandleFunc("/healthcheck", controllers.HealthCheck).Methods("GET")
	AddVehicleRouter(s)

	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
