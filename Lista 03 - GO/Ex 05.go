package main

import "fmt"

func main() {

    var N, temp int
	var (
		posicao    []int
		maiorvalor []int
	)
	rodada := 0

	for {
		fmt.Scan(&N)

		if N != 0 {
            
            maiorvalortemp := 0
	        posicaotemp := 0
            
            var temporario []int
            
			for i := 0; i < N; i++ {

				fmt.Scan(&temp)
				temporario = append(temporario, temp)

				if temp >= maiorvalortemp {
					maiorvalortemp = temp
					posicaotemp = i
				}
			}

			maiorvalor = append(maiorvalor, maiorvalortemp)
			posicao = append(posicao, posicaotemp)

		} else {
		    break
		}
		rodada++
	}

	for j := 0; j < rodada; j++ {
		fmt.Printf("\n%d %d", posicao[j], maiorvalor[j])
	}
}
