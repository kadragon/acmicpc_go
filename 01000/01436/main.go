package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	wr *bufio.Writer
)

func init() {
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	p := []string{"Yonsei", "Korea"}
	fmt.Fprintf(wr, "%s\n", p[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)])
}
