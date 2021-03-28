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