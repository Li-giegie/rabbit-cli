package main

import (
	"bufio"
	"fmt"
	"github.com/Li-giegie/rabbit-cli/example/cmd"
	"os"
)

func main() {
	Execute()
}

func Execute() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print(">>")
	for s.Scan() {
		if s.Text() != "" {
			_, err := cmd.Group.ExecuteCmdLine(s.Text())
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print(">>")
	}
}
