## 汇编分享

**本仓库全程参考 <https://github.com/cch123/asmshare> 内容**

只为自身学习，不要打我~

```
Something I hope you know before go into the coding~
* First, please watch or star this repo, I'll be more happy if you follow me.
* Bug report, questions and discussion are welcome, you can post an issue or pull a request.
```

* [汇编 is so easy](docs/layout.md)

## 参考

x86 体系的汇编、调用规约

https://github.com/0xAX/asm

Go plan9 汇编官方文档(很不详细)

https://golang.org/doc/asm

补充文档，作者是原来 intel 的工程师

https://quasilyte.dev/blog/post/go-asm-complementary-reference/#external-resources

法国的一个系统工程师的汇编教程，之前有一些疏漏，不过已经被网友(其中包括我(笑)纠正了

https://github.com/teh-cmc/go-internals/tree/master/chapter1_assembly_primer

我在学习汇编的时间总结的笔记：

https://github.com/cch123/golang-notes/blob/master/assembly.md

柴总和我编写的 《Go 高级编程》中的汇编章节：

https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/readme.md

本书将于 19 年四月由人民邮电出版社出版，到时候希望大家支持！

## faq

Q: 为什么要学习汇编

Q: 有人说 new(xxxStruct) 比 &xxxStruct{} 牛逼，因为前者只生成一个堆上变量。后者要生成变量，再做引用，所以效率低

那么，我们怎么样确定 Go 这两种写法到底是一样，还是不一样呢？

Q: 他们老是说这个东西在 runtime 里对应 xxx 函数，我怎么知道他们没骗我呢？

对编译原理不太熟，也不太想看 Go 的编译器代码，代码量有点大(50w+)。想节省点时间。

Q: 什么是寄存器

Q: 什么是函数栈

Q: 为什么函数调用会 Stack Overflow

Q: 为什么 Go 可以实现函数多返回值

Q: 一个参数，有的人说是值传递，有的人说是引用传递，我到底该相信谁的

Q: 为什么要有伪寄存器的概念

Q: 怎么判断我看到的 plan9 汇编代码中的寄存器是物理寄存器，还是伪寄存器？

Q: plan9 汇编和 intel 汇编有啥异同？

Q: 汇编结果中的 gobuf_sp() 表示的是别名，还是特殊含义

Q: 我看社区里有好多人搞 goroutine id 呀，我怎么能实现这么一个工具呢

Q: SIMD 指令要怎么用
