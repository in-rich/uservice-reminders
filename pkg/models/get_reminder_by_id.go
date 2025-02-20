package models

type GetReminderByID struct {
	AuthorID   string `json:"authorID" validate:"required,max=255"`
	ReminderID string `json:"reminderID" validate:"required"`
}
