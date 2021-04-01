package searchalg

import(
	//	"fmt"
		"testing"
		"github.com/stretchr/testify/assert"
	)

/**********************************************
*	Binary Search
***********************************************/
func TestBinarySearch(t *testing.T) {
	array := []int{0,1,21,33,45,45,61,71,72,73}
	r := BinarySearch(array,33)
	assert.Equal(t, array[r], 33)
	r = BinarySearch(array,100)
	assert.Equal(t, r, -1);
}

/**********************************************
*	Find Three Largest Numbers
***********************************************/
func TestFindThreeLargestNumbers(t *testing.T) {
	array := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	result := FindThreeLargestNumbers(array)
	expected := []int{18,141,541}
	for i:= range(result){
		assert.Equal(t, result[i], expected[i]);
	}
}