package models

import "time"

type Task struct {
	Name        string
	Status      string
	Description string
	CreatedDate time.Time
	CreatedBy   string
	UpdatedDate time.Time
	UpdatedBy   string
}

type User struct {
	Name  string
	Email string
}
