package main

import (
	"fmt"
	"github.com/neefrankie/bump/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
