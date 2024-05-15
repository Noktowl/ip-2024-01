package main
import "fmt"

func main() {
    var N, temp int
    var lista [] int
    
    fmt.Scan(&N)
    
    for i:=0; i<N; i++{
        fmt.Scan(&temp)
        lista = append(lista, temp)
    }
    
    for i:=0; i<N; i++{
        for j:=0; j<N; j++{
            if j>i{
                if lista[i]>lista[j]{
                    temp = lista[i]
                    lista[i] = lista[j]
                    lista[j] = temp
                }
            }
        }
    }
    
    for i:=0; i<N; i++{
        fmt.Printf("%d\n", lista[i])
    }

}