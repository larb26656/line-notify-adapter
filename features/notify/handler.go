package notify

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/larb26656/line-notify-adapter/utils"
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
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		c.JSON(http.StatusUnauthorized, &SendNotifyRes{
			Status:  401,
			Message: "Missing authorization header",
		})
		return
	}

	token, err := utils.ExtractBearerToken(authorization)

	if err != nil {
		c.JSON(http.StatusBadRequest, &SendNotifyRes{
			Status:  401,
			Message: "Missing Bearer",
		})
		return
	}

	message := c.PostForm("message")

	if message == "" {
		c.JSON(http.StatusBadRequest, &SendNotifyRes{
			Status:  400,
			Message: "message: must not be empty",
		})
		return
	}

	notifyRes, err := h.NotifyService.SendNotify(token, message)

	if err != nil {
		log.Printf("Error sending notification: %v", err)
		c.JSON(http.StatusInternalServerError, &SendNotifyRes{
			Status:  500,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, notifyRes)
}
