package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Li-giegie/rabbit-cli/example/cmd"
	"os"
)

func main() {
	Execute()
}

var m = make(map[string]string)

func Execute() {
	ctx := context.WithValue(context.Background(), "m", m)
	s := bufio.NewScanner(os.Stdin)
	fmt.Print(">>")
	for s.Scan() {
		if s.Text() != "" {
			_, err := cmd.Group.ExecuteCmdLineContext(ctx, s.Text())
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print(">>")
	}
}
