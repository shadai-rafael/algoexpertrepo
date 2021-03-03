package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNthFib(t *testing.T) {
	//according to algoexpert GetNthFib(2) = F1
	assert.Equal(t, 5, GetNthFib(6))
}
