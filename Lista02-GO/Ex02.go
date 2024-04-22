//exercico 02

package main

import "fmt"

func main() {
	var A, B float64
	fmt.Scan(&A, &B)

	if A > 0 && B > 0 {
		if A >= B {
			fmt.Printf("Valor Inválido")
		}
		var anos int
		for A <= B {
			A = A * 1.03
			B = B * 1.015

			anos++
		}
			fmt.Printf("ANOS = %d", anos)
	} else {
		fmt.Printf("Valor Inválido")
	}
}

