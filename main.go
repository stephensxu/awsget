package main

import (
	"github.com/urfave/cli"
	"os/exec"
	"fmt"
	"os"
	"encoding/json"
)

type AmazonResponse struct {
	Reservations map[string]interface{} `json:"Reservations"`
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

// type Response struct {
// 	SimplifiedResults []SimplifiedResult
// }

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

				instances := amazonResponse.Reservations

				fmt.Println(len(instances))

				if err != nil {
					fmt.Println(err)
				}

				resp := []SimplifiedResult{}

				// for _, instance := range instances {
					// fmt.Println("hi")
					// simplifiedResult := &SimplifiedResult{}
					// simplifiedResult.InstanceName = instance.Tags[0].Value
					// simplifiedResult.PublicDnsName = instance.PublicDnsName

					// fmt.Println(string(simplifiedResult.PublicDnsName))
					// fmt.Println(string(instance.PublicDnsName))
				// }

				final, err := json.Marshal(resp)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(final))
				return nil
			},
		},
	}
	app.Run(os.Args)
}
