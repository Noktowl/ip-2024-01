//Exercicio 03

/*OBS: o programa deixa de funcionar em valores altos,
devito a falta de capacidade da variavel do tipo "int" 
expressar valores desta magnitude;*/

package main

import "fmt"

func main() {
	var n, x int
	h := 1

	fmt.Scan(&n)

	x = n

	for i := 1; i < n; i++ {
		x = x * (n - h)
		h++
	}
	fmt.Printf("%d! = %d", n, x)
}
