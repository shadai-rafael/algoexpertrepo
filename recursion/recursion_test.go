package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNthFib(t *testing.T) {
	//according to algoexpert GetNthFib(2) = F1
	assert.Equal(t, 5, GetNthFib(6))
}

func TestProductSum(t *testing.T) {
	var array SpecialArray
	array = append(array, 5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4})
	assert.Equal(t, 12, ProductSum(array))
}
