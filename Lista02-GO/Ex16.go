package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
    menorcateto:= 0
    
	for i := 1; i <= n; i++ {
		
		for j := 1; j <= i; j++ {
		    
		        for cateto2:= i; cateto2 >= 1; cateto2--{
			        cat := j*j + cateto2*cateto2

			        if i*i == cat {
			            
			            if menorcateto <= cateto2 {
				            fmt.Printf("hipotenusa = %d, catetos %d e %d\n", i, j, cateto2)
				            menorcateto = cateto2
			            }
			        }
            }
		}
	}
}

