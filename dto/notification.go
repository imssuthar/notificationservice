package dto

type Notification_status_response struct {
	Id      string
	Status  NotificationStatus
	NotType Notification_type
}

type Notification_type string

const (
	SMS   Notification_type = "SMS"
	EMAIL Notification_type = "EMAIL"
)

type NotificationStatus string

const (
	QUEUED NotificationStatus = "QUEUED"
	SENT   NotificationStatus = "SENT"
)

type NotificationPriority int

const (
	LOW NotificationPriority = iota
	MED
	HIGH
	CRITICAL
)

type NotificationRequest struct {
	NotType  Notification_type
	ToId     string
	Content  string
	Priority NotificationPriority
}

type NotificationEntity struct {
	Id        string               `bson:"_id"`
	NotType   Notification_type    `bson:"notType"`
	ToId      string               `bson:"toId"`
	Content   string               `bson:"content"`
	Priority  NotificationPriority `bson:"priority"`
	Status    NotificationStatus   `bson:"status"`
	CreatedAt int64                `bson:"createdAt"`
}
