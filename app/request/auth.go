package request

// RegistationParams is struct to read the request body
type RegistationParams struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// LoginParams is struct to read the request body
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// FacebookUserDetails is struct used for user details
type FacebookUserDetails struct {
	ID    string
	Name  string
	Email string
}
