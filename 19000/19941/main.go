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

	var n, k, ans int
	fmt.Fscan(rd, &n, &k)

	var s string
	fmt.Fscan(rd, &s)

	d := make([]bool, len(s))

	for i, v := range s {
		if v == 'P' {
			st, en := 0, len(s)-1
			if i-k > st {
				st = i - k
			}
			if i+k < en {
				en = i + k
			}

			for j := st; j <= en; j++ {
				if !d[j] && s[j] == 'H' {
					ans++
					d[j] = true
					break
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
