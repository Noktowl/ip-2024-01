package main

import "fmt"

func main(){
    var valores[10] float64
    
    for i:= 0; i<10; i++{
        fmt.Scan(&valores[i])
    }
    
    for i:=9; i>=0; i--{
        fmt.Printf(" %.2f", valores[i])
    }
    
}