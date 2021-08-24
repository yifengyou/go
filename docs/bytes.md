<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [bytes](#bytes)   
   - [type Buffer](#type-buffer)   
      - [func (b *Buffer) Bytes() []byte](#func-b-buffer-bytes-byte)   
      - [func (b *Buffer) String() string](#func-b-buffer-string-string)   
      - [func NewBuffer(buf []byte) *Buffer](#func-newbufferbuf-byte-buffer)   
      - [func NewBufferString(s string) *Buffer](#func-newbufferstrings-string-buffer)   
      - [func (b *Buffer) String() string](#func-b-buffer-string-string)   
      - [func (b *Buffer) ReadByte() (c byte, err error)](#func-b-buffer-readbyte-c-byte-err-error)   
      - [func (b *Buffer) WriteByte(c byte) error](#func-b-buffer-writebytec-byte-error)   
      - [func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)](#func-b-buffer-readfromr-ioreader-n-int64-err-error)   
      - [func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)](#func-b-buffer-writetow-iowriter-n-int64-err-error)   
   - [type Reader](#type-reader)   
      - [func NewReader(b []byte) *Reader](#func-newreaderb-byte-reader)   
   - [参考](#参考)   

<!-- /MDTOC -->




# bytes

```
import "bytes"
```

* bytes包实现了操作[]byte的常用函数。
* 函数和strings包的函数相当类似。



## type Buffer

```
type Buffer struct {
  type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}
}
```

* Buffer 是一个实现了读写方法的可变大小的字节缓冲。
* Buffer 类型的零值是一个空的可用于读写的缓冲。
* 从 bytes.Buffer 读取数据后，**被成功读取的数据仍保留在原缓冲区**，只是无法被使用(坑)，因为缓冲区的可见数据从偏移 off 开始，即buf[off : len(buf)]

### func (b *Buffer) Bytes() []byte

* 返回未读取部分字节数据的切片，len(b.Bytes()) == b.Len()。
* 如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容

### func (b *Buffer) String() string

* 将未读取部分的字节数据作为字符串返回
* 如果b是nil指针，会返回"<nil>"。











### func NewBuffer(buf []byte) *Buffer

```
// NewBuffer creates and initializes a new Buffer using buf as its
// initial contents. The new Buffer takes ownership of buf, and the
// caller should not use buf after this call. NewBuffer is intended to
// prepare a Buffer to read existing data. It can also be used to set
// the initial size of the internal buffer for writing. To do that,
// buf should have the desired capacity but a length of zero.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.
func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }
```

* NewBuffer使用buf作为初始内容创建并初始化一个Buffer。
* 本函数用于创建一个用于读取已存在数据的buffer；
* 也用于指定用于写入的内部缓冲的大小，此时，buf应为一个具有指定容量但长度为0的切片。
* buf会被作为返回值的底层缓冲切片。

大多数情况下，new(Buffer)（或只是声明一个Buffer类型变量）就足以初始化一个Buffer了。



### func NewBufferString(s string) *Buffer

```
// NewBufferString creates and initializes a new Buffer using string s as its
// initial contents. It is intended to prepare a buffer to read an existing
// string.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.
func NewBufferString(s string) *Buffer {
	return &Buffer{buf: []byte(s)}
}
```

func (b *Buffer) Reset()
func (b *Buffer) Len() int
func (b *Buffer) Bytes() []byte

### func (b *Buffer) String() string



```
func (b *Buffer) Truncate(n int)
func (b *Buffer) Grow(n int)
func (b *Buffer) Read(p []byte) (n int, err error)
func (b *Buffer) Next(n int) []byte
func (b *Buffer) UnreadByte() error
func (b *Buffer) ReadRune() (r rune, size int, err error)
func (b *Buffer) UnreadRune() error
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
func (b *Buffer) ReadString(delim byte) (line string, err error)
func (b *Buffer) Write(p []byte) (n int, err error)
func (b *Buffer) WriteString(s string) (n int, err error)
func (b *Buffer) WriteRune(r rune) (n int, err error)
```


### func (b *Buffer) ReadByte() (c byte, err error)
### func (b *Buffer) WriteByte(c byte) error

* WriteByte将字节c写入缓冲中，如必要会增加缓冲容量，返回值总是nil，但仍保留以匹配bufio.Writer的WriteByte方法。
* 如果缓冲太大，WriteByte会采用错误值ErrTooLarge引发panic。
* ReadBytes读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的切片。
* 如果ReadBytes方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
* 当且仅当ReadBytes方法返回的切片不以delim结尾时，会返回一个非nil的错误。




### func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)


示例：

```
func main() {
    file, _ := os.Open("./test.txt")    
    buf := bytes.NewBufferString("Hello bytes ")    
    buf.ReadFrom(file)              //将text.txt内容追加到缓冲器的尾部    
    fmt.Println(buf.String())
}
```



### func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)










## type Reader

### func NewReader(b []byte) *Reader

```
func (r *Reader) Len() int
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadByte() (b byte, err error)
func (r *Reader) UnreadByte() error
func (r *Reader) ReadRune() (ch rune, size int, err error)
func (r *Reader) UnreadRune() error
func (r *Reader) Seek(offset int64, whence int) (int64, error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```






## 参考

* <https://dablelv.blog.csdn.net/article/details/94456479>
