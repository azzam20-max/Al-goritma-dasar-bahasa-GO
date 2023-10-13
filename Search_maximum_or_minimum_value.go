package main
import "fmt"

type array[10]int

func isi_array(T *array,n int){ //procedur pengisian array/procedure input array
    for i:=0;i <n;i++{
        fmt.Scan(&T[i])
    }
    fmt.Println("nilai array berhasil di input")
}

func max(T array,n int)int { //fungsi untuk mencari nilai tertinggi/function for search max value
    var max int = 1
    for i:=0;i < n;i++{
        if T[i] > T[max]{ //jika ingin mengganti pencarian nilai terkecil maka gunakan < / if you want to search minimum value use this < 
            max = i
        }
    }
    return T[max]
    
}

func main(){
    var T array
    var n int
    fmt.Print("berapa banyak nilai array yang ingin dimasukan ? ")
    fmt.Scan(&n)
    isi_array(&T,n)
    fmt.Println("nilai terbesar adalah =",max(T,n))
