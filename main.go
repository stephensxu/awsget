package main

import (
	"github.com/urfave/cli"
	"os/exec"
	"fmt"
	"os"
	"github.com/fatih/color"
)

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

				result := string(out)
				color.Green(result)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
