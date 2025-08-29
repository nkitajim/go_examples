# 考察
funcを引数にした場合でも処理速度は落ちていない
for, gotoともに速度に変化はない
```
go test -bench ... -benchmem
goos: darwin
goarch: arm64
pkg: github.com/go_examples/ch05/exercise_solutions/sp1
cpu: Apple M3
BenchmarkBubbleSort-8     	50057308	        24.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort2-8    	49678470	        23.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSortEx-8   	46616576	        24.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort-8           	89050776	        13.29 ns/op	       0 B/op	       0 allocs/op
```
