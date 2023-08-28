package logger

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func GinCustomLogger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["level"] = "info"
			log["time"] = params.TimeStamp.UTC().Format(time.RFC3339)
			log["caller"] = "gin"
			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
		Output:    gin.DefaultWriter,
		SkipPaths: []string{"/health"},
	})
}
