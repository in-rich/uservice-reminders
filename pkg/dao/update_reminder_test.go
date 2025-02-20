package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/entities"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var updateReminderFixtures = []*entities.Reminder{
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		AuthorID:         "author-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestUpdateReminder(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name             string
		authorID         string
		publicIdentifier string
		target           entities.Target
		data             *dao.UpdateReminderData
		expect           *entities.Reminder
		expectErr        error
	}{
		{
			name:             "UpdateReminder",
			authorID:         "author-id-1",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetUser,
			data: &dao.UpdateReminderData{
				Content:   "new-content",
				ExpiredAt: lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &entities.Reminder{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
				Content:          "new-content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:             "ReminderNotFound",
			authorID:         "author-id-1",
			publicIdentifier: "public-identifier-2",
			target:           entities.TargetUser,
			data: &dao.UpdateReminderData{
				Content:   "new-content",
				ExpiredAt: lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectErr: dao.ErrReminderNotFound,
		},
	}

	stx := BeginTX(db, updateReminderFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewUpdateReminderRepository(tx)
			reminder, err := repo.UpdateReminder(context.TODO(), tt.authorID, tt.target, tt.publicIdentifier, tt.data)

			if reminder != nil {
				// Since UpdatedAt is random, nullify it for comparison.
				reminder.UpdatedAt = nil
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, reminder)
		})
	}
}
