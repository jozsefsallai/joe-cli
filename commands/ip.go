package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/urfave/cli"
)

func getIPv4() string {
	resp, err := http.Get("https://api.ipify.org?format=text")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	ip4, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(ip4)
}

func getIPv6() string {
	resp, err := http.Get("https://api6.ipify.org?format=text")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	ip6, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(ip6)
}

// IPCommand adds IP address fetching functionality to the CLI
var IPCommand = cli.Command{
	Name:  "ip",
	Usage: "Get the IP address of this machine",
	Action: func(ctx *cli.Context) error {
		fmt.Println("IPv4:", getIPv4())
		fmt.Println("IPv6:", getIPv6())
		return nil
	},
	Subcommands: []cli.Command{
		{
			Name:  "v4",
			Usage: "Get the IPv4 address of this machine",
			Action: func(ctx *cli.Context) error {
				fmt.Println(getIPv4())
				return nil
			},
		},
		{
			Name:  "v6",
			Usage: "Get the IPv6 address of this machine",
			Action: func(ctx *cli.Context) error {
				fmt.Println(getIPv6())
				return nil
			},
		},
	},
}
