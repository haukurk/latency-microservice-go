package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int
}

type LatencyService struct{}

func (s *LatencyService) Run(cfg Config) error {
	latencyResource := &LatencyResource{}
	router := gin.Default()
	router.GET("/latency/:host", latencyResource.LatencyHost)
	router.Run(fmt.Sprint(":", cfg.Port))

	return nil
}
