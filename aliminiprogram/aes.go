package aliminiprogram

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

////pkcs7Padding 填充
//func pkcs7Padding(data []byte, blockSize int) []byte {
//	//判断缺少几位长度。最少1，最多 blockSize
//	padding := blockSize - len(data)%blockSize
//	//补足位数。把切片[]byte{byte(padding)}复制padding个
//	padText := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(data, padText...)
//}
//
////pkcs7UnPadding 填充的反向操作
//func pkcs7UnPadding(data []byte) ([]byte, error) {
//	length := len(data)
//	if length == 0 {
//		return nil, errors.New("加密字符串错误！")
//	}
//	//获取填充的个数
//	unPadding := int(data[length-1])
//	return data[:(length - unPadding)], nil
//}

//AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs5Padding(data, blockSize)
	iv := make([]byte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(data []byte, key []byte) (string, error) {
	res, err := AesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

//AesCBCDecrypt 解密
func AesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	// 创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := make([]byte, blockSize)
	if len(encryptData)%blockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}
	//使用cbc
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	// 判断是否越界问题
	if len(encryptData) == 0 {
		return nil, fmt.Errorf("crypt block error")
	}
	encryptData = PKCS5UnPadding(encryptData)
	return encryptData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
