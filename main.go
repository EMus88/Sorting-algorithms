package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	arr := newArray(30000)
	arr2 := make([]int, 0, len(arr))
	arr2 = append(arr2, arr...)

	bubbleSort(arr)
	shakerSort(arr2)

}

func newArray(size int) []int {
	arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(99))
	}
	return arr
}

//сортировка методом пузырька
func bubbleSort(arr []int) {
	log.Println("Bubble sort started")
	start := time.Now()
	var count int
	var changes int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changes++
			}
			count++
		}
	}
	exTime := time.Since(start)
	log.Printf("Execution time: %s", exTime)
	log.Printf("Iteration count: %d", count)
	log.Printf("Count of changes: %d", changes)
	log.Println("Bubble sort stopped")
}

//шейкерная сортировка
func shakerSort(arr []int) {
	log.Println("Shaker sort started")
	start := time.Now()
	var count int
	var changes int
	left := 0
	right := len(arr) - 1
	for left <= right {

		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				changes++

			}
			count++

		}
		left++
		for i := left + 1; i < right; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				changes++
			}
			count++
		}
		right--
	}
	exTime := time.Since(start)
	log.Printf("Execution time: %s", exTime)
	log.Printf("Iteration count: %d", count)
	log.Printf("Count of changes: %d", changes)
	log.Println("Shaker sort stopped")
}
