package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var stdio = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

func main() {
	defer stdio.Flush()
	for {
		var a, b, c int64
		fmt.Fscan(stdio, &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}

		arr := []int64{a, b, c}
		sort.Slice(arr, func(idx1, idx2 int) bool {
			return arr[idx1] < arr[idx2]
		})

		squaredA := arr[0] * arr[0]
		squaredB := arr[1] * arr[1]
		squaredC := arr[2] * arr[2]

		if squaredA+squaredB == squaredC {
			fmt.Println("right")
			continue
		}
		fmt.Println("wrong")
	}
}
