package response

// ErrorResponse is struct for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
	Code     int
	Message  string
	Response any
}
