<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [os](#os)   
   - [func Create(name string) (*File, error)](#func-createname-string-file-error)   
   - [func Open(name string) (*File, error)](#func-openname-string-file-error)   
   - [func OpenFile(name string, flag int, perm FileMode) (*File, error)](#func-openfilename-string-flag-int-perm-filemode-file-error)   
   - [func (f *File) Read(b []byte) (n int, err error)](#func-f-file-readb-byte-n-int-err-error)   
   - [func ReadFile(name string) ([]byte, error)](#func-readfilename-string-byte-error)   
   - [func (f *File) write(b []byte) (n int, err error)](#func-f-file-writeb-byte-n-int-err-error)   
   - [func WriteFile(name string, data []byte, perm FileMode) error](#func-writefilename-string-data-byte-perm-filemode-error)   
   - [func (f *File) Close() error](#func-f-file-close-error)   
   - [func IsPermission(err error) bool](#func-ispermissionerr-error-bool)   
   - [参考](#参考)   

<!-- /MDTOC -->

# os




## func Create(name string) (*File, error)

```
func main() {
    newFile, err = os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
}
```

## func Open(name string) (*File, error)
## func OpenFile(name string, flag int, perm FileMode) (*File, error)


```
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}
```






## func (f *File) Read(b []byte) (n int, err error)
## func ReadFile(name string) ([]byte, error)








## func (f *File) write(b []byte) (n int, err error)
## func WriteFile(name string, data []byte, perm FileMode) error

* write是File的方法，WriteFile是独立的函数
* WriteFile底层还是调用Write方法，根据name文件名打开文件，写入data，

```
func WriteFile(name string, data []byte, perm FileMode) error {
	f, err := OpenFile(name, O_WRONLY|O_CREATE|O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
```


## func (f *File) Close() error


## func IsPermission(err error) bool

```
func main() {
    // 这个例子测试写权限，如果没有写权限则返回error。
    // 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
    if err != nil {
        if os.IsPermission(err) {
            log.Println("Error: Write permission denied.")
        }
    }
    file.Close()
    // 测试读权限
    file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
    if err != nil {
        if os.IsPermission(err) {
            log.Println("Error: Read permission denied.")
        }
    }
    file.Close()
}
```












## func Stat(name string) (FileInfo, error)
## func IsNotExist(err error) bool














## 参考

* <https://colobu.com/2016/10/12/go-file-operations/#%E5%88%9B%E5%BB%BA%E7%A9%BA%E6%96%87%E4%BB%B6>


---
