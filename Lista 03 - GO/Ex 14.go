package main
import "fmt"

func main(){
    var N int
    var mediana float64
    
    fmt.Scan(&N)
    
    lista:= lervetor(N)
    lista = ordenar(N, lista)

    
    if N%2 == 0{
        mediana = (lista[N/2 - 1] + lista[N/2])/2.0
    }else{
        mediana = lista[N/2]
    }
    
    fmt.Printf("%.2f", mediana)
}

func ordenar(tamanho int, vetor[] float64) []float64 {
    
    var temp float64
    
    for i:=0; i<tamanho; i++{
        for j:=0; j<tamanho; j++{
            if j>i{
                if vetor[i]>vetor[j]{
                    temp = vetor[i]
                    vetor[i] = vetor[j]
                    vetor[j] = temp
                }
            }
        }
    }
    return vetor
}

func lervetor(tamanho int) [] float64 {
    
    var temp float64
    var lista[] float64
    
        for i:=0; i<tamanho; i++{
        fmt.Scan(&temp)
        lista = append(lista, temp)
    }
    return lista
}