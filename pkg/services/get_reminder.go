package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/in-rich/uservice-reminders/pkg/models"
)

type GetReminderService interface {
	Exec(ctx context.Context, selector *models.GetReminder) (*models.Reminder, error)
}

type getReminderServiceImpl struct {
	getReminderRepository dao.GetReminderRepository
}

func (s *getReminderServiceImpl) Exec(ctx context.Context, selector *models.GetReminder) (*models.Reminder, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(selector); err != nil {
		return nil, errors.Join(ErrInvalidReminderSelector, err)
	}

	reminder, err := s.getReminderRepository.GetReminder(ctx, selector.AuthorID, entities.Target(selector.Target), selector.PublicIdentifier)
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

func NewGetReminderService(getReminderRepository dao.GetReminderRepository) GetReminderService {
	return &getReminderServiceImpl{
		getReminderRepository: getReminderRepository,
	}
}
