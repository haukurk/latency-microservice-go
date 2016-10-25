package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/haukurk/latency-microservice-go/client"
)

func main() {

	app := cli.NewApp()
	app.Name = "latency cli"
	app.Usage = "cli to work with the latency service"
	app.Version = "0.2.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Value: "http://localhost:8080", Usage: "Hostname for latency API", EnvVar: "APP_HOST"},
		cli.StringFlag{Name: "remotehost", Value: "8.8.8.8", Usage: "Host to check latency to", EnvVar: "APP_REMOTE_HOST"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "latency",
			Usage: "(title description) check latency to a remote ip",
			Action: func(c *cli.Context) {

				host := c.GlobalString("host")
				remotehost := c.GlobalString("remotehost")

				client := client.LatencyClient{Host: host, RemoteHost: remotehost}

				latencyData, err := client.PingHost()

				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Printf("%+v\n", latencyData)
			},
		},
	}

	app.Run(os.Args)
}
