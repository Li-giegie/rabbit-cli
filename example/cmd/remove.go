package cmd

import (
	"fmt"
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var remove = &rabbit_cli.Cmd{
	Name:        "remove",
	Description: "删除一个实例键值对",
	RunE: func(c *rabbit_cli.Cmd, args []string) error {
		m := c.Context().Value("m").(map[string]string)
		k, err := c.Flags().GetString("k")
		if err != nil {
			return err
		}
		if _, ok := m[k]; !ok {
			return fmt.Errorf("key \"%s\" not exist", k)
		}
		delete(m, k)
		return nil
	},
}

func init() {
	remove.AddSub(&rabbit_cli.Cmd{
		Name:        "help",
		Description: "帮助命令",
		Run: func(c *rabbit_cli.Cmd, args []string) {
			remove.Usage()
		},
	})
	remove.Flags().BoolFunc("help", "", func(s string) error {
		remove.Usage()
		return nil
	})
	remove.Flags().BoolFunc("h", "", func(s string) error {
		remove.Usage()
		return nil
	})
	remove.Flags().String("k", "", "")
	Group.AddCmdMust(remove)
}
