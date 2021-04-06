package inmemory

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
)

func setupPubSubClient(t *testing.T) (c pubsub.Client) {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	os.Setenv("GOOGLE_CLOUD_PROJECT", "local")
	os.Setenv("PUBSUB_TOPIC", "topic_local")
	os.Setenv("PUBSUB_SUB", "sub_local")
	client, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}

	topicName := mustGetenv("PUBSUB_TOPIC")
	topic := client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %v doesn't exist - creating it", topicName)
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("created topic")
	subID := mustGetenv("PUBSUB_SUB")
	sub := client.Subscription(subID)
	exists, err = sub.Exists(ctx)
	if err != nil {
		log.Printf("error of getting the subscript exist does not exists")
	}
	if !exists {
		log.Printf("Sub %v doesn't exist - creating it", subID)
		_, err = client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
			Topic: topic,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return *client
}

func TestAddJob(t *testing.T) {
	client := setupPubSubClient(t)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	sub := client.Subscription(mustGetenv("PUBSUB_SUB"))

	cases := []struct {
		name     string
		expected string
	}{
		{
			name:     "hej",
			expected: "hello",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("running test %s", c.name)
			// Consume 10 messages.
			err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
				var jsonMessage map[string]interface{}
				json.Unmarshal(msg.Data, &jsonMessage)
				log.Printf("received %v", jsonMessage)
			})
			if err != nil {
				t.Errorf("recieved error while receiving from subscription with error %s", err)
			}
		})
	}
}
