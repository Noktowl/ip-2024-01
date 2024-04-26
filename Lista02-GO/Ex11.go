package main

import "fmt"

func main() {
    
    var n, e float64
    raiz:= 1.0
    
    fmt.Scan(&n, &e)
    
    erro:= n-raiz*raiz
    
    for erro > e{
        
        erro = n-raiz*raiz
        raiz = (raiz + (n/raiz))/2
        
        if erro < 0{
            
            erro = -erro
            
        }
        
        if erro != 1 {
            fmt.Printf("r: %.9f, erro: %.9f\n", raiz, erro)
        }
    }
    
}
