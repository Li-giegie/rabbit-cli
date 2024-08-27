package cmd

import (
	"errors"
	"fmt"
	rabbit_cli "github.com/Li-giegie/rabbit-cli"
)

var m = make(map[string]string)

var create = &rabbit_cli.Cmd{
	Name:        "create",
	Description: "创建一个实例键值对",
	RunE: func(c *rabbit_cli.Cmd, args []string) error {

		c.Usage()
		return nil
		k := c.Flag().Lookup("key")
		if k == nil {
			return errors.New("flag \"key\" not exist")
		}
		v := c.Flag().Lookup("value")
		if v == nil {
			return errors.New("flag \"value\" not exist")
		}
		if k.Value.String() == "" {
			return errors.New("key can't be empty")
		}
		_, ok := m[k.Value.String()]
		if ok {
			return fmt.Errorf("%s exist", v)
		}
		m[k.Value.String()] = v.Value.String()
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
	create.Flag().String("key", "", "")
	create.Flag().String("value", "", "")
	Group.AddCmdMust(create)
}
