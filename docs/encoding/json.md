<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [json](#json)   
   - [func Unmarshal(data []byte, v interface{}) error](#func-unmarshaldata-byte-v-interface-error)   
   - [func Marshal(v interface{}) ([]byte, error)](#func-marshalv-interface-byte-error)   
   - [func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)](#func-marshalindentv-interface-prefix-indent-string-byte-error)   
   - [type Encoder](#type-encoder)   
      - [func NewEncoder(w io.Writer) *Encoder](#func-newencoderw-iowriter-encoder)   
      - [func (enc *Encoder) Encode(v interface{}) error](#func-enc-encoder-encodev-interface-error)   
      - [func (enc *Encoder) SetEscapeHTML(on bool)](#func-enc-encoder-setescapehtmlon-bool)   
   - [func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error](#func-indentdst-bytesbuffer-src-byte-prefix-indent-string-error)   

<!-- /MDTOC -->

# json

* json包实现了json对象的编解码，参见RFC 4627
* Json对象和go类型的映射关系请参见Marshal和Unmarshal函数的文档


## func Unmarshal(data []byte, v interface{}) error

将data反序列化存储到v对象，返回error

实例：

```
func main() {
	var jsonBlob = []byte ( ` [
        { "Name" : "nicyou" , "Age" : 28, "Company" : "ctyun" } ,
        { "Name" : "jackzhao" , "Age" : 30, "Company" : "Tencent"}
    ] ` )
  // 多余的Company不影响。如果数据少了，会在反序列化过程中用类型零值替代，
  // 所以一般建议不省略，除非你很清楚自己在干啥
	type Person struct {
		Name string
		Age  uint
	}
	var person []Person
	err := json.Unmarshal(jsonBlob, &person)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", person) // 输出 [{nicyou 28} {jackzhao 30}]
}
```

```
func main() {
	var jsonBlob = []byte ( ` ` ) // 需要提供json规范的切片，常规错误就是不规范
	type Person struct {
		Name string
		Age  uint
	}
	var person []Person
	err := json.Unmarshal(jsonBlob, &person)
	if err != nil {
		fmt.Println("error:", err) // error: unexpected end of JSON input
		return
	}
	fmt.Printf("%v", person)
}
```

```
func main() {
	var jsonBlob = []byte ( `[{"naMe":"123"}]` ) // 反序列化与字段大小写无关
	type Person struct {
		Name string
		Age  uint
	}
	var person []Person
	err := json.Unmarshal(jsonBlob, &person)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%v", person) // 输出 [{123 0}]
}
```

```
func main() {
	var jsonBlob = []byte ( `[{"aGe":"123"}]` )
  // 字符串不会自动转为uint类型
	type Person struct {
		Name string
		Age  uint
	}
	var person []Person
	err := json.Unmarshal(jsonBlob, &person)
	if err != nil {
		fmt.Println("error:", err)
    // error: json: cannot unmarshal string into Go struct field Person.Age of type uint
		return
	}
	fmt.Printf("%v", person)
}
```


## func Marshal(v interface{}) ([]byte, error)
## func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

```
func main ( ) {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors [ ] string
    }
    group := ColorGroup {
        ID :     1 ,
        Name :   "Reds" ,
        Colors : [ ] string { "Crimson" , "Red" , "Ruby" , "Maroon" } ,
    }
    b , err := json. Marshal ( group )
    if err != nil {
        fmt. Println ( "error:" , err )
    }
    os. Stdout . Write ( b )
}
```


## type Encoder

```
// An Encoder writes JSON values to an output stream.
type Encoder struct {
	w          io.Writer
	err        error
	escapeHTML bool

	indentBuf    *bytes.Buffer
	indentPrefix string
	indentValue  string
}
```


### func NewEncoder(w io.Writer) *Encoder

NewEncoder创建一个将数据写入w的*Encoder。

### func (enc *Encoder) Encode(v interface{}) error

Encode将v的json编码写入输出流，并会写入一个换行符，参见Marshal函数的文档获取细节信息

### func (enc *Encoder) SetEscapeHTML(on bool)

* 默认情况下，使用json.Marshal，会转换html标签。但在某些情况下需要关闭
* 修改Encoder结构体中的escapeHTML标志位
* json.Marshal 默认 escapeHtml 为true,会转义 <、>、&


```
// SetEscapeHTML specifies whether problematic HTML characters
// should be escaped inside JSON quoted strings.
// The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
// to avoid certain safety problems that can arise when embedding JSON in HTML.
//
// In non-HTML settings where the escaping interferes with the readability
// of the output, SetEscapeHTML(false) disables this behavior.
func (enc *Encoder) SetEscapeHTML(on bool) {
	enc.escapeHTML = on
}
```

实例：

* json.Marshal默认会转义<、>、&

```
Content := "http://tencent.com?id=123&test=1&str=<abc>"
jsonByte, _ := json.Marshal(Content)
fmt.Println(string(jsonByte))
// "http://tencent.com?id=123\u0026test=1\u0026str=\u003cabc\u003e"
```

可以用字符串替换处理

```
content = strings.Replace(content, "\\u003c", "<", -1)
content = strings.Replace(content, "\\u003e", ">", -1)
content = strings.Replace(content, "\\u0026", "&", -1)
```

使用json.Encoder

```
Content := "http://tencent.com?id=123&test=1&str=<abc>"
bf := bytes.NewBuffer([]byte{})
jsonEncoder := json.NewEncoder(bf)
jsonEncoder.SetEscapeHTML(false)
jsonEncoder.Encode(Content)
fmt.Println(bf.String())
// "http://tencent.com?id=123&test=1&str=<abc>"
```

## func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error

* Indent函数将json编码的调整缩进之后写入dst。
* 每一个json元素/数组都另起一行开始，以prefix为起始，一或多个indent缩进（数目看嵌套层数）。
* prefix一般为""
* indent一般为"\t"
* 写入dst的数据起始没有prefix字符，也没有indent字符，最后也不换行，因此可以更好的嵌入其他格式化后的json数据里。

## func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

* MarshalIndent类似Marshal但会使用缩进将输出格式化。
* prefix一般为""
* indent一般为"\t"




---
