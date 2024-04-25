package main

import "fmt"

func main() {
    
    var m, n int
    var linha[] int
    var coluna[] int
    
    fmt.Scan(&m, &n)
    
    for i:=0; i<m; i++{
        
        linha = append(linha, i+1)
        
        for j:=0; j<n; j++{
            
            coluna = append(coluna, j+1)
            
            if j < i {
                
                fmt.Printf("(%d,%d)", linha[i], coluna[j])
                            
            }
            
            if j < i && j+1 != i{
                
                fmt.Printf(" - ")
                
            }
            
        }
                
        fmt.Printf("\n")
        
    }
    
}