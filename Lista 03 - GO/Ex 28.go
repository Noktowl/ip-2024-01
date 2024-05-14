package main
import "fmt"

func main(){

//  ler o tamanho dos conjuntos
    tA:= tamanhovalido()
    tB:= tamanhovalido()
    
//  le os valores dos conjuntos sem repetir elementos dentro do conjunto
    A:= conjuntosemrepeticao(tA)
    B:= conjuntosemrepeticao(tB)
    
//  faz a uniao e a interseção
    uniao:= uniao(A, B, tA, tB)
    intersecao:= intersecao(A, B, tA, tB)
    
//  imprimir a uniao
    tuniao:= tamanhouniao(A, B, tA, tB)
    
    if tuniao > 1{
        fmt.Printf("(%d, ", uniao[0])
        for i:=1; i<tuniao-1; i++{
            fmt.Printf("%d, ", uniao[i])
        }
        fmt.Printf("%d)\n", uniao[tuniao-1])
    }else{
        fmt.Printf("(%d)", intersecao[0])
    }
    
//  imprimir a interseção
    tintersecao:= tamanhointersecao(A, B, tA, tB)
    
    if tintersecao > 1{
        fmt.Printf("(%d, ", intersecao[0])
        for i:=1; i<tintersecao-1; i++{
            fmt.Printf("%d, ", intersecao[i])
        }
        fmt.Printf("%d)\n", intersecao[tintersecao-1])
    }
    if tintersecao == 1{
        fmt.Printf("(%d)", intersecao[0])
    }
    if tintersecao == 0{
        fmt.Print("()")
    }
}

func tamanhovalido() int{
    
    var tamanho int
    
    for{
        fmt.Scan(&tamanho)
        if tamanho>=1 && tamanho<=100{
            break
        }
    }
    
    return tamanho
}

func conjuntosemrepeticao(tamanhoconjunto int) [] int{
    
    var temp int
    var vetor[] int
    
    for i:=0; i<tamanhoconjunto; i++{
        v:=-5 // verifica se um numero foi repetido
        
        fmt.Scan(&temp) // le o elemento
        
        for j:=0; j<i; j++{ // ve se o elemento já está o conjunto
            if temp == vetor[j]{ // se o elemento está, ele faz v ter o valor de i
                v = i
            }
        }
        
        if v != i{ // se o elemento não tem valor de i, ele não esta no conjunto e pode ser agregado
            vetor = append(vetor, temp)
        }else{// se tem valor de i, entao retorna i para o valor anterior, para refazer a atual fase
            i = i-1
        }
    }
    return vetor
}

func uniao(vetor1, vetor2 []int, tamanhov1, tamanhov2 int) []int {
    
    tamanhodovetor:= tamanhov1
    
    for i:=0; i<tamanhov2; i++{
        
        v:=0 // verificador; reseta a cada iteração
        
        for j:=0; j<tamanhov1; j++{ // ve se o elemento já está o conjunto
            if vetor2[i] == vetor1[j]{ // se o elemento está, ele faz v somar +1
                v++
            }
        }
        
        if v == 0{ // se o elemento não tem valor de i, ele não esta no conjunto e pode ser agregado
            vetor1 = append(vetor1, vetor2[i])
            tamanhodovetor++
        }
        
    }
    return vetor1
}

func intersecao(vetor1, vetor2 []int, tamanhov1, tamanhov2 int) []int{
    
    var Cintersecao [] int
    tamanhodoconjunto:= 0
    
    for i:=0; i<tamanhov2; i++{
        
        v:=0 // verificador; reseta a cada iteração
        
        for j:=0; j<tamanhov1; j++{ // ve se o elemento já está o conjunto
            if vetor2[i] == vetor1[j]{ // se o elemento está, ele faz v somar +1
                v++
            }
        }
        
        if v != 0{ // se o elemento não tem valor de i, ele não esta no conjunto e pode ser agregado
            Cintersecao = append(Cintersecao, vetor2[i])
            tamanhodoconjunto++
        }
        
    }
    return Cintersecao
}

func tamanhouniao(vetor1, vetor2 []int, tamanhov1, tamanhov2 int) int {
    
    tamanhodovetor:= tamanhov1
    
    for i:=0; i<tamanhov2; i++{
        
        v:=0 // verificador; reseta a cada iteração
        
        for j:=0; j<tamanhov1; j++{ // ve se o elemento já está o conjunto
            if vetor2[i] == vetor1[j]{ // se o elemento está, ele faz v somar +1
                v++
            }
        }
        
        if v == 0{ // se o elemento não tem valor de i, ele não esta no conjunto e pode ser agregado
            vetor1 = append(vetor1, vetor2[i])
            tamanhodovetor++
        }
        
    }
    return tamanhodovetor
}

func tamanhointersecao(vetor1, vetor2 []int, tamanhov1, tamanhov2 int) int{
    
    var Cintersecao [] int
    tamanhodoconjunto:= 0
    
    for i:=0; i<tamanhov2; i++{
        
        v:=0 // verificador; reseta a cada iteração
        
        for j:=0; j<tamanhov1; j++{ // ve se o elemento já está o conjunto
            if vetor2[i] == vetor1[j]{ // se o elemento está, ele faz v somar +1
                v++
            }
        }
        
        if v != 0{ // se o elemento não tem valor de i, ele não esta no conjunto e pode ser agregado
            Cintersecao = append(Cintersecao, vetor2[i])
            tamanhodoconjunto++
        }
        
    }
    return tamanhodoconjunto
}