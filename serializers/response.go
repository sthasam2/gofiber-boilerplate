package serializers

// Response object as HTTP response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"body"`
}

// ErrorBody object
type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// FiberErrorResponse object
type FiberErrorResponse struct {
	Error []*ErrorBody `json:"errors"`
}

// ErrorResponse object
type ErrorResponse struct {
	Error []*Response `json:"errors"`
}
