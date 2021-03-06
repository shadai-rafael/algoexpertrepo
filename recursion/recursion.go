package recursion

/**********************************************
*	Nth Fibonacci
***********************************************/
func GetNthFib(n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return GetNthFib(n-1) + GetNthFib(n-2)
}

/**********************************************
*	Product Sum
***********************************************/
type SpecialArray []interface{}

func ProductSum(array []interface{}) int {

	return productSumHelper(array, 1)

}
func productSumHelper(array SpecialArray, d int) int {
	var sum int
	for _, elem := range array {
		switch elem.(type) {
		case SpecialArray:
			sum += (d + 1) * productSumHelper(elem.(SpecialArray), d+1)
		case int:
			sum += elem.(int)
		}
	}
	return sum
}
