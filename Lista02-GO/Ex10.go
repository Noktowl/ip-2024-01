package main

import "fmt"

func main() {

	var (
		matricula [99999]int
		horas     [99999]float64
		valorH    [99999]float64
		salario   [99999]float64
		v         = 0
	)

	for i := 0; i >= 0; i++ {
		fmt.Scanf("%d %f %f", &matricula[i], &horas[i], &valorH[i])

		if matricula[i] != 0 {
			salario[i] = horas[i] * valorH[i]
			v++
		} else {
			for j := 0; j < v; j++ {
				fmt.Printf("%d %.2f\n", matricula[j], salario[j])
			}
		}
	}
}