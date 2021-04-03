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
		return binarySearchHelper(array, target, ll, index)
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

/**********************************************
*	Search In Sorted Matrix
***********************************************/
func SearchInSortedMatrix(matrix [][]int, target int) []int {
	totalRows, totalColumns := len(matrix)-1, len(matrix[0])-1
	endOfIt := totalColumns
	var rowIt, colIt int
	result := []int{-1,-1}
	if totalRows > totalColumns{
		endOfIt = totalRows
	}
	for it:=0; it < endOfIt ;it++{
		if rowIt <= totalRows{
			result = binarySearchRow(matrix, target, rowIt, totalRows, colIt )
			if result[0] != -1{
				return result
			}
		}
		if colIt <= totalColumns{
			result = binarySearchColumn(matrix, target, colIt, totalColumns, rowIt)
			if result[0] != -1{
				return result
			}
		}
		if rowIt < totalRows{
			rowIt++
		}
		if colIt < totalColumns{
			colIt++
		}
	} 
	return result
}

func binarySearchRow(array [][]int, target int, rll int, rul int, column int) []int {
	var row int
	if (rul-rll) == 0 {
		if array[rll][column] == target{
			return []int{rll,column}
		}else{
			return []int{-1,-1}
		}
	}
	row = ((rul - rll)/2)+rll
	if array[row][column] == target{
		return []int{row,column}
	}

	if array[row][column] > target {
		return binarySearchRow(array, target, rll, row, column)
	}else{
		return binarySearchRow(array, target, row + 1, rul, column)
	}
}

func binarySearchColumn(array [][]int, target int, cll int, cul int, row int) []int {
	var column int
	if (cul-cll) == 0 {
		if array[row][cll] == target{
			return []int{row,cll}
		}else{
			return []int{-1,-1}
		}
	}
	column = ((cul - cll)/2)+cll
	if array[row][column] == target{
		return []int{row,column}
	}
	
	if array[row][column] > target {
		return binarySearchColumn(array, target, cll, column, row)
	}else{
		return binarySearchColumn(array, target, column + 1, cul, row)
	}
}