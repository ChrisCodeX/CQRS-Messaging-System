package events

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/ChrisCodeX/Event-Architecture-CQRS-Go/models"
	"github.com/nats-io/nats.go"
)

/*Concrete implementation of NATS*/

type NatsEventStore struct {
	// Connection
	conn *nats.Conn
	// Subscription which will use the feed struct to subscribe to an event
	feedCreatedSub *nats.Subscription
	//
	feedCreatedChan chan CreatedFeedMessage
}

/*Methods of NatsEventStore*/
// Constructor
func NewNats(url string) (*NatsEventStore, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{conn: conn}, nil
}

// Close Method
func (n *NatsEventStore) Close() {
	// Verify if the connection exists
	if n.conn != nil {
		// Close the connection with the server
		n.conn.Close()
	}
	// Verify if the subscription exists
	if n.feedCreatedSub != nil {
		// Unsubscribe from event
		n.feedCreatedSub.Unsubscribe()
	}
	// Close channel of transmition of feeds
	close(n.feedCreatedChan)
}

// Message encoder to bytes
func (n *NatsEventStore) encodeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Publish Created Feed to services connected to NATS
func (n *NatsEventStore) PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	msg := CreatedFeedMessage{
		Id:          feed.Id,
		Title:       feed.Title,
		Description: feed.Description,
		CreatedAt:   feed.CreatedAt,
	}

	data, err := n.encodeMessage(msg)

	if err != nil {
		return err
	}

	//
	return n.conn.Publish(msg.Type(), data)
}