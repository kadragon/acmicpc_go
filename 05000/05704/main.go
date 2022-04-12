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

	for {
		s, _, _ := rd.ReadLine()
		if s[0] == '*' {
			break
		}

		c := make([]int, 26)
		for _, v := range s {
			if v == ' ' {
				continue
			}
			c[v-'a']++
		}

		checked := false
		for i := 0; i < 26; i++ {
			if c[i] == 0 {
				checked = true
				break
			}
		}

		if checked {
			fmt.Fprintf(wr, "N\n")
		} else {
			fmt.Fprintf(wr, "Y\n")
		}
	}
}
