# 考察
数が小さいのか意外とpointer渡しも、argsでの渡しも大きな処理速度の差がない
若干pointerが有利
```
go test -bench ... -benchmem
goos: darwin
goarch: arm64
pkg: github.com/go_examples/ch06/exercise_solutions/sp1
cpu: Apple M3
BenchmarkSliceAdd-8   	45095953	        26.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkArgsAdd-8    	40406761	        29.70 ns/op	       0 B/op	       0 allocs/op
PASS
```
