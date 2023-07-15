package gin

import (
	"github.com/fluxstack/fluxworks/log"
	"github.com/gin-gonic/gin"
	"time"
)

type Config struct {
	TimeFormat   string
	SkipPaths    []string
	DefaultLevel string
	UTC          bool
}

func SetLogger(logger *log.Logger, conf *Config) gin.HandlerFunc {
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

			fields := log.Field("status", c.Writer.Status()).
				Field("method", c.Request.Method).
				Field("path", path).
				Field("query", query).
				Field("ip", c.ClientIP()).
				Field("user-agent", c.Request.UserAgent()).
				Field("latency", latency)

			if conf.TimeFormat != "" {
				fields = fields.Field("time", end.Format(conf.TimeFormat))
			}

			if len(c.Errors) > 0 {
				for _, e := range c.Errors.Errors() {
					f := fields.Field("error", e)
					logger.Errorw(f)
				}
			} else {
				logger.Infow(fields)
			}
		}

	}

}
