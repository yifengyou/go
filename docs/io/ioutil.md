# ioutil

```
import "io/ioutil"
```
* 实现一些IO操作的工具函数


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







---
