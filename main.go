package main

import (
	"github.com/imanhodjaev/confetti/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
