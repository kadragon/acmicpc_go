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
	for {
		s, _, _ := rd.ReadLine()
		if len(s) == 0 {
			break
		}
		n++
	}

	fmt.Fprintf(wr, "%d\n", n)
}
