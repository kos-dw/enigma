# Enigma

Enigmaプロジェクトは、暗号化および復号化のアルゴリズムを実装したCLIツールです。可逆性のある暗号化を行うために、AES（Advanced Encryption Standard）を使用しています。このツールは、ユーザーが指定したテキストを暗号化または復号化することができます。

## 使用方法
1. 環境変数をセットします:
```bash
export AES_KEY=＄{暗号鍵}
export AES_IV=＄{初期化ベクトル}
```

2. buildします
```bash
go build -o bin/enigma cmd/enigma/main.go
```
3. 暗号化または復号化したいテキストを入力し、設定を調整してください。
```bash
bin/enigma -e "暗号化したいテキスト"
bin/enigma -d "復号化したいテキスト"
```


