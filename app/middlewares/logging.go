package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
)


func LogApiRequests() gin.HandlerFunc {
	return func (context *gin.Context) {
		t := time.Now()
		log.Println("Request received at timestamp", t)
		context.Next()
		latency := time.Since(t)
		log.Println("Time To execure", latency)
	}
}