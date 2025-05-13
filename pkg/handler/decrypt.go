package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/kos-dw/enigma/pkg/utils"
)

// AES-CBCモードで暗号化されたBase64文字列を復号化する
func Decrypt(base64Data string) (string, error) {
	key, iv, err := utils.Env()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Base64文字列をデコード
	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// 復号化されたデータを格納するバッファを作成
	plainText := make([]byte, len(decodedData))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, decodedData)

	// パディングを削除して元の文字列を返す
	return utils.Unpad(string(plainText)), nil
}
