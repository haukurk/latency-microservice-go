package main

import (
	"fmt"
	"testing"

	"github.com/haukurk/latency-microservice-go/client"
)

func TestPingIPEndpointSuccess(t *testing.T) {

	// given
	client := client.LatencyClient{Host: "http://localhost:7801", RemoteHost: "8.8.8.8"}

	// when
	latencyResponse, err := client.PingHost()

	//then
	if err != nil {
		t.Error(err)
	} else {
		if latencyResponse.STATUS != "ok" {
			fmt.Printf("%+v\n", latencyResponse)
			t.Error("Status not ok.")
		}
	}
}

func TestPingHostnameEndpointSuccess(t *testing.T) {

	// given
	client := client.LatencyClient{Host: "http://localhost:7801", RemoteHost: "google.com"}

	// when
	latencyResponse, err := client.PingHost()

	//then
	if err != nil {
		t.Error(err)
	} else {
		if latencyResponse.STATUS != "ok" {
			fmt.Printf("%+v\n", latencyResponse)
			t.Error("Status not ok.")
		}
	}
}

func TestPingBogusHostnameEndpointFailure(t *testing.T) {

	// given
	client := client.LatencyClient{Host: "http://localhost:7801", RemoteHost: "somehostname.that.does.not.exist.com"}

	// when
	_, err := client.PingHost()

	//then
	if err == nil {
		t.Error("No error encountered")
	}
}

func TestPingBogusIPEndpointFailure(t *testing.T) {

	// given
	client := client.LatencyClient{Host: "http://localhost:7801", RemoteHost: "8.8.-1.-1"}

	// when
	_, err := client.PingHost()

	//then
	if err == nil {
		t.Error("No error encountered")
	}

}
