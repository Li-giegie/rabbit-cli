package rabbit_cli

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	root := &Cmd{
		Name:        "root",
		Description: "初始命令",
		Run: func(c *Cmd, args []string) {
			fmt.Println("args", args)
		},
	}
	_, err := root.Execute([]string{"1"})
	if err != nil {
		fmt.Println(err)
	}
	_, err = root.Execute([]string{"root"})
	if err != nil {
		fmt.Println(err)
	}
}

func TestGroupCmd(t *testing.T) {
	g := &GroupCmd{
		Description: "command group",
	}
	g.AddCmd(&Cmd{
		Name:        "help",
		Description: "",
		Run: func(c *Cmd, args []string) {
			g.Usage()
		},
	})
	_, err := g.ExecuteCmdLine("help")
	if err != nil {
		fmt.Println(err)
	}

}
