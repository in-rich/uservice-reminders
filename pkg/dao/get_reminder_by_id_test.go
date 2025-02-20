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

var getReminderByIDFixtures = []*entities.Reminder{
	// User 1
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		AuthorID:         "author-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		UpdatedAt:        lo.ToPtr(time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		AuthorID:         "author-id-1",
		PublicIdentifier: "public-identifier-2",
		Target:           entities.TargetUser,
		Content:          "content-2",
		UpdatedAt:        lo.ToPtr(time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		AuthorID:         "author-id-1",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetCompany,
		Content:          "content-1",
		UpdatedAt:        lo.ToPtr(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000004")),
		AuthorID:         "author-id-1",
		PublicIdentifier: "public-identifier-2",
		Target:           entities.TargetCompany,
		Content:          "content-2",
		UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},

	// User 2
	{
		ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000005")),
		AuthorID:         "author-id-2",
		PublicIdentifier: "public-identifier-1",
		Target:           entities.TargetUser,
		Content:          "content-1",
		UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestGetReminderByID(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		authorID   string
		reminderID uuid.UUID

		expect    *entities.Reminder
		expectErr error
	}{
		{
			name:       "GetReminderByID",
			authorID:   "author-id-1",
			reminderID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			expect: &entities.Reminder{
				ID:               lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				AuthorID:         "author-id-1",
				PublicIdentifier: "public-identifier-1",
				Target:           entities.TargetUser,
				Content:          "content-1",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:       "ListRemindersByAuthor/NoResult",
			authorID:   "author-id-1",
			reminderID: uuid.MustParse("00000000-0000-0000-0000-000000000006"),
			expectErr:  dao.ErrReminderNotFound,
		},
	}

	stx := BeginTX(db, getReminderByIDFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			repo := dao.NewGetReminderByIDRepository(stx)
			reminder, err := repo.GetReminderByID(context.TODO(), tt.authorID, tt.reminderID)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, reminder)
		})
	}
}
