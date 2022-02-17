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

	var n, i int64
	fmt.Fscan(rd, &n)

	var ans int64
	for i = 1; i < n; i++ {
		ans += n*i + i
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
