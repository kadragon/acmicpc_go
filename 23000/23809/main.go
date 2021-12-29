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

	at := strings.Repeat("@", n)
	bl := strings.Repeat(" ", n)

	for i := 0; i < 5; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == 4 {
				fmt.Fprintf(wr, "%s%s%s%s%s\n", at, bl, bl, bl, at)
			} else if i == 1 || i == 3 {
				fmt.Fprintf(wr, "%s%s%s%s\n", at, bl, bl, at)
			} else {
				fmt.Fprintf(wr, "%s%s%s\n", at, at, at)
			}
		}
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
