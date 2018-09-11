package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server-routing/logger"

	"github.com/gorilla/mux"
)

type methodResponse struct {
	Code             int                    `json:"code"`
	Method           string                 `json:"method"`
	Path             string                 `json:"path"`
	Data             string                 `json:"data"`
	UnstructuredData map[string]interface{} `json:"unstruct"`
}

const currentPath = "/method"

func Provision(router *mux.Router) {
	router.HandleFunc(currentPath, getMethodHandler).Methods("GET")
	router.HandleFunc(currentPath, postMethodHandler).Methods("POST")
	router.HandleFunc(currentPath, putMethodHandler).Methods("PUT")
	router.HandleFunc(currentPath, deleteMethodHandler).Methods("DELETE")
	router.HandleFunc(currentPath, optionsMethodHandler).Methods("OPTIONS")
}

func getMethodHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET " + currentPath))
}

func postMethodHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
	}
	var body map[string]interface{}
	w.Header().Set("Content-Type", "application/json")
	if err := json.Unmarshal(requestBody, &body); err != nil {
		json.Unmarshal([]byte(`{"error": "invalid json"}`), &body)
		json, err := json.Marshal(methodResponse{
			http.StatusBadRequest,
			"POST",
			currentPath,
			"",
			body,
		})
		if err != nil {
			logger.Error(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		json, err := json.Marshal(methodResponse{
			http.StatusOK,
			"POST",
			currentPath,
			"",
			body,
		})
		if err != nil {
			logger.Error(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func putMethodHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Fatal(err)
	}
	json, err := json.Marshal(methodResponse{
		200,
		"PUT",
		currentPath,
		string(body),
		nil,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func deleteMethodHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE " + currentPath))
}

func optionsMethodHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OPTIONS " + currentPath))
}
