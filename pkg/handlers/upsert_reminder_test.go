package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
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

func TestUpsertReminder(t *testing.T) {
	testData := []struct {
		name string

		in *reminders_pb.UpsertReminderRequest

		upsertResponse   *models.Reminder
		upsertIDResponse string
		upsertErr        error

		expect     *reminders_pb.UpsertReminderResponse
		expectCode codes.Code
	}{
		{
			name: "UpsertReminder",
			in: &reminders_pb.UpsertReminderRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				Content:          "content-1",
				ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			upsertResponse: &models.Reminder{
				ID:               "reminder-id",
				PublicIdentifier: "public-identifier-1",
				AuthorID:         "author-id-1",
				Target:           "company",
				Content:          "content-1",
				UpdatedAt:        lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				ExpiredAt:        lo.ToPtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			upsertIDResponse: "reminder-id",
			expect: &reminders_pb.UpsertReminderResponse{
				Id: "reminder-id",
				Reminder: &reminders_pb.Reminder{
					ReminderId:       "reminder-id",
					PublicIdentifier: "public-identifier-1",
					AuthorId:         "author-id-1",
					Target:           "company",
					Content:          "content-1",
					UpdatedAt:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
					ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "DeleteReminder",
			in: &reminders_pb.UpsertReminderRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				Content:          "content-1",
				ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			upsertIDResponse: "reminder-id",
			expect: &reminders_pb.UpsertReminderResponse{
				Id: "reminder-id",
			},
		},
		{
			name: "InvalidArgument",
			in: &reminders_pb.UpsertReminderRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				Content:          "content-1",
				ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			upsertErr:  services.ErrInvalidReminderUpdate,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "Internal",
			in: &reminders_pb.UpsertReminderRequest{
				Target:           "company",
				PublicIdentifier: "public-identifier-1",
				AuthorId:         "author-id-1",
				Content:          "content-1",
				ExpiredAt:        timestamppb.New(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			upsertErr:  errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockUpsertReminderService(t)
			service.On("Exec", context.TODO(), &models.UpsertReminder{
				Target:           tt.in.GetTarget(),
				PublicIdentifier: tt.in.GetPublicIdentifier(),
				AuthorID:         tt.in.GetAuthorId(),
				Content:          tt.in.GetContent(),
				ExpiredAt:        lo.ToPtr(tt.in.GetExpiredAt().AsTime()),
			}).Return(tt.upsertResponse, tt.upsertIDResponse, tt.upsertErr)

			handler := handlers.NewUpsertReminderHandler(service, monitor.NewDummyGRPCLogger())

			resp, err := handler.UpsertReminder(context.TODO(), tt.in)

			require.Equal(t, tt.expect, resp)
			RequireGRPCCodesEqual(t, err, tt.expectCode)

			service.AssertExpectations(t)
		})
	}
}
