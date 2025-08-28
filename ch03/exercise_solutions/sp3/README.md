# 考察
validatorのサンプル
validatorは外部公開するもの(一が大文字)のものしかチェックしない
```
{1 Noboru Kitajima}
Validate safe
{0 hogehogehogehogehogehoge Kitajima}
ID gt 0
FirstName max 16
{0 Noboru }
ID gt 0
LastName required 
{0  }
Validate safe
```
