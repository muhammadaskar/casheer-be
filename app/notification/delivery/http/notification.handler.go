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

func (h *NotificationHandler) Update(c *gin.Context) {
	var inputID notification.GetInputID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update notification", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData notification.InputUpdateNotification

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update notification", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateNotification, err := h.notificationUseCase.Update(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to update notification", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to update notification", http.StatusOK, "success", updateNotification)
	c.JSON(http.StatusOK, response)
}
