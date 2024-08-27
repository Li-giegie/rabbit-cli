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
			triggerCmd, err := cmd.Group.ExecuteCmdStr(s.Text())
			if err != nil {
				fmt.Println(err)
				if triggerCmd != nil {
					triggerCmd.Usage()
				} else {
					cmd.Group.Usage()
				}
			}
		}
		fmt.Print(">>")
	}
}
