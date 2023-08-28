package errors

// APIMessage represents a generic message returned by an API
type APIMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// BusinessError represents the standard error structure for the services and use cases error.
type BusinessError struct {
	Msg   string
	Err   string
	Fatal bool
}

// APIError represents the standard error structure for the HTTP responses.
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Err     string `json:"error"`
}
