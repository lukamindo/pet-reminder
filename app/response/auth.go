package response

// SuccessfulLoginResponse is struct to send the request response
type SuccessfulLoginResponse struct {
	Email     string `json:"email"`
	AuthToken string `json:"token"`
}
