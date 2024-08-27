package cmd

import (
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var m = make(map[string]string)

var create = &rabbit_cli.Cmd{
	Name:        "create",
	Description: "创建一个实例键值对",
	RunE: func(c *rabbit_cli.Cmd, args []string) error {
		k := c.Flags().Lookup("k").Value.String()
		v := c.Flags().Lookup("v").Value.String()
		m[k] = v
		return nil
	},
}

func init() {
	create.AddSub(&rabbit_cli.Cmd{
		Name:        "help",
		Description: "帮助命令",
		Run: func(c *rabbit_cli.Cmd, args []string) {
			create.Usage()
		},
	})
	create.Flags().String("k", "", "")
	create.Flags().String("v", "", "")
	Group.AddCmdMust(create)
}
