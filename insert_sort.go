package main

import "fmt"

func main() {
	var data = []int{5, 3, 6, 1}

	fmt.Println(insertSort(data))
}

func insertSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
