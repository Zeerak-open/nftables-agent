package payloads

type ErrorPayload struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Extra   string `json:"extra"`
}
