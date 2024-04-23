package main

import "fmt"

func main() {
	var sequencia [99999]int
	par := 0
	impar := 0
	tpar := 0
	timpar := 0

	for i := 0; i >= 0; i++ {
		fmt.Scan(&sequencia[i])
		if sequencia[i] != 0 {
			if sequencia[i]%2 == 0 {
				par++
				tpar = tpar + sequencia[i]
			} else {
				impar++
				timpar = timpar + sequencia[i]
			}
		} else {
			fmt.Printf("MEDIA PAR: %.2f \nMEDIA IMPAR: %.2f", float64(tpar)/float64(par), float64(timpar)/float64(impar))
		}
	}
}