package google

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type PubSubService interface {
	GetInstance() *pubsub.Client
	CreateTopic(string) (*pubsub.Topic, error)
}

type gcpubsub struct {
	client *pubsub.Client
}

func NewPubSubClient(projectID string) (PubSubService, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating Google Pub/Sub client: %v", err))
	}

	return &gcpubsub{client}, nil
}

func (g *gcpubsub) GetInstance() *pubsub.Client {
	return g.client 
}

func (g *gcpubsub) CreateTopic(topicID string) (*pubsub.Topic, error) {
	ctx := context.Background()

	topic, err := g.client.CreateTopic(ctx, topicID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating topic: %v", err))
	}

	return topic, nil
}