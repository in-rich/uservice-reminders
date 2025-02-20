package dao

import (
	"context"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/uptrace/bun"
)

type DeleteReminderRepository interface {
	DeleteReminder(ctx context.Context, author string, target entities.Target, publicIdentifier string) (*entities.Reminder, error)
}

type deleteReminderRepositoryImpl struct {
	db bun.IDB
}

func (r *deleteReminderRepositoryImpl) DeleteReminder(
	ctx context.Context, author string, target entities.Target, publicIdentifier string,
) (*entities.Reminder, error) {
	reminder := &entities.Reminder{}

	_, err := r.db.NewDelete().
		Model(reminder).
		Where("author_id = ?", author).
		Where("public_identifier = ?", publicIdentifier).
		Where("target = ?", target).
		Returning("id").
		Exec(ctx)

	return reminder, err
}

func NewDeleteReminderRepository(db bun.IDB) DeleteReminderRepository {
	return &deleteReminderRepositoryImpl{
		db: db,
	}
}
