package shared

type RegisterAgentBody struct {
	HwId       string
	Name       string
	OsType     string
	Version    string
	CustomerId string
}
