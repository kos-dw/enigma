package utils

import (
	"os"
)

// 環境変数からAES暗号化に必要なキーとIVを読み込む（16進文字列形式）
func Env() ([]byte, []byte, error) {
	keyHex := os.Getenv("AES_KEY")
	ivHex := os.Getenv("AES_IV")

	key, err := decodehex(keyHex, "AES_KEY")
	if err != nil {
		return nil, nil, err
	}
	iv, err := decodehex(ivHex, "AES_IV")
	if err != nil {
		return nil, nil, err
	}
	return key, iv, nil
}
