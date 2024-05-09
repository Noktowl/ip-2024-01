package main

import "fmt"

func main() {
    
    var n, temporario int
    var vetor[] int
    
    fmt.Scan(&n)
    
    for i:=0; i<n; i++{
        fmt.Scan(&temporario)
        vetor = append(vetor, temporario)
    }
    
    for i:=0; i<n; i++{
        fmt.Print(vetor[i]," ")
    }
    
}
