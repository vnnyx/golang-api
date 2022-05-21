package entity

import "time"

type Customer struct {
	Id        int       `json:"id"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
