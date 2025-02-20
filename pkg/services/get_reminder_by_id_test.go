package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-reminders/pkg/dao/mocks"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetReminderByID(t *testing.T) {
	testData := []struct {
		name string

		selector *models.GetReminderByID

		shouldCallGetReminderByID bool
		getReminderByIDResponse   *entities.Reminder
		getReminderByIDError      error

		expect    *models.Reminder
		expectErr error
	}{
		{
			name: "ListReminders",
			selector: &models.GetReminderByID{
				AuthorID:   "author-id",
				ReminderID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetReminderByID: true,
			getReminderByIDResponse: &entities.Reminder{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				PublicIdentifier: "public-identifier",
				AuthorID:         "author-id",
				Target:           entities.Target("target"),
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Reminder{
				ID:               "00000000-0000-0000-0000-000000000001",
				PublicIdentifier: "public-identifier",
				AuthorID:         "author-id",
				Target:           "target",
				Content:          "content",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "ListRemindersError",
			selector: &models.GetReminderByID{
				AuthorID:   "author-id",
				ReminderID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetReminderByID: true,
			getReminderByIDError:      FooErr,
			expectErr:                 FooErr,
		},
		{
			name: "InvalidSelector",
			selector: &models.GetReminderByID{
				AuthorID:   "author-id",
				ReminderID: "",
			},
			expectErr: services.ErrInvalidReminderSelector,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getReminder := daomocks.NewMockGetReminderByIDRepository(t)

			if tt.shouldCallGetReminderByID {
				getReminder.
					On("GetReminderByID", context.TODO(), tt.selector.AuthorID, uuid.MustParse(tt.selector.ReminderID)).
					Return(tt.getReminderByIDResponse, tt.getReminderByIDError)
			}

			service := services.NewGetReminderByIDService(getReminder)

			result, err := service.Exec(context.TODO(), tt.selector)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, result)

			getReminder.AssertExpectations(t)
		})
	}
}
