package main

import "fmt"

func main(){
    
    var n, temp int
    var valores[] int
    
    fmt.Scan(&n)
    
    for i:= 0; i<n; i++{
        fmt.Scan(&temp)
        valores = append(valores, temp)
    }
    
    for i:=n-1; i>=0; i--{
        fmt.Printf("%d ", valores[i])
    }
    
}
