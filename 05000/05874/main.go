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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var s string
	fmt.Fscan(rd, &s)

	l, ans := 0, 0

	for i := 1; i < len(s); i++ {
		if s[i] == '(' && s[i-1] == '(' {
			l++
		}
		if s[i] == ')' && s[i-1] == ')' {
			ans += l
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
