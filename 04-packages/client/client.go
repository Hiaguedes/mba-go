package client

import "github.com/google/uuid"

type Client struct {
	id uuid.UUID
	name string
	isDumb bool
}

func NewClient(name string) *Client {
	return &Client{
		id: uuid.New(),
		name: name,
		isDumb: false,
	}
}

func (client *Client) GetID() uuid.UUID {
	return client.id
}

func (client *Client) SetDumb(isDumb bool) {
	client.isDumb = isDumb
}

func (client *Client) GetName() string {
	return client.name
}

func (client *Client) GetIsDumb() bool {
	return client.isDumb
}

