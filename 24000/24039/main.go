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

	primeNuber := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

	var n int
	fmt.Fscan(rd, &n)

	for i := range primeNuber {
		tmp := primeNuber[i] * primeNuber[i+1]
		if tmp > n {
			fmt.Fprintf(wr, "%d\n", tmp)
			break
		}
	}
}
