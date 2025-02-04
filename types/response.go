package types

type Response struct {
	Status  int    `json:"status"`
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}
