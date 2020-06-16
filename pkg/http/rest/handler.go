package rest

import (
	"encoding/json"
	"go-skeleton/pkg/adding"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Handler returns handler for creators request
func Handler(a adding.Service) http.Handler {
	router := httprouter.New()
	router.GET("/", HealthHandler(a))
	router.POST("/creators", AddCreatorsHandler(a))
	return router
}

//HealthHandler does handle for checking health using GET
func HealthHandler(a adding.Service) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		s := a.CheckHealth()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s)
	}
}

//AddCreatorsHandler does handle for POST request
func AddCreatorsHandler(a adding.Service) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("New beer review added.")
	}
}