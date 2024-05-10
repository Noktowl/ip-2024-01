package main

import "fmt"

func main() {
    
    var temp int; var N int // numero de pontos
    fmt.Scan(&N)
    
    var vetor1[] float64
    
    for i:=0; i<3; i++{ // le apenas o primeiro vetor
        fmt.Scan(&temp)
        vetor1 = append(vetor1, float64(temp))
    }
    
    for i:=1; i<N; i++{ // repete para a quantidade de vetores
    
        var vetor2[] float64 //define um vetor2 vazio a cada iteração
        soma:= 0.0 // define a soma como 0 a cada iteração
        
        for j:=0; j<3; j++{ // le todo 2° vetor
            fmt.Scan(&temp)
            vetor2 = append(vetor2, float64(temp))
        }
        
        for k:=0; k<3; k++{ //soma os quadrados da diferença de cada entrada dos pontos
            soma = soma + (vetor1[k]-vetor2[k])*(vetor1[k]-vetor2[k])
        }
        
        distancia:= raiz(soma)
        fmt.Printf("%.2f\n", distancia)
        
        for v:=0; v<3; v++{ // reseta o vetor1 para os valores do antigo vetor2
            vetor1[v] = vetor2[v]
        }
    }
}

func raiz(n float64) float64{
    
    raiz:= 1.0
    e:= 0.000000001
    
    erro:= n-raiz*raiz
    
    for erro > e{
        
        erro = n-raiz*raiz
        raiz = (raiz + (n/raiz))/2
        
        if erro < 0{
            
            erro = -erro
            
        }
    }

    return raiz
}

