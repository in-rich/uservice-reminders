package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/handlers"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	servicesmocks "github.com/in-rich/uservice-reminders/pkg/services/mocks"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGetUserData(t *testing.T) {
	testData := []struct {
		name string

		in *reminders_pb.GetReminderRequest

		getResponse *models.Reminder
		getErr      error

		expect     *reminders_pb.Reminder
		expectCode codes.Code
	}{
		{
			name: "GetReminder",
			in: &reminders_pb.GetReminderRequest{
				Target:           "user",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
			},
			getResponse: &models.Reminder{
				PublicIdentifier: "public-identifier-1",
				AuthorID:         "author-id-1",
				Target:           "user",
				Content:          "content-1",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &reminders_pb.Reminder{
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				Target:           "user",
				Content:          "content-1",
				UpdatedAt:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "ReminderNotFound",
			in: &reminders_pb.GetReminderRequest{
				Target:           "user",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
			},
			getErr:     dao.ErrReminderNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "InvalidArgument",
			in: &reminders_pb.GetReminderRequest{
				Target:           "user",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
			},
			getErr:     services.ErrInvalidReminderSelector,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &reminders_pb.GetReminderRequest{
				Target:           "user",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
			},
			getErr:     errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetReminderService(t)
			service.
				On("Exec", context.TODO(), &models.GetReminder{
					Target:           tt.in.Target,
					PublicIdentifier: tt.in.PublicIdentifier,
					AuthorID:         tt.in.AuthorId,
				}).
				Return(tt.getResponse, tt.getErr)

			handler := handlers.NewGetReminderHandler(service, monitor.NewDummyGRPCLogger())
			resp, err := handler.GetReminder(context.TODO(), tt.in)

			require.Equal(t, tt.expect, resp)
			RequireGRPCCodesEqual(t, err, tt.expectCode)

			service.AssertExpectations(t)
		})
	}
}
