package main

import "fmt"

func main() {
    
    var N int; var temp float64
    fmt.Scan(&N)
    
    var vetor1[] float64
    for i:=0; i<3; i++{ // le o vetor 1
        fmt.Scan(&temp)
        vetor1 = append(vetor1, temp)
    }
    
    for i:=0; i<N; i++{
        
        var vetor2[] float64
        var coordenadas[] float64
        
        for j:=0; j<3; j++{ // le o vetor 2
            fmt.Scan(&temp)
            vetor2 = append(vetor2, temp)
        }
        
        for j:=0; j<3; j++{
            coordenadas =append(coordenadas, vetor1[j] - vetor2[j])
        }
        
        maiorj:= 0
        maiorcoord:= 0.0
        
        for j:=0; j<3; j++{
            if coordenadas[j] < 0{
                coordenadas[j] = -coordenadas[j]
            }
            if coordenadas[j] >= coordenadas[maiorj]{
                maiorcoord = coordenadas[j]
                maiorj = j
            }
        }
        
        fmt.Printf("%.2f\n", maiorcoord)
        
        for j:=0; j<3; j++{ // redefine o vetor 1 com os valores do vetor 2
            vetor1[j] = vetor2[j]
        }
    }
}