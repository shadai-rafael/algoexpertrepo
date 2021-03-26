package recursion

import (
//	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

/**********************************************
*	Nth Fibonacci
***********************************************/
func TestGetNthFib(t *testing.T) {
	//according to algoexpert GetNthFib(2) = F1
	assert.Equal(t, 5, GetNthFib(6))
}

/**********************************************
*	Product Sum
***********************************************/
func TestProductSum(t *testing.T) {
	var array SpecialArray
	array = append(array, 5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4})
	assert.Equal(t, 12, ProductSum(array))
}
/**********************************************
*	Get Permutation
***********************************************/
func TestGetPermutations(t *testing.T) {
	array := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	result := GetPermutations(array);
	for i:=0;i<len(expected); i++{
		for j:=0;j<len(expected[i]);j++{
			assert.Equal(t,expected[i][j], result[i][j]);
		}
	}
}

/**********************************************
*	Power Set
***********************************************/
func TestPowerset(t *testing.T) {
	array := []int{1,2,3,4,5}
	res := Powerset(array)
	expected := [][]int{
	{},
	{1},
	{2},
	{1,2},
	{3},
	{1,3},
	{2,3},
	{1,2,3},
	{4},
	{1,4},
	{2,4},
	{1,2,4},
	{3,4},
	{1,3,4},
	{2,3,4},
	{1,2,3,4},
	{5},
	{1,5},
	{2,5},
	{1,2,5},
	{3,5},
	{1,3,5},
	{2,3,5},
	{1,2,3,5},
	{4,5},
	{1,4,5},
	{2,4,5},
	{1,2,4,5},
	{3,4,5},
	{1,3,4,5},
	{2,3,4,5},
	{1,2,3,4,5},
	}
	for i:=0;i<len(expected); i++{
		for j:=0;j<len(expected[i]);j++{
			assert.Equal(t,expected[i][j], res[i][j]);
		}
	}
}