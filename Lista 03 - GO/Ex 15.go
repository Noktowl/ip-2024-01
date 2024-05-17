package main

import "fmt"

func main() {

	var T int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N int

		fmt.Scan(&N)

		vetor := lervetor(N)
		menordist := menordistancia(N, vetor)
		ntestes := testesnecessarios(N)

		fmt.Printf("%v %v\n", menordist, ntestes)
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

func menordistancia(tamanhovetor int, vetor []int) int {

	menordist := modulo(vetor[0] - vetor[1])

	for i := 0; i < tamanhovetor; i++ {
		for j := i + 1; j < tamanhovetor; j++ {
			distancia := vetor[i] - vetor[j]
			distancia = modulo(distancia)
			if distancia < menordist {
				menordist = distancia
			}
		}
	}

	return menordist
}

func testesnecessarios(N int) int {

	a1 := N - 1
	testes := (a1 * N) / 2

	return testes
}

func modulo(valor int) int {
	if valor < 0 {
		valor = -valor
	}
	return valor
}
