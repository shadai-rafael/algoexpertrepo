package stringalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
/**********************************************
*	Palindrome Check
***********************************************/
func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, IsPalindrome("abba"))
	assert.Equal(t, true, IsPalindrome("abiba"))
	assert.Equal(t, false, IsPalindrome("rafael"))
}
/**********************************************
*	Caesar Cipher Encryptor
***********************************************/
func TestCaesarCipherEncryptor(t *testing.T) {
	str := "xyz"
	nstr := CaesarCipherEncryptor(str,2)
	assert.Equal(t, "zab", nstr);
}

