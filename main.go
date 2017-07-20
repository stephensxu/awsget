package main

import (
	"github.com/urfave/cli"
	"os/exec"
	"fmt"
	"os"
	"encoding/json"
)

func main() {
	type InstanceList struct {
		List string `json:list`
	}

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

				// results := string(out)

				data := &InstanceList{}
				err = json.Unmarshal([]byte(out), data)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(data)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
