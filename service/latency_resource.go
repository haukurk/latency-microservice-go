package service

import (
	"github.com/gin-gonic/gin"
	"github.com/haukurk/latency-microservice-go/api"
	"github.com/tatsushid/go-fastping"
	"net"
	"net/http"
	"time"
)

type LatencyResource struct{}

func (lr *LatencyResource) LatencyHost(context *gin.Context) {

	var latencyResp api.Latency

	ip := context.Params.ByName("host")
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)

	if err != nil {
		context.JSON(http.StatusNotFound, api.NewError("cannot resolve remote address"))
	} else {
		p := fastping.NewPinger()
		p.AddIPAddr(ra)

		// Only Single ICMP packet here.
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			latencyResp.IP = addr.String()
			latencyResp.RTT = rtt.Seconds() * float64(time.Second/time.Millisecond)
			latencyResp.UNIT = "ms"
			latencyResp.STATUS = "ok"
		}

		p.OnIdle = func() {
		}

		err = p.Run()

		if err != nil {
			context.JSON(http.StatusBadRequest, api.NewError(err.Error()))
		} else if latencyResp.STATUS != "ok" {
			context.JSON(http.StatusBadRequest, api.NewError("remote peer not answering or address bogus"))
		} else {
			context.JSON(http.StatusOK, latencyResp)
		}

	}

}
