package http

import (
	"net/http"
	"github.com/hongkailiu/svt-go/log"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	infoP := GetInfo()
	if json, error := json.Marshal(infoP); error != nil {
		log.Fatal(error)
		http.Error(w, error.Error(), 500)
	} else {
		w.Write(json)
	}
}


type Server struct {
	Port  int
}


func (s Server) Run() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", rootHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), r))
}

