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

	fmt.Fprintf(wr, "%d\n", n)
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
