package main

import (
	"testing"

	"github.com/haukurk/latency-microservice-go/client"
)

func TestPingEndpointSuccess(t *testing.T) {

	// given
	client := client.LatencyClient{Host: "http://localhost:7801", RemoteHost: "8.8.8.8"}

	// when
	latenctyObj, err := client.PingHost()

	//then
	if err != nil {
		t.Error(err)
	}

	if latenctyObj.STATUS != "ok" {
		t.Error("Status not ok.")
	}

}
