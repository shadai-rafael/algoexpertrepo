package searchalg
import(
//		"fmt"
)
/**********************************************
*	Binary Search
***********************************************/
func BinarySearch(array []int, target int) int {
	return binarySearchHelper(array,target,0, len(array)-1)
}

func binarySearchHelper(array []int, target int, ll int, ul int) int {
	index := 0
	if (ul-ll) == 0 {
		if array[ll] == target{
			return ll
		}else{
			return -1
		}
	}
	index = ((ul - ll)/2)+ll
	if array[index] == target{
		return index
	}
	
	if array[index]>target {
		return binarySearchHelper(array, target, ll, index-1)
	}else{
		return binarySearchHelper(array, target, index+1, ul)
	}

}

/**********************************************
*	Find Three Largest Numbers
***********************************************/
func FindThreeLargestNumbers(array []int) []int {
	result := make([]int,3)
	copy(result,array[:3])
	sortingThreeLenghtArray(&array)
	for i := 2; i<len(array);i++{
		if array[i]>=result[2]{
			result[0] = result[1]
			result[1] = result[2]
			result[2] = array[i]
		}else if array[i]>=result[1]{
			result[0] = result[1]
			result[1] = array[i]
		}else if array[i]>result[0]{
			result[0] = array[i]
		}
	}
	return result
}

func sortingThreeLenghtArray(array *[]int){
	for i := range(*array){
		min := i
		for j := i+1; j < len(*array); j++ {
			if (*array)[min] > (*array)[j]{
				min = j
			}
		}
		(*array)[i], (*array)[min] = (*array)[min], (*array)[i]
	}
}