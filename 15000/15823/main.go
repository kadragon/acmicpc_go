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

	var n, m, d int
	fmt.Fscan(rd, &n, &m)

	dt := make([]int, n+1)
	check := make([]int, 500001)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d)
		if dt[i-1]+1 < i-check[d] {
			dt[i] = dt[i-1] + 1
		} else {
			dt[i] = i - check[d]
		}

		check[d] = i
	}

	dp := make([]int, n+2)
	for i := 1; i <= n; i++ {
		dp[dt[i]]++
	}

	for i := n; i >= 1; i-- {
		dp[i] += dp[i+1]
		if dp[i] >= m {
			cnt := 0
			for j := 1; j <= n; j++ {
				if dt[j] >= i {
					j = j + i - 1
					cnt++
				}
			}

			if cnt >= m {
				fmt.Fprintf(wr, "%d\n", i)
				return
			}
		}
	}
}
