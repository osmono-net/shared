package nammu

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
