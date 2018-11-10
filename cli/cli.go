package cli

import (
	"fmt"
	"os"

	"github.com/inabajunmr/gattai/html/mashup"
)

func Run(args []string) int {
	fmt.Println(args)
	if len(args) != 3 {
		fmt.Println("error")
		os.Exit(1)
	}

	fmt.Println("call")

	file, err := os.Create("./gattai.html")
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	file.Write(([]byte)(mashup.Gattai(args[1], args[2])))
	return 0
}
