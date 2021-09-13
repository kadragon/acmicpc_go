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

	var a, b string
	var la, lb int
	fmt.Fscan(rd, &a, &b)

	la = len(a)
	lb = len(b)

	for i := 0; i < la*lb; i++ {
		if a[i%la] != b[i%lb] {
			fmt.Fprintf(wr, "0")
			return
		}
	}
	fmt.Fprintf(wr, "1")
}
