package response

// SuccessfulLoginResponse is struct to send the request response
type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
}
