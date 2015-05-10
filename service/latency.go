package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thoas/stats"
)

var Stats = stats.New()

type Config struct {
	Port int
}

type LatencyService struct{}

func (s *LatencyService) Run(cfg Config) error {
	latencyResource := &LatencyResource{}

	router := gin.Default()

	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			beginning := time.Now()
			c.Next()
			Stats.End(beginning, c.Writer)
		}
	}())

	router.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, Stats.Data())
	})

	router.GET("/latency/:host", latencyResource.LatencyHost)

	router.Run(fmt.Sprint(":", cfg.Port))

	return nil
}
