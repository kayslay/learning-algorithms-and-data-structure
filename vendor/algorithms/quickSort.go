package algo

func quickSort(list []int) []int {
	return divide(list)
}

func conquer(pivot int, list []int) []int {
	var l, e, g []int
	e = append(e, pivot)
	for _, val := range list {
		if val < pivot {
			l = append(l, val)
		} else if val > pivot {
			g = append(g, val)
		} else {
			e = append(e, val)
		}
	}
	return combine(<-divChan(divide, l), e, <-divChan(divide, g))
}

func divide(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[len(list)-1]
	list = list[:len(list)-1]
	return conquer(pivot, list)
}

func combine(l, e, g []int) []int {
	return append(l, append(e, g...)...)
}

func divChan(a func([]int) []int, l []int) <-chan []int {
	out := make(chan []int)
	go func() { out <- a(l) }()
	return out
}
