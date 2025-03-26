package notify

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyHandler interface {
	SendNotify(c *gin.Context)
}

type notifyHandler struct {
	NotifyService NotifyService
}

func NewNotifyHandler(notifyService NotifyService) NotifyHandler {
	return &notifyHandler{NotifyService: notifyService}
}

func (h *notifyHandler) SendNotify(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, &SendNotifyRes{
			Status:  401,
			Message: "Missing authorization header",
		})
		return
	}

	notifyRes, err := h.NotifyService.SendNotify(authHeader)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &SendNotifyRes{
			Status:  500,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, notifyRes)
}
