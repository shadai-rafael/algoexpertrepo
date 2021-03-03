package stack

/**********************************************
*	Min Max Stack construction
***********************************************/
type MinMaxTag string

const (
	Min MinMaxTag = "min"
	Max           = "max"
)

type Node struct {
	Value int
	Next  *Node
}

type MinMaxStack struct {
	Top    *Node
	MinMax map[MinMaxTag]*Node
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		Top:    nil,
		MinMax: make(map[MinMaxTag]*Node),
	}
}

func (stack *MinMaxStack) Peek() int {
	if stack.isEmpty() {
		return -1
	}
	return stack.Top.Value
}

func (stack *MinMaxStack) findMax() *Node {
	maxNode := stack.Top
	for node := stack.Top; node != nil; node = node.Next {
		if maxNode.Value < node.Value {
			maxNode = node
		}
	}
	return maxNode
}

func (stack *MinMaxStack) findMin() *Node {
	minNode := stack.Top
	for node := stack.Top; node != nil; node = node.Next {
		if minNode.Value > node.Value {
			minNode = node
		}
	}
	return minNode
}

func (stack *MinMaxStack) Pop() int {
	if stack.isEmpty() {
		return -1
	}
	removedNode := stack.Top
	stack.Top = removedNode.Next
	if stack.isEmpty() {
		stack.MinMax[Max] = nil
		stack.MinMax[Min] = nil
	} else if stack.Top.Next == nil {
		stack.MinMax[Max] = stack.Top
		stack.MinMax[Min] = stack.Top
	} else if stack.MinMax[Max] == removedNode {
		stack.MinMax[Max] = stack.findMax()
	} else if stack.MinMax[Min] == removedNode {
		stack.MinMax[Min] = stack.findMin()
	}
	return removedNode.Value
}

func (stack *MinMaxStack) Push(number int) {
	node := Node{
		Value: number,
		Next:  stack.Top,
	}
	if stack.isEmpty() {
		stack.MinMax[Min] = &node
		stack.MinMax[Max] = &node
	} else {
		if stack.MinMax[Min].Value > node.Value {
			stack.MinMax[Min] = &node
		} else if stack.MinMax[Max].Value < node.Value {
			stack.MinMax[Max] = &node
		}
	}
	stack.Top = &node
}

func (stack *MinMaxStack) GetMin() int {
	if stack.isEmpty() {
		return -1
	}
	return stack.MinMax[Min].Value
}

func (stack *MinMaxStack) GetMax() int {
	if stack.isEmpty() {
		return -1
	}
	return stack.MinMax[Max].Value
}

func (stack *MinMaxStack) isEmpty() bool {
	if stack.Top == nil {
		return true
	}
	return false
}

/**********************************************
*	Balanced Brackets
***********************************************/
func getMatchBracketMap() map[rune]rune {
	return map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		')': '(',
		'}': '{',
		']': '[',
	}
}

func getBracketToPush() map[rune]bool {
	return map[rune]bool{
		'(': true,
		'{': true,
		'[': true,
		')': false,
		'}': false,
		']': false,
	}
}

func BalancedBrackets(s string) bool {
	var stack []rune
	for _, r := range s {
		if _, ok := getBracketToPush()[r]; ok {
			if getBracketToPush()[r] {
				stack = append(stack, r)
			} else {
				if 0 == len(stack) {
					return false
				} else if getMatchBracketMap()[r] == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	if 0 != len(stack) {
		return false
	}
	return true
}

/**********************************************
*	Sunset View
***********************************************/
/*it sucks but works*/
func SunsetViews(buildings []int, direction string) []int {
	r := []int{}
	if len(buildings) == 0 {
		return r
	} else if direction == "WEST" {
		hstb := buildings[0]
		r = append(r, 0)
		for i := 1; i < len(buildings); i += 1 {
			if buildings[i] > buildings[r[len(r)-1]] && buildings[i] > hstb {
				r = append(r, i)
				if buildings[i] > hstb {
					hstb = buildings[i]
				}
			}
		}
	} else if direction == "EAST" {
		hstb := buildings[len(buildings)-1]
		r = append([]int{len(buildings) - 1}, r...)
		for i := len(buildings) - 2; i >= 0; i -= 1 {
			if buildings[i] > buildings[r[len(r)-1]] && buildings[i] > hstb {
				r = append([]int{i}, r...)
				if buildings[i] > hstb {
					hstb = buildings[i]
				}
			}
		}
	}
	return r
}

/**********************************************
*	Shorten Path
***********************************************/
/*func ShortenPath(path string) string {
	// Write your code here.
	return ""
}*/
