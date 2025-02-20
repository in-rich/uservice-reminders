package services

import "errors"

var (
	ErrInvalidReminderSelector     = errors.New("invalid reminder selector")
	ErrInvalidReminderUpdate       = errors.New("invalid reminder update")
	ErrRemindersUpdateLimitReached = errors.New("reminders update limit reached")
)
