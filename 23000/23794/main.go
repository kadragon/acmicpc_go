package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func main() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	n := nextInt()

	d := [105][105]byte{}

	for i := 0; i < n+2; i++ {
		for j := 0; j < n+2; j++ {
			if i == 0 || i == n+1 || j == 0 || j == n+1 {
				d[i][j] = '@'
			} else {
				d[i][j] = ' '
			}
		}
	}

	for i := 0; i < n+2; i++ {
		fmt.Fprintf(wr, "%s\n", string(d[i][:n+2]))
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
