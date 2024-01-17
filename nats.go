package shared

type AgentMessage[T any] struct {
	Id      string
	Payload T
}
