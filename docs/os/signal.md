<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [signal](#signal)   
   - [func Notify(c chan<- os.Signal, sig ...os.Signal)](#func-notifyc-chan-ossignal-sig-ossignal)   
   - [func Stop(c chan<- os.Signal)](#func-stopc-chan-ossignal)   

<!-- /MDTOC -->

# signal

```
import "os/signal"
```

* signal包实现了对输入信号的访问

## func Notify(c chan<- os.Signal, sig ...os.Signal)

* Notify函数让signal包将输入信号转发到c。
* 如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。

signal包不会为了向c发送信息而阻塞（就是说如果发送时c阻塞了，signal包会直接放弃）：调用者应该保证c有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的通道，缓存为1就足够了。

可以使用同一通道多次调用Notify：每一次都会扩展该通道接收的信号集。唯一从信号集去除信号的方法是调用Stop。可以使用同一信号和不同通道多次调用Notify：每一个通道都会独立接收到该信号的一个拷贝。

```
// Set up channel on which to send signal notifications.
// We must use a buffered channel or risk missing the signal
// if we're not ready to receive when the signal is sent.
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt, os.Kill)
// Block until a signal is received.
s := <-c
fmt.Println("Got signal:", s)
```


## func Stop(c chan<- os.Signal)

* Stop函数让signal包停止向c转发信号。它会取消之前使用c调用的所有Notify的效果
* 当Stop返回后，会保证c不再接收到任何信号















---
