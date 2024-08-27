package rabbit_cli

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	p := new(GroupCmd)
	root := &Cmd{
		Name:        "root",
		Description: "初始命令",
		Run:         nil,
	}
	root.Flag().String("a", "123", "属性a")
	root.Flag().String("b", "123", "属性b")
	list := &Cmd{
		Name: "list",
		Run: func(c *Cmd, args []string) {
			fmt.Println("root.list", args)
			fmt.Println(c.Flag().Lookup("b"))
		},
	}
	list.Flag().String("b", "234", "b")
	root.AddSub(list)
	root.AddSub(&Cmd{
		Name:        "create",
		Description: "创建一个实例",
	})
	if !p.AddCmd(root) {
		log.Fatalln("cmd exist")
	}
	_, err := p.ExecuteCmdStr("-h")
	if err != nil {
		fmt.Println(err)
		p.Usage()
	}
}

func TestFlagSet(t *testing.T) {
	f := new(FlagSet)
	f.FlagSet = flag.NewFlagSet("test", flag.ContinueOnError)
	n := f.String("n", "", "")
	a := f.Int("a", 0, "0")
	b := f.Bool("b", false, "")
	d := f.Duration("abc", time.Second, "")
	err := f.Parse(strings.Fields("-n 123 -a 10 -b=true -abc 1ms"))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(*n, *a, *b, *d)
	f.Reset()
	fmt.Println(*n, *a, *b, *d)
}
