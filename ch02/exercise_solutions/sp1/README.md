# 考察
globalとmainのスコープ内での変数の読み取り操作に速度差は感じられない
書き込みに関してはglobal スコープの場合影響をかなり受ける(まあそんなことしないけども)

```
% ./sp1
time: 167ns
normal time: 666.25µs
cig time: 499.875µs
vig time: 506.042µs
cil time: 502.041µs
vil time: 507.417µs
vil add time: 500.583µs
vig add time: 2.632667ms
```
