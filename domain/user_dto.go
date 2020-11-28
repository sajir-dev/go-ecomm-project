package domain

import "time"

type User struct {
	UserID   string    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Dob      time.Time `json:"dob"`
	Phone    string    `json:"phone"`
}
