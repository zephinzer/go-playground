package router

import (
	"encoding/json"
	"net/http"
	"path"
	"server-routing/logger"

	"github.com/gorilla/mux"
)

const currentPath = "/param"

func Provision(router *mux.Router) {
	router.HandleFunc(path.Join(currentPath, "/{id}"), idParamHandler)
}

func idParamHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger.Infow("request", "request", params)
	response, _ := json.Marshal(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
