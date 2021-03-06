package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxStackPushPop(t *testing.T) {

	minMaxStack := NewMinMaxStack()
	minMaxStack.Push(5)
	assert.Equal(t, 5, minMaxStack.GetMin())
	assert.Equal(t, 5, minMaxStack.GetMax())
	assert.Equal(t, 5, minMaxStack.Peek())
	minMaxStack.Push(7)
	assert.Equal(t, 5, minMaxStack.GetMin())
	assert.Equal(t, 7, minMaxStack.GetMax())
	assert.Equal(t, 7, minMaxStack.Peek())
	minMaxStack.Push(2)
	assert.Equal(t, 7, minMaxStack.GetMax())
	assert.Equal(t, 2, minMaxStack.Peek())
	assert.Equal(t, 2, minMaxStack.Pop())
	assert.Equal(t, 7, minMaxStack.Pop())
	assert.Equal(t, 5, minMaxStack.GetMin())
	assert.Equal(t, 5, minMaxStack.GetMax())
	assert.Equal(t, 5, minMaxStack.Peek())
}
func TestMinMaxStackPushPop2(t *testing.T) {

	minMaxStack := NewMinMaxStack()
	minMaxStack.Push(2)
	assert.Equal(t, 2, minMaxStack.GetMin())
	assert.Equal(t, 2, minMaxStack.GetMax())
	assert.Equal(t, 2, minMaxStack.Peek())
	minMaxStack.Push(7)
	assert.Equal(t, 2, minMaxStack.GetMin())
	assert.Equal(t, 7, minMaxStack.GetMax())
	assert.Equal(t, 7, minMaxStack.Peek())
	minMaxStack.Push(1)
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 7, minMaxStack.GetMax())
	assert.Equal(t, 1, minMaxStack.Peek())
	minMaxStack.Push(8)
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 8, minMaxStack.GetMax())
	assert.Equal(t, 8, minMaxStack.Peek())
	minMaxStack.Push(3)
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 8, minMaxStack.GetMax())
	assert.Equal(t, 3, minMaxStack.Peek())
	minMaxStack.Push(9)
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 9, minMaxStack.GetMax())
	assert.Equal(t, 9, minMaxStack.Peek())
	assert.Equal(t, 9, minMaxStack.Pop())
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 8, minMaxStack.GetMax())
	assert.Equal(t, 3, minMaxStack.Peek())
	assert.Equal(t, 3, minMaxStack.Pop())
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 8, minMaxStack.GetMax())
	assert.Equal(t, 8, minMaxStack.Peek())
	assert.Equal(t, 8, minMaxStack.Pop())
	assert.Equal(t, 1, minMaxStack.GetMin())
	assert.Equal(t, 7, minMaxStack.GetMax())
	assert.Equal(t, 1, minMaxStack.Peek())
	assert.Equal(t, 1, minMaxStack.Pop())
	assert.Equal(t, 2, minMaxStack.GetMin())
	assert.Equal(t, 7, minMaxStack.GetMax())
	assert.Equal(t, 7, minMaxStack.Peek())
	assert.Equal(t, 7, minMaxStack.Pop())
	assert.Equal(t, 2, minMaxStack.GetMin())
	assert.Equal(t, 2, minMaxStack.GetMax())
	assert.Equal(t, 2, minMaxStack.Peek())
	assert.Equal(t, 2, minMaxStack.Pop())
}

func TestBalancedBrackets(t *testing.T) {
	r := BalancedBrackets("aafwgaga()[]a{}{gggg")
	assert.Equal(t, false, r)
	r = BalancedBrackets("(((((([[[[[[{{{{{{{{{{{{()}}}}}}}}}}}}]]]]]]))))))((([])({})[])[])[]([]){}(())")
	assert.Equal(t, true, r)
	r = BalancedBrackets("(((((({{{{{safaf[[[[[([)]safsafsa)]]]]]}}}gawga}}))))))")
	assert.Equal(t, false, r)
}
func TestSunsetViews(t *testing.T) {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	expected := []int{0, 1}
	r := SunsetViews(buildings, "WEST")
	for i, index := range r {
		assert.Equal(t, expected[i], index)
	}
	expected = []int{1, 3, 6, 7}
	r = SunsetViews(buildings, "EAST")
	for i, index := range r {
		assert.Equal(t, expected[i], index)
	}
}
