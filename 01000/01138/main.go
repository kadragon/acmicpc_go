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

	var n, t int
	fmt.Fscan(rd, &n)

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &t)
		for j := 1; j <= n; j++ {
			if t == 0 && a[j] == 0 {
				a[j] = i
				break
			}
			if a[j] == 0 {
				t--
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(wr, "%d ", a[i])
	}
}
