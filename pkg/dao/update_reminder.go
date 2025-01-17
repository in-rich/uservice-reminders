package dao

import (
	"context"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
	"time"
)

type UpdateReminderData struct {
	Content   string
	ExpiredAt *time.Time
}

type UpdateReminderRepository interface {
	UpdateReminder(ctx context.Context, author string, target entities.Target, publicIdentifier string, data *UpdateReminderData) (*entities.Reminder, error)
}

type updateReminderRepositoryImpl struct {
	db bun.IDB
}

func (r *updateReminderRepositoryImpl) UpdateReminder(
	ctx context.Context, author string, target entities.Target, publicIdentifier string, data *UpdateReminderData,
) (*entities.Reminder, error) {
	reminder := &entities.Reminder{
		Content:   data.Content,
		UpdatedAt: lo.ToPtr(time.Now()),
		ExpiredAt: data.ExpiredAt,
	}

	res, err := r.db.NewUpdate().
		Model(reminder).
		Column("content", "updated_at", "expired_at").
		Where("author_id = ?", author).
		Where("public_identifier = ?", publicIdentifier).
		Where("target = ?", target).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrReminderNotFound
	}

	return reminder, nil
}

func NewUpdateReminderRepository(db bun.IDB) UpdateReminderRepository {
	return &updateReminderRepositoryImpl{
		db: db,
	}
}
