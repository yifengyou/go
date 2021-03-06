<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [strconv](#strconv)   
   - [func ParseInt(s string, base int, bitSize int) (i int64, err error)](#func-parseints-string-base-int-bitsize-int-i-int64-err-error)   
   - [func ParseBool(str string) (bool, error)](#func-parseboolstr-string-bool-error)   
   - [func Atoi(s string) (i int, err error)](#func-atois-string-i-int-err-error)   

<!-- /MDTOC -->

# strconv

```
import "strconv"
```

strconv包实现了基本数据类型和 字符串表示 的相互转换。

## func ParseInt(s string, base int, bitSize int) (i int64, err error)

* 返回字符串表示的整数值，接受正负号。
* base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
* bitSize指定结果必须能无溢出赋值的整数类型，值为：0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。


## func ParseBool(str string) (bool, error)


* 返回字符串表示的bool值。
* 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。



## func Atoi(s string) (i int, err error)

* Atoi是ParseInt(s, 10, 0)的简写
* 支持正负数
* 最好判断返回值的error

实例：

```
s := "abc"
value, err := strconv.Atoi(s)
if err != nil {
	fmt.Println("convert err", err)
	return
}
fmt.Println(value)
```

输出：

```
convert err strconv.Atoi: parsing "abc": invalid syntax
```

```
s := "-123"
value, err := strconv.Atoi(s)
if err != nil {
	fmt.Println("convert err", err)
	return
}
fmt.Println(value) // 输出-123
```

---
