## FindUnspentTransactions
圖解
```shell
                              0(spend)
Tx0 Gensis: -> 10:A VOuts:{[10, A]}

                         0       1(spend)
Tx1 A:10 -> 3:B VOuts:{[3, B], [7, A]}

                        0       1
Tx2 A:7 -> 2:B VOuts:{[2, B], [5, A]}
```
此時
```go
$ balance -a A
Balance of 'A': 5 // only Tx2 VOuts [5, A] is not spend
$ balance -a B
Balance of 'B': 5 // 3+2=5
```
再加一筆
```shell
                              0(spend)
Tx0 Gensis: -> 10:A VOuts:{[10, A]}

                         0(spend)  1(spend)
Tx1 A:10 -> 3:B VOuts:{[3, B], [7, A]}

                        0(spend)  1
Tx2 A:7 -> 2:B VOuts:{[2, B], [5, A]}
                        0       1
Tx3 B:5 -> 4:C VOuts:{[4, C], [1, B]}
```
此時
```go
$ balance -a B
Balance of 'B': 1 // only Tx3 VOuts [1, B] is not spend
$ balance -a C
Balance of 'C': 4
```