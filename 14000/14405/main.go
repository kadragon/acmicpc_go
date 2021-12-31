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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var s string
	fmt.Fscan(rd, &s)

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && (s[i:i+2] == "pi" || s[i:i+2] == "ka") {
			i += 1
		} else if i+2 < len(s) && s[i:i+3] == "chu" {
			i += 2
		} else {
			fmt.Fprintf(wr, "%s\n", "NO")
			return
		}
	}

	fmt.Fprintf(wr, "%s\n", "YES")
}
