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
	array = []int{0,1,21,33,45,45,61,71,72}
	r = BinarySearch(array,33)
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

/**********************************************
*	Search In Sorted Matrix
***********************************************/
func TestSearchInSortedMatrix(t *testing.T) {
	matrix := make([][]int,5)
	matrix[0] = []int{1}
	matrix[1] = []int{3}
	matrix[2] = []int{6}
	matrix[3] = []int{8}
	matrix[4] = []int{10}

	result := SearchInSortedMatrix(matrix,8)

	assert.Equal(t, result[0], 3);
	assert.Equal(t, result[1], 0);

	matrix = make([][]int,6)
	matrix[0] = []int{1}
	matrix[1] = []int{3}
	matrix[2] = []int{6}
	matrix[3] = []int{7}
	matrix[4] = []int{8}
	matrix[5] = []int{10}

	result = SearchInSortedMatrix(matrix,8)

	assert.Equal(t, result[0], 4);
	assert.Equal(t, result[1], 0);

	result = SearchInSortedMatrix(matrix,100)

	assert.Equal(t, result[0], -1);
	assert.Equal(t, result[1], -1);

	matrix1 := make([][]int,1)
	matrix1[0] = []int{1,3,6,8,10}

	result = SearchInSortedMatrix(matrix1,8)

	assert.Equal(t, result[0], 0);
	assert.Equal(t, result[1], 3);

	matrix1[0] = append(matrix1[0],12)

	result = SearchInSortedMatrix(matrix1,10)

	assert.Equal(t, result[0], 0);
	assert.Equal(t, result[1], 4);

	matrix = make([][]int,5)
	matrix[0] = []int{1,4,7,12,15,1000}
	matrix[1] = []int{2,5,19,31,32,1001}
	matrix[2] = []int{3,8,24,33,35,1002}
	matrix[3] = []int{40,41,42,44,45,1003}
	matrix[4] = []int{99,100,103,106,128,1004}

	result = SearchInSortedMatrix(matrix,44)

	assert.Equal(t, result[0], 3)
	assert.Equal(t, result[1], 3)

	result = SearchInSortedMatrix(matrix,4);
	assert.Equal(t, result[0], 0);
	assert.Equal(t, result[1], 1);
}