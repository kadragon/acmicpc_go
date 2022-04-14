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

	var n, tmp, total, max int
	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &tmp)
		total += tmp
		if max < tmp {
			max = tmp
		}
	}

	fmt.Fprintf(wr, "%d\n", total+max*(n-2))
}
