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

	var t int
	fmt.Fscanf(rd, "%d\n", &t)

	for ; t > 0; t-- {
		var p int
		var name, kind string
		fmt.Fscanf(rd, "%d\n", &p)

		d := map[string]int{}

		for ; p > 0; p-- {
			fmt.Fscanf(rd, "%s %s\n", &name, &kind)
			d[kind]++
		}

		ans := 1
		for _, value := range d {
			ans *= value + 1
		}

		fmt.Fprintf(wr, "%d\n", ans-1)
	}
}
