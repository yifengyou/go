<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [fmt](#fmt)   
   - [func Print(a ...interface{}) (n int, err error)](#func-printa-interface-n-int-err-error)   
   - [func Println(a ...interface{}) (n int, err error)](#func-printlna-interface-n-int-err-error)   
   - [func Printf(format string, a ...interface{}) (n int, err error)](#func-printfformat-string-a-interface-n-int-err-error)   
   - [func Sprint(a ...interface{}) string](#func-sprinta-interface-string)   
   - [func Sprintln(a ...interface{}) string](#func-sprintlna-interface-string)   
   - [func Sprintf(format string, a ...interface{}) string](#func-sprintfformat-string-a-interface-string)   
   - [func Fprint(w io.Writer, a ...interface{}) (n int, err error)](#func-fprintw-iowriter-a-interface-n-int-err-error)   
   - [func Fprintln(w io.Writer, a ...interface{}) (n int, err error)](#func-fprintlnw-iowriter-a-interface-n-int-err-error)   
   - [func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)](#func-fprintfw-iowriter-format-string-a-interface-n-int-err-error)   

<!-- /MDTOC -->
# fmt

* 区别于内置函数print、println

## func Print(a ...interface{}) (n int, err error)
## func Println(a ...interface{}) (n int, err error)

```
  向io.Stdout打印a切片内容，各元素用空格分隔开来
  a : 目标切片
  返回写入字节数，错误
```

* 打印a切片内容，多个参数空格分开
* 末尾自带换行
* 定义简单，直接列出
* Print与Println区别有两个，一个是结尾换行，另外**ln还会在参数中间加空格**
* 底层调用在doPrint、doPrintln有所差异

```
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}
```

用例：

```
fmt.Fprintln(os.Stdout, "abc", "def")  # abc def
fmt.Println("a", "b", "c") # a b c
```


## func Printf(format string, a ...interface{}) (n int, err error)

```
  根据format，格式化a切片内容
  format string : 格式化方式
  a : 目标切片
  返回写入字节数，错误
```

* format string参数必须有，不能生路。a切片参数可以省略
  - ```fmt.Printf("This is print content\n")
* 与C语言printf兼容，**需要自行处理换行**
* 最麻烦的就是格式化标记，都是中规中矩，记住常用的


| 格式化标记 |	格式化含义 |
| ----  | ----  |
| %v    | 值的默认格式表示  |
| %+v   | 值的默认格式表示 + 字段名称（仅有字段时显示，一般为struct） |
| %#v   | 值的默认格式表示 + 字段名称（仅有字段时显示，一般为struct） + go类型标识 |
| %%    | 百分号  |
| %T    | 输出对应的类型  |
| %t	  | 单词true或false，**常规记忆%b是bool，然而%b是二进制表示**  |
| %b	| 表示为二进制 |
| %d	| 表示为十进制 |
| %o	| 表示为八进制 |
| %X	| 表示为十六进制，使用A-F|
| %x	| 表示为十六进制，使用a-f|
| %c	| 该值对应的unicode码值|
| %p	| 十六进制表示，前缀 0x, 不存在%P|
| %s  | 	直接输出字符串或者[]byte|

罕见操作：

| 格式化标记 |	格式化含义 |
| ----  | ----  |
| %w    | 包裹error类型，Go1.13新增  |

实例：

```
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "nic", Age: 35}

	fmt.Printf("p %%v is %v\n", p)   // p %v is   {nic 35}
	fmt.Printf("p %%+v is %+v\n", p) // p %+v is  {Name:nic Age:35}
	fmt.Printf("p %%#v is %#v\n", p) // p %#v is  main.Person{Name:"nic", Age:35}
	fmt.Printf("p type is %T\n", p)  // p type is main.Person

	a := 10
	fmt.Printf("a %%#v is %#v, a type is %T\n", a, a) // a %#v is 10, a type is int

  s := "hello world"
	fmt.Printf("s %%#v is %#v, s type is %T\n", s, s) // s %#v is "hello world", s type is string

  c := []int{1, 2, 3, 4}
	fmt.Printf("c %%v is %v, c type is %T\n", c, c)   // c %v is [1 2 3 4], c type is []int
	fmt.Printf("c %%#v is %#v, c type is %T\n", c, c) // c %#v is []int{1, 2, 3, 4}, c type is []int

  var b bool
	fmt.Printf("b %%v is %v, %%#v is %#v, %%T is %T\n", b, b, b) 	// b %v is false, %#v is false, %T is bool

  var i int = 15
	fmt.Printf("i %%#v is %#v\n", i)  // i %#v is 15
	fmt.Printf("i %%b is %b\n", i)    // i %b is 1111
	fmt.Printf("i %%c is %c\n", i)    // i %c is    // 不可见字符
	fmt.Printf("i %%d is %d\n", i)    // i %d is 15
	fmt.Printf("i %%o is %o\n", i)    // i %o is 17
	fmt.Printf("i %%q is %q\n", i)    // i %q is '\x0f'
	fmt.Printf("i %%x is %x\n", i)    // i %x is f
	fmt.Printf("i %%X is %X\n", i)    // i %X is F
	fmt.Printf("i %%U is %U\n", i)    // i %U is U+000F

	var pointer string = "abc"
	fmt.Printf("pointer %%p is %p,  %%P is %P\n", &pointer, &pointer)
	// pointer %p is 0xc000044260,  %P is %!P(*string=0xc000044260)
}
```


## func Sprint(a ...interface{}) string
## func Sprintln(a ...interface{}) string
## func Sprintf(format string, a ...interface{}) string

```
  根据format，格式化a切片内容，返回字符串
  与Printf区别在于返回字符串，而Printf输出到io.Writer返回输出字符数以及错误
  format string : 格式化方式
  a : 目标切片
  返回字符串
```

* 用于多个数据类型根据format规则拼接转为字符串
* 返回字符串，不会写入io.Writer
* 格式化规则与fmt.Printf一致
* Sprint相比Sprintln，少了 末尾换行 + 字符串间空格
* Sprintln是加空格是仅在前后类型都是字符串的时候加，否则不加，代码如下：

```
		// Add a space between two non-string arguments.
		if argNum > 0 && !isString && !prevString {
			p.buf.writeByte(' ')
		}
```




示例：

```
func main() {
  st1 := fmt.Sprintf("%d + %d", 1, 2)
  fmt.Println(st1)
}
```

```
func main() {
	v := 123
	fmt.Println(fmt.Sprintln("abc", v, "def"), "end")
	/*
	abc 123 def
	end
	 */
	fmt.Println(fmt.Sprint("abc", v, "def"), "end")
	// abc123def end
}
```

## func Fprint(w io.Writer, a ...interface{}) (n int, err error)
## func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
## func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

Fprintf 根据 format 参数生成格式化的字符串并写入 w 。返回写入的字节数和遇到的任何错误。

示例：

输出到终端

```
	fmt.Fprintf(os.Stdout, "abc") // abc
	fmt.Fprintf(os.Stderr, "def") // abc def
```


```
	v := 123
	fmt.Fprint(os.Stdout,"abc", v, "end") // abc123end
	fmt.Fprintln(os.Stdout,"abc", v, "end")
	/*
	abc 123 end

	 */
```

输出到文件

```
file, _ := os.OpenFile("./xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
//用逻辑或操作"|"实现不同的组合
defer file.Close()
fmt.Fprintf(file, "hahahhahah\n")
```


其中io.Writer是个接口，实现了Write方法就行

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

## func Errorf(format string, a ...interface{}) error

* Errorf函数根据format参数生成格式化字符串并返回一个包含该error错误类型（根本就是字符串）

```
// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
```

```
func Sprintf(format string, a ...interface{}) string
```

* Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error
* 通常使用这种方式来自定义错误类型，例如：

```
err := fmt.Errorf("这是一个错误")
Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
e := errors.New("原始错误e")
w := fmt.Errorf("Wrap了一个错误%w", e)
```




## 参考

* <https://zhuanlan.zhihu.com/p/111204359>


---
