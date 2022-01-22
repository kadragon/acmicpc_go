package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var s string
	var ans int
	conv := []int{26, 10}

	fmt.Fscan(rd, &s)

	ans = conv[s[0]-'c']
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			ans *= conv[s[i]-'c'] - 1
		} else {
			ans *= conv[s[i]-'c']
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
