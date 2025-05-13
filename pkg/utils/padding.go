package utils

import (
	"bytes"
	"crypto/aes"
)

// データをAESブロックサイズに揃えるためにパディングを追加する
func Pad(s string) string {
	padding := aes.BlockSize - len(s)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return s + string(padText)
}

// パディングを削除して元のデータを復元する
func Unpad(s string) string {
	length := len(s)
	unpadding := int(s[length-1])
	return s[:(length - unpadding)]
}
