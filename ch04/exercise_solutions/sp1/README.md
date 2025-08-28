# 考察
caseが少ない場合、if文が有利に働いてしまっている
10個まで増やした条件の場合、差は縮まったが依然ifの方が若干早くて悩むなぁ

関数であればほぼ速度は変わらない
最適化もうちょい頑張ってほしいなぁ
```
% ./sp1 
switch: 2698124 7.122416ms
switch multi key: 2298528 9.645917ms
if 2698124 4.453625ms
simple switch: 1799190 2.920375ms
simple if 1799190 2.039292ms
simple 10 switch: 4495750 7.49125ms
simple 10 if 4495750 5.8725ms
func switch_return 2698124 5.050583ms
func if_return 2698124 4.823666ms
func switch_return_10 4495750 6.065334ms
func if_return_10 4495750 5.322625ms
```
