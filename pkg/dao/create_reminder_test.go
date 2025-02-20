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

var createReminderFixtures = []*entities.Reminder{
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		AuthorID:         "author-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		UpdatedAt:        lo.ToPtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestCreateReminder(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name             string
		authorID         string
		publicIdentifier string
		target           entities.Target
		data             *dao.CreateReminderData
		expect           *entities.Reminder
		expectErr        error
	}{
		{
			name:             "CreateReminder/SameTarget/DifferentIdentifier/SameAuthor",
			authorID:         "author-id-1",
			publicIdentifier: "public-identifier-2",
			target:           entities.TargetUser,
			data: &dao.CreateReminderData{
				Content:   "new-content",
				ExpiredAt: lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &entities.Reminder{
				AuthorID:         "author-id-1",
				PublicIdentifier: "public-identifier-2",
				Target:           entities.TargetUser,
				Content:          "new-content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:             "CreateReminder/DifferentTarget/SameIdentifier/SameAuthor",
			authorID:         "author-id-1",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetCompany,
			data: &dao.CreateReminderData{
				Content:   "new-content",
				ExpiredAt: lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &entities.Reminder{
				AuthorID:         "author-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetCompany,
				Content:          "new-content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:             "CreateReminder/SameTarget/SameIdentifier/DifferentAuthor",
			authorID:         "author-id-2",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetUser,
			data: &dao.CreateReminderData{
				Content:   "new-content",
				ExpiredAt: lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &entities.Reminder{
				AuthorID:         "author-id-2",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
				Content:          "new-content",
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:             "CreateReminder/SameTarget/SameIdentifier/SameAuthor",
			authorID:         "author-id-1",
			publicIdentifier: "public-identifier-1",
			target:           entities.TargetUser,
			data: &dao.CreateReminderData{
				Content:   "new-content",
				ExpiredAt: lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectErr: dao.ErrReminderAlreadyExists,
		},
	}

	stx := BeginTX(db, createReminderFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewCreateReminderRepository(tx)
			reminder, err := repo.CreateReminder(context.TODO(), tt.authorID, tt.target, tt.publicIdentifier, tt.data)

			if reminder != nil {
				// Since ID and UpdatedAt are random, nullify them for comparison.
				reminder.ID = nil
				reminder.UpdatedAt = nil
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, reminder)
		})
	}
}
