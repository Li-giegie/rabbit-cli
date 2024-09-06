package rabbit_cli

import (
	"flag"
	"fmt"
	"testing"
	"time"
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

func TestFlagSetGet(t *testing.T) {
	f := new(FlagSet)
	f.FlagSet = flag.NewFlagSet("", flag.ContinueOnError)
	f.String("a", "default", "")
	f.Int("b", -1, "")
	f.Uint("c", 0, "")
	f.Float64("d", 0.0000009, "")
	f.Bool("e", false, "")
	f.Duration("f", time.Duration(time.Now().UnixNano()), "")
	fmt.Println(f.GetString("a"))
	fmt.Println(f.GetInt("b"))
	fmt.Println(f.GetUint("c"))
	fmt.Println(f.GetFloat64("d"))
	fmt.Println(f.GetBool("e"))
	fmt.Println(f.GetDuration("f"))
	fmt.Println(f.GetDuration("g"))
}
