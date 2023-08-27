package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/weflux/fluxworks/logging"
	"github.com/weflux/fluxworks/types"
	"time"
)

type Config struct {
	TimeFormat   string
	SkipPaths    []string
	DefaultLevel string
	UTC          bool
}

func SetLogger(logger *logging.Logger, conf *Config) gin.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.Query()
		c.Next()

		if _, exists := skipPaths[path]; !exists {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			fields := types.M{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       path,
				"query":      query,
				"ip":         c.ClientIP(),
				"user-agent": c.Request.UserAgent(),
				"latency":    latency,
			}

			if conf.TimeFormat != "" {
				fields["time"] = end.Format(conf.TimeFormat)
			}

			if len(c.Errors) > 0 {
				for _, e := range c.Errors {
					logger.Error("", e, fields)
				}
			} else {
				logger.Info("", fields)
			}
		}

	}

}
