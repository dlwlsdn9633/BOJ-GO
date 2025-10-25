package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdio = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

func main() {
	defer stdio.Flush()
	var a, b, c int64
	fmt.Fscan(stdio, &a, &b, &c)
	ret := recursive(a, b, c)
	fmt.Println(ret)
}

func recursive(a, b, c int64) (ret int64) {
	if b == 1 {
		return a % c
	}
	ret = recursive(a, b/2, c)
	ret = (ret * ret) % c
	if b%2 != 0 {
		ret = (ret * a) % c
	}
	return
}
