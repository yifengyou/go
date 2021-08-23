# rune



## 字符串、[]byte 和 []rune关系


```
var name = "游益锋abc"
fmt.Printf("%#v\n", name)
// "游益锋abc"
fmt.Printf("%#v\n", []byte(name))
//[]byte{0xe6, 0xb8, 0xb8, 0xe7, 0x9b, 0x8a, 0xe9, 0x94, 0x8b, 0x61, 0x62, 0x63}
fmt.Printf("%#v\n", []rune(name))
// []int32{28216, 30410, 38155, 97, 98, 99}
fmt.Printf("%s\n", string([]rune(name)[0]))
// 游
fmt.Printf("%s\n", []rune(name)[0])
// %!s(int32=28216)
```

* go中字符串用utf-8编码，属于**变长**unicode编码
* stirng默认下表访问到的与byte访问一致，都是底层字节，类型固定为```type byte = uint8```
* len(string) 与 len([]byte) 一致，但不是实际单次数量。**因为golang中string底层是通过byte数组实现的。所以len一致**
* len([]rune) 为实际有效单词数量
* rune则是据utf-8解析后拆分的字符，类型是```type rune = int32```但实际不一定用那么多
* 只要知道unicode编码，就能string转为实际字符。但rune本身为int32，就只是数值
* rune的骚操作就是能**根据utf8规则划分字符边界**，string不能划分、[]byte也不行，只有rune可以

```
a := rune(28216)
fmt.Println(string(a)) // 游
a = 30410
fmt.Println(string(a)) // 益
a = 38155
fmt.Println(string(a)) // 锋
```


获取有效字符长度的方法有两种：

1. 使用utf8.RuneCountInString
2. 转为[]rune后len

```
//golang中的unicode/utf8包提供了用utf-8获取长度的方法
fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

//通过rune类型处理unicode字符
fmt.Println("rune:", len([]rune(str)))
```

* byte 等同于uint8，常用来处理ascii字符
* rune 等同于int32,常用来处理unicode或utf-8字符
* rune理解为 **一个 可以表示unicode 编码的值int 的值，称为码点（code point）**。只不过go语言把这个码点抽象为rune





---
