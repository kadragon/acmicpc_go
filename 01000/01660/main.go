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

	var n int
	fmt.Fscan(rd, &n)

	a := []int{0}
	for i := 1; a[i-1] <= n; i++ {
		a = append(a, a[i-1]+i)
	}

	b := []int{0}
	for i := 1; b[i-1]+a[i] <= n; i++ {
		b = append(b, b[i-1]+a[i])
	}

	dp := make([]int, 300001)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		k := 300001
		for j := 1; j < len(b) && b[j] <= i; j++ {
			if dp[i-b[j]] < k {
				k = dp[i-b[j]]
			}
		}
		dp[i] = k + 1
	}

	fmt.Fprintf(wr, "%d\n", dp[n])
}
