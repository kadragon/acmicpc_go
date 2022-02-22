package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	s = strings.ReplaceAll(s, "XXXX", "AAAA")
	s = strings.ReplaceAll(s, "XX", "BB")

	if strings.Count(s, "X") > 0 {
		fmt.Fprintf(wr, "-1")
	} else {
		fmt.Fprintf(wr, "%s\n", s)
	}
}
