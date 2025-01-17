package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/in-rich/uservice-reminders/pkg/models"
)

type UpsertReminderService interface {
	Exec(ctx context.Context, reminder *models.UpsertReminder) (*models.Reminder, error)
}

type upsertReminderServiceImpl struct {
	updateReminderRepository dao.UpdateReminderRepository
	createReminderRepository dao.CreateReminderRepository
	deleteReminderRepository dao.DeleteReminderRepository
}

func (s *upsertReminderServiceImpl) Exec(ctx context.Context, reminder *models.UpsertReminder) (*models.Reminder, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(reminder); err != nil {
		return nil, errors.Join(ErrInvalidReminderUpdate, err)
	}

	// Delete reminder if content is empty.
	if reminder.Content == "" {
		deletedReminder, err := s.deleteReminderRepository.DeleteReminder(ctx, reminder.AuthorID, entities.Target(reminder.Target), reminder.PublicIdentifier)
		if err != nil {
			return nil, err
		}
		return &models.Reminder{
			ID: deletedReminder.ID.String(),
		}, nil
	}

	// Attempt to create a reminder.
	createdReminder, err := s.createReminderRepository.CreateReminder(
		ctx,
		reminder.AuthorID,
		entities.Target(reminder.Target),
		reminder.PublicIdentifier,
		&dao.CreateReminderData{
			Content:   reminder.Content,
			ExpiredAt: reminder.ExpiredAt,
		},
	)

	// Reminder was successfully created.
	if err == nil {
		return &models.Reminder{
			ID:               createdReminder.ID.String(),
			PublicIdentifier: createdReminder.PublicIdentifier,
			AuthorID:         createdReminder.AuthorID,
			Target:           string(createdReminder.Target),
			Content:          createdReminder.Content,
			UpdatedAt:        createdReminder.UpdatedAt,
			ExpiredAt:        createdReminder.ExpiredAt,
		}, nil
	}

	if !errors.Is(err, dao.ErrReminderAlreadyExists) {
		return nil, err
	}

	// Reminder already exists. Update it.
	updatedReminder, err := s.updateReminderRepository.UpdateReminder(
		ctx,
		reminder.AuthorID,
		entities.Target(reminder.Target),
		reminder.PublicIdentifier,
		&dao.UpdateReminderData{
			Content:   reminder.Content,
			ExpiredAt: reminder.ExpiredAt,
		},
	)

	if err != nil {
		return nil, err
	}

	return &models.Reminder{
		ID:               updatedReminder.ID.String(),
		PublicIdentifier: updatedReminder.PublicIdentifier,
		AuthorID:         updatedReminder.AuthorID,
		Target:           string(updatedReminder.Target),
		Content:          updatedReminder.Content,
		UpdatedAt:        updatedReminder.UpdatedAt,
		ExpiredAt:        updatedReminder.ExpiredAt,
	}, nil
}

func NewUpsertReminderService(
	updateReminderRepository dao.UpdateReminderRepository,
	createReminderRepository dao.CreateReminderRepository,
	deleteReminderRepository dao.DeleteReminderRepository,
) UpsertReminderService {
	return &upsertReminderServiceImpl{
		updateReminderRepository: updateReminderRepository,
		createReminderRepository: createReminderRepository,
		deleteReminderRepository: deleteReminderRepository,
	}
}
