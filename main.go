package main

import (
	"os"

	"github.com/ebadfd/helm-repo-html/cmd"
)

var version = "0.0.1"

func main() {
	if err := cmd.Execute(version); err != nil {
		os.Exit(1)
	}
}
