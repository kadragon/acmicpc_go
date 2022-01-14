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

	var n int64
	var ans string = "SK"
	fmt.Fscan(rd, &n)

	if n%7 == 2 || n%7 == 0 {
		ans = "CY"
	}

	fmt.Fprintf(wr, "%s\n", ans)
}
