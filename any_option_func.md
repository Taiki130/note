## 関数のオプション引数

- Goの文法に可変長引数はあるが、引数の数は変更できてもそれぞれの型は同一である必要がある
- Pythonのようなキーワード引数といったオプション引数のための文法はない
- JavaやC++のような関数のオーバーロードもない

```go
type Portion int

const (
	Regular Portion = iota // 普通
	Small                  // 小盛り
	Large                  // 大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// 麺の量、油揚げ、海老天の有無でインスタンス作成
func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

// 海老天2本入りの大盛り
var tempuraUdon = NewUdon(Large, false, 2)
```

### 別名の関数によるオプション引数

- 一番簡単に使うことができ、コードも読みやすい

```go
func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitsuneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}
```

### 構造体を利用したオプション引数

- 比較的少ないコード量でオプションが大量にある柔軟性のある機能を実現できる
- 欠点はゼロ値の対策やデフォルトの実装がやや面倒

```go
type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(opt Option) *Udon {
	// ゼロ値に対するデフォルト値処理は関数内で行う
	// 朝食時間は海老天1本無料
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}

```

### ビルダーを利用したオプション引数

- 構造体を利用するオプション引数のそれぞれにメソッドを用意する
    - その分だけコード量が増える
- メリットはFluentインターフェース形式のAPIを利用することで、コード補完が賢いエディタではスムーズにコードが書ける

```go
// fluent-option
type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion) *fluentOpt {
	// デフォルトはコンストラクタ関数で設定
	// 必須オプションはここに付与可能
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func useFluentInterface() {
	oomoriKitsune := NewUdon(Large).Aburaage().Order()
}

```

### Functional Optionパターンを使ったオプション引数

- メリットはビルダー形式に似ている
    - パッケージ作成者以外がオプションを自作できる
- デメリットはパッケージ外部から利用する場合に、いちいちパッケージ名を書かなければならない

```go
type OptFunc func(r *Udon)

func NewUdon(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) { r.men = p }
}

func OptAburaage() OptFunc {
	return func(r *Udon) { r.aburaage = true }
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) { r.ebiten = n }
}

func useFuncOption() {
	tokuseiUdon := NewUdon(OptAburaage(), OptEbiten(3))
}
```

## どの実装方法を利用すべきか？

- おすすめはコード量の少ない構造体パターンを実装
    - それを土台に必要になってから他のパターンを実装
