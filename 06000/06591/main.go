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

	for {
		var n, k int
		fmt.Fscan(rd, &n, &k)
		if n == 0 && k == 0 {
			break
		}

		if k > n-k {
			k = n - k
		}

		ans := 1
		for i := 0; i < k; i++ {
			ans = ans * (n - i) / (i + 1)
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
