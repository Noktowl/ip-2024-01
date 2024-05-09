package main

import "fmt"

func main() {
    
    var N, K, temporario int
    var vetor[] int
    
    for{
        fmt.Scan(&N)
        if N>0 && N<=1000{
            break
        }
    }
    
    for i:=0; i<N; i++{
        fmt.Scan(&temporario)
        vetor = append(vetor, temporario)
    }
    
    fmt.Scan(&K)
    
    v:= 0 //verificador
    
    for j:=0; j<N; j++{
        if vetor[j] >= K{
            v++
        }
    }
    
    fmt.Printf("\n%d", v)
    
}
