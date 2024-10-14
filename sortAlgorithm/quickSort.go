package main

import "fmt"

// Quick Sort function
func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    left, right := 0, len(arr)-1
    pivot := len(arr) / 2
    // Move the pivot to the right
    arr[pivot], arr[right] = arr[right], arr[pivot]
    // Partitioning
    for i := range arr {
        if arr[i] < arr[right] {
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }
    // Place the pivot in its correct position
    arr[left], arr[right] = arr[right], arr[left]
    // Recursively sort the subarrays
    quickSort(arr[:left])
    quickSort(arr[left+1:])
    return arr
}

func main() {
    arr := []int{10, 7, 8, 9, 1, 5}
    fmt.Println("Unsorted array:", arr)
    sortedArr := quickSort(arr)
    fmt.Println("Sorted array:  ", sortedArr)
}

/*
Quick Sort
Deskripsi: Algoritma berbasis divide and conquer yang memilih elemen sebagai pivot, dan mempartisi array di sekitar pivot. Quick sort kemudian diterapkan secara rekursif pada dua subarray yang dihasilkan.
Kelebihan: Sangat cepat pada rata-rata kasus dengan kompleksitas waktu O(n log n).
Kekurangan: Pada kasus terburuk, waktu eksekusinya bisa menjadi O(n²) (misalnya, jika pivot terpilih buruk). Namun, ini dapat diminimalkan dengan memilih pivot secara acak.
Penggunaan: Quick sort banyak digunakan karena kecepatan dan efisiensinya pada sebagian besar kasus.
*/
