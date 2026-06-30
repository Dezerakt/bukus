package entity

import "time"

type Book struct {
	ID        string    `json:"Id"`
	Name      string    `json:"name" example:"Crime And Punishment"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"createdAt" example:"2020-01-01T00:00:00+00:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2020-01-01T00:00:00+00:00"`
}
