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

	p := [26][14]int{}
	q := [26]int{}

	for i := 0; i < 26; i++ {
		q[i] = 13
	}

	var c rune
	var d int

	for {
		fmt.Fscanf(rd, "%c%d", &c, &d)
		if c == '\n' {
			break
		}

		p[c-'A'][d]++
		q[c-'A']--

		if p[c-'A'][d] > 1 {
			fmt.Fprintln(wr, "GRESKA")
			return
		}
	}

	for _, v := range "PKHT" {
		fmt.Fprintf(wr, "%d ", q[v-'A'])
	}
}
