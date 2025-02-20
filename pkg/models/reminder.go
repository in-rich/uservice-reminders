package models

import "time"

type Reminder struct {
	ID               string
	Target           string
	PublicIdentifier string
	AuthorID         string
	Content          string
	UpdatedAt        *time.Time
	ExpiredAt        *time.Time
}
