package main

import "fmt"

func main() {

	var T int
	var mensagem []string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {

		cpf := lervetor(11)
		b1 := digitob1(cpf)
		b2 := digitob2(cpf)

		if cpf[9] == b1 && cpf[10] == b2 {
			mensagem = append(mensagem, "CPF valido")
		} else {
			mensagem = append(mensagem, "CPF invalido")
		}
	}

	for i := 0; i < T; i++ {
		fmt.Println(mensagem[i])
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

func digitob1(cpf []int) int {

	soma := 0

	for i := 0; i < 9; i++ {
		soma = soma + (cpf[i] * (i + 1))
	}

	b1 := soma % 11

	if b1 == 10 {
		b1 = 0
	}

	return b1
}

func digitob2(cpf []int) int {

	soma := 0
	multiplicador := 9

	for i := 0; i < 9; i++ {
		soma = soma + (cpf[i] * multiplicador)
		multiplicador--
	}

	b2 := soma % 11

	if b2 == 10 {
		b2 = 0
	}

	return b2
}
