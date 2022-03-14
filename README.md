# go

常用go代码片段

---

## 目录

* [go命令](docs/go命令.md)
* [常用内置函数](docs/常用内置函数.md)
* [基本数据类型](docs/基本数据类型.md)
    * [string](docs/基本数据类型/string.md)
    * [map](docs/基本数据类型/map.md)
    * [rune](docs/基本数据类型/rune.md)
* [基本写法](docs/基本写法.md)
    * [range范围](docs/基本写法/range范围.md)
* [archive](docs/archive.md)
    * [tar](docs/archive/tar.md)
    * [zip](docs/archive/zip.md)
* [bufio](docs/bufio.md)
* [bytes](docs/bytes.md)
* [compress](docs/compress.md)
    * [bzip2](docs/compress/bzip2.md)
    * [flate](docs/compress/flate.md)
    * [gzip](docs/compress/gzip.md)
    * [lzw](docs/compress/lzw.md)
    * [zlib](docs/compress/zlib.md)
* [container](docs/container.md)
    * [heap](docs/container/heap.md)
    * [list](docs/container/list.md)
    * [ring](docs/container/ring.md)
* [context](docs/context.md)
* [crypto](docs/crypto.md)
    * [aes](docs/crypto/aes.md)
    * [cipher](docs/crypto/cipher.md)
    * [des](docs/crypto/des.md)
    * [dsa](docs/crypto/dsa.md)
    * [ecdsa](docs/crypto/ecdsa.md)
    * [elliptic](docs/crypto/elliptic.md)
    * [md5](docs/crypto/md5.md)
    * [rand](docs/crypto/rand.md)
    * [sha1](docs/crypto/sha1.md)
    * [sha256](docs/crypto/sha256.md)
    * [sha512](docs/crypto/sha512.md)
    * [tls](docs/crypto/tls.md)
* [database](docs/database.md)
    * [sql](docs/database/sql.md)
* [debug](docs/debug.md)
    * [elf](docs/debug/elf.md)
    * [dwarf](docs/debug/dwarf.md)
    * [pe](docs/debug/pe.md)
* [encoding](docs/encoding.md)
    * [json](docs/encoding/json.md)
    * [binary](docs/encoding/binary.md)
    * [csv](docs/encoding/csv.md)
    * [gob](docs/encoding/gob.md)
    * [xml](docs/encoding/xml.md)
    * [hex](docs/encoding/hex.md)
    * [base32](docs/encoding/base32.md)
    * [base64](docs/encoding/base64.md)
    * [anscii85](docs/encoding/anscii85.md)
* [errors](docs/errors.md)
* [expar](docs/expar.md)
* [flag](docs/flag.md)
* [fmt](docs/fmt.md)
* [go](docs/go.md)
    * [build](docs/go/build.md)
    * [format](docs/go/format.md)
    * [importer](docs/go/importer.md)
    * [parser](docs/go/parser.md)
    * [printer](docs/go/printer.md)
    * [scanner](docs/go/scanner.md)
* [hash](docs/hash.md)
    * [crc32](docs/hash/crc32.md)
    * [crc64](docs/hash/crc64.md)
* [html](docs/html.md)
    * [template](docs/html/template.md)
* [image](docs/image.md)
    * [jpeg](docs/image/jpeg.md)
    * [png](docs/image/png.md)
    * [gif](docs/image/gif.md)
* [io](docs/io.md)
    * [ioutil](docs/io/ioutil.md)
* [log](docs/log.md)
    * [syslog](docs/log/syslog.md)
* [math](docs/math.md)
    * [rand](docs/math/rand.md)
* [mime](docs/mime.md)
    * [multipart](docs/mime/multipart.md)
    * [quotedprintal](docs/mime/quotedprintal.md)
* [net](docs/net.md)
    * [rpc](docs/net/rpc.md)
    * [http](docs/net/http.md)
        * [cgi](docs/net/http/cgi.md)
        * [fcgi](docs/net/http/fcgi.md)
    * [url](docs/net/url.md)
    * [smtp](docs/net/smtp.md)
* [os](docs/os.md)
    * [signal](docs/os/signal.md)
    * [exec](docs/os/exec.md)
    * [user](docs/os/user.md)
* [path](docs/path.md)
    * [filepath](docs/path/filepath.md)
* [reflect](docs/reflect.md)
* [regexp](docs/regexp.md)
    * [synatax](docs/regexp/synatax.md)
* [runtime](docs/runtime.md)
    * [cgo](docs/runtime/cgo.md)
    * [pprof](docs/runtime/pprof.md)
    * [race](docs/runtime/race.md)
    * [trace](docs/runtime/trace.md)
* [sort](docs/sort.md)
* [strconv](docs/strconv.md)
* [strings](docs/strings.md)
* [sync](docs/sync.md)
    * [atomic](docs/sync/atomic.md)
* [syscall](docs/syscall.md)
* [testing](docs/testing.md)
    * [iotest](docs/testing/iotest.md)
    * [quick](docs/testing/quick.md)
* [text](docs/text.md)
    * [scanner](docs/text/scanner.md)
    * [tabwriter](docs/text/tabwriter.md)
* [time](docs/time.md)
* [unicode](docs/unicode.md)
    * [utf8](docs/unicode/utf8.md)
    * [utf16](docs/unicode/utf16.md)
* [unsafe](docs/unsafe.md)


---

* [go设计模式](src/go设计模式.md)
* [创建型模式](src/创建型模式.md)
    * [1.单例模式](src/创建型模式/Singleton/单例模式.md)
    * [2.工厂方法模式](src/创建型模式/Factory/工厂方法模式.md)
    * [3.原型模式](src/创建型模式/Prototype/原型模式.md)
    * [4.抽象工厂模式](src/创建型模式/Abstract_Factory/抽象工厂模式.md)
    * [5.创建者模式](src/创建型模式/Builder/创建者模式.md)
    * [6.对象池模式](src/创建型模式/Object_Pool/对象池模式.md)
* [结构型模式](src/结构型模式.md)
    * [1.适配器模式](src/结构型模式/Adapter/适配器模式.md)
    * [2.桥接模式](src/结构型模式/Bridge/桥接模式.md)
    * [3.代理模式](src/结构型模式/Proxy/代理模式.md)
    * [4.装饰模式](src/结构型模式/Decorator/装饰模式.md)
    * [5.组合模式](src/结构型模式/Composite/组合模式.md)
    * [6.享元模式](src/结构型模式/Flyweight/享元模式.md)
    * [7.门面模式](src/结构型模式/Facade/门面模式.md)
* [行为型模式](src/行为型模式.md)
    * [1.模板方法模式](src/行为型模式/template/模板方法模式.md)
    * [2.观察者模式](src/行为型模式/Observer/观察者模式.md)
    * [3.状态模式](src/行为型模式/State/状态模式.md)
    * [4.策略模式](src/行为型模式/Strategy/策略模式.md)
    * [5.职责链模式](src/行为型模式/Responsibility_Chain/职责链模式.md)
    * [6.命令模式](src/行为型模式/Command/命令模式.md)
    * [7.访问者模式](src/行为型模式/Visitor/访问者模式.md)
    * [8.调停者模式](src/行为型模式/Mediator/调停者模式.md)
    * [9.备忘录模式](src/行为型模式/Memento/备忘录模式.md)
    * [10.迭代器模式](src/行为型模式/Iterator/迭代器模式.md)
    * [11.解释器模式](src/行为型模式/Interpreter/解释器模式.md)

---



* [常用框架](docs/常用框架.md)
    * [gorm](docs/常用框架/gorm.md)
    * [gin](docs/常用框架/gin.md)

---

## 设计模式分类

总体来说设计模式分为三大类：

A. 创建型模式（5种）

创建型模式，就是创建对象的模式，抽象了实例化的过程。它帮助一个系统独立于如何创建、组合和表示它的那些对象。关注的是对象的创建，创建型模式将创建对象的过程进行了抽象，也可以理解为将创建对象的过程进行了封装，作为客户程序仅仅需要去使用对象，而不再关系创建对象过程中的逻辑。

社会化的分工越来越细，自然在软件设计方面也是如此，因此对象的创建和对象的使用分开也就成为了必然趋势。因为对象的创建会消耗掉系统的很多资源，所以单独对对象的创建进行研究，从而能够高效地创建对象就是创建型模式要探讨的问题。这里有6个具体的创建型模式可供研究，它们分别是：

  - 简单工厂模式（Simple Factory）
  - 工厂方法模式（Factory Method）
  - 抽象工厂模式（Abstract Factory）
  - 创建者模式（Builder）
  - 原型模式（Prototype）
  - 单例模式（Singleton）

B. 结构型模式（7种）

结构型模式是为解决怎样组装现有的类，设计它们的交互方式，从而达到实现一定的功能目的。结构型模式包容了对很多问题的解决。例如：扩展性（外观、组成、代理、装饰）、封装（适配器、桥接）。

在解决了对象的创建问题之后，对象的组成以及对象之间的依赖关系就成了开发人员关注的焦点，因为如何设计对象的结构、继承和依赖关系会影响到后续程序的维护性、代码的健壮性、耦合性等。对象结构的设计很容易体现出设计人员水平的高低，这里有7个具体的结构型模式可供研究，它们分别是：

  - 外观模式/门面模式（Facade门面模式）
  - 适配器模式（Adapter）
  - 代理模式（Proxy）
  - 装饰模式（Decorator）
  - 桥梁模式/桥接模式（Bridge）
  - 组合模式（Composite）
  - 享元模式（Flyweight）

C. 行为型模式（11种）

行为型模式涉及到算法和对象间职责的分配，行为模式描述了对象和类的模式，以及它们之间的通信模式，行为模式刻划了在程序运行时难以跟踪的复杂的控制流可分为行为类模式和行为对象模式。1. 行为类模式使用继承机制在类间分派行为。2. 行为对象模式使用对象聚合来分配行为。一些行为对象模式描述了一组对等的对象怎样相互协作以完成其中任何一个对象都无法单独完成的任务。

在对象的结构和对象的创建问题都解决了之后，就剩下对象的行为问题了，如果对象的行为设计的好，那么对象的行为就会更清晰，它们之间的协作效率就会提高，这里有11个具体的行为型模式可供研究，它们分别是：

  - 模板方法模式（Template Method）
  - 观察者模式（Observer）
  - 状态模式（State）
  - 策略模式（Strategy）
  - 职责链模式（Chain of Responsibility）
  - 命令模式（Command）
  - 访问者模式（Visitor）
  - 调停者模式（Mediator）
  - 备忘录模式（Memento）
  - 迭代器模式（Iterator）
  - 解释器模式（Interpreter）

核心主要是上述三类，23种设计模式。但是，不够，掌握高并发需要了解并发行模式
---


D. 并发型模式：

并发模式是指I/O（输入/输出）处理单元和多个逻辑单元之间协调完成任务的方法，

在服务器上主要有两种并发编程模式：半同步/半异步模式 和 领导者/追随者模式


多线程编程的12种设计模式：

  - Immutable Object模式（不可变对象）
  - Guarded Suspension模式（保护性暂挂）
  - Two-phase Termination模式（两阶段终止）
  - Promise模式（承诺）
  - Producer-Consumer模式（生产者-消费者）
  - Active Object模式（主动对象）
  - Thread Pool模式（线程池）
  - Thread Specific Storage模式（线程特有存储）
  - Serial Thread Confinement模式（串行线程封闭模式）
  - Pipeline模式（流水线）
  - Half-sync/Half-async模式（半同步/半异步）
  - Master-Slave模式（主仆）


## 创建型、结构型、行为型三者之间的区别和联系

**创建型模式提供生存环境，结构型模式提供生存理由，行为型模式提供如何生存。**

**创建型模式为其他两种模式使用提供了环境。**

**结构型模式侧重于接口的使用，它做的一切工作都是对象或是类之间的交互，提供一个门。**

**行为型模式顾名思义，侧重于具体行为，所以概念中才会出现职责分配和算法通信等内容。**


## 设计模式六大原则

1. 开闭原则（Open Close Principle）
   -  对扩展开放，对修改关闭
2. 里氏代换原则（Liskov Substitution Principle）
   -  子类继承父类，单独完全可以运行
3. 依赖倒置原则（Dependence Inversion Principle）
   -  对扩展开放，引用一个对象，如果这个对象有底层类型，直接引用底层类型
4. 接口隔离原则（Interface Segregation Principle）
   -  每一个接口应该是一种角色
5. 迪米特法则（最少知道原则）（Demeter Principle）
   -  新的对象应使用一些已有的对象，使之成为新对象的一部分
6. 单一职责原则（Single-Responsibility-Principle）
   -  一个对象应对其他对象有尽可能少的了解

---


* [go爬虫框架-colly](docs/colly.md)


---


## 参考

* <https://github.com/PacktPublishing/Go-Design-Patterns>
* <https://github.com/tmrts/go-patterns>
* <https://golangbyexample.com/all-design-patterns-golang/>

---
