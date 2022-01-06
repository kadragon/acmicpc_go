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

	var s, t string
	fmt.Fscan(rd, &s, &t)

	p := make([]byte, len(s))

	q, lent := 0, len(t)
	for i := 0; i < len(s); i++ {
		p[q] = s[i]
		q++

		if q >= lent && string(p[q-lent:q]) == t {
			q -= lent
		}
	}

	if q == 0 {
		fmt.Fprintf(wr, "FRULA\n")
	} else {
		fmt.Fprintf(wr, "%s\n", string(p[:q]))
	}
}
