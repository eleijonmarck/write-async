package inmemory

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"write-async/internal/pkg/writeasync"

	"cloud.google.com/go/pubsub"
)

type DatabaseWriter interface {
	AddJob(string) error
}

type Database struct {
	filepath string
}

func NewDatabase(filepath string) *Database {
	tempdir := os.TempDir()
	_, err := os.OpenFile(tempdir+"/"+filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("could not delete the database wither error: %s", err)
	}
	return &Database{
		filepath: filepath,
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

func (d *Database) AddJob(s string, jt string, payload writeasync.AddJobPayload) error {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	os.Setenv("GOOGLE_CLOUD_PROJECT", "local")
	os.Setenv("PUBSUB_TOPIC", "topic_local")
	os.Setenv("PUBSUB_SUB", "sub_local")
	client, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("created sub")
	res := client.Topic(mustGetenv("PUBSUB_SUB")).Publish(ctx, &pubsub.Message{Data: []byte("{\"greeting\" : \"hello\"}")})
	log.Printf("published")
	log.Printf("%v", res)
	if err != nil {
		return fmt.Errorf("receive: %v", err)
	}
	return nil
}
