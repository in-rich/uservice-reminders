package dao

import "errors"

var (
	ErrReminderAlreadyExists = errors.New("reminder already exists")
	ErrReminderNotFound      = errors.New("reminder not found")
)
