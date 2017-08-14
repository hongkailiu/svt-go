package http

import (
	"net/http"
	"github.com/hongkailiu/svt-go/log"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"path/filepath"
	"os"
)

var (
	server *http.Server
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	infoP := GetInfo()
	if json, error := json.Marshal(infoP); error != nil {
		log.Critical(error)
		http.Error(w, error.Error(), http.StatusInternalServerError)
	} else {
		w.Write(json)
	}
}

type logBody struct {
	Line string `json:"line"`
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%d - Log entries created.", http.StatusCreated)))
	decoder := json.NewDecoder(r.Body)
	var myLogBody logBody
	if err := decoder.Decode(&myLogBody); err != nil {
		log.Critical(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Info(myLogBody.Line)
	}
	defer r.Body.Close()
}

func foldersHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("p")
	log.Info("path: " + path)
	if path == "" {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		log.Info("dir: " + dir)
		path = dir
	}

	result := []string{}
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			result = append(result, err.Error())
		} else {
			basename := filepath.Base(path)
			result = append(result, basename)
		}
		return err
	})

	if json, error := json.Marshal(result); error != nil {
		log.Critical(error)
		http.Error(w, error.Error(), http.StatusInternalServerError)
	} else {
		w.Write(json)
	}
}

type Server struct {
	Port int
}

func (s Server) Run() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/logs", logsHandler).Methods("POST").HeadersRegexp("Content-Type", "application/(json)")
	r.HandleFunc("/folders", foldersHandler).Methods("GET")

	// Bind to a port and pass our router in
	server = &(http.Server{Addr: fmt.Sprintf(":%d", s.Port), Handler: r})
	log.Fatal(server.ListenAndServe())
}

func (s Server) Stop() {
	if server!=nil {
		server.Close()
		server = nil
	}
}

