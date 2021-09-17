<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [unsafe](#unsafe)   
   - [type Pointer](#type-pointer)   
   - [type Pointer *ArbitraryType](#type-pointer-arbitrarytype)   
   - [func Sizeof(v ArbitraryType) uintptr](#func-sizeofv-arbitrarytype-uintptr)   
   - [func Alignof(v ArbitraryType) uintptr](#func-alignofv-arbitrarytype-uintptr)   
   - [func Offsetof(v ArbitraryType) uintptr](#func-offsetofv-arbitrarytype-uintptr)   

<!-- /MDTOC -->
# unsafe

```
import "unsafe"
```

* unsafe包提供了一些跳过go语言类型安全限制的操作。
* uintptr是go的内置类型，用于指针运算，其底层基于int类型。
* uintptr不是指针，GC会回收uintptr类型的对象。



## type Pointer

## type Pointer *ArbitraryType

Pointer类型用于表示任意类型的指针。有4个特殊的只能用于Pointer类型的操作：

1) 任意类型的指针可以转换为一个Pointer类型值
2) 一个Pointer类型值可以转换为任意类型的指针
3) 一个uintptr类型值可以转换为一个Pointer类型值
4) 一个Pointer类型值可以转换为一个uintptr类型值
因此，Pointer类型允许程序绕过类型系统读写任意内存。使用它时必须谨慎。


## func Sizeof(v ArbitraryType) uintptr

* Sizeof返回类型v本身数据所占用的字节数。
* 返回值是“顶层”的数据占有的字节数。例如，若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存的大小。
* unsafe.Sizeof接受任意类型的值（表达式），返回其占用的字节数

```
var p float64 = 99
fmt.Println(reflect.TypeOf(unsafe.Sizeof(p))) // uintptr
fmt.Println(unsafe.Sizeof(p))  // 8
```


## func Alignof(v ArbitraryType) uintptr

Alignof返回类型v的对齐方式（即类型v在内存中占用的字节数）；若是结构体类型的字段的形式，它会返回字段f在该结构体中的对齐方式。


## func Offsetof(v ArbitraryType) uintptr

Offsetof返回类型v所代表的结构体字段在结构体中的偏移量，它必须为结构体类型的字段的形式。换句话说，它返回该结构起始处与该字段起始处之间的字节数。

---
