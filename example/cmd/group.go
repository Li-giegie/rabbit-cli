package cmd

import rabbit_cli "github.com/Li-giegie/rabbit-cli"

var Group = rabbit_cli.GroupCmd{
	Description: "command group",
}

func init() {
	Group.AddCmdMust(&rabbit_cli.Cmd{
		Name:        "help",
		Description: "帮助信息",
		Run: func(c *rabbit_cli.Cmd, args []string) {
			Group.Usage()
		},
	})
}
