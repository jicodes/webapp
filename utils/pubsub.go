package utils

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func PublishMessage(projectID, topicID, msg string) (string, error) {
    ctx := context.Background()
    client, err := pubsub.NewClient(ctx, projectID)
    if err != nil {
        return "", fmt.Errorf("pubsub: NewClient: %w", err)
    }
    defer client.Close()

    t := client.Topic(topicID)
    result := t.Publish(ctx, &pubsub.Message{
        Data: []byte(msg),
    })
    id, err := result.Get(ctx)
    if err != nil {
        return "", fmt.Errorf("pubsub: result.Get: %w", err)
    }

    return id, nil
}