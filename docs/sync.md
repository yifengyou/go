
<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [sync](#sync)   
   - [type Locker](#type-locker)   
   - [type Once](#type-once)   
      - [func (o *Once) Do(f func())](#func-o-once-dof-func)   
   - [type Mutex](#type-mutex)   
      - [func (m *Mutex) Lock()](#func-m-mutex-lock)   
      - [func (m *Mutex) Unlock()](#func-m-mutex-unlock)   
   - [type RWMutex](#type-rwmutex)   
      - [func (rw *RWMutex) Lock()](#func-rw-rwmutex-lock)   
      - [func (rw *RWMutex) Unlock()](#func-rw-rwmutex-unlock)   
      - [func (rw *RWMutex) RLock()](#func-rw-rwmutex-rlock)   
      - [func (rw *RWMutex) RUnlock()](#func-rw-rwmutex-runlock)   
      - [func (rw *RWMutex) RLocker() Locker](#func-rw-rwmutex-rlocker-locker)   
   - [type Cond](#type-cond)   
      - [func NewCond(l Locker) *Cond](#func-newcondl-locker-cond)   
      - [func (c *Cond) Broadcast()](#func-c-cond-broadcast)   
      - [func (c *Cond) Signal()](#func-c-cond-signal)   
      - [func (c *Cond) Wait()](#func-c-cond-wait)   
   - [type WaitGroup](#type-waitgroup)   
      - [func (wg *WaitGroup) Add(delta int)](#func-wg-waitgroup-adddelta-int)   
      - [func (wg *WaitGroup) Done()](#func-wg-waitgroup-done)   
      - [func (wg *WaitGroup) Wait()](#func-wg-waitgroup-wait)   
   - [type Pool](#type-pool)   
      - [func (p *Pool) Get() interface{}](#func-p-pool-get-interface)   
      - [func (p *Pool) Put(x interface{})](#func-p-pool-putx-interface)   
   - [参考](#参考)   

<!-- /MDTOC -->
# sync

sync包提供了基本的同步基元，如互斥锁。

除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些

## type Locker

## type Once
### func (o *Once) Do(f func())

## type Mutex

Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。

底层定义为：

```
// A Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.
type Mutex struct {
	state int32
	sema uint32
}
```

```
	demo := sync.Mutex{}
	fmt.Printf("%T - %#v\n", demo, demo)
	// sync.Mutex - sync.Mutex{state:0, sema:0x0}
```

### func (m *Mutex) Lock()
### func (m *Mutex) Unlock()

* Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。
* sync.Mutex一旦被锁住，其它的Lock()操作就无法再获取它的锁，只有通过Unlock()释放锁之后才能通过Lock()继续获取锁。
* 已有的锁会导致其它申请Lock()操作的goroutine被阻塞，且只有在Unlock()的时候才会解除阻塞。
* sync.Mutex不区分读写锁，只有Lock()与Lock()之间才会导致阻塞的情况，如果在一个地方Lock()，在另一个地方不Lock()而是直接修改或访问共享数据，这对于sync.Mutex类型来说是允许的，因为mutex不会和临界区进行关联
* 因此就需要用户代码把握好，访问临界区一定要加锁
* Lock()和Unlock()之间的代码段称为资源的临界区(critical section)，在这一区间内的代码是严格被Lock()保护的，是线程安全的，任何一个时间点都只能有一个goroutine执行这段区间的代码

```
	demo.Lock()
	// ....
	defer demo.Unlock()
```





## type RWMutex



读写锁，务必记住，适用于读多写少的情况。基本遵循两大原则：

1. 可以随便读，多个goroutine同时读。
2. 写的时候，啥也不能干。不能读也不能写

```
type RWMutex struct {
	w           Mutex  // 互斥锁，写锁协程获取该锁后，其他写锁处于阻塞等待
	writerSem   uint32 // 写入等待信号量，最后一个读取协程释放锁时会释放信号量
	readerSem   uint32 // 读取等待信号量，持有写锁协程释放后会释放信号量
	readerCount int32  // 读锁的个数
	readerWait  int32  // 写操作时，需要等待读操作的个数
}
```

```
	demo := sync.RWMutex{}
	fmt.Printf("%T \n %#v\n", demo, demo)
	// sync.RWMutex
	// sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:0, readerWait:0}
```

* RWMutex是读写互斥锁。
* 该锁可以被同时多个读取者持有或唯一个写入者持有。
* RWMutex可以创建为其他结构体的字段；零值为解锁状态。
* RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。


### func (rw *RWMutex) Lock()
### func (rw *RWMutex) Unlock()
### func (rw *RWMutex) RLock()
### func (rw *RWMutex) RUnlock()

都是方法，其中Lock、Unlock为rw锁定，Rlock、R**U**nlock为读锁定

* Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入。
* Unlock方法解除rw的写入锁状态，如果m未加写入锁会导致运行时错误。
* R**L**ock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取。
* R**U**nlock方法解除rw的读取锁状态，如果m未加读取锁会导致运行时错误。


### func (rw *RWMutex) RLocker() Locker


// 返回一个Locker对象，用得较少

```
func (rw *RWMutex) RLocker() Locker {
	return (*rlocker)(rw)
}
```

## type Cond
### func NewCond(l Locker) *Cond
### func (c *Cond) Broadcast()
### func (c *Cond) Signal()
### func (c *Cond) Wait()





## type WaitGroup

```
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
type WaitGroup struct {
	noCopy noCopy

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state, and the other 4 as storage
	// for the sema.
	state1 [3]uint32
}
```

* WaitGroup用于等待一组线程的结束。
* 父线程调用Add方法来设定应等待的线程的数量。
* 每个被等待的线程在结束时应调用Done方法。
* 同时，主线程里可以调用Wait方法阻塞至所有线程结束。


实例：

```
var wg sync.WaitGroup
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
    "http://www.somestupidname.com/",
}
for _, url := range urls {
    // Increment the WaitGroup counter.
    wg.Add(1)
    // Launch a goroutine to fetch the URL.
    go func(url string) {
        // Decrement the counter when the goroutine completes.
        defer wg.Done()
        // Fetch the URL.
        http.Get(url)
    }(url)
}
// Wait for all HTTP fetches to complete.
// 如果不wait，主进程退出，协程也都要结束掉。甚至可能都还没开始执行
wg.Wait()
```

### func (wg *WaitGroup) Add(delta int)

* Add方法向内部计数加上delta，delta可以是负数；
* 如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，
* 如果计数器小于0，方法panic。
* 注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。
* 一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。
* 提前比较需要wait的计数

### func (wg *WaitGroup) Done()

* Done方法减少WaitGroup计数器的值，应在协程的最后执行


### func (wg *WaitGroup) Wait()

* Wait方法阻塞直到WaitGroup计数器减为0








## type Pool
### func (p *Pool) Get() interface{}
### func (p *Pool) Put(x interface{})





## 参考

* <https://blog.csdn.net/guichenglin/article/details/105729923>



---
