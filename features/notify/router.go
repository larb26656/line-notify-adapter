package notify

import "github.com/gin-gonic/gin"

type NotifyRouter interface {
	Setup(r *gin.Engine)
}

type notifyRouter struct {
	NotifyHandler NotifyHandler
}

func NewNotifyRouter(notifyHandler NotifyHandler) NotifyRouter {
	return &notifyRouter{NotifyHandler: notifyHandler}
}

func (ur *notifyRouter) Setup(r *gin.Engine) {
	r.POST("/api/v1/notify", ur.NotifyHandler.SendNotify)
}
