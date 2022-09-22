package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	const max = 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
}

func main() {
	ans := []int{0}
	var binary []string = strings.Split(Reverse(fmt.Sprintf("%b", StrToInt(NextLine(sc)))), "")
	for i := 0; i < len(binary); i++ {
		if binary[i] == "1" {
			l := ans
			var plus int = int(math.Pow(float64(2), float64(i)))

			for j := 0; j < len(l); j++ {
				ans = append(ans, l[j]+plus)
			}
		}
	}

	for i := 0; i < len(ans); i++ {
		fmt.Println(int(ans[i]))
	}

}

// Reverse 文字列を反転
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// NextLine buinfo.Scanのポインタを渡し、標準入力の次の行を読み込み
// ex. sc := buinfo.NewScanner(os.stdin)
//      GetNextLine(sc)
func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}

// StrToInt string型をint型に変換
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
