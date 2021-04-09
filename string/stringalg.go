package stringalg

import(
//	"fmt"
)

/**********************************************
*	Palindrome Check
***********************************************/
func IsPalindrome(str string) bool {
	mid := len(str) / 2
	isPal := true
	var l, h int
	if len(str) % 2 == 0 {
		l = mid - 1
		h = mid
	}else{
		l = mid - 1
		h = mid + 1
	}

	for i:= 0 ;i < mid;i++{
		if str[l-i] != str[h+i]{
			isPal = false
			break
		}
	}

	return isPal
}

/**********************************************
*	Caesar Cipher Encryptor
***********************************************/
func CaesarCipherEncryptor(str string, key int) string {
	var nstr string
	for _, r :=range(str){
		nr := 97 + (r-97+rune(key)) % (26)
		nstr = nstr + string(nr)
	}
	return nstr
}