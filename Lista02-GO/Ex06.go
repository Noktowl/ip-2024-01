//Exercicio 06

package main    
import "fmt"

func main(){
    var n, armazem int
    var elementos[99999] int
    
    fmt.Scan(&n)
    
    if n>0{
        for i:=0; i<n; i++{
            fmt.Scan(&elementos[i])
        }
        v:= 1
        for i:=0; i<n; i++{
            if elementos[i] == elementos[i+1]{
                v++
                if v>armazem{
                    armazem = v
                }
            }else{
                v = 1
            }
        }
    }else{
        fmt.Println("INSIRA VALORES VÁLIDOS")
    }
    fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.", armazem)
}