package response

import "time"

type (
	// Player model
	Player struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}

	// TODO: ლოგინი რო გაკეთდება ჩასახედია აქ
	SuccessfulLoginResponse struct {
		Email     string `json:"email"`
		AuthToken string `json:"token"`
	}
)
