# binary

```
import "encoding/binary"
```

* binary包实现了简单的数字与字节序列的转换以及变长值的编解码。
* 数字翻译为定长值来读写，一个定长值，要么是固定长度的数字类型（int8, uint8, int16, float32, complex64, ...）或者只包含定长值的结构体或者数组。
* 变长值是使用一到多个字节编码整数的方法，绝对值较小的数字会占用较少的字节数。详情请参见：http://code.google.com/apis/protocolbuffers/docs/encoding.html。
* 本包**相对于效率更注重简单**。如果需要高效的序列化，特别是数据结构较复杂的，请参见更高级的解决方法，例如encoding/gob包，或者采用协议缓存。





## func Read(r io.Reader, order ByteOrder, data interface{}) error


* 从r中读取binary编码的数据并赋给data，data必须是一个指向定长值的指针或者定长值的切片。
* 从r读取的字节使用order指定的字节序解码并写入data的字段里当写入结构体是，名字中有'_'的字段会被跳过，这些字段可用于填充（内存空间）。


## func Write(w io.Writer, order ByteOrder, data interface{}) error


* 将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针。
* order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。


## 参考

* <https://studygolang.com/pkgdoc>




---
