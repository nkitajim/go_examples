# 考察
MakePersonPointer から返される &Person{} はヒープにエスケープします。関数からポインタが返されるたびに、ポインタはスタック上に返されますが、それが指す値はヒープに格納される必要があります。

驚くべきことに、 p もヒープにエスケープするとも書かれています。これは、 p が fmt.Println に渡されるからです。これは、 fmt.Println の引数が ...any であるためです。現在の Go コンパイラは、インターフェース型の引数を介して関数に渡される値をすべてヒープに移動します。(インターフェースについては第7章で説明します。)

```
% go build -gcflags="-m"
# github.com/go_examples/ch06/exercise_solutions/ex1
./main.go:13:6: can inline MakePerson
./main.go:21:6: can inline MakePersonPointer
./main.go:30:18: inlining call to MakePerson
./main.go:31:25: inlining call to MakePersonPointer
./main.go:33:13: inlining call to fmt.Println
./main.go:34:13: inlining call to fmt.Println
./main.go:13:17: leaking param: firstName to result ~r0 level=0
./main.go:13:35: leaking param: lastName to result ~r0 level=0
./main.go:21:24: leaking param: firstName
./main.go:21:42: leaking param: lastName
./main.go:22:9: &Person{...} escapes to heap
./main.go:31:25: &Person{...} does not escape
./main.go:33:13: ... argument does not escape
./main.go:33:14: p1 escapes to heap
./main.go:34:13: ... argument does not escape
./main.go:34:14: *p2 escapes to heap
```

## can inline XXX
inline化可能

## inlining call to XXX
inline化した

## leaking param
関数が終わっても変数が生き続ける

## X argument does not escape
変数 X は関数内でしか使用されていない（からスタックに積めるよ）

## X escapes to heap
変数が関数外で使用されているからスタックに積めない（からヒープに移動するよ）

## inline化の条件
https://go.dev/wiki/CompilerOptimizations
```
関数は十分に単純である必要があり、ASTノードの数は予算(80)未満でなければなりません。
関数には、クロージャ、defer、recover、select などの複雑なものは含まれていません。
関数の先頭に go:noinline; が付いていません
関数には go:uintptrescapes がプレフィックスとして付きません。インライン展開中にエスケープ情報が失われるためです。
関数には本体があります。
```
