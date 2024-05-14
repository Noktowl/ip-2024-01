package main

import "fmt"

func main() {
    
    var N, temp int
    var notas[] int
    var maiornota, posicao, v int
    
    fmt.Scan(&N)
    
    for i:=0; i<N; i++ { // le as notas e ao mesmo tempo verifica a maior e sua posicao
        fmt.Scan(&temp)
        notas = append(notas, temp)
        if notas[i] > maiornota{
            maiornota = notas[i]
            posicao = i
        }
    }
    
    for i:=N-1; i>=0; i--{
        if notas[N-1] == notas[i]{
            v++ // quantidade de Vezes que a ultima nota aparece
        }
    }
    
    fmt.Printf("Nota %d, %d vezes \nNota %d, indice %d \n", notas[N-1], v, maiornota, posicao)
}