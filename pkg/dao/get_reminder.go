package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/uptrace/bun"
)

type GetReminderRepository interface {
	GetReminder(ctx context.Context, author string, target entities.Target, publicIdentifier string) (*entities.Reminder, error)
}

type getReminderRepositoryImpl struct {
	db bun.IDB
}

func (r *getReminderRepositoryImpl) GetReminder(
	ctx context.Context, author string, target entities.Target, publicIdentifier string,
) (*entities.Reminder, error) {
	reminder := new(entities.Reminder)

	err := r.db.NewSelect().
		Model(reminder).
		Where("author_id = ?", author).
		Where("public_identifier = ?", publicIdentifier).
		Where("target = ?", target).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrReminderNotFound
		}

		return nil, err
	}

	return reminder, nil
}

func NewGetReminderRepository(db bun.IDB) GetReminderRepository {
	return &getReminderRepositoryImpl{
		db: db,
	}
}
