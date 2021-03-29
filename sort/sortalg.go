package sortalg

import(
//		"fmt"
)

/**********************************************
*	Bubble Sort
***********************************************/
func BubbleSort(array []int) []int {
	for i := range(array){
		for j := 1; j < len(array)-i; j++ {
			if array[j-1] > array[j]{
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
	return array;
}

/**********************************************
*	Insertion Sort
***********************************************/
func InsertionSort(array []int) []int {
	for i := range(array){
		for j := i; j > 0; j-- {
			if array[j-1] > array[j]{
				array[j-1], array[j] = array[j], array[j-1]
			}else{
				break
			}
		}
	}
	return array;
}

/**********************************************
*	Selection Sort
***********************************************/
func SelectionSort(array []int) []int {
	for i := range(array){
		min := i
		for j := i+1; j < len(array); j++ {
			if array[min] > array[j]{
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
	return array
}

/**********************************************
*	Merge Sort
***********************************************/
func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	lh := make([]int, len(array[:len(array)/2]))
	copy(lh, array[:len(array)/2] )
	hh := make([]int, len(array[len(array)/2:]))
	copy(hh, array[len(array)/2:] )
	lh = MergeSort(lh)
	hh = MergeSort(hh)
	i, j, k:= 0, 0, 0; 
	for{
		if i == len(lh) && j == len(hh){
			break
		}else if i == len(lh){
			array[k] = hh[j]
			j++
		}else if j == len(hh){
			array[k] = lh[i]
			i++
		}else if lh[i] > hh[j] {
			array[k] = hh[j]
			j++
		}else{
			array[k] = lh[i]
			i++
		}
		k++
	}
	return array
}
