package services_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	daomocks "github.com/in-rich/uservice-reminders/pkg/dao/mocks"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUpsertReminder(t *testing.T) {
	testData := []struct {
		name string

		reminder *models.UpsertReminder

		shouldCallCountUpdates bool
		countUpdatesResponse   int
		countUpdatesError      error

		shouldCallDeleteReminder bool
		deleteReminderResponse   *entities.Reminder
		deleteReminderError      error

		shouldCallCreateReminder bool
		createReminderResponse   *entities.Reminder
		createReminderError      error

		shouldCallUpdateReminder bool
		updateReminderResponse   *entities.Reminder
		updateReminderError      error

		expect    *models.Reminder
		expectID  string
		expectErr error
	}{
		{
			name: "UpdateReminder",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			shouldCallCreateReminder: true,
			createReminderError:      dao.ErrReminderAlreadyExists,
			shouldCallUpdateReminder: true,
			updateReminderResponse: &entities.Reminder{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id",
				Target:           entities.TargetCompany,
				PublicIdentifier: "public-identifier",
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Reminder{
				ID:               "00000000-0000-0000-0000-000000000001",
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectID: "00000000-0000-0000-0000-000000000001",
		},
		{
			name: "CreateReminder",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			shouldCallCreateReminder: true,
			createReminderResponse: &entities.Reminder{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id",
				Target:           entities.TargetCompany,
				PublicIdentifier: "public-identifier",
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Reminder{
				ID:               "00000000-0000-0000-0000-000000000001",
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectID: "00000000-0000-0000-0000-000000000001",
		},
		{
			name: "DeleteReminder",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			shouldCallDeleteReminder: true,
			deleteReminderResponse: &entities.Reminder{
				ID: lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
			},
			expectID: "00000000-0000-0000-0000-000000000001",
		},
		{
			name: "UpdateReminderError",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			shouldCallCreateReminder: true,
			createReminderError:      dao.ErrReminderAlreadyExists,
			shouldCallUpdateReminder: true,
			updateReminderError:      FooErr,
			expectErr:                FooErr,
		},
		{
			name: "CreateReminderError",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			shouldCallCreateReminder: true,
			createReminderError:      FooErr,
			expectErr:                FooErr,
		},
		{
			name: "DeleteReminderError",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "company",
				PublicIdentifier: "public-identifier",
				Content:          "",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			shouldCallDeleteReminder: true,
			deleteReminderError:      FooErr,
			expectErr:                FooErr,
		},
		{
			name: "InvalidTarget",
			reminder: &models.UpsertReminder{
				AuthorID:         "author-id",
				Target:           "invalid",
				PublicIdentifier: "public-identifier",
				Content:          "content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectErr: services.ErrInvalidReminderUpdate,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			deleteReminder := daomocks.NewMockDeleteReminderRepository(t)
			createReminder := daomocks.NewMockCreateReminderRepository(t)
			updateReminder := daomocks.NewMockUpdateReminderRepository(t)

			if tt.shouldCallDeleteReminder {
				deleteReminder.
					On("DeleteReminder", context.TODO(), tt.reminder.AuthorID, entities.Target(tt.reminder.Target), tt.reminder.PublicIdentifier).
					Return(tt.deleteReminderResponse, tt.deleteReminderError)
			}

			if tt.shouldCallCreateReminder {
				createReminder.
					On(
						"CreateReminder",
						context.TODO(),
						tt.reminder.AuthorID,
						entities.Target(tt.reminder.Target),
						tt.reminder.PublicIdentifier,
						&dao.CreateReminderData{Content: tt.reminder.Content, ExpiredAt: tt.reminder.ExpiredAt},
					).
					Return(tt.createReminderResponse, tt.createReminderError)
			}

			if tt.shouldCallUpdateReminder {
				updateReminder.
					On(
						"UpdateReminder",
						context.TODO(),
						tt.reminder.AuthorID,
						entities.Target(tt.reminder.Target),
						tt.reminder.PublicIdentifier,
						&dao.UpdateReminderData{Content: tt.reminder.Content, ExpiredAt: tt.reminder.ExpiredAt},
					).
					Return(tt.updateReminderResponse, tt.updateReminderError)
			}

			service := services.NewUpsertReminderService(updateReminder, createReminder, deleteReminder)

			reminder, reminderID, err := service.Exec(context.TODO(), tt.reminder)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, reminder)
			require.Equal(t, tt.expectID, reminderID)

			deleteReminder.AssertExpectations(t)
			createReminder.AssertExpectations(t)
			updateReminder.AssertExpectations(t)
		})
	}
}
