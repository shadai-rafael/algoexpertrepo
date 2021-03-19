package recursion

//import "fmt"

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

/**********************************************
*	Get Permutation
***********************************************/
func GetPermutations(array []int) [][]int {
	var result [][]int;
	getPermutationsHelper(array, make([]int, 0), &result)
	return result;
}

func getPermutationsHelper(array []int, permBulider []int, perms *[][]int){
	if len(array) == 0 && len(permBulider) != 0{
		*perms = append(*perms,permBulider)
		return
	}
	for i := 0; i < len(array); i++ {
		tempArray := make([]int, i)
		copy(tempArray, array[:i])
		tempArray = append(tempArray, array[i+1:]...)
		newPermBuilder := make([]int ,len(permBulider))
		copy(newPermBuilder,permBulider)
		newPermBuilder = append(newPermBuilder, array[i])
		getPermutationsHelper(tempArray, newPermBuilder, perms)
	}
	return
}
