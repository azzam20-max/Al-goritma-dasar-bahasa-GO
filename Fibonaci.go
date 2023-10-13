package main
import "fmt"

func main(){
    var angka_1,angka_2,number,hasil int
    fmt.Print("masukan number ? ")
    fmt.Scan(&number)
    
    angka_1 = 0
    angka_2 = 1
    
    fmt.Print(angka_1, angka_2," ")
    
    for i := 2 ; i< number;i++{
        hasil = angka_1 + angka_2
        fmt.Print(hasil," ")
        angka_1 = angka_2
        angka_2 = hasil 
        
    }
}
