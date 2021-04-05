package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"write-async/internal/pkg/storage/inmemory"
)

func TestService(t *testing.T) {
	cases := []struct {
		name string
	}{
		{
			name: "should return pong from health",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("runnign test %s", c.name)
			db := inmemory.NewDatabase("tests.txt")
			s := NewServer(db)
			s.routes()
			req := httptest.NewRequest("GET", "/health", nil)
			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)
			if !reflect.DeepEqual(w.Result().StatusCode, http.StatusOK) {
				t.Errorf("unexpected status code from health")
			}
		})
	}
}

func TestAddJob(t *testing.T) {
	p := struct {
		Name string `json:"name"`
	}{
		Name: "Mat Ryer",
	}
	cases := []struct {
		name    string
		Payload struct {
			Name string `json:"name"`
		}
	}{
		{
			name:    "should be able to add a job",
			Payload: p,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("runnign test %s", c.name)
			db := inmemory.NewDatabase("tests_job.txt")
			s := NewServer(db)
			s.routes()

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(c.Payload)
			if err != nil {
				t.Errorf("errroo while jsoning %s", err)
			}
			log.Printf("before http handling")

			req := httptest.NewRequest(http.MethodPost, "/job", &buf)
			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)
			t.Logf("statuscode %s", w.Result().Status)
			if !reflect.DeepEqual(w.Result().StatusCode, http.StatusOK) {
				t.Errorf("not qequal")
			}
		})
	}
}
