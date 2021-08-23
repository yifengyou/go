# utf8

```
import "unicode/utf8"
```


utf8包实现了对utf-8文本的常用函数和常数的支持，包括rune和utf-8编码byte序列之间互相翻译的函数。

## 常量


```
const (
    RuneError = '\uFFFD'     // 错误的Rune或"Unicode replacement character"
    RuneSelf  = 0x80         // 低于RunSelf的字符用代表单字节的同一值表示
    MaxRune   = '\U0010FFFF' // 最大的合法unicode码值
    UTFMax    = 4            // 最大的utf-8编码的unicode字符的长度
)
```

## func RuneCountInString(s string) (n int)

计算字符串中有效地UTF8字符数量
// RuneCountInString is like RuneCount but its input is a string.

示例：

```
name := "游益锋"
fmt.Printf("%T- %#v\n", name, name)
// string- "游益锋"
fmt.Println(len(name))
// 9
fmt.Println(utf8.RuneCountInString(name))
// 3
fmt.Println(len([]byte(name)))
// 9
fmt.Println(len([]rune(name)))
// 3
```


## func EncodeRune(p []byte, r rune) int


* encode将rune转为[]byte，返回值为处理字节数，只处理一个码点
* EncodeRune将符文的UTF-8编码写入p(必须足够大)。
* 如果符文超出范围，它将写入RuneError的编码。它返回写入的字节数。


示例：

```
r := "游益锋"
fmt.Println([]byte(r))
//	[230 184 184 231 155 138 233 148 139]
fmt.Printf("%T - %#v\n", r, r)
//	string - "游益锋"
result := []byte{}

for _, b := range r {
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, b)
	result = append(result, buf...)
}
fmt.Println(result)
//	[230 184 184 231 155 138 233 148 139]
```





## func DecodeRune(p []byte) (r rune, size int)

* 将[]byte转为rune，只处理一个码点
* DecodeRune在p中解压缩第一个UTF-8编码，并返回符文及其宽度(以字节为单位)。
* 如果p为空，则返回(RuneError，0)。
* 否则，如果编码无效，则返回(RuneError，1)。
* 如果编码不正确，则它是无效的UTF-8，对超出范围或不是该值的最短UTF-8编码的符文进行编码。不执行其他任何验证。

示例：

```
b := []byte("游益锋")
for len(b) > 0 {
	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c %v\n", r, size)
	b = b[size:]
}
```

输出

```
游 3
益 3
锋 3
```














---
