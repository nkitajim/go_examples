# 考察
pointerなのでlist, mapのshallow copyは注意が必要
deep copyは重い
```
% ./sp2     
-----------
li: [1 2 3 4 5 6 7 8 9]
msi: map[a:1 b:2 c:3 d:4 e:5 f:6 g:7 h:8 i:9]
ssi: {1 2 3 4 5 6 7 8 9}

li: [999 2 3 4 5 6 7 8 9]
msi: map[a:999 b:2 c:3 d:4 e:5 f:6 g:7 h:8 i:9]
ssi: {1 2 3 4 5 6 7 8 9}
li1: [999 2 3 4 5 6 7 8 9]
msi1: map[a:999 b:2 c:3 d:4 e:5 f:6 g:7 h:8 i:9]
ssi1: {999 2 3 4 5 6 7 8 9}
-----------
shallow copy list time: 4.923917ms
shallow copy map time: 4.292625ms
copy struct time: 6.373417ms
-----------
deep copy list time: 27.824375ms
deep copy map time: 1.379906834s
copy struct time: 2.49225ms
```
