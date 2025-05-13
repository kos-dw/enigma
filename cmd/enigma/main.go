// このコードは、コマンドライン引数を使用して文字列を暗号化または復号化するためのものです。
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kos-dw/enigma/pkg/handler"
)

func main() {
	// 暗号化する文字列を指定するフラグ
	enc := flag.String("e", "none", "Specifies the string to encrypt.")
	// 復号化する文字列を指定するフラグ
	dec := flag.String("d", "none", "Specifies the string to decrypted.")

	flag.Parse()

	// フラグの入力を検証して処理を実行
	if *enc != "none" && *dec == "none" {
		encrypted, err := handler.Encrypt(*enc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encrypted)
	} else if *dec != "none" && *enc == "none" {
		decrypted, err := handler.Decrypt(*dec)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(decrypted)

	} else {
		log.Fatal("[Error] Missing flags: specify either -e or -d")
	}
}
