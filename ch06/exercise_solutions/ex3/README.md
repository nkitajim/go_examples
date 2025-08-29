# 考察
容量を確保せずにappendした場合GCが大量に発生する
p5が早いケースがあるがこの時にはmemory容量がすでに多いため早い
len 0の場合もallocationは行われている
```
nkitajim@macbook ex3 % time GOGC=off ./ex3
make p1 time 257.208µs 10000000
make p2 time 163.792µs 0
make p3 time 321.203958ms 10000000
make p4 time 410.757417ms 10000000
make p5 time 49.990125ms 10000000
GOGC=off ./ex3  0.26s user 0.43s system 84% cpu 0.821 total

nkitajim@macbook ex3 % time GOGC=1 ./ex3
make p1 time 152.458µs 10000000
make p2 time 48.611833ms 0
make p3 time 381.556708ms 10000000
make p4 time 540.15075ms 10000000
make p5 time 157.867666ms 10000000
GOGC=1 ./ex3  2.31s user 0.25s system 222% cpu 1.146 total

nkitajim@macbook ex3 % time GOGC=1 GODEBUG=gctrace=1 ./ex3
gc 1 @0.001s 1%: 0.015+0.19+0.005 ms clock, 0.12+0.088/0.12/0.043+0.043 ms cpu, 0->0->0 MB, 0 MB goal, 0 MB stacks, 0 MB globals, 8 P
make p1 time 196.25µs 10000000
gc 2 @0.002s 4%: 0.009+0.16+0.069 ms clock, 0.076+0.077/0.13/0.082+0.55 ms cpu, 381->381->0 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 3 @0.003s 6%: 0.007+0.36+0.054 ms clock, 0.060+0/0.37/0.060+0.43 ms cpu, 381->381->381 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 8 P
make p2 time 94.582417ms 0
gc 4 @0.098s 0%: 0.028+0.42+0.002 ms clock, 0.23+0/0.20/0.27+0.018 ms cpu, 385->385->1 MB, 385 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 5 @0.098s 0%: 0.011+0.43+0.004 ms clock, 0.095+0/0.42/0+0.035 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 6 @0.099s 0%: 0.023+0.64+0.013 ms clock, 0.18+0/0.51/0+0.10 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 7 @0.099s 1%: 0.96+0.57+0.23 ms clock, 7.6+0/0.25/0.002+1.8 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 8 @0.101s 1%: 0.42+0.72+0.074 ms clock, 3.3+0/0.27/0+0.59 ms cpu, 5->5->4 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 9 @0.103s 2%: 0.020+2.1+0.14 ms clock, 0.16+0/0.62/0+1.1 ms cpu, 6->10->8 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 10 @0.105s 2%: 0.020+0.35+0.77 ms clock, 0.16+0/0.48/0.11+6.1 ms cpu, 13->13->8 MB, 13 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 11 @0.107s 3%: 0.90+1.1+0.001 ms clock, 7.2+0/1.2/0+0.010 ms cpu, 13->13->10 MB, 13 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 12 @0.109s 3%: 0.27+1.8+0 ms clock, 2.1+0/1.3/0.87+0.007 ms cpu, 16->16->12 MB, 17 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 13 @0.111s 3%: 0.018+2.2+0.002 ms clock, 0.14+0/0.70/0.32+0.016 ms cpu, 21->21->15 MB, 21 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 14 @0.113s 3%: 0.018+0.68+0.032 ms clock, 0.15+0/0.87/0.53+0.25 ms cpu, 26->26->19 MB, 26 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 15 @0.115s 3%: 0.014+0.46+0.017 ms clock, 0.11+0/0.75/0.95+0.14 ms cpu, 33->33->24 MB, 33 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 16 @0.117s 3%: 0.015+3.6+0.017 ms clock, 0.12+0/1.7/3.2+0.14 ms cpu, 41->41->30 MB, 41 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 17 @0.120s 4%: 0.022+0.76+0.019 ms clock, 0.17+0/1.2/1.8+0.15 ms cpu, 51->51->38 MB, 51 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 18 @0.123s 4%: 0.016+5.9+0.008 ms clock, 0.13+0/2.4/5.0+0.064 ms cpu, 64->64->47 MB, 64 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 19 @0.129s 4%: 0.018+1.2+0.016 ms clock, 0.14+0/2.3/1.3+0.13 ms cpu, 80->80->59 MB, 80 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 20 @0.132s 4%: 0.020+1.1+0.010 ms clock, 0.16+0/1.8/3.2+0.081 ms cpu, 101->101->74 MB, 101 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 21 @0.136s 4%: 0.030+11+0.003 ms clock, 0.24+0/11/0.10+0.029 ms cpu, 126->126->93 MB, 126 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 22 @0.149s 6%: 0.028+11+2.0 ms clock, 0.22+0/6.9/13+16 ms cpu, 157->157->116 MB, 157 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 23 @0.162s 6%: 0.026+16+0.004 ms clock, 0.21+0/8.2/15+0.032 ms cpu, 197->197->145 MB, 197 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 24 @0.182s 6%: 0.015+20+0.003 ms clock, 0.12+0/12/22+0.026 ms cpu, 246->246->181 MB, 246 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 25 @0.205s 6%: 0.015+3.2+0.016 ms clock, 0.12+0/6.2/9.2+0.13 ms cpu, 308->308->227 MB, 308 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 26 @0.227s 6%: 0.022+32+1.1 ms clock, 0.18+0/14/34+9.4 ms cpu, 385->385->284 MB, 385 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 27 @0.261s 7%: 0.014+50+0.002 ms clock, 0.11+5.5/56/26+0.023 ms cpu, 481->481->197 MB, 481 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 28 @0.313s 7%: 0.016+40+2.5 ms clock, 0.13+0/11/26+20 ms cpu, 444->444->444 MB, 444 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 29 @0.360s 9%: 0.015+92+0.005 ms clock, 0.12+0/107/77+0.046 ms cpu, 752->752->555 MB, 752 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 30 @0.454s 9%: 0.014+6.8+0.019 ms clock, 0.11+0/13/31+0.15 ms cpu, 940->940->694 MB, 940 MB goal, 0 MB stacks, 0 MB globals, 8 P
make p3 time 416.427292ms 10000000
gc 31 @0.514s 8%: 0.020+0.60+0.002 ms clock, 0.16+0/0.71/0.10+0.020 ms cpu, 700->700->1 MB, 701 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 32 @0.514s 8%: 0.019+0.77+0.010 ms clock, 0.15+0/0.90/0.095+0.084 ms cpu, 3->3->1 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 33 @0.515s 8%: 0.029+1.0+0 ms clock, 0.23+0/1.1/0.011+0.007 ms cpu, 4->4->2 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 34 @0.516s 8%: 0.018+1.1+0.001 ms clock, 0.14+0/1.3/0.089+0.009 ms cpu, 5->5->2 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 35 @0.518s 8%: 0.016+1.2+0.001 ms clock, 0.13+0/1.3/0.53+0.008 ms cpu, 6->6->3 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 36 @0.519s 8%: 0.017+1.3+0.001 ms clock, 0.14+0/1.2/0.60+0.009 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 37 @0.520s 8%: 0.016+1.5+0.001 ms clock, 0.13+0/0.52/1.4+0.011 ms cpu, 13->13->10 MB, 13 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 38 @0.522s 8%: 0.016+2.0+0.001 ms clock, 0.12+0/1.2/0.99+0.010 ms cpu, 16->16->12 MB, 17 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 39 @0.524s 8%: 0.018+2.4+0.002 ms clock, 0.15+0/1.3/1.8+0.017 ms cpu, 21->21->15 MB, 21 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 40 @0.527s 8%: 0.018+3.2+0.001 ms clock, 0.14+0/1.8/1.8+0.013 ms cpu, 26->26->19 MB, 26 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 41 @0.530s 8%: 0.017+6.7+0.001 ms clock, 0.13+0/4.6/0.80+0.008 ms cpu, 33->33->24 MB, 33 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 42 @0.537s 8%: 0.023+4.7+0.001 ms clock, 0.18+0/3.4/1.8+0.014 ms cpu, 41->41->30 MB, 41 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 43 @0.542s 8%: 0.019+5.9+0.002 ms clock, 0.15+0/2.5/3.9+0.022 ms cpu, 51->51->38 MB, 51 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 44 @0.548s 8%: 0.017+8.0+0.002 ms clock, 0.13+0/4.5/5.2+0.021 ms cpu, 64->64->47 MB, 64 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 45 @0.556s 8%: 0.015+9.3+0.002 ms clock, 0.12+0/5.7/5.8+0.023 ms cpu, 80->80->59 MB, 80 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 46 @0.565s 8%: 0.017+10+1.2 ms clock, 0.13+0/5.3/8.8+9.6 ms cpu, 101->101->74 MB, 101 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 47 @0.577s 8%: 0.029+14+0.013 ms clock, 0.23+0/10/4.3+0.11 ms cpu, 126->126->93 MB, 126 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 48 @0.592s 8%: 0.018+18+0.004 ms clock, 0.14+0/19/0.072+0.038 ms cpu, 157->157->116 MB, 157 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 49 @0.612s 8%: 0.014+20+0.007 ms clock, 0.11+0/10/18+0.063 ms cpu, 197->197->145 MB, 197 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 50 @0.632s 8%: 0.015+23+0.25 ms clock, 0.12+0/11/25+2.0 ms cpu, 246->246->181 MB, 246 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 51 @0.658s 8%: 0.016+31+1.5 ms clock, 0.12+0/25/12+12 ms cpu, 308->308->227 MB, 308 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 52 @0.692s 8%: 0.016+37+2.9 ms clock, 0.13+0/16/40+23 ms cpu, 385->385->284 MB, 385 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 53 @0.739s 9%: 0.015+42+10 ms clock, 0.12+0/22/52+80 ms cpu, 481->481->355 MB, 481 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 54 @0.799s 9%: 0.015+61+3.9 ms clock, 0.12+0/31/55+31 ms cpu, 602->602->444 MB, 602 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 55 @0.868s 9%: 0.016+67+1.0 ms clock, 0.12+0/32/80+8.2 ms cpu, 752->752->555 MB, 752 MB goal, 0 MB stacks, 0 MB globals, 8 P
gc 56 @0.943s 9%: 0.017+13+0.020 ms clock, 0.13+0/26/64+0.16 ms cpu, 940->940->694 MB, 940 MB goal, 0 MB stacks, 0 MB globals, 8 P
make p4 time 498.985458ms 10000000
gc 57 @1.016s 9%: 0.024+0.20+0.019 ms clock, 0.19+0/0.15/0.20+0.15 ms cpu, 1075->1075->381 MB, 1075 MB goal, 0 MB stacks, 0 MB globals, 8 P
make p5 time 37.320375ms 10000000
GOGC=1 GODEBUG=gctrace=1 ./ex3  1.82s user 0.27s system 195% cpu 1.068 total
```
