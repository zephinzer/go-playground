package main

import (
	"basic-server/config"
	"bytes"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	http.HandleFunc("/", defaultResponse)
	var PORT = viper.GetString("PORT")
	if len(PORT) == 0 {
		PORT = "9090"
	}
	var port bytes.Buffer
	port.WriteString(":")
	port.WriteString(PORT)
	fmt.Println("Listening at http://0.0.0.0" + port.String())
	http.ListenAndServe("0.0.0.0"+port.String(), nil)
}

func defaultResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
