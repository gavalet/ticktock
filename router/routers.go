package router

import (
	"cmd/ticktock/ctrls"
	"cmd/ticktock/utils/logger"

	"github.com/gorilla/mux"
)

func apiRoutes(router *mux.Router) {

	log := logger.NewLogger("Initialise")
	log.Info("Start API routes")

	router.HandleFunc("/ptlist", ctrls.GetTimestamps).Methods("GET")

}

func LoadRoutes() *mux.Router {
	routes := mux.NewRouter()
	apiRoutes(routes)

	return routes
}
