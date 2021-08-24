<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [json](#json)   
   - [func Unmarshal(data []byte, v interface{}) error](#func-unmarshaldata-byte-v-interface-error)   
   - [func Marshal(v interface{}) ([]byte, error)](#func-marshalv-interface-byte-error)   
   - [func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)](#func-marshalindentv-interface-prefix-indent-string-byte-error)   
   - [type Encoder](#type-encoder)   
      - [func NewEncoder(w io.Writer) *Encoder](#func-newencoderw-iowriter-encoder)   
      - [func (enc *Encoder) Encode(v interface{}) error](#func-enc-encoder-encodev-interface-error)   
      - [func (enc *Encoder) SetEscapeHTML(on bool)](#func-enc-encoder-setescapehtmlon-bool)   

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





---
