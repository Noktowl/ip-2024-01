package main

import "fmt"

func main() {
    
    var n, temp int; soma:= 0
    var vetor[] int
    
    fmt.Scan(&n)
    
    for i:=0; i<n; i++{
        
        fmt.Scan(&temp)
        vetor = append(vetor, temp)
        soma += vetor[i]
        
    }
    
    fmt.Printf("%v", soma)
    
}
