package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetReminderHandler struct {
	reminders_pb.GetReminderServer
	service services.GetReminderService
	logger  monitor.GRPCLogger
}

func (h *GetReminderHandler) getReminder(ctx context.Context, in *reminders_pb.GetReminderRequest) (*reminders_pb.Reminder, error) {
	reminder, err := h.service.Exec(ctx, &models.GetReminder{
		Target:           in.GetTarget(),
		PublicIdentifier: in.GetPublicIdentifier(),
		AuthorID:         in.GetAuthorId(),
	})
	if err != nil {
		if errors.Is(err, dao.ErrReminderNotFound) {
			return nil, status.Error(codes.NotFound, "reminder not found")
		}
		if errors.Is(err, services.ErrInvalidReminderSelector) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid reminder selector: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to get reminder: %v", err)
	}

	return &reminders_pb.Reminder{
		PublicIdentifier: reminder.PublicIdentifier,
		AuthorId:         reminder.AuthorID,
		Target:           reminder.Target,
		Content:          reminder.Content,
		UpdatedAt:        TimeToTimestampProto(reminder.UpdatedAt),
		ExpiredAt:        TimeToTimestampProto(reminder.ExpiredAt),
	}, nil
}

func (h *GetReminderHandler) GetReminder(ctx context.Context, in *reminders_pb.GetReminderRequest) (*reminders_pb.Reminder, error) {
	res, err := h.getReminder(ctx, in)
	h.logger.Report(ctx, "GetReminder", err)
	return res, err
}

func NewGetReminderHandler(service services.GetReminderService, logger monitor.GRPCLogger) *GetReminderHandler {
	return &GetReminderHandler{
		service: service,
		logger:  logger,
	}
}
