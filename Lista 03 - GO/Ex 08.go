package main

import "fmt"

func main(){
    
    var N, Npordois, tamanhodobinario int
    
    for{
        
        var binario[] int // define o vetor que armazena os algarismos do binario e o reseta a cada loop
        
        tamanhodobinario = 0
        fmt.Scan(&N)
        Npordois = N
        
        for Npordois != 1 && Npordois != 0{ // fica em loop até N/2 ser igual a 1 ou 0
            binario = append(binario, Npordois%2) // aloca o resto da divisao por 2 no vetor
            Npordois = Npordois/2 // divide resultado da divisão por 2, preparando pra proxima iteração
            tamanhodobinario++
        }
        
        if Npordois == 0 {
            fmt.Print(Npordois)
        }
        
        if Npordois == 1 {
            fmt.Print(Npordois)
        }
        
        for i := tamanhodobinario-1; i>=0; i--{ // imprime os valores do vetor na ordem inversa
            fmt.Print(binario[i])
        }
        
        fmt.Print("\n") // quebra de linha para proxima iteração
    }
}