package main

import "fmt"

func main(){
    
    var (
        valoringresso[99999] float64
        lucro[99999] float64
        qtingresso[99999] float64
        valorinicial, valorfinal float64
        )
    
    fmt.Scan(&valoringresso[0], &valorinicial, &valorfinal)
    
    if valorinicial >= valorfinal{
        
        fmt.Print("INTERVALO INVALIDO.")
        
    }else{
        
        intervalo := valorfinal-valorinicial
        v:= valorinicial
        
        for i:= 0; i<=int(intervalo); i++{
            
            diferenca:= valoringresso[0]-v
            
            if diferenca > 0{
                
                qtingresso[i] = 120+(diferenca)*50
                
            }else{
                
                qtingresso[i] = 120+(diferenca)*60
                
            }
            
            valoringresso[i+1] = v
            lucro[i] = (v*qtingresso[i]) - (200+0.05*qtingresso[i])
            
            fmt.Printf("V: %.2f, N: %d, L: %.2f\n", v, int(qtingresso[i]), lucro[i])
            
            v++
            
        }
        
        var melhorvalor float64
        var melhorlucro int
        
        for i:=0; i<=int(intervalo); i++{
            
            if lucro[i] > lucro[melhorlucro]{
                
                melhorvalor = valorinicial
                melhorlucro = i
                
            }
            
            valorinicial++
            
        }
        if lucro[melhorlucro] > 0{
            
        fmt.Printf("Melhor valor final: %.2f \nLucro: %.2f \nNumero de ingressos: %d", melhorvalor, lucro[melhorlucro], int(qtingresso[melhorlucro]))
        
        }else{
            
        fmt.Printf("Melhor valor final: 0.00 \nLucro: 0.00 \nNumero de ingressos: 0")
            
        }
        
    }
}