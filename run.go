package main

import (
	algo "algorithms"
	"fmt"
	"os"
	"strconv"
)

var (
	list []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}
)

func main() {
	//project.MainPooler()
	//project.GetNetwork()
	i, _ := strconv.Atoi(os.Args[1])
	fmt.Println(algo.BinarySearchWithPosition(list, i))
	fmt.Println(algo.BinarySearch(list, i))
}
