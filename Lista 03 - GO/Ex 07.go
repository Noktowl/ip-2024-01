package main

import "fmt"

func main() {
    
    for{
        
        var tamanho, temp, maiorelemento int
        var vetor[] int
        posicaodomaior := 0;
        
        fmt.Scan(&tamanho)
        
        if tamanho != 0{
            
            for i:=0; i<tamanho; i++{
                fmt.Scan(&temp)
                vetor = append(vetor, temp)
                if vetor[i] >= vetor[posicaodomaior]{
                    maiorelemento = vetor[i]
                    posicaodomaior = i
                }
            }
            
            for i:=0; i<=maiorelemento; i++{
                qdelementos := 0
                for j:=0; j<tamanho; j++{
                    if vetor[j] <= i {
                        qdelementos ++
                    }
                }
                fmt.Printf("(%d) %d\n", i, qdelementos)
            }
            
        }else{break}
    }
}