# 設定項目

## [gibo](https://github.com/simonwhitaker/gibo)で.gitignoreを作成

```bash
gibo dump macOS Go >> .gitignore
```

## [Release Please](https://github.com/googleapis/release-please)で、0.1.0から発番するためのコミットを作成する

```bash
git commit --allow-empty -m "chore: set initial version to 0.1.0" -m "Release-As: 0.1.0"
```
