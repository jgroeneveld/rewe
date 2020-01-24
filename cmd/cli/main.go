package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := rootCommand(os.Stdout)

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
