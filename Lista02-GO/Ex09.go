package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N > 0 {
		v := 0
		for i := 1; i <= N; i++ {
			if N%i == 0 {
				v++
			}
		}
		if v == 2 {
			fmt.Print("PRIMO")
		} else {
			fmt.Print("NAO PRIMO")
		}
	} else {
		fmt.Print("INSIRA UM VALOR VÁLIDO!")
	}
}