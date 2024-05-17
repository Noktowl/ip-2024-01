package main

import "fmt"

func main() {

	var N, K int
	var posicao []int
	qpresentes := 0

	fmt.Scanf("%v %v", &N, &K)

	chegada := lervetor(N)

	for i := 0; i < N; i++ {
		if chegada[i] <= 0 {
			qpresentes++
			posicao = append(posicao, i+1)
		}
	}

	if qpresentes < K {
		fmt.Print("SIM\n")
	} else {
		fmt.Print("NAO\n")
		for i := qpresentes - 1; i >= 0; i-- {
			fmt.Print(posicao[i], "\n")
		}
	}

}

func lervetor(tamanho int) []int {

	var temp int
	var vetor []int

	for i := 0; i < tamanho; i++ {
		fmt.Scan(&temp)
		vetor = append(vetor, temp)
	}

	return vetor
}
