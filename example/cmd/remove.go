package cmd

import (
	"errors"
	"fmt"
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var remove = &rabbit_cli.Cmd{
	Name:        "remove",
	Description: "删除一个实例键值对",
	RunE: func(c *rabbit_cli.Cmd, args []string) error {
		k := c.Flag().Lookup("key")
		if k == nil {
			return errors.New("flag \"key\" not exist")
		}
		if _, ok := m[k.Value.String()]; !ok {
			return fmt.Errorf("%s key not exist", k)
		}
		delete(m, k.Value.String())
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
	remove.Flag().String("key", "", "")
	Group.AddCmdMust(remove)
}
