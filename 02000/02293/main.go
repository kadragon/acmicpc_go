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

	var n, k, v int
	fmt.Fscan(rd, &n, &k)

	dt := make([]int, k+1)
	dt[0] = 1

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &v)
		for j := 0; j <= k; j++ {
			if j >= v {
				dt[j] += dt[j-v]
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[k])
}
