package sortalg

import(
//	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)
/**********************************************
*	Bubble Sort
***********************************************/
func TestBubbleSort(t *testing.T) {
	array := []int{3,2,1,0}
	array = BubbleSort(array)
	expected := []int{0,1,2,3}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
	array = []int{7,4,2,6}
	array = BubbleSort(array)
	expected = []int{2,4,6,7}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}

/**********************************************
*	Insertion Sort
***********************************************/
func TestInsertionSort(t *testing.T) {
	array := []int{3,2,1,0}
	array = InsertionSort(array)
	expected := []int{0,1,2,3}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
	array = []int{7,4,2,6}
	array = InsertionSort(array)
	expected = []int{2,4,6,7}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}

/**********************************************
*	Selection Sort
***********************************************/
func TestSelectionSort(t *testing.T) {
	array := []int{3,2,1,0}
	array = SelectionSort(array)
	expected := []int{0,1,2,3}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
	array = []int{7,4,2,6}
	array = SelectionSort(array)
	expected = []int{2,4,6,7}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}

/**********************************************
*	Merge Sort
***********************************************/
func TestMergeSort(t *testing.T) {
	array := []int{3,2,1,0}
	array = MergeSort(array)
	expected := []int{0,1,2,3}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
	array = []int{7,4,2,6}
	array = MergeSort(array)
	expected = []int{2,4,6,7}
	for i := range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}