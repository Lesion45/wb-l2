package main

import "fmt"

type SortStrategy interface {
	Sort(arr []int) []int
}

type SortContext struct {
	strategy SortStrategy
}

func (sc *SortContext) SetStrategy(strategy SortStrategy) {
	sc.strategy = strategy
}

func (sc *SortContext) ExecuteStrategy(arr []int) []int {
	return sc.strategy.Sort(arr)
}

type InsertionSort struct{}

func (i *InsertionSort) Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

type QuickSort struct{}

func (q *QuickSort) Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	return append(append(q.Sort(left), pivot), q.Sort(right)...)
}

func main() {
	arr := []int{3, 5, 2, 1, 4, 6}

	context := &SortContext{}

	context.SetStrategy(&InsertionSort{})
	sortedArr := context.ExecuteStrategy(arr)
	fmt.Println("InsertionSort:", sortedArr)

	context.SetStrategy(&QuickSort{})
	sortedArr = context.ExecuteStrategy(arr)
	fmt.Println("QuickSort:", sortedArr)
}
