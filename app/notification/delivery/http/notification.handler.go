package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/notification"
	"github.com/muhammadaskar/casheer-be/app/notification/usecase"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type NotificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
}

func NewNotificationHandler(notificationUseCase usecase.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{notificationUseCase}
}

func (h *NotificationHandler) FindAll(c *gin.Context) {
	notifications, err := h.notificationUseCase.FindAll()
	if err != nil {
		response := customresponse.APIResponse("Error to get notifications", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := customresponse.APIResponse("List of notifications", http.StatusOK, "success", notification.NotificationsFormatter(notifications))
	c.JSON(http.StatusOK, response)
}
