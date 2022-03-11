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

	var a, b, c int
	fmt.Fscanf(rd, "%d + %d = %d", &a, &b, &c)

	if a+b == c {
		fmt.Fprint(wr, "YES\n")
	} else {
		fmt.Fprint(wr, "NO\n")
	}
}
