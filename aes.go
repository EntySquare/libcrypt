package libcrypt

type IAesCipher interface {
	Encrypt(origData []byte) []byte
	Decrypt(origData []byte) []byte
}

func NewAesCipher(key, iv []byte) IAesCipher {
	return NewAesCipher128(key, iv)
}

//加密
func AesEncrypt(key, iv []byte, origData []byte) []byte {
	aesCipher := NewAesCipher(key, iv)
	return aesCipher.Encrypt(origData)
}

//解密
func AesDecrypt(key, iv []byte, encrData []byte) []byte {
	aesCipher := NewAesCipher(key, iv)
	return aesCipher.Decrypt(encrData)
}
