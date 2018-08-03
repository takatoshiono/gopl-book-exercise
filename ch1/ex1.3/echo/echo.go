// Echoは2種類のバージョンのEchoの実行時の差をベンチマークして結果を表示します
package echo

import (
	"os"
	"strings"
)

// Echo1は非効率バージョン
func Echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// ベンチマーク実行時に出力するとベンチマークの結果出力と混ざるので値を返すだけにする
	//fmt.Println(s)
	return s
}

// Echo3は効率的なバージョン
func Echo3() string {
	// ベンチマーク実行時に出力するとベンチマークの結果出力と混ざるので値を返すだけにする
	//fmt.Println(strings.Join(os.Args[1:], " "))
	return strings.Join(os.Args[1:], " ")
}
