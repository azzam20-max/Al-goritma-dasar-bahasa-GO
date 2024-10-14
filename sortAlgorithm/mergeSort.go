package main

import "fmt"

// Merge Sort function
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}

// Merge two sorted subarrays
func merge(left, right []int) []int {
    result := []int{}
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}

func main() {
    arr := []int{38, 27, 43, 3, 9, 82, 10}
    fmt.Println("Unsorted array:", arr)
    sortedArr := mergeSort(arr)
    fmt.Println("Sorted array:  ", sortedArr)
}







/*
Merge Sort
Deskripsi: Algoritma berbasis divide and conquer yang membagi array menjadi dua subarray, mengurutkan keduanya, dan kemudian menggabungkannya. Ini adalah algoritma rekursif.
Kelebihan: Kompleksitas waktu O(n log n) bahkan pada kasus terburuk. Stabil dan efisien.
Kekurangan: Membutuhkan ruang tambahan O(n) untuk penggabungan (merge).
Penggunaan: Cocok untuk dataset besar yang membutuhkan stabilitas dan efisiensi tinggi.
*/
