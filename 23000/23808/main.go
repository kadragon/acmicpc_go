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

	d := [510][510]byte{}
	for i := 0; i < n*5; i++ {
		for j := 0; j < n*5; j++ {
			if j < n || j > n*5-(n+1) {
				d[i][j] = '@'
			} else if i/n == 2 || i/n == 4 {
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
