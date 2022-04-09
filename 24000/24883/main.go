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

	var s string
	fmt.Fscan(rd, &s)

	if s == "N" || s == "n" {
		fmt.Fprintln(wr, "Naver D2")
	} else {
		fmt.Fprintln(wr, "Naver Whale")
	}
}
