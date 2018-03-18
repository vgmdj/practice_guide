package main

import (
	"log"
)

//快速排序
func quickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		log.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[i-1] = mid
	quickSort(values[:i-1])
	quickSort(values[i:])

}

func main() {
	nums := []int{7, 6, 5, 4, 3, 2, 1}

	quickSort(nums)

	log.Println(nums)

}
