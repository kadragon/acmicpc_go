package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()

	n := nextInt()
	d := [2000005]int{}

	for i := 0; i < n; i++ {
		d[nextInt()+1000000]++
	}

	for i, v := range d {
		for ; v > 0; v-- {
			fmt.Fprintf(wr, "%d\n", i-1000000)
		}
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
