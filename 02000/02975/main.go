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

	var a, b int
	var s string
	for {
		fmt.Fscanf(rd, "%d %s %d\n", &a, &s, &b)
		if a == 0 && b == 0 && s == "W" {
			break
		}

		if s == "W" {
			if a-b >= -200 {
				fmt.Fprintf(wr, "%d\n", a-b)
			} else {
				fmt.Fprintf(wr, "Not allowed\n")
			}
		} else {
			fmt.Fprintf(wr, "%d\n", a+b)
		}
	}
}
