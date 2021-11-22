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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var k, d, a int
	fmt.Fscanf(rd, "%d/%d/%d", &k, &d, &a)

	if k+a < d || d == 0 {
		fmt.Fprintf(wr, "hasu")
	} else {
		fmt.Fprintf(wr, "gosu")
	}
}
