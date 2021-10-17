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

	var t, a int
	fmt.Fscan(rd, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(rd, &a)
		fmt.Fprintf(wr, "%d\n", a*23)
	}
}
