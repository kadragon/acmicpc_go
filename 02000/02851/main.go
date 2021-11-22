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

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	d := make([]int, 11)

	for i := 1; i <= 10; i++ {
		fmt.Fscan(rd, &d[i])
		d[i] += d[i-1]
	}

	ans := 0

	for i := 10; i >= 1; i-- {
		if abs(100, ans) > abs(100, d[i]) {
			ans = d[i]
		}	
	}

	fmt.Fprintf(wr, "%#v\n", ans)
}
