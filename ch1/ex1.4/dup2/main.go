// Dup2は入力に2回以上現れた行の数とその行のテキストを表示します。
// 標準入力から読み込むか、名前が指定されたファイルの一覧から読み込みます。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileCounts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			files := make([]string, len(fileCounts[line]))
			i := 0
			for file := range fileCounts[line] {
				files[i] = file
				i++
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(files, ","))
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if fileCounts[input.Text()] == nil {
			fileCounts[input.Text()] = make(map[string]int)
		}
		fileCounts[input.Text()][f.Name()]++
	}
}
