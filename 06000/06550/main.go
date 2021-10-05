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
		s, err := rd.ReadString(' ')
		if err != nil {
			break
		}
		t, _ := rd.ReadString('\n')

		var sc int

		for i := 0; i < len(t) && sc < len(s); i++ {
			if t[i] == s[sc] {
				sc++
			}
		}

		if sc == len(s)-1 {
			fmt.Fprintf(wr, "Yes\n")
		} else {
			fmt.Fprintf(wr, "No\n")
		}
	}
}
