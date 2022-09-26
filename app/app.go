package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IPs and Servers"
	app.Usage = "Search for ips on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "juanpetrus.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "search address ips in intenert",
			Flags:  flags,
			Action: SearchIps,
		},
		{
			Name:   "server",
			Usage:  "search servers in intenert",
			Flags:  flags,
			Action: SearchServers,
		},
	}
	return app
}

func SearchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func SearchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servers := range servers {
		fmt.Println(servers.Host)
	}
}
