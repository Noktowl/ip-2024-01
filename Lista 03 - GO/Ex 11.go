package main

import "fmt"

func main(){
    
    var N, temp int
    var V[] int
    
    fmt.Scan(&N)
    
    for i:=0; i<N; i++{
        fmt.Scan(&temp)
        V = append(V, temp)
    }
    
    W:= inverte(N, V)
    maiorelementoV:= maiorelemento(N, V)
    menorelementoW:= menorelemento(N, W)
    
    for i:=0; i<N; i++{
        fmt.Printf("%d ", V[i])
    }
    fmt.Println()
    for i:=0; i<N; i++{
        fmt.Printf("%d ", W[i])
    }
    fmt.Print("\n", maiorelementoV)
    fmt.Print("\n", menorelementoW)
}

func inverte(tamanho int, vetor[] int) []int{

    var W[] int
    
    for i:=tamanho-1; i>=0; i--{
        W = append(W, vetor[i])
    }
    return W
}

func maiorelemento(tamanhovetor int, vetor[] int) int{
    
    maior:= vetor[0]
    
    for i:=1; i<tamanhovetor; i++{
        if vetor[i] > maior{
            maior = vetor[i]
        }
    }
    return maior
}

func menorelemento(tamanhovetor int, vetor[] int) int{
    
    menor:= vetor[0]
    
    for i:=1; i<tamanhovetor; i++{
        if vetor[i] < menor{
            menor = vetor[i]
        }
    }
    return menor
}