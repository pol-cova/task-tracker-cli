package main

import (
	"github.com/pol-cova/task-tracker-cli/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
