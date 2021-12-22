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

	fmt.Fprintf(wr, "%d\n", n)
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
