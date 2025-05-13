package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/kos-dw/enigma/pkg/utils"
)

// 指定された文字列をAES-CBCモードで暗号化する
func Encrypt(data string) (string, error) {
	key, iv, err := utils.Env()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// パディングを追加してデータをブロックサイズに揃える
	paddedData := utils.Pad(data)

	// 暗号化されたデータを格納するバッファを作成
	cipherText := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, []byte(paddedData))

	// 暗号化されたデータをBase64エンコードして返す
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
