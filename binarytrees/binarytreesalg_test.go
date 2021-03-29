package binarytreesalg

import(
//		"fmt"
		"testing"
		"github.com/stretchr/testify/assert"
	)

/**********************************************
*	Branch Sums
***********************************************/
/*
"tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "4", "left": "8", "right": "9", "value": 4},
      {"id": "5", "left": "10", "right": null, "value": 5},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "8", "left": null, "right": null, "value": 8},
      {"id": "9", "left": null, "right": null, "value": 9},
      {"id": "10", "left": null, "right": null, "value": 10}
    ],
    "root": "1"
  }
}*/
func TestBranchSums(t *testing.T) {

	id10 := BinaryTree{10,nil,nil}
	id9 := BinaryTree{9,nil,nil}
	id8 := BinaryTree{8,nil,nil}
	id7 := BinaryTree{7,nil,nil}
	id6 := BinaryTree{6,nil,nil}
	id5 := BinaryTree{5,&id10,nil}
	id4 := BinaryTree{4,&id8,&id9}
	id3 := BinaryTree{3,&id6,&id7}
	id2 := BinaryTree{2,&id4,&id5}
	id1 := BinaryTree{1,&id2,&id3}

	r:= BranchSums(&id1)

	e :=[]int{15,16,18,10,11}

	for i, n := range e {
		assert.Equal(t, r[i], n)
	}

}