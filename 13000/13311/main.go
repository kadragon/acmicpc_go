package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wr *bufio.Writer
)

func init() {
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	fmt.Fprintf(wr, "%d", -1)
}
