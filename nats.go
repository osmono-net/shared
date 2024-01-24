package shared

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type AgentMessage[T any] struct {
	Id      string
	Payload T
}

func Send(url, channel string, T any) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}
	defer nc.Drain()
	jb, err := json.Marshal(T)
	if err != nil {
		return err
	}

	err = nc.Publish(channel, jb)
	if err != nil {
		return err
	}

	return nil
}
