package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"write-async/internal/pkg/storage/inmemory"
	"write-async/internal/pkg/writeasync"

	"github.com/gorilla/mux"
)

type server struct {
	db          *inmemory.Database
	router      *mux.Router
	writerasync writeasync.WriterAsyncer
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func NewServer(db *inmemory.Database) *server {
	return &server{
		router: NewRouter(),
		db:     db,
	}
}

type JobType string

const (
	sayHi JobType = "say_hi"
)

func (s *server) Run() {
	port := "8000"
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	srv := &http.Server{
		Handler: s.router,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s.routes()
	log.Printf("started server at %s", port)
	log.Fatal(srv.ListenAndServe())
}

func (s *server) HandleAddJob() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("handling add job")
		var p writeasync.AddJobPayload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Printf("error while handling job %s", err)
		}
		log.Printf("calling addjob")

		s.db.AddJob("Hi", string(sayHi), p)
	}
}

func (s *server) HandleHealth() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("handling health")
		rw.Write([]byte("pong"))
	}
}

func (s *server) HandleIndex() http.HandlerFunc {
	log.Printf("hit index")
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("wtf"))
	}
}
