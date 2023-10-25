package response

import "time"

type (
	Pet struct {
		ID          int        `json:"id"`
		Name        string     `json:"name"`
		OwnerID     int        `json:"owner_id"`
		DateOfBirth *time.Time `json:"date_of_birth"`
		CreatedAt   time.Time  `json:"created_at"`
	}
	Pets []Pet
)
