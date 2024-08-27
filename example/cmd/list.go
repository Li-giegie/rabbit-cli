package cmd

import (
	"fmt"
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var list = &rabbit_cli.Cmd{
	Name:        "list",
	Description: "输出实例键值对",
	Run: func(c *rabbit_cli.Cmd, args []string) {
		for s, s2 := range m {
			fmt.Printf("%s \"%s\"", s, s2)
		}
	},
}

func init() {
	list.AddSub(&rabbit_cli.Cmd{
		Name:        "help",
		Description: "帮助命令",
		Run: func(c *rabbit_cli.Cmd, args []string) {
			list.Usage()
		},
	})
	Group.AddCmdMust(list)
}
