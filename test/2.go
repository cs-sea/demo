package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	key := []byte("aaaaabbbbbeeeeea")
	data := []byte("data1")
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("错误 -" + err.Error())
		panic(err)
	}
	blockMode := cipher.NewCBCEncrypter(block, key)

	data = PKCS7Padding(data, block.BlockSize())
	crypted := make([]byte, len(data))

	blockMode.CryptBlocks(crypted, data)

	text := base64.StdEncoding.EncodeToString(crypted)
	fmt.Println(text)
	decode(text)
}

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func decode(text string) {

	decode_data, err := base64.StdEncoding.DecodeString(text)

	key := []byte("aaaaabbbbbeeeeea")
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("错误 -" + err.Error())
	}
	blockMode := cipher.NewCBCDecrypter(block, key)

	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	fmt.Println(string(unpad(origin_data)))

}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
