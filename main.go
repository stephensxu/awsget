package main

import (
	"github.com/urfave/cli"
	"os/exec"
	"fmt"
	"os"
	"encoding/json"
	"github.com/fatih/color"
)

type AmazonResponse struct {
	Reservations []Reservation `json:"Reservations"`
}

type Reservation struct {
	Instances []Instance `json: "Instances"`
}

type Instance struct {
	PublicDnsName string `json:"PublicDnsName"`
	Tags []Tag `json:"Tags"`
}

type Tag struct {
	Value string `json: "Value"`
	Key string `json: "Key"`
}

type SimplifiedResult struct {
	InstanceName string
	PublicDnsName string
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name: "ssh-ips",
			Usage: "get list of PublicDnsName of ec2 instances",
			Action: func(c *cli.Context) error {
				cmd := "aws"
				args := []string{"ec2", "describe-instances"}

				out, err := exec.Command(cmd, args...).Output()

				if err != nil {
					fmt.Println(err)
				}

				amazonResponse := &AmazonResponse{}
				json.Unmarshal(out, amazonResponse)
				reservations := amazonResponse.Reservations

				results := []SimplifiedResult{}

				for _, reservation := range reservations {
					instance := reservation.Instances[0]
					if instance.PublicDnsName == "" || len(instance.Tags) == 0 {
						continue
					}
					simplifiedResult := SimplifiedResult{}
					simplifiedResult.PublicDnsName = instance.PublicDnsName
					simplifiedResult.InstanceName = instance.Tags[0].Value
					results = append(results, simplifiedResult)
				}

				for _, item := range results {
					color.Green(string(item.InstanceName))
					color.Green(string(item.PublicDnsName))
					fmt.Println("")
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
