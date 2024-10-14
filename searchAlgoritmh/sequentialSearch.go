// Sequential search function
func sequentialSearch(arr []int, target int) int {
    for i, v := range arr {
        if v == target {
            return i
        }
    }
    return -1 // return -1 if target is not found
}

/*
Sequential Search adalah algoritma pencarian sederhana yang bekerja dengan cara memeriksa 
setiap elemen dalam array atau daftar secara berurutan, mulai dari elemen pertama hingga 
elemen terakhir, untuk menemukan nilai yang dicari (target).
*/
