<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [ioutil](#ioutil)   
   - [var Discard io.Writer = devNull(0)](#var-discard-iowriter-devnull0)   
   - [func ReadFile(filename string) ([]byte, error)](#func-readfilefilename-string-byte-error)   
   - [func WriteFile(filename string, data []byte, perm os.FileMode) error](#func-writefilefilename-string-data-byte-perm-osfilemode-error)   
   - [func NopCloser(r io.Reader) io.ReadCloser](#func-nopcloserr-ioreader-ioreadcloser)   
   - [func ReadAll(r io.Reader) ([]byte, error)](#func-readallr-ioreader-byte-error)   
   - [func ReadDir(dirname string) ([]os.FileInfo, error)](#func-readdirdirname-string-osfileinfo-error)   
   - [func TempDir(dir, prefix string) (name string, err error)](#func-tempdirdir-prefix-string-name-string-err-error)   
   - [func TempFile(dir, prefix string) (f *os.File, err error)](#func-tempfiledir-prefix-string-f-osfile-err-error)   

<!-- /MDTOC -->
# ioutil

```
import "io/ioutil"
```
* 实现一些IO操作的工具函数

## var Discard io.Writer = devNull(0)

* Discard是一个io.Writer接口，对它的所有Write调用都会无实际操作的成功返回

## func ReadFile(filename string) ([]byte, error)
## func WriteFile(filename string, data []byte, perm os.FileMode) error

```
// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
//
// As of Go 1.16, this function simply calls os.ReadFile.
```

* ReadFile **读取文件的所有内容**
* ReadFile 从filename指定的文件中读取数据并返回文件的内容。
* 成功的调用返回的err为nil而非EOF。因为本函数定义为**读取整个文件**，它不会将读取返回的EOF视为应报告的错误。
* WriteFile函数向filename指定的文件中写入数据。
* 如果文件**不存在将按给出的权限创建文件**，
* 如果文件**存在则在写入数据之前清空文件**。


## func NopCloser(r io.Reader) io.ReadCloser
## func ReadAll(r io.Reader) ([]byte, error)
## func ReadDir(dirname string) ([]os.FileInfo, error)
## func TempDir(dir, prefix string) (name string, err error)
## func TempFile(dir, prefix string) (f *os.File, err error)




---
