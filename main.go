package main

import (
	"encoding/json"
	"fmt"
	"github.com/OpenIoTHub/getip/iputils"
	"github.com/OpenIoTHub/getip/models"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sort"
)

func main() {
	ipv4 := iputils.GetMyPublicIpv4()
	ipv6 := iputils.GetMyPublicIpv6()
	output := models.Output{
		Ipv4Addr: ipv4,
		Ipv6Addr: ipv6,
	}
	outputJson := func(c *cli.Context) error {
		jsonBytes, _ := json.Marshal(output)
		fmt.Println(string(jsonBytes))
		return nil
	}
	outputYaml := func(c *cli.Context) error {
		yamlBytes, _ := yaml.Marshal(output)
		fmt.Print(string(yamlBytes))
		return nil
	}
	app := &cli.App{
		Name:  "getip",
		Usage: "Get my public ip",
		Commands: []*cli.Command{
			{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "output formate:json",
				Action:  outputJson,
			},
			{
				Name:    "yaml",
				Aliases: []string{"y"},
				Usage:   "output formate:yaml",
				Action:  outputYaml,
			},
		},
		Action: outputJson,
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
