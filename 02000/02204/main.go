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

	for {
		var n int
		fmt.Fscan(rd, &n)

		if n == 0 {
			break
		}

		s := make([]string, n)
		p := make([]string, n)
		var a int
		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &s[i])
			p[i] = strings.ToLower(s[i])

			if p[a] > p[i] {
				a = i
			}
		}

		fmt.Fprintf(wr, "%s\n", s[a])
	}
}
