package stringalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**********************************************
*	Caesar Cipher Encryptor
***********************************************/
func TestCaesarCipherEncryptor(t *testing.T) {
	str := "xyz"
	nstr := CaesarCipherEncryptor(str,2)
	assert.Equal(t, "zab", nstr);
}

