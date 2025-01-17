package dao

import (
	"context"
	"errors"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

type CreateReminderData struct {
	Content   string
	ExpiredAt *time.Time
}

type CreateReminderRepository interface {
	CreateReminder(ctx context.Context, author string, target entities.Target, publicIdentifier string, data *CreateReminderData) (*entities.Reminder, error)
}

type createReminderRepositoryImpl struct {
	db bun.IDB
}

func (r *createReminderRepositoryImpl) CreateReminder(
	ctx context.Context, author string, target entities.Target, publicIdentifier string, data *CreateReminderData,
) (*entities.Reminder, error) {
	reminder := &entities.Reminder{
		PublicIdentifier: publicIdentifier,
		AuthorID:         author,
		Target:           target,
		Content:          data.Content,
		ExpiredAt:        data.ExpiredAt,
	}

	if _, err := r.db.NewInsert().Model(reminder).Returning("*").Exec(ctx); err != nil {
		var pgErr pgdriver.Error
		if errors.As(err, &pgErr) && pgErr.IntegrityViolation() {
			return nil, ErrReminderAlreadyExists
		}

		return nil, err
	}

	return reminder, nil
}

func NewCreateReminderRepository(db bun.IDB) CreateReminderRepository {
	return &createReminderRepositoryImpl{
		db: db,
	}
}
