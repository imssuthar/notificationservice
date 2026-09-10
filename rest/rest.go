package rest

import (
	"encoding/json"
	"errors"
	"imssuthar/notificationsystem/dao"
	"imssuthar/notificationsystem/dto"
	"imssuthar/notificationsystem/service"
	"log"
	"net/http"
	"strings"
)

type Status struct {
	Name string
}

func writeJSON(res http.ResponseWriter, status int, body any) {
	bytes, err := json.Marshal(body)
	if err != nil {
		http.Error(res, "failed to encode response", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(bytes)
}

func getNotification(res http.ResponseWriter, req *http.Request) {
	notId := req.PathValue("notId")
	if (strings.TrimSpace(notId)) == "" {
		http.Error(res, "notId is required", http.StatusBadRequest)
		return
	}
	notification, err := service.GetNotification(req.Context(), notId)
	if errors.Is(err, dao.ErrNotFound) {
		http.Error(res, "notification not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("get notification %s: %v", notId, err)
		http.Error(res, "failed to fetch notification", http.StatusInternalServerError)
		return
	}
	writeJSON(res, http.StatusOK, notification)
}

func getNotifications(res http.ResponseWriter, req *http.Request) {
	notifications, err := service.GetNotifications(req.Context())
	if err != nil {
		log.Printf("list notifications: %v", err)
		http.Error(res, "failed to fetch notifications", http.StatusInternalServerError)
		return
	}
	writeJSON(res, http.StatusOK, notifications)
}

func createNotification(res http.ResponseWriter, req *http.Request) {
	var notification dto.NotificationRequest
	if err := json.NewDecoder(req.Body).Decode(&notification); err != nil {
		http.Error(res, "invalid request body", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(notification.ToId) == "" || strings.TrimSpace(notification.Content) == "" {
		http.Error(res, "toId and content are required", http.StatusBadRequest)
		return
	}
	created, err := service.CreateNotification(req.Context(), notification)
	if err != nil {
		log.Printf("create notification: %v", err)
		http.Error(res, "failed to create notification", http.StatusInternalServerError)
		return
	}
	writeJSON(res, http.StatusCreated, created)
}

func myhandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/notification/{notId}", getNotification)
	mux.HandleFunc("POST /api/v1/notification", createNotification)
	mux.HandleFunc("GET /api/v1/notification", getNotifications)

	return mux
}

func Attach() error {
	s := &http.Server{
		Addr:    ":8080",
		Handler: myhandler(),
	}
	return s.ListenAndServe()
}
