package cli

import (
	"fmt"
	"os"

	"github.com/inabajunmr/gattai/html/mashup"
)

func Run(args []string) int {
	if len(args) != 3 {
		fmt.Println("You must specify 2 urls as arguments.")
		os.Exit(1)
	}

	file, err := os.Create("./gattai.html")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	file.Write(([]byte)(mashup.Gattai(args[1], args[2])))
	return 0
}
