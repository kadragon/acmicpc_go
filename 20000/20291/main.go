package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func main() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	n := nextInt()
	d := map[string]int{}

	for i := 0; i < n; i++ {
		t := nextString()
		d[strings.Split(t, ".")[1]]++
	}

	extList := []string{}
	for key := range d {
		extList = append(extList, key)
	}

	sort.Strings(extList)

	for _, ext := range extList {
		fmt.Fprintf(wr, "%s %d\n", ext, d[ext])
	}
}

// FastIO
func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}

func nextString() string {
	sc.Scan()

	return sc.Text()
}
