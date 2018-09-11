package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

const currentPath = "/static"

func Provision(router *mux.Router) {
	router.PathPrefix(currentPath).Handler(
		http.StripPrefix(
			"/static",
			http.FileServer(
				http.Dir("./static"),
			),
		),
	)
}
