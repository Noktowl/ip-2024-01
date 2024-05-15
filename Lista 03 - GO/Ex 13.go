package main
import "fmt"

func main(){
    var N, temp, maisfrequente int
    var lista[] int
    maxfreq:= 0
    
    fmt.Scan(&N)
    
    for i:=0; i<N; i++{
        fmt.Scan(&temp)
        lista = append(lista, temp)
    }
    
    for i:=0; i<N; i++{
        frequencia:=0
        for j:=0; j<N; j++{
            if lista[i] == lista[j]{
                frequencia++
            }
            if frequencia > maxfreq{
                maxfreq = frequencia
                maisfrequente = lista[i]
            }
        }
    }
    
    fmt.Printf("%d\n%d", maisfrequente, maxfreq)
    
}