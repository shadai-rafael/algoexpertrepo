package sort

/**********************************************
*	BubbleSort
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