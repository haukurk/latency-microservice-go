package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haukurk/latency-microservice-go/api"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type LatencyResource struct{}

func (lr *LatencyResource) LatencyHost(context *gin.Context) {

	var latencyResp api.Latency

	ip := context.Params.ByName("host")
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.AddIPAddr(ra)

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		latencyResp.IP = addr.String()
		latencyResp.RTT = rtt
		latencyResp.UNIT = "ms"
		latencyResp.STATUS = "ok"
	}

	p.OnIdle = func() {
		log.Printf("log idle")
	}

	err = p.Run()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err})
	}

	context.JSON(http.StatusOK, latencyResp)
}
