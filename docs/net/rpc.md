<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [rpc](#rpc)   
   - [简单例子](#简单例子)   
   - [type Server](#type-server)   
      - [func NewServer() *Server](#func-newserver-server)   
      - [func (server *Server) Register(rcvr interface{}) error](#func-server-server-registerrcvr-interface-error)   
      - [func (server *Server) RegisterName(name string, rcvr interface{}) error](#func-server-server-registernamename-string-rcvr-interface-error)   
      - [func (server *Server) Accept(lis net.Listener)](#func-server-server-acceptlis-netlistener)   
      - [func (server *Server) ServeConn(conn io.ReadWriteCloser)](#func-server-server-serveconnconn-ioreadwritecloser)   
   - [func Dial(network, address string) (*Client, error)](#func-dialnetwork-address-string-client-error)   
   - [func ServeConn(conn io.ReadWriteCloser)](#func-serveconnconn-ioreadwritecloser)   
   - [type Client](#type-client)   
      - [func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error](#func-client-client-callservicemethod-string-args-interface-reply-interface-error)   
      - [func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call](#func-client-client-goservicemethod-string-args-interface-reply-interface-done-chan-call-call)   
      - [func (client *Client) Close() error](#func-client-client-close-error)   

<!-- /MDTOC -->
# rpc

```
import "net/rpc"
```

## 简单例子

* 服务端


```
type Args struct {
	A, B int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.RegisterName("App", arith)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Accept error:", e)
		}
		rpc.ServeConn(conn)
	}
}
```

* 客户端

```
type Args struct {
	A, B int
}
func main() {
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Conn error:", err)
	}
	var result int
	args := Args{9, 9}
	err = conn.Call("App.Multiply", args, &result)
	if err != nil {
		log.Fatal("Call error:", err)
	}
	fmt.Println("Call success :", result)
}
```


## type Server

```
type Server struct {
    // 内含隐藏或非导出字段
}
```
Server代表RPC服务端。

### func NewServer() *Server

NewServer创建并返回一个*Server。

### func (server *Server) Register(rcvr interface{}) error

Register在server注册并公布rcvr的方法集中满足如下要求的方法：

- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值

如果rcvr不是一个导出类型的值，或者该类型没有满足要求的方法，Register会返回错误。Register也会使用log包将错误写入日志。客户端可以使用格式为"Type.Method"的字符串访问这些方法，其中Type是rcvr的具体类型。

### func (server *Server) RegisterName(name string, rcvr interface{}) error

RegisterName类似Register，但使用提供的name代替rcvr的具体类型名作为服务名。

### func (server *Server) Accept(lis net.Listener)

Accept接收监听器l获取的连接，然后服务每一个连接。Accept会阻塞，调用者应另开线程："go server.Accept(l)"

### func (server *Server) ServeConn(conn io.ReadWriteCloser)

ServeConn在单个连接上执行server。ServeConn会阻塞，服务该连接直到客户端挂起。调用者一般应另开线程调用本函数："go server.ServeConn(conn)"。ServeConn在该连接使用gob（参见encoding/gob包）有线格式。要使用其他的编解码器，可调用ServeCodec方法。











## func Dial(network, address string) (*Client, error)

* Dial在指定的网络和地址与RPC服务端连接。

## func ServeConn(conn io.ReadWriteCloser)

ServeConn在单个连接上执行DefaultServer。ServeConn会阻塞，服务该连接直到客户端挂起。调用者一般应另开线程调用本函数："go ServeConn(conn)"。ServeConn在该连接使用gob（参见encoding/gob包）有线格式。要使用其他的编解码器，可调用ServeCodec方法


## type Client
type Client struct {
    // 内含隐藏或非导出字段
}
Client类型代表RPC客户端。同一个客户端可能有多个未返回的调用，也可能被多个go程同时使用。

### func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error

Call调用指定的方法，等待调用返回，将结果写入reply，然后返回执行的错误状态。

### func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call

Go异步的调用函数。本方法Call结构体类型指针的返回值代表该次远程调用。通道类型的参数done会在本次调用完成时发出信号（通过返回本次Go方法的返回值）。如果done为nil，Go会申请一个新的通道（写入返回值的Done字段）；如果done非nil，done必须有缓冲，否则Go方法会故意崩溃。

### func (client *Client) Close() error

---
