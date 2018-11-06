package cli

import (
	"fmt"
	"os"
)

func Run(args []string) int {
	if len(args) != 2 {
		fmt.Println("error")
		os.Exit(1)
	}

	fmt.Println("call")

	return 0
}
