package handlers

import (
	"context"
	"errors"
	reminders_pb "github.com/in-rich/proto/proto-go/reminders"
	"github.com/in-rich/uservice-reminders/pkg/dao"
	"github.com/in-rich/uservice-reminders/pkg/models"
	"github.com/in-rich/uservice-reminders/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetReminderByIDHandler struct {
	reminders_pb.GetReminderByIDServer
	service services.GetReminderByIDService
}

func (h *GetReminderByIDHandler) GetReminderByID(ctx context.Context, in *reminders_pb.GetReminderByIDRequest) (*reminders_pb.Reminder, error) {
	reminder, err := h.service.Exec(ctx, &models.GetReminderByID{
		AuthorID:   in.GetAuthorId(),
		ReminderID: in.GetReminderId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidReminderSelector) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid reminder selector: %v", err)
		}
		if errors.Is(err, dao.ErrReminderNotFound) {
			return nil, status.Errorf(codes.NotFound, "reminder not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get reminder: %v", err)
	}

	return &reminders_pb.Reminder{
		Target:           reminder.Target,
		PublicIdentifier: reminder.PublicIdentifier,
		Content:          reminder.Content,
		AuthorId:         reminder.AuthorID,
		ReminderId:       reminder.ID,
		UpdatedAt:        TimeToTimestampProto(reminder.UpdatedAt),
		ExpiredAt:        TimeToTimestampProto(reminder.ExpiredAt),
	}, nil
}

func NewGetReminderByIDHandler(service services.GetReminderByIDService) *GetReminderByIDHandler {
	return &GetReminderByIDHandler{
		service: service,
	}
}
