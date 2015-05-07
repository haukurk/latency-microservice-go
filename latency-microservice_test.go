package main

import (
	"fmt"
	"testing"

	"github.com/haukurk/latency-microservice-go/client"
)

func TestPingEndpointSuccess(t *testing.T) {

	// given
	client := client.LatencyClient{Host: "http://localhost:7801", RemoteHost: "8.8.8.8"}

	// when
	latencyResponse, err := client.PingHost()

	//then
	if err != nil {
		t.Error(err)
	}

	if latencyResponse.STATUS != "ok" {
		fmt.Printf("%+v\n", latencyResponse)
		t.Error("Status not ok.")
	}

}
