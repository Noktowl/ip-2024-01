package main

import "fmt"

func main() {
    
    var T, A, B int
    
    fmt.Scan(&T)
    
    for i := 0; i < T; i++ {
        
        fmt.Scan(&A, &B)

        A = (A % 10 * 100) + (A % 100 / 10 * 10) + (A / 100)
        B = (B % 10 * 100) + (B % 100 / 10 * 10) + (B / 100)

        if A > B {
            
            fmt.Printf("%03d\n", A)
            
        } else {
            
            fmt.Printf("%03d\n", B)
            
        }
        
    }
    
}
