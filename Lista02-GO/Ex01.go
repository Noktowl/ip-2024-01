//Exercicio 01

package main

import "fmt"

func main() {

	var (
		m                              int
		n1, n2, n3, n4, n5, n6, n7, n8 float32
		l1, l2, l3, l4, l5             float32
		ntf                            float32

		mp, ml, ntd, presenca float32
	)

	fmt.Scanf("%d", &m)
	fmt.Scanf("%f %f %f %f %f %f %f %f", &n1, &n2, &n3, &n4, &n5, &n6, &n7, &n8)
	fmt.Scanf("%f %f %f %f %f", &l1, &l2, &l3, &l4, &l5)
	fmt.Scanf("%f", &ntf)
	fmt.Scanf("%f", &presenca)

	if m < 0 {
		fmt.Printf("Programa finalizado")
	} else {
		mp = (n1 + n2 + n3 + n4 + n5 + n6 + n7 + n8) / 8
		ml = (l1 + l2 + l3 + l4 + l5) / 5
		ntd = 0.7*mp + 0.15*ml + 0.15*ntf

		if ntd < 6 {
			if presenca < 128/100*75 {
				fmt.Printf("Matricula: %d, Nota Final: %f, Situação Final: REPROVADO POR NOTA E POR FREQUENCIA", m, ntd)
			} else {
				fmt.Printf("Matricula: %d, Nota Final: %f, Situação Final: REPROVADO POR NOTA", m, ntd)

			}
		} else {
			if presenca < 128/100*75 {
				fmt.Printf("Matricula: %d, Nota Final: %f, Situação Final: REPROVADO POR FREQUENCIA", m, ntd)
			} else {
				fmt.Printf("Matricula: %d, Nota Final: %f, Situação Final: APROVADO", m, ntd)

			}
		}
	}
}
