package api

type Latency struct {
	IP     string  `json:"ip"`
	RTT    float64 `json:"rtt" binding:"required"`
	STATUS string  `json:"status"`
	UNIT   string  `json:"unit"`
}
