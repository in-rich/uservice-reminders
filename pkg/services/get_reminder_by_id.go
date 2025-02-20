package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/models"
)

type GetReminderByIDService interface {
	Exec(ctx context.Context, selector *models.GetReminderByID) (*models.Reminder, error)
}

type getReminderByIDServiceImpl struct {
	getReminderByIDRepository dao.GetReminderByIDRepository
}

func (s *getReminderByIDServiceImpl) Exec(ctx context.Context, selector *models.GetReminderByID) (*models.Reminder, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidReminderSelector, err)
	}

	reminderID, err := uuid.Parse(selector.ReminderID)
	if err != nil {
		return nil, errors.Join(ErrInvalidReminderSelector, err)
	}

	reminder, err := s.getReminderByIDRepository.GetReminderByID(ctx, selector.AuthorID, reminderID)
	if err != nil {
		return nil, err
	}

	return &models.Reminder{
		ID:               reminder.ID.String(),
		PublicIdentifier: reminder.PublicIdentifier,
		AuthorID:         reminder.AuthorID,
		Target:           string(reminder.Target),
		Content:          reminder.Content,
		UpdatedAt:        reminder.UpdatedAt,
		ExpiredAt:        reminder.ExpiredAt,
	}, nil
}

func NewGetReminderByIDService(getReminderByIDRepository dao.GetReminderByIDRepository) GetReminderByIDService {
	return &getReminderByIDServiceImpl{
		getReminderByIDRepository: getReminderByIDRepository,
	}
}
