# io


## func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

* ReadAtLeast从r至少读取min字节数据填充进buf。
* 函数返回写入的字节数和错误（如果没有读取足够的字节）。
* 只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。
* 如果min比buf的长度还大，函数会返回ErrShortBuffer。只有返回值err为nil时，返回值n才会不小于min。

## func ReadFull(r Reader, buf []byte) (n int, err error)

* ReadFull从r精确地读取len(buf)字节数据填充进buf。
* 函数返回写入的字节数和错误（如果没有读取足够的字节）。
* 只有没有读取到字节时才可能返回EOF；
* 如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。 只有返回值err为nil时，返回值n才会等于len(buf)。
* 其中Reader为接口，只要实现了Read方法都能作为Reader

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```






## 参考

* <https://blog.csdn.net/chenxun_2010/article/details/73896502>



















---
