package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpsertReminderHandler struct {
	reminders_pb.UpsertReminderServer
	service services.UpsertReminderService
	logger  monitor.GRPCLogger
}

func (h *UpsertReminderHandler) upsertReminder(ctx context.Context, in *reminders_pb.UpsertReminderRequest) (*reminders_pb.UpsertReminderResponse, error) {
	reminder, err := h.service.Exec(ctx, &models.UpsertReminder{
		Target:           in.GetTarget(),
		PublicIdentifier: in.GetPublicIdentifier(),
		AuthorID:         in.GetAuthorId(),
		Content:          in.GetContent(),
		ExpiredAt:        lo.ToPtr(in.GetExpiredAt().AsTime()),
	})
	if err != nil {
		if errors.Is(err, services.ErrRemindersUpdateLimitReached) {
			return nil, status.Error(codes.ResourceExhausted, "reminder update limit reached")
		}
		if errors.Is(err, services.ErrInvalidReminderUpdate) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid reminder update: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to upsert reminder: %v", err)
	}

	if reminder == nil {
		return &reminders_pb.UpsertReminderResponse{}, nil
	}

	return &reminders_pb.UpsertReminderResponse{
		Reminder: &reminders_pb.Reminder{
			ReminderId:       reminder.ID,
			Target:           reminder.Target,
			PublicIdentifier: reminder.PublicIdentifier,
			Content:          reminder.Content,
			UpdatedAt:        TimeToTimestampProto(reminder.UpdatedAt),
			ExpiredAt:        TimeToTimestampProto(reminder.ExpiredAt),
			AuthorId:         reminder.AuthorID,
		},
	}, nil
}

func (h *UpsertReminderHandler) UpsertReminder(ctx context.Context, in *reminders_pb.UpsertReminderRequest) (*reminders_pb.UpsertReminderResponse, error) {
	res, err := h.upsertReminder(ctx, in)
	h.logger.Report(ctx, "UpsertReminder", err)
	return res, err
}

func NewUpsertReminderHandler(service services.UpsertReminderService, logger monitor.GRPCLogger) *UpsertReminderHandler {
	return &UpsertReminderHandler{
		service: service,
		logger:  logger,
	}
}
