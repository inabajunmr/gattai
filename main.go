package main

import (
	"os"

	"github.com/inabajunmr/gattai/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
