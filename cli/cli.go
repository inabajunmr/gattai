package cli

import (
	"fmt"
	"os"

	"github.com/inabajunmr/gattai/html/mashup"
)

func Run(args []string) int {
	fmt.Println(args)
	if len(args) != 2 {
		fmt.Println("error")
		os.Exit(1)
	}

	fmt.Println("call")

	file, err := os.Create("./gattai.html")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	file.Write(([]byte)(mashup.Gattai(args[0], args[0])))
	return 0
}
