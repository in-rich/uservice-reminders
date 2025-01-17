package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/uptrace/bun"
)

type GetReminderByIDRepository interface {
	GetReminderByID(ctx context.Context, authorID string, reminderID uuid.UUID) (*entities.Reminder, error)
}

type getRemindersByIDRepositoryImpl struct {
	db bun.IDB
}

func (r *getRemindersByIDRepositoryImpl) GetReminderByID(ctx context.Context, authorID string, reminderID uuid.UUID) (*entities.Reminder, error) {
	reminder := &entities.Reminder{}

	err := r.db.NewSelect().
		Model(reminder).
		Where("author_id = ?", authorID).
		Where("id = ?", reminderID).
		Order("updated_at DESC").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrReminderNotFound
		}

		return nil, err
	}

	return reminder, nil
}

func NewGetReminderByIDRepository(db bun.IDB) GetReminderByIDRepository {
	return &getRemindersByIDRepositoryImpl{
		db: db,
	}
}
