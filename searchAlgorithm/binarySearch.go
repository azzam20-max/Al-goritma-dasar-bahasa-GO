package main

// Binary search function (iterative version)
func binarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1

    for low <= high {
        mid := (low + high) / 2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1 // return -1 if target is not found
}

/**
Binary Search adalah algoritma pencarian yang sangat efisien untuk menemukan nilai dalam daftar atau array yang sudah terurut. 
Binary search bekerja dengan prinsip divide and conquer untuk membagi ruang pencarian menjadi dua bagian pada setiap langkah.
/
