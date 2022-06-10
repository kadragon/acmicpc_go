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

	a := [2][9]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(rd, &a[i][j])
		}
	}

	s := [2]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 2; j++ {
			s[j] += a[j][i]
			if s[0] > s[1] {
				fmt.Fprintf(wr, "Yes")
				return
			}
		}

	}

	fmt.Fprintf(wr, "No\n")
}
