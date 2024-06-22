package main

import (
	"fmt"
)

func main() {
	n := 10
	workArray := make([]uint8, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&workArray[i])
	}
	counter := 0
	for counter < 3 {
		var a, b uint8
		_, _ = fmt.Scan(&a, &b)
		temp := workArray[a]
		workArray[a] = workArray[b]
		workArray[b] = temp
		counter++
	}
	for k := 0; k < n; k++ {
		fmt.Print(workArray[k], " ")
	}
}
