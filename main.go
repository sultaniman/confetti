package main

import (
	"github.com/imanhodjaev/getout/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
