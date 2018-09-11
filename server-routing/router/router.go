package router

import (
	"net/http"
	"os"
	"server-routing/logger"
	methodRoute "server-routing/router/method"
	staticRoute "server-routing/router/static"
	"time"

	"github.com/gorilla/mux"
)

// Create the router and return it
func Create() *mux.Router {
	router := mux.NewRouter()
	staticRoute.Provision(router)
	methodRoute.Provision(router)
	router.Use(routerLogger)
	return router
}

type logFormat struct {
	Method        string    `json:"method"`
	Host          string    `json:"host"`
	Path          string    `json:"path"`
	RemoteAddr    string    `json:"remote-addr"`
	UserAgent     string    `json:"user-agent"`
	ContentLength int64     `json:"content-length"`
	Referer       string    `json:"referer"`
	LocalHostname string    `json:"local-hostname"`
	Origin        string    `json:"origin"`
	Timestmap     time.Time `json:"timestamp"`
}

func routerLogger(next http.Handler) http.Handler {
	localHostname, _ := os.Hostname()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("access", "log", logFormat{
			r.Method,
			r.Host,
			r.RequestURI,
			r.RemoteAddr,
			r.UserAgent(),
			r.ContentLength,
			r.Referer(),
			localHostname,
			r.Header.Get("origin"),
			time.Now().UTC(),
		})
		next.ServeHTTP(w, r)
	})
}
