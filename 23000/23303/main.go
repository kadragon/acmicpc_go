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

	s, _ := rd.ReadString('\n')
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] == '2' && (s[i] == 'd' || s[i] == 'D') {
			fmt.Fprintf(wr, "D2\n")
			return
		}
	}

	fmt.Fprintf(wr, "unrated\n")
}
