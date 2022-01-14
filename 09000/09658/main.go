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
	var ans string
	fmt.Fscan(rd, &n)

	if n%7 == 3 || n%7 == 1 {
		ans = "CY"
	} else {
		ans = "SK"
	}

	fmt.Fprintf(wr, "%s\n", ans)
}
