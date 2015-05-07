package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/haukurk/latency-microservice-go/service"
)

func getConfig(c *cli.Context) (service.Config, error) {
	jsonPath := c.GlobalString("config")
	config := service.Config{}

	file, _ := os.Open(jsonPath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("error:", err)
	}
	return config, err
}

func main() {

	app := cli.NewApp()
	app.Name = "latency-go"
	app.Usage = "Analyse latency with Go"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"config", "config.json", "config file (json format) to use", "APP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the HTTP server",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := service.LatencyService{}

				if err = svc.Run(cfg); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)

}
