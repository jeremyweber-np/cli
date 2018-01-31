package client

const (
	EventSourceType           = "eventSource"
	EventSourceFieldComponent = "component"
	EventSourceFieldHost      = "host"
)

type EventSource struct {
	Component string `json:"component,omitempty"`
	Host      string `json:"host,omitempty"`
}
