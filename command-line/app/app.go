package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func getIps(c *cli.Context) {
	host := c.String("address")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func App() *cli.App {
	application := cli.NewApp()
	application.Name = "Command Line Application"
	application.Usage = "Search ip address on the internet"

	application.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search ip address on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "address",
					Value: "mercadolivre.com.br",
				},
			},
			Action: getIps,
		},
	}

	return application
}