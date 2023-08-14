package main

import "fmt"

func main() {
	var data = []int{5, 3, 6, 1}

	fmt.Println(bubbleSort(data))
}

func bubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
