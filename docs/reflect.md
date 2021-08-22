# reflect

* reflect包实现了运行时反射，允许程序操作任意类型的对象。
* 典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。
* 调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。
* Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。

参见"The Laws of Reflection"获取go反射的介绍：http://golang.org/doc/articles/laws_of_reflection.html




## type Type

* 获取Type类型，则类型信息基本都能获取到

```
type Type interface {
    // Kind返回该接口的具体分类
    Kind() Kind
    // Name返回该类型在自身包内的类型名，如果是**未命名类型会返回空字符串**
    Name() string
    // PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
    // 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
    PkgPath() string
    // 返回类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
    // 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较。
    String() string
    // 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
    Size() uintptr
    // 返回当从内存中申请一个该类型值时，会对齐的字节数
    Align() int
    // 返回当该类型作为结构体的字段时，会对齐的字节数
    FieldAlign() int
    // 如果该类型实现了u代表的接口，会返回真
    Implements(u Type) bool
    // 如果该类型的值可以直接赋值给u代表的类型，返回真
    AssignableTo(u Type) bool
    // 如该类型的值可以转换为u代表的类型，返回真
    ConvertibleTo(u Type) bool
    // 返回该类型的字位数。如果该类型的Kind不是Int、Uint、Float或Complex，会panic
    Bits() int
    // 返回array类型的长度，如非数组类型将panic
    Len() int
    // 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
    Elem() Type
    // 返回map类型的键的类型。如非映射类型将panic
    Key() Type
    // 返回一个channel类型的方向，如非通道类型将会panic
    ChanDir() ChanDir
    // 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
    NumField() int
    // 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
    Field(i int) StructField
    // 返回索引序列指定的嵌套字段的类型，
    // 等价于用索引中每个值链式调用本方法，如非结构体将会panic
    FieldByIndex(index []int) StructField
    // 返回该类型名为name的字段（会查找匿名字段及其子字段），
    // 布尔值说明是否找到，如非结构体将panic
    FieldByName(name string) (StructField, bool)
    // 返回该类型第一个字段名满足函数match的字段，布尔值说明是否找到，如非结构体将会panic
    FieldByNameFunc(match func(string) bool) (StructField, bool)
    // 如果函数类型的最后一个输入参数是"..."形式的参数，IsVariadic返回真
    // 如果这样，t.In(t.NumIn() - 1)返回参数的隐式的实际类型（声明类型的切片）
    // 如非函数类型将panic
    IsVariadic() bool
    // 返回func类型的参数个数，如果不是函数，将会panic
    NumIn() int
    // 返回func类型的第i个参数的类型，如非函数或者i不在[0, NumIn())内将会panic
    In(i int) Type
    // 返回func类型的返回值个数，如果不是函数，将会panic
    NumOut() int
    // 返回func类型的第i个返回值的类型，如非函数或者i不在[0, NumOut())内将会panic
    Out(i int) Type
    // 返回该类型的方法集中方法的数目
    // 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
    // 匿名字段导致的歧义方法会滤除
    NumMethod() int
    // 返回该类型方法集中的第i个方法，i不在[0, NumMethod())范围内时，将导致panic
    // 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
    // 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
    Method(int) Method
    // 根据方法名返回该类型方法集中的方法，使用一个布尔值说明是否发现该方法
    // 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
    // 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
    MethodByName(string) (Method, bool)
    // 内含隐藏或非导出方法
}
```

```
func TypeOf(i interface{}) Type
func PtrTo(t Type) Type
func SliceOf(t Type) Type
func MapOf(key, elem Type) Type
func ChanOf(dir ChanDir, t Type) Type
```

### func TypeOf(i interface{}) Type

* TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。
* 参数是空接口，所以所有类型都可以作为参数，获取到Type之后，根据Type的方法，获取信息，进一步处理


示例：

```
	fmt.Println(reflect.TypeOf(func(a, b string) {}).Name(), " == ",
		reflect.TypeOf(func(a, b, c, d, e, f string) {}).Kind()) //     ==  func

	i := 123
	fmt.Println(reflect.TypeOf(i).Name(),
		reflect.TypeOf(i).Kind()) //   int int
	s := "abc"
	fmt.Println(reflect.TypeOf(s).Name(),
		reflect.TypeOf(s).Kind()) //   int int
	f := 1.2
	fmt.Println(reflect.TypeOf(f).Name(),
		reflect.TypeOf(f).Kind()) //   float64 float64
	type demo_struct struct {
		name string
		age  uint32
	}
	st := demo_struct{
		"nicyou", 28,
	}
	fmt.Println(reflect.TypeOf(st).Name(),
		reflect.TypeOf(st).Kind()) //   demo_struct struct
	m := make(map[int]string, 20)
	fmt.Println(reflect.TypeOf(m).Name(), " == ",
		reflect.TypeOf(m).Kind()) //     ==  map
```








---
