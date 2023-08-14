package main

import "fmt"

func main() {
	var data = []int{5, 3, 6, 1}

	fmt.Println(selectSort(data))
}

func selectSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		minElement := data[i]
		minIdx := i

		for j := i + 1; j < len(data); j++ {
			if data[j] < minElement {
				minElement = data[j]
				minIdx = j
			}
		}

		data[i], data[minIdx] = data[minIdx], data[i]

	}
	return data
}
