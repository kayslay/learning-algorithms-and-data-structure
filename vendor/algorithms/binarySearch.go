package algo

import (
	"fmt"
)

func BinarySearch(list []int, item int) bool {
	searchSpace := quickSort(list)
	for {
		midPoint := len(searchSpace) / 2
		if len(searchSpace) == 1 {

			return (searchSpace[midPoint] == item)
		}

		if len(searchSpace) == 2 {
			if searchSpace[midPoint] == item {
				return true
			} else {
				return searchSpace[midPoint-1] == item
			}
		}

		switch {
		case searchSpace[midPoint] == item:
			return true
		case searchSpace[midPoint] > item:
			searchSpace = searchSpace[:midPoint]
			continue
		case searchSpace[midPoint] < item:
			searchSpace = searchSpace[midPoint+1:]
			continue

		}
	}

}

func BinarySearchWithPosition(list []int, item int) (int, bool) {
	searchSpace := quickSort(list)
	start, mid, end := 0, len(searchSpace)/2, len(searchSpace)-1
	for {
		if (start - end) == 0 {
			fmt.Println(start)
			if searchSpace[start] == item {
				return start, true
			}
			return -1, false
		}

		if (end - start) == 1 {
			if searchSpace[start] == item {
				return start, true
			}

			if searchSpace[end] == item {
				return end, true
			}
			return -1, false
		}

		if searchSpace[mid] == item {
			return mid, true
		}
		if item < searchSpace[mid] {
			end = mid
			mid = start + (end-start)/2
			continue
		}
		if item > searchSpace[mid] {
			start = mid + 1
			mid = start + (end-start)/2
			continue
		}
	}
}
