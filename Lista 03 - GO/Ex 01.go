package main

import "fmt"

func main() {

	var N, M, temporario int
	var vetor []int
	var vetorM []int

	fmt.Scan(&N)

	for i := 0; i < N; i++ { // le todos os valores do vetor
		fmt.Scan(&temporario)
		vetor = append(vetor, temporario)
	}

	fmt.Scan(&M)

	for i := 0; i < M; i++ { // le todos os valores a serem verificados se sao elementos do vetor
		fmt.Scan(&temporario)
		vetorM = append(vetorM, temporario)
	}

	for i := 0; i < M; i++ { //verifica todos os M elementos

		v := 0 //verificador

		for k := 0; k < N; k++ { //compara cada um dos M valores com cada elemento do vetor
			if vetorM[i] == vetor[k] {
				fmt.Print("ACHEI\n")
				v++
			}
		}

		if v == 0 {
			fmt.Print("NAO ACHEI\n")
		}
	}
}

