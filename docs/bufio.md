<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [bufio](#bufio)   
   - [type Reader](#type-reader)   
      - [func NewReader(rd io.Reader) *Reader](#func-newreaderrd-ioreader-reader)   
   - [type Writer](#type-writer)   
      - [func NewWriter(w io.Writer) *Writer](#func-newwriterw-iowriter-writer)   
      - [func (b *Writer) Flush() error](#func-b-writer-flush-error)   

<!-- /MDTOC -->

# bufio

## type Reader

```
// Reader implements buffering for an io.Reader object.
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}
```


### func NewReader(rd io.Reader) *Reader

* NewReader创建一个具有默认大小缓冲、从r读取的*Reader



```
func NewReaderSize(rd io.Reader, size int) *Reader
func (b *Reader) Reset(r io.Reader)
func (b *Reader) Buffered() int
func (b *Reader) Peek(n int) ([]byte, error)
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) ReadByte() (c byte, err error)
func (b *Reader) UnreadByte() error
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) UnreadRune() error
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
func (b *Reader) ReadString(delim byte) (line string, err error)
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
```


## type Writer

```
// Writer implements buffering for an io.Writer object.
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}
```

### func NewWriter(w io.Writer) *Writer

* NewWriter创建一个具有默认大小缓冲、写入w的*Writer

### func (b *Writer) Flush() error

* 先写入缓存后Flush才是正在底层调用Write写入

```
func NewWriter(w io.Writer) *Writer
func NewWriterSize(w io.Writer, size int) *Writer
func (b *Writer) Reset(w io.Writer)
func (b *Writer) Buffered() int
func (b *Writer) Available() int
func (b *Writer) Write(p []byte) (nn int, err error)
func (b *Writer) WriteString(s string) (int, error)
func (b *Writer) WriteByte(c byte) error
func (b *Writer) WriteRune(r rune) (size int, err error)
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
```

---
