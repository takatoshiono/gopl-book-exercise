// Echoはコマンドライン引数をインデックスと一緒に1行ごとに表示します
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
