package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	data []int
	dp   [][]int

	ans int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	ans = -1
}

func f(n, s, m, end int) {
	if n > end {
		if ans < s {
			ans = s
		}
		return
	}

	if s+data[n] <= m {
		if dp[n][s+data[n]] == 0 {
			dp[n][s+data[n]] = 1
			f(n+1, s+data[n], m, end)
		}
	}
	if s-data[n] >= 0 {
		if dp[n][s-data[n]] == 0 {
			dp[n][s-data[n]] = 1
			f(n+1, s-data[n], m, end)
		}
	}
}

func main() {
	defer wr.Flush()

	var n, s, m int
	fmt.Fscan(rd, &n, &s, &m)

	data = make([]int, n+1)
	dp = make([][]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, 1001)
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &data[i])
	}

	f(1, s, m, n)

	wr.WriteString(fmt.Sprintf("%#v\n", ans))
}
