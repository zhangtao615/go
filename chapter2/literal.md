字面量是值的一种标记法。

### 常用字变量

1. 用于表示基础数据类型的值的各种字变量，例如表示浮点类型的值12E-3
2. 用于构造各种自定义的复合数据类型的类型字面量。
```
type Name struct {
  Forename  String
  Surname   String
}
```
3. 用于表示复合数据类型的值的复合自变量，可以用来构造struct、array、slice、map类型的值。
```
Name{ Forename: "Robert", Surname: "Hao" }
```
