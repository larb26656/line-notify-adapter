package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/larb26656/line-notify-adapter/features/notify"
)

func setupRoute(r *gin.Engine) {
	// Service
	notifyService := notify.NewNotifyService()

	// Handler
	notifyHandler := notify.NewNotifyHandler(notifyService)

	// Router
	notifyRouter := notify.NewNotifyRouter(notifyHandler)

	// Setup routes
	notifyRouter.Setup(r)
}

func main() {
	r := gin.Default()

	setupRoute(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
