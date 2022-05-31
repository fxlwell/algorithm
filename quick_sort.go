package main

import "fmt"

func main() {
	data := []int{5, 3, 1, 7, 2, 9, 4, 6, 8}
	quickSort(0, len(data)-1, data)
	fmt.Println(data)
}

func quickSort(start, end int, data []int) {
	if len(data) == 1 {
		return
	}

	s := start
	e := end

	v := data[start]
	for {

		for {
			if data[end] >= v {
				end--
			} else {
				break
			}
		}

		if start >= end {
			data[start] = v
			break
		}

		data[start] = data[end]

		start++

		for {
			if data[start] < v {
				start++
			} else {
				break
			}
		}

		if start >= end {
			data[start] = v
			break
		}

		data[end] = data[start]
	}

	quickSort(s, start-1, data)
	quickSort(start+1, e, data)

}
