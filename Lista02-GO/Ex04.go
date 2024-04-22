package main

import "fmt"

func main(){
    var K int
    var n, i, s float64
    
    fmt.Scan(&n, &i, &K, &s)
    
    if n>=0 && n<=9 && K>=0{
        
        fmt.Println("Tabuada de soma:")
        for v:=0; v<K; v++{
            fmt.Printf("%.2f + %.2f = %.2f\n", n, i+s*float64(v), n+i+s*float64(v))}
            
        fmt.Println("Tabuada de subtracao:")
        for v:=0; v<K; v++{
            fmt.Printf("%.2f - %.2f = %.2f\n", n, i+s*float64(v), (n-i+s*float64(v)))}
            
        fmt.Println("Tabuada de multiplicacao:")
        for v:=0; v<K; v++{
            fmt.Printf("%.2f * %.2f = %.2f\n", n, i+s*float64(v), (n*i+s*float64(v)))}
            
        fmt.Println("Tabuada de divisao:")
        for v:=0; v<K; v++{
            fmt.Printf("%.2f / %.2f = %.2f\n", n, i+s*float64(v), n/(i+s*float64(v)))}
            
    }else{
        fmt.Print("Digite valores válidos") 
    }
}