package api

import "time"

type Latency struct {
	IP     string        `json:"ip"`
	RTT    time.Duration `json:"rtt" binding:"required"`
	STATUS string        `json:"status"`
	UNIT   string        `json:"unit"`
}
