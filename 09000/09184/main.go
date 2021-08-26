package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dp [21][21][21]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func w(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	} else if a > 20 || b > 20 || c > 20 {
		return w(20, 20, 20)
	}

	if dp[a][b][c] != 0 {
		return dp[a][b][c]
	}

	if a < b && b < c {
		dp[a][b][c] = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
	} else {
		dp[a][b][c] = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
	}
	return dp[a][b][c]
}

func main() {
	defer wr.Flush()

	var a, b, c int

	for {
		fmt.Fscan(rd, &a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		fmt.Fprintf(wr, "w(%d, %d, %d) = %d\n", a, b, c, w(a, b, c))
	}
}
