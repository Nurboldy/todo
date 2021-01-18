package main

import (
	"fmt"

	"github.com/Nurboldy/todo/cmd"
)

func main() {
	cmd.RootCmd.Execute()

	fmt.Println("hello func", cmd.Hello())
}
