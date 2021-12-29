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
	d := [501][501]byte{}
	for i := 0; i < n*5; i++ {
		for j := 0; j < n*5; j++ {
			if j/n%2 != 1 || (i/n == 0 && j/n == 1) ||
				(i/n == 4 && j/n == 3) {
				d[i][j] = '@'
			} else {
				d[i][j] = ' '
			}
		}
	}

	for i := 0; i < n*5; i++ {
		fmt.Fprintf(wr, "%s\n", string(d[i][:n*5]))
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
