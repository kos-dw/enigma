package utils

import (
	"encoding/hex"
	"errors"
)

// 16進文字列をデコードし、エラー時に適切なメッセージを返す
func decodehex(input, fieldName string) ([]byte, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return nil, errors.New("invalid " + fieldName)
	}
	return decoded, nil
}
