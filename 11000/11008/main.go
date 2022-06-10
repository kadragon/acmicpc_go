package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var t int
	var s, p string

	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		ans := 0
		fmt.Fscan(rd, &s, &p)

		for {
			t := strings.Index(s, p)
			if t == -1 {
				ans += len(s)
				break
			} else {
				ans += t + 1
				s = s[t+len(p):]
			}
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
