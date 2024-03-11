
# ファイルioについて

https://www.kwbtblog.com/entry/2020/05/07/014836

- 入出力はioパッケージとして抽象化されている
- ファイルに限らず、他の入出力であっても、io.Reader・io.Writerインターフェイスを持っていれば、そのまま使える
- ioでのデータのやり取りは、全てbyteのスライスを介して行われる

## ファイル読み込み（全部まとめて書き込み）

```go
f, err := os.Open("read_data.txt")
defer f.Close()
if err != nil {
    return err
}

content, err := os.ReadAll(f) // 全部読み込んでくれる
if err != nil {
    return err
}

fmt.Println(string(content))
```

## ファイル読み込み（ファイルオープンから全てまとめて）

```go
content, err := os.ReadFile("read_data.txt")
if err != nil {
    return err
}
fmt.Println(string(content))
```

## ファイル書き込み

```go
f, err := os.Create("write_data.txt")
defer f.Close()
if err != nil {
    return err
}

content := "TEST\ntest\nテスト\nてすと"

_, err = f.Write([]byte(content))
if err != nil {
    return err
}
```

## ファイル書き込み（バッファ）
- 大量のデータを書き込んだ時など
- バッファーサイズまでデータが溜まったら、溜まった分を書き込むので、書き込みデータの最後まで来たら、Writer.Flash()で、バッファーに溜まっていた残りデータを書き込む必要がある

```go
f, err := os.Create("write_data.txt")
defer f.Close()
if err != nil {
    return err
}

fw := bufio.NewWriter(f)
content := "TEST\ntest\nテスト\nてすと"

_, err = fw.Write([]byte(content))
if err != nil {
    return err
}

err = fw.Flush()
if err != nil {
    return err
}
```

## ファイル書き込み（ファイルオープンから全てまとめて）
- ファイルオープンから書き込み、そしてファイルのクローズまで

```go
content := "TEST\ntest\nテスト\nてすと"

err := os.WriteFile("write_data.txt", []byte(content), 0644)
if err != nil {
    return err
}
```

## ファイルから１行づつ読み込み
- データを１行づつ読み込みたい

```go
f, err := os.Open("read_data.txt")
defer f.Close()
if err != nil {
    return err
}

fr := bufio.NewScanner(f)

for fr.Scan() {
    fmt.Println(fr.Text())
}

if err := fr.Err(); err != nil {
    return err
}
```

## メモリ上にioを作る

```go
b := bytes.NewBuffer([]byte("TEST"))

content, err := os.ReadAll(b) // io.Reader
if err != nil {
    return err
}
fmt.Println(string(content)) // TEST

fmt.Println(b.String()) // TEST
```

## io.Reaerとio.Writerを直接つなぐ（ストリーム）
- io.Readerからデータを読み取り、そのままio.Writerに書き込みたい時
- 全てのデータを受け取ってからデータを出力すると、全てのデータを格納するメモリ領域が必要になるので効率が悪くなる
- Read()で部分的に読み取って、それをWrite()で書き込むという作業を、io.ReaderがEOFになるまで続けるようにすれば、メモリはRead()で読み取った分だけで済むので、 効率がよくなる

```go
r, err := Open("read_data.txt")
if err != nil{
    return err
}
defer r.Close()

w, err := Create("write_data.txt")
if err != nil{
    return err
}
defer w.Close()

// rから読み取ってwに書き出すのを、rのデータが最後になるまでやってくれる
_, err = io.Copy(w, r)
if err != nil{
    return err
}
```

## io.Writerとio.Readerを直接つなぐ（パイプ）
- io.Writeにデータを書き込んで、その書き込まれたデータを、io.Readerを引数に取る関数に渡したい時は、io.Pipe()を使う

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
```
