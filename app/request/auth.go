package request

// RegistationParams is struct to read the request body
type RegistationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginParams is struct to read the request body
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserDetails is struct used for user details
type UserDetails struct {
	Name     string
	Email    string
	Password string
}
