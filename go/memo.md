# Goが利用される場面
- CLI
- TUI
- Webアプリ

## 理由
- 標準でUTF-8をサポート
    - UTF-8をサポートしているため、ソースコードがUTF-8で記述されていたとしても、Windowsでの出力やファイルI/Oで問題が起こるということはない
- マルチプラットフォーム
    - 異なる環境向けでも、開発者がプラットフォームの違いをあまり意識することない
- 並行処理の扱いやすさ
    - goroutineというスレッドよりも軽い並行処理の仕組みがある
    - 並行処理で問題となる処理も簡単に実装できる機能を提供している
    - Goはマルチコアが活かされるようにランタイムが設計されていて、CPUを効率良く使うことができる
- ストリーム指向
    - io.Readerやio.Writerというインターフェースを使ってファイルや通信接続を抽象的に扱える
    - メモリを無駄に使わない
- シングルバイナリ
    - 標準パッケージがサードパーティライブラリに依存していない

# なぜ使われるか
- コンパイルの速さ
    - 簡素な文法で設計されているため、コンパイルが速い
    - 開発サイクルの加速
- レビューのしやすさ
    - シンプルな言語仕様のため、難しい慣用句がない
    - gofmt
- 周辺ツールの充実
    - 有志の手によって便利に実行できるように拡張機能が提供されたり、その相乗効果によってプログラミング言語の質が向上
- パッケージ公開の簡単さ
    - VCSで簡単に公開できる
- libc非依存
    - libcを使うということは、実行ファイルを生成するためにC言語のコンパイラやリンカが必要となる
        - これによりクロスコンパイルの難易度が高くなる
    - libc非依存なため、クロスコンパイルを容易に行うことができる
- 共同開発でのスキル差
    - シンプルで誰が書いても同じようになるため、引き継ぎ時のリスクは低くなる

## 疑問
- アプリケーションプロジェクト構成
    - https://zenn.dev/foxtail88/articles/824c5e8e0c6d82
    - https://github.com/golang-standards/project-layout/blob/master/README_ja.md
    - https://future-architect.github.io/articles/20200528/
- channel
- sync.Mutex
- goroutine
- context
- ブランクimport

# filepath.Walk

- 第一引数の `root` に指定されたディレクトリから再帰的にファイルやディレクトリを列挙し、第二引数の `fn` でフィルタリングやエラーハンドリングする
    - https://zenn.dev/akinobufujii/articles/c65a91061fdc7f

# exec.Command

- 外部コマンドを実行する
    - https://tokitsubaki.com/go-exec-command/665/
- 結果を受け取る場合は、Outputを使用する
```go
result, err := exec.Command("date").Output()
fmt.Println(string(result)) // -> 2021年  1月  8日 金曜日 10:38:40 JST
```

# for … range

https://golang.hateblo.jp/entry/2019/10/07/171630#foreach-for--range

- foreachのようなもの
- 以下 i にはループ回数が、 v には配列の値が順番に代入される
- 回数が必要ない場合は `for _, v := range` とアンダーバーにして破棄する

```go
for i, v := range []string{"foo", "bar", "baz"} {
    fmt.Println(i, v)
    // 0 foo
    // 1 bar
    // 2 baz
}
```

- map(連想配列)の場合は以下の k にキーが, v に値が順番に代入される。
```go
for k, v := range map[string]int{"key-1": 100, "key-2": 200, "key-3": 300} {
    fmt.Println(k, v)
    // key-2 200
    // key-3 300
    // key-1 100
}
```

# bufio.Newscanner

https://programing-school.work/golang-file-line-bufio/

- ファイルなどを一行ずつ読み込みたいときに使う

```go
// スキャナーを作成
    scanner := bufio.NewScanner(file)

    // ファイルを1行ずつ読み込む
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
```

# バッファ

https://note.crohaco.net/2019/golang-buffer

- `bytes` 標準パッケージには[Buffer](https://golang.org/pkg/bytes/#Buffer) という構造体が定義されていて、これをバッファとして使うことができる
    - *bytes.Bufferはio.Readerとio.Writerを実装している
    
    ```go
    // ここではdataをバイト列と仮定
    b := new(bytes.Buffer)
    b.Write(data)
    
    uploader := &s3manager.UploadInput{
    	// Bodyはio.Readerを受け入れる
    	Body:        b,
    	以下つづく...
    }
    ```
    
- `bufio` はIOをラップしてそこに対する入出力をバッファリングする標準パッケージ
- IOアクセス処理で透過的にバッファを利用する

# sync.WaitGroupの使い方
- go routine使用時など、すべての処理が完了するまで待機するときに使う

```go
var wg.WaitGroup
for i := 0; i < 10 ; i++ {
	wg.Add(1)
	go func() {
		/* do something */
		wg.Done()
	}()
}

wg.Wait()
```

# パッケージ

- すべてのソースコードがパッケージに属する
- フォルダ内に含まれるファイルは基本的に同じパッケージ名である必要がある
    - ただし<package>_testだけは共存可能
- 慣習として、パッケージ名とフォルダ名を同一にすることが多い
- エントリーポイント`func main()`は、必ずmainパッケージに含まれている必要がある
    - この場合のフォルダ名はアプリケーションの実行ファイル名と同一にする


# 構造体

- 名前が大文字始まりなら他のパッケージから読める「公開」状態となり、それ以外なら「非公開」となる
- フィールド名も同様に大文字なら「公開」それ以外なら「非公開」となる

# 外部コマンドを実行するos.Exec()のStdoutPipeからaws-sdk-goのs3managerでuploadする方法

https://github.com/purecloudlabs/gprovision/blob/b79a0d8da330cf23243c651e8d923e9450eeccde/pkg/corer/stream/stream.go#L82

- S3アップロードはイントロスペクションを行い、Seek()メソッドが存在すればそれを見つける。
- パイプはファイルなのでSeekメソッドがあるが、それを使うことはできない。これを使うと、s3アップロードが失敗する。
- これを避けるには、io.Reader用のRead()メソッドを持つ独自の型を作成する。
```go
type unseekableReader struct {
	rdr io.Reader
}

func (u *unseekableReader) Read(p []byte) (int, error) { return u.rdr.Read(p) }

var _ io.Reader = &unseekableReader{}
```

