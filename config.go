package shared

import "github.com/google/uuid"

type Config struct {
	Id        string
	PatchMgmt bool
	// Time interval between heartbeat requests
	// back to the C2 server, in seconds.
	HeartbeatInterval uint

	EnableLogging bool
	EnableEvents  bool
	EnableRdp     bool
	EnablePing    bool
	EnableSsh     bool
}

func NewConfig() Config {
	return Config{
		Id:                uuid.New().String(),
		HeartbeatInterval: 10,
	}
}
