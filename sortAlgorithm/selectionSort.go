package main

import "fmt"

// Selection Sort function
func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Swap arr[i] with the element at minIdx
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Unsorted array:", arr)
    selectionSort(arr)
    fmt.Println("Sorted array:  ", arr)
}

/*
Selection Sort
Deskripsi: Algoritma yang menemukan elemen terkecil (atau terbesar) dalam array dan menukarnya dengan elemen di posisi pertama. Proses ini diulangi untuk seluruh elemen array.
Kelebihan: Tidak membutuhkan banyak pertukaran elemen (swap).
Kekurangan: Seperti bubble sort, memiliki kompleksitas waktu O(n²), sehingga tidak efisien untuk dataset besar.
Penggunaan: Digunakan ketika jumlah swap ingin diminimalkan.
*/
