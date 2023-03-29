package response

import "time"

type (
	// User model
	User struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}

	// TODO: ლოგინი რო გაკეთდება ჩასახედია აქ
	SuccessfulLoginResponse struct {
		User      User   `json:"user"`
		AuthToken string `json:"token"`
	}
)
