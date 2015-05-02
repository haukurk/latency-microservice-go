package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tatsushid/go-fastping"
)

type Configuration struct {
	Port int
}

func main() {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("error:", err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello I'm a latency analyser.")
	})
	router.GET("/ping/:host", func(context *gin.Context) {
		ip := context.Params.ByName("host")

		p := fastping.NewPinger()
		ra, err := net.ResolveIPAddr("ip4:icmp", ip)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			context.JSON(200, gin.H{"status": "ok", "ip": addr.String(), "rtt": rtt / 1000000.0, "unit": "ms"})
		}
		p.OnIdle = func() {
			log.Printf("log idle")
		}

		err = p.Run()
		if err != nil {
			context.JSON(200, gin.H{"status": "ok", "message": err})
		}

	})
	router.Run(fmt.Sprint(":", configuration.Port))
}
