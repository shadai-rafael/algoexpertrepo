package stringalg

import(
//	"fmt"
)

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