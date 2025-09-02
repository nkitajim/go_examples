# 考察
5倍程度早くなっている
Mutexパターンが一番遅く、Atomicとchannelの速度はほぼ変わらなかった(atomicの方が早そうなんだけどなぁ)
```
% ./sp3
normal 4999995000000 10.976834ms
go routine mutex 4999995000000 2.057166ms
go routine atomic 653067456 1.801833ms
go routine channel 4999995000000 1.80125ms
```
