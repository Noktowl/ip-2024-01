package main

import "fmt"

func main() {
	var N int
	v := 1

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			fmt.Printf("final %d: time%d x time%d\n", v, i+1, j+1)
			v++
		}
	}
}