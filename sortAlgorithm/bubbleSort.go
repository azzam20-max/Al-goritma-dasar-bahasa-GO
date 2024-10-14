package main

import "fmt"

// Bubble Sort function
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap arr[j] and arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Unsorted array:", arr)
    bubbleSort(arr)
    fmt.Println("Sorted array:  ", arr)
}

/*
Bubble Sort
Deskripsi: Algoritma yang bekerja dengan membandingkan elemen-elemen yang berdekatan dan menukar posisinya jika salah urut. Proses ini diulang hingga tidak ada lagi elemen yang perlu ditukar.
Kelebihan: Implementasi sederhana.
Kekurangan: Tidak efisien untuk dataset besar karena memiliki kompleksitas waktu O(n²).
Penggunaan: Cocok untuk dataset kecil atau ketika kesederhanaan lebih diutamakan daripada efisiensi.
*/
