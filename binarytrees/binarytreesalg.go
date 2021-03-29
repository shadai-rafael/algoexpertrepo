package binarytreesalg

/**********************************************
*	Branch Sums
***********************************************/
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	branches := make([]int,0)
	sum := 0
	branchSumsHelper(root, sum, &branches)
	return branches
}

func branchSumsHelper(root *BinaryTree, sum int, branches *[]int) {
	sumlocal := sum + root.Value

	if nil == root.Right && nil == root.Left {
		*branches = append(*branches,sumlocal)
		return
	}

	if nil != root.Left{
		branchSumsHelper(root.Left, sumlocal, branches)
	}

	if nil != root.Right{
		branchSumsHelper(root.Right, sumlocal, branches)
	}
		
}
