package main

import "fmt"

// Heapify a subtree rooted at index i
func heapify(arr []int, n, i int) {
    largest := i    // Initialize largest as root
    left := 2*i + 1 // left child
    right := 2*i + 2 // right child

    // If left child is larger than root
    if left < n && arr[left] > arr[largest] {
        largest = left
    }
    // If right child is larger than the largest so far
    if right < n && arr[right] > arr[largest] {
        largest = right
    }
    // If largest is not root
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        // Recursively heapify the affected subtree
        heapify(arr, n, largest)
    }
}

// Heap Sort function
func heapSort(arr []int) {
    n := len(arr)
    // Build heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    // Extract elements from heap one by one
    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0] // Move current root to end
        heapify(arr, i, 0)              // call max heapify on the reduced heap
    }
}

func main() {
    arr := []int{12, 11, 13, 5, 6, 7}
    fmt.Println("Unsorted array:", arr)
    heapSort(arr)
    fmt.Println("Sorted array:  ", arr)
}


/*
Heap Sort
Deskripsi: Algoritma berbasis heap yang membangun binary heap (max-heap) dari array dan kemudian menukar elemen terbesar (root) ke akhir array, mengurangi ukuran heap, dan mengulangi proses.
Kelebihan: Kompleksitas waktu O(n log n). Tidak membutuhkan ruang tambahan seperti merge sort.
Kekurangan: Lebih lambat dibanding quick sort pada sebagian besar implementasi karena overhead dari heapify.
Penggunaan: Berguna untuk lingkungan di mana konsumsi memori harus diminimalkan.
*/
