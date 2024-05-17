package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	vetor := lervetor(n)
	unicos := elementosunicos(n, vetor)

	fmt.Print(unicos)

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

func elementosunicos(tamanhovetor int, vetor []int) int {

	qunicos := 0

	for i := 0; i < tamanhovetor; i++ {
		v := 0 // contador
		for j := 0; j < tamanhovetor; j++ {
			if vetor[i] == vetor[j] {
				v++
			}
		}
		if v == 1 {
			qunicos++
		}
	}
	return qunicos
}
