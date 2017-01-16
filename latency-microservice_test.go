package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/haukurk/latency-microservice-go/client"
	"github.com/tatsushid/go-fastping"
)

func TestPinger(t *testing.T) {

	p := fastping.NewPinger()
	p.Network("ip")

	if err := p.AddIP("127.0.0.1"); err != nil {
		t.Fatalf("AddIP failed: %v", err)
	}

	if err := p.AddIP("1.1.1.1"); err != nil {
		t.Fatalf("AddIP failed: %v", err)
	}

	found1, found100 := false, false
	called, idle := false, false
	p.OnRecv = func(ip *net.IPAddr, d time.Duration) {
		called = true
		if ip.String() == "127.0.0.1" {
			found1 = true
		} else if ip.String() == "127.0.0.100" {
			found100 = true
		} /*else if ip.String() == "::1" {
			foundv6 = true
		}*/
	}

	p.OnIdle = func() {
		idle = true
	}

	err := p.Run()
	if err != nil {
		t.Fatalf("Pinger returns error: %v", err)
	}
	if !called {
		t.Fatalf("Pinger didn't get any responses")
	}
	if !idle {
		t.Fatalf("Pinger didn't call OnIdle function")
	}
	if !found1 {
		t.Fatalf("Pinger `127.0.0.1` didn't respond")
	}
	if found100 {
		t.Fatalf("Pinger `127.0.0.100` responded")
	}
	/*if !foundv6 {
		t.Fatalf("Pinger `::1` didn't responded")
	}*/

}

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
