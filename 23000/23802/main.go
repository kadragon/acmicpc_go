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

	d := strings.Repeat("@", n*5)

	for i := 0; i < n*5; i++ {
		if i < n {
			fmt.Fprintf(wr, "%s\n", d)
		} else {
			fmt.Fprintf(wr, "%s\n", d[:n])
		}
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
