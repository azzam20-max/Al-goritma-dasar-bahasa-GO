package main

import "fmt"

// Insertion Sort function
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        // Move elements greater than key to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

func main() {
    arr := []int{12, 11, 13, 5, 6}
    fmt.Println("Unsorted array:", arr)
    insertionSort(arr)
    fmt.Println("Sorted array:  ", arr)
}

/*
Insertion Sort
Deskripsi: Algoritma yang membangun hasil sorting satu elemen pada satu waktu dengan cara menyisipkan elemen di tempat yang sesuai di subarray yang sudah terurut.
Kelebihan: Cepat untuk dataset yang hampir terurut atau dataset kecil.
Kekurangan: Kompleksitas waktu O(n²) pada kasus terburuk.
Penggunaan: Sangat efisien untuk dataset kecil atau ketika elemen sudah sebagian besar terurut.
*/
