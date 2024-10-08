package cmd

import (
	"fmt"
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var list = &rabbit_cli.Cmd{
	Name:        "list",
	Description: "输出实例键值对",
	Run: func(c *rabbit_cli.Cmd, args []string) {
		m := c.Context().Value("m").(map[string]string)
		for s, s2 := range m {
			fmt.Printf("%s \"%s\"\n", s, s2)
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
	list.Flags().BoolFunc("help", "", func(s string) error {
		list.Usage()
		return nil
	})
	list.Flags().BoolFunc("h", "", func(s string) error {
		list.Usage()
		return nil
	})
	Group.AddCmdMust(list)
}
