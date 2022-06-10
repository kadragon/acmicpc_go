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
	s := []string{"ITMO", "PetrSU, ITMO", "SPbSU"}

	var n, ans int
	fmt.Fscan(rd, &n)

	switch n {
	case 2006:
		ans = 1
	case 1996, 1997, 2000, 2007, 2008, 2013, 2018:
		ans = 2
	}

	fmt.Fprintf(wr, "%s\n", s[ans])
}
