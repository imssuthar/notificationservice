package dao

import (
	"context"
	"errors"
	"imssuthar/notificationsystem/dto"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrNotFound = errors.New("notification not found")

func GetNotificationById(ctx context.Context, id string) (dto.NotificationEntity, error) {
	var not dto.NotificationEntity
	err := notificationCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&not)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return dto.NotificationEntity{}, ErrNotFound
	}
	if err != nil {
		return dto.NotificationEntity{}, err
	}
	return not, nil
}

func GetNotifications(ctx context.Context) ([]dto.NotificationEntity, error) {
	cur, err := notificationCollection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	nots := []dto.NotificationEntity{}
	if err := cur.All(ctx, &nots); err != nil {
		return nil, err
	}
	return nots, nil
}

func CreateNotification(ctx context.Context, notification dto.NotificationRequest) (dto.NotificationEntity, error) {
	not := dto.NotificationEntity{
		Id:        bson.NewObjectID().Hex(),
		NotType:   notification.NotType,
		ToId:      notification.ToId,
		Content:   notification.Content,
		Priority:  notification.Priority,
		Status:    dto.QUEUED,
		CreatedAt: time.Now().UnixMilli(),
	}
	if _, err := notificationCollection().InsertOne(ctx, not); err != nil {
		return dto.NotificationEntity{}, err
	}
	return not, nil
}
