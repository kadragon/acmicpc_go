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
	fmt.Fscanf(rd, "%d\n", &n)

	for ; n > 0; n-- {
		s, _, _ := rd.ReadLine()
		if string(s[:10]) == "Simon says" {
			fmt.Fprintf(wr, "%s\n", string(s[10:]))
		}
	}
}
