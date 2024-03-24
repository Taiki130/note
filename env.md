# 環境変数の扱い方

https://zenn.dev/kurusugawa/articles/golang-env-lib

- 簡単な処理の場合はosパッケージを使う
    - `os.Getenv`で環境変数を取得できる
    
    ```go
    hoge := os.Getenv("HOGE")
    ```
    
    - 環境変数が設定されているかどうかを判定する場合は、`os.LookupEnv`を使う
    
    ```go
    hoge, ok = os.LookupEnv("HOGE")
    if !ok {
        fmt.Println("HOGE is not set")
    }
    ```
    
- 普通のコードのときは外部ライブラリ`caarlos0/env`を使う（３以上とか？）

```go
type config struct {
    Host string `env:"HOST"`
    Port string `env:"PORT"`
}

// 読み込み
var cfg config
if err := env.Parse(&cfg); err != nil {
    fmt.Println(err)
}

// 結果を利用
fmt.Println(cfg.Host)
fmt.Println(cfg.Port)
```

- .envファイルを読み込みたいときは**`joho/godotenv`を使う**
