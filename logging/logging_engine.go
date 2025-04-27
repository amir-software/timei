package logging

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logging(context *gin.Context) {
	// a midlleware that log every request info and response time
	url := context.Request.URL.String()
	Start := time.Now()

	context.Next()
	duration := time.Since(Start)
	addLog(url, duration.Seconds())

}
