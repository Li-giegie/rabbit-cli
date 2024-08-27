package cmd

import (
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var help = &rabbit_cli.Cmd{
	Name:        "help",
	Description: "输出实例键值对",
	Run: func(c *rabbit_cli.Cmd, args []string) {
		Group.Usage()
	},
}

func init() {
	Group.AddCmdMust(help)
}
