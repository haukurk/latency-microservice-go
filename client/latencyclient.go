package client

import (
	"github.com/haukurk/latency-microservice-go/api"
)

type LatencyClient struct {
	Host       string
	RemoteHost string
}

func (tc *LatencyClient) PingHost() (api.Latency, error) {
	var respLatency api.Latency

	url := tc.Host + "/latency/" + tc.RemoteHost
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respLatency, err
	}
	err = processResponseEntity(r, &respLatency, 201)
	return respLatency, err
}
