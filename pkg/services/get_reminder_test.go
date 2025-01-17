package services_test

import (
	"context"
	daomocks "github.com/in-rich/uservice-reminders/pkg/dao/mocks"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetReminderService(t *testing.T) {
	testData := []struct {
		name string

		selector *models.GetReminder

		shouldCallGetReminder bool
		getReminderResponse   *entities.Reminder
		getReminderError      error

		expect    *models.Reminder
		expectErr error
	}{
		{
			name: "GetReminder",
			selector: &models.GetReminder{
				AuthorID:         "author-id",
				Target:           "user",
				PublicIdentifier: "public-identifier",
			},
			shouldCallGetReminder: true,
			getReminderResponse: &entities.Reminder{
				PublicIdentifier: "public-identifier",
				AuthorID:         "author-id",
				Target:           entities.Target("target"),
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Reminder{
				PublicIdentifier: "public-identifier",
				AuthorID:         "author-id",
				Target:           "target",
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "GetReminderError",
			selector: &models.GetReminder{
				AuthorID:         "author-id",
				Target:           "user",
				PublicIdentifier: "public-identifier",
			},
			shouldCallGetReminder: true,
			getReminderError:      FooErr,
			expectErr:             FooErr,
		},
		{
			name: "InvalidSelector",
			selector: &models.GetReminder{
				AuthorID:         "author-id",
				Target:           "invalid",
				PublicIdentifier: "public-identifier",
			},
			expectErr: services.ErrInvalidReminderSelector,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getReminder := daomocks.NewMockGetReminderRepository(t)

			if tt.shouldCallGetReminder {
				getReminder.
					On("GetReminder", context.TODO(), tt.selector.AuthorID, entities.Target(tt.selector.Target), tt.selector.PublicIdentifier).
					Return(tt.getReminderResponse, tt.getReminderError)
			}

			service := services.NewGetReminderService(getReminder)

			reminder, err := service.Exec(context.TODO(), tt.selector)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, reminder)

			getReminder.AssertExpectations(t)
		})
	}
}
