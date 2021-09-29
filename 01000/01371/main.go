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
	d := make([]int, 26)
	var max int
	for {
		s, err := rd.ReadString('\n')
		if err != nil {
			break
		}

		for _, t := range s {
			if 'a' <= t && t <= 'z' {
				d[t-'a']++
			}
		}
	}

	for _, v := range d {
		if v > max {
			max = v
		}
	}

	for i, v := range d {
		if v == max {
			fmt.Fprintf(wr, "%c", i+'a')
		}
	}
}
