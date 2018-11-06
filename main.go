package main

import (
	"os"

	"github.com/inabajunmr/kimera-chang/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
