package rabbit_cli

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type FlagSet struct {
	*flag.FlagSet
}

func (f *FlagSet) Reset() {
	f.VisitAll(func(f *flag.Flag) {
		err := f.Value.Set(f.DefValue)
		if err != nil {
			fmt.Println(f.Usage, f.Name, err)
		}
	})
}

// Cmd 对命令的定义，没有任何默认行为
type Cmd struct {
	Name        string
	Description string
	Run         func(c *Cmd, args []string)
	RunE        func(c *Cmd, args []string) error
	ctx         context.Context
	sub         map[string]*Cmd
	flag        *FlagSet
}

func (c *Cmd) Context() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}

func (c *Cmd) AddSub(cmd *Cmd) bool {
	if c.sub == nil {
		c.sub = make(map[string]*Cmd)
	}
	if cmd.Name == "" || cmd.Name[0] == '-' {
		panic("invalid command name \"" + cmd.Name + "\"")
	}
	_, ok := c.sub[cmd.Name]
	if ok {
		return false
	}
	c.sub[cmd.Name] = cmd
	return true
}

func (c *Cmd) AddSubMust(cmd *Cmd) {
	ok := c.AddSub(cmd)
	if !ok {
		panic(cmd.Name + " sub command already exists")
	}
}

// Usage 调用UsageInfo输出用法信息
func (c *Cmd) Usage() {
	_, _ = os.Stdout.WriteString(c.UsageInfo())
}

// UsageInfo 用法信息
func (c *Cmd) UsageInfo() string {
	buf := bytes.NewBuffer(make([]byte, 0, 127))
	if c.Description != "" {
		buf.WriteString(c.Description + "\n")
	}
	buf.WriteString("Usage:\n  " + c.Name + "  [command]\n")
	if len(c.sub) > 0 {
		buf.WriteString("Sub commands:\n")
		for s, cmd := range c.sub {
			buf.WriteString("  " + s)
			if cmd.Description != "" {
				buf.WriteString("  " + cmd.Description)
			}
			buf.WriteByte(10)
		}
	}
	if c.flag != nil && c.flag.FlagSet != nil {
		buf.WriteString("Flags:\n")
		c.flag.FlagSet.VisitAll(func(f *flag.Flag) {
			buf.WriteString("  -" + f.Name)
			if f.Usage != "" {
				buf.WriteString("   " + f.Usage)
			}
			if f.DefValue != "" {
				buf.WriteString("   (default \"" + f.DefValue + "\")")
			}
			buf.WriteByte(10)
		})
	}
	return buf.String()
}

func (c *Cmd) Flags() *FlagSet {
	if c.flag == nil {
		c.flag = new(FlagSet)
		c.flag.FlagSet = flag.NewFlagSet(c.Name, flag.ContinueOnError)
		c.flag.Usage = func() {}
		c.flag.SetOutput(io.Discard)
	}
	return c.flag
}

var ErrNotRun = errors.New("command is executed, but the Run method is nil")

func (c *Cmd) Execute(args []string) (cmd *Cmd, err error) {
	defer func() {
		if err == nil {
			if c.Run == nil && c.RunE == nil {
				err = ErrNotRun
				return
			}
			if c.RunE != nil {
				err = c.RunE(c, args)
				return
			}
			c.Run(c, args)
		}
	}()
	if len(args) == 0 {
		return c, nil
	}
	if c.flag == nil {
		return c, nil
	}
	err = c.flag.Parse(args)
	args = c.flag.Args()
	return c, err
}

func (c *Cmd) ExecuteContext(ctx context.Context, args []string) (cmd *Cmd, err error) {
	c.ctx = ctx
	return c.Execute(args)
}

// GroupCmd 一组命令的定义，没有任何默认行为
type GroupCmd struct {
	Description string
	m           map[string]*Cmd
}

func (p *GroupCmd) Usage() {
	_, _ = os.Stdout.WriteString(p.UsageInfo())
}

// UsageInfo 用法信息
func (p *GroupCmd) UsageInfo() string {
	buf := bytes.NewBuffer(make([]byte, 0, 127))
	if p.Description != "" {
		buf.WriteString(p.Description)
		buf.WriteByte(10)
	}
	if p.m != nil {
		buf.WriteString("Usage:\n")
		for s, cmd := range p.m {
			buf.WriteString("  " + s + " [command]")
			if cmd.Description != "" {
				buf.WriteString("  " + cmd.Description)
			}
			buf.WriteByte(10)
		}
	}
	return buf.String()
}

func (p *GroupCmd) AddCmd(cmds ...*Cmd) bool {
	if p.m == nil {
		p.m = make(map[string]*Cmd)
	}
	for _, cmd := range cmds {
		if cmd.Name == "" || cmd.Name[0] == '-' {
			panic("invalid command name \"" + cmd.Name + "\"")
		}
		if _, ok := p.m[cmd.Name]; ok {
			return false
		}
		p.m[cmd.Name] = cmd
	}
	return true
}

func (p *GroupCmd) AddCmdMust(cmds ...*Cmd) {
	for _, cmd := range cmds {
		ok := p.AddCmd(cmd)
		if !ok {
			panic(cmd.Name + " command already exists")
		}
	}
}

// ExecuteCmdLine s 输入的命令行字符串，返回值*Cmd如果不为nil，则代表错误发生在执行的命令中，等于nil则代表GroupCmd执行错误
func (p *GroupCmd) ExecuteCmdLine(s string) (*Cmd, error) {
	return p.Execute(strings.Fields(s))
}

// ExecuteCmdLineContext 同上但支持传递ctx 参数
func (p *GroupCmd) ExecuteCmdLineContext(ctx context.Context, s string) (*Cmd, error) {
	return p.ExecuteContext(ctx, strings.Fields(s))
}

// Execute args 输入的命令参数，返回值*Cmd如果不为nil，则代表错误发生在执行的命令中，等于nil则代表GroupCmd执行错误
func (p *GroupCmd) Execute(args []string) (*Cmd, error) {
	cmd, _args, err := p.queryCmd(args)
	if err != nil {
		return nil, err
	}
	return cmd.Execute(_args)
}

func (p *GroupCmd) ExecuteContext(ctx context.Context, args []string) (*Cmd, error) {
	cmd, _args, err := p.queryCmd(args)
	if err != nil {
		return nil, err
	}
	return cmd.ExecuteContext(ctx, _args)
}

func (p *GroupCmd) queryCmd(args []string) (*Cmd, []string, error) {
	if len(args) == 0 {
		return nil, nil, fmt.Errorf("args is empty")
	}
	var cmd, temp *Cmd
	var ok bool
	var m = p.m
	var index int
	for i, arg := range args {
		if len(arg) == 0 {
			continue
		}
		if arg[0] == '-' {
			break
		}
		temp, ok = m[arg]
		if ok {
			index = i
			cmd = temp
			m = temp.sub
		} else {
			break
		}
	}
	if cmd == nil {
		return nil, nil, fmt.Errorf("Error \"%s\" command does not exist", args[0])
	}
	return cmd, args[index+1:], nil
}

func (p *GroupCmd) VisitAll(fn func(c *Cmd)) {
	if p.m == nil {
		return
	}
	for _, cmd := range p.m {
		fn(cmd)
	}
}
