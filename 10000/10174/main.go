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
		fmt.Fprintln(wr, solve())
	}
}

func solve() string {
	s, _, _ := rd.ReadLine()

	l := len(s)
	for i, v := range s {
		if 'a' <= v && v <= 'z' {
			s[i] = v - 32
		}
	}

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return "No"
		}
	}
	return "Yes"
}
