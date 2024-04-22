//Exercicio 03

package main

import "fmt"

func main() {
	var n int
	var sequencia [9999999]int
	armazem := 0

	fmt.Scan(&n)

	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Scan(&sequencia[i])
		}
		v := 0
		for i := 0; i < n; i++ {

			if sequencia[i+1] > sequencia[i] {
				v++
				if v > armazem {
					armazem = v

				}

			} else {
				v = 0
			}
		}
		fmt.Print("O comprimento do segmento crescente maximo e: ", armazem)
	} else {
		fmt.Print("DIGITE UM NÚMERO VÁLIDO")
	}

}
