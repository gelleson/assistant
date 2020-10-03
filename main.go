package main

import (
	"alfred/internal/cmd"
	"fmt"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
