package routers

import (
	"github.com/gorilla/mux"
	"vehicle-search/controllers"
)

func AddVehicleRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/vehicle/").Subrouter()
	s.HandleFunc("/get/{Placa}", controllers.GetPlaca).Methods("GET")
	return s
}
