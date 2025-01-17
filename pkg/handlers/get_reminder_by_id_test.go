package handlers_test

import (
	"context"
	"errors"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
	"github.com/in-rich/uservice-reminders/pkg/handlers"
	"github.com/in-rich/uservice-reminders/pkg/models"
	servicesmocks "github.com/in-rich/uservice-reminders/pkg/services/mocks"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGetReminderByID(t *testing.T) {
	testData := []struct {
		name string

		in *reminders_pb.GetReminderByIDRequest

		getResponse *models.Reminder
		getErr      error

		expect     *reminders_pb.Reminder
		expectCode codes.Code
	}{
		{
			name: "GetReminderByID",
			in: &reminders_pb.GetReminderByIDRequest{
				AuthorId:   "author-id-1",
				ReminderId: "reminder_id-1",
			},
			getResponse: &models.Reminder{
				ID:               "reminder_id-1",
				PublicIdentifier: "public-identifier-1",
				AuthorID:         "author-id-1",
				Target:           "company",
				Content:          "content-1",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &reminders_pb.Reminder{
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				Target:           "company",
				Content:          "content-1",
				UpdatedAt:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
				ReminderId:       "reminder_id-1",
			},
		},
		{
			name: "ListError",
			in: &reminders_pb.GetReminderByIDRequest{
				AuthorId:   "author-id-1",
				ReminderId: "reminder_id-1",
			},
			getErr:     errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetReminderByIDService(t)

			service.
				On("Exec", context.TODO(), &models.GetReminderByID{
					AuthorID:   tt.in.AuthorId,
					ReminderID: tt.in.ReminderId,
				}).
				Return(tt.getResponse, tt.getErr)

			handler := handlers.NewGetReminderByIDHandler(service)
			resp, err := handler.GetReminderByID(context.TODO(), tt.in)

			require.Equal(t, tt.expect, resp)
			RequireGRPCCodesEqual(t, err, tt.expectCode)

			service.AssertExpectations(t)
		})
	}
}
