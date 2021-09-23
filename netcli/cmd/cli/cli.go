package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:      "Website Lookup CLI",
		Usage:     "Let's you query IPs, CNAMEs and Name Servers!",
		UsageText: "cli [global options] command [command options] [arguments...]",
		Version:   "0.0.0",
		Commands: []*cli.Command{
			{
				Name:  "ns",
				Usage: "Looks Up the NameServers for a Particular Host",
				Action: func(c *cli.Context) error {
					ns, err := net.LookupNS(c.Args().First())
					for _, elem := range ns {
						fmt.Println(elem.Host)

					}
					return err
				},
			},
			{
				Name:  "cname",
				Usage: "Looks up the CNAME for a particular host",
				Action: func(c *cli.Context) error {
					cname, err := net.LookupCNAME(c.Args().First())
					fmt.Println(cname)
					return err
				},
			},
			{
				Name:  "ip",
				Usage: "Looks up the IP addresses for a particular host",
				Action: func(c *cli.Context) error {
					ip, err := net.LookupIP(c.Args().First())
					for _, elem := range ip {
						fmt.Println(elem)
					}
					return err
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
