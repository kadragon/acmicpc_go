package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	at := strings.Repeat("@", n*5)
	bt := strings.Repeat("@", n) + strings.Repeat(" ", n*3) + strings.Repeat("@", n)

	for i := 0; i < n*5; i++ {
		if i/n == 1 || i/n == 3 {
			fmt.Fprintf(wr, "%s\n", at)
		} else {
			fmt.Fprintf(wr, "%s\n", bt)
		}
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
