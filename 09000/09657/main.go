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
	dt := [1001]int{0, 1, 0, 1, 1}

	for i := 5; i <= n; i++ {
		if dt[i-1]+dt[i-3]+dt[i-4] < 3 {
			dt[i] = 1
		}
	}

	ans := []string{"CY", "SK"}
	fmt.Fprintf(wr, "%s\n", ans[dt[n]])
}
