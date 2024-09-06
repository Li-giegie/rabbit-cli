package cmd

import (
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var create = &rabbit_cli.Cmd{
	Name:        "create",
	Description: "创建一个实例键值对",
	RunE: func(c *rabbit_cli.Cmd, args []string) error {
		m := c.Context().Value("m").(map[string]string)
		k, err := c.Flags().GetString("k")
		if err != nil {
			return err
		}
		v, err := c.Flags().GetString("v")
		if err != nil {
			return err
		}
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
	create.Flags().BoolFunc("help", "", func(s string) error {
		create.Usage()
		return nil
	})
	create.Flags().BoolFunc("h", "", func(s string) error {
		create.Usage()
		return nil
	})
	create.Flags().String("k", "", "")
	create.Flags().String("v", "", "")
	Group.AddCmdMust(create)
}
