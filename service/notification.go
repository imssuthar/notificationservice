package service

import (
	"context"
	"imssuthar/notificationsystem/dao"
	"imssuthar/notificationsystem/dto"
)

func GetNotification(ctx context.Context, notId string) (dto.Notification_status_response, error) {
	not, err := dao.GetNotificationById(ctx, notId)
	if err != nil {
		return dto.Notification_status_response{}, err
	}
	return dto.Notification_status_response{Id: not.Id, Status: not.Status, NotType: not.NotType}, nil
}

func GetNotifications(ctx context.Context) ([]dto.NotificationEntity, error) {
	return dao.GetNotifications(ctx)
}

func CreateNotification(ctx context.Context, not dto.NotificationRequest) (dto.NotificationEntity, error) {
	return dao.CreateNotification(ctx, not)
}
