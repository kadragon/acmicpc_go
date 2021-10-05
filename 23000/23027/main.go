package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var s string
	fmt.Fscan(rd, &s)

	c := make([]int, 26)
	for _, v := range s {
		c[v-'A']++
	}

	conv := []string{"B", "C", "D", "F"}

	if c[0] > 0 {
		for _, w := range conv {
			s = strings.Replace(s, w, "A", -1)
		}
	} else if c[1] > 0 {
		for _, w := range conv[1:] {
			s = strings.Replace(s, w, "B", -1)
		}
	} else if c[2] > 0 {
		for _, w := range conv[2:] {
			s = strings.Replace(s, w, "C", -1)
		}
	} else {
		for i := 0; i < len(s); i++ {
			fmt.Fprintf(wr, "A")
		}
		return
	}
	
	fmt.Fprintf(wr, "%s\n", s)
}
