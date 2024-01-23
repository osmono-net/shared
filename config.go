package shared

type Config struct {
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
		HeartbeatInterval: 10,
		PatchMgmt:         true,
		EnableSsh:         true,
		EnableRdp:         true,
		EnablePing:        true,
		EnableEvents:      true,
		EnableLogging:     true,
	}
}
