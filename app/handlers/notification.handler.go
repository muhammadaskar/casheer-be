package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/helper"
	"github.com/muhammadaskar/casheer-be/app/notification"
)

type NotificationHandler struct {
	notificationService notification.Service
}

func NewNotificationHandler(notificationService notification.Service) *NotificationHandler {
	return &NotificationHandler{notificationService}
}

func (h *NotificationHandler) FindAll(c *gin.Context) {
	notifications, err := h.notificationService.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get notifications", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of notifications", http.StatusOK, "success", notification.NotificationsFormatter(notifications))
	c.JSON(http.StatusOK, response)
}
