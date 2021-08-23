<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [net](#net)   
   - [通用](#通用)   
      - [func Listen(network, address string) (Listener, error)](#func-listennetwork-address-string-listener-error)   
      - [func Dial(network, address string) (Conn, error)](#func-dialnetwork-address-string-conn-error)   
      - [func (c *conn) Read(b []byte) (int, error)](#func-c-conn-readb-byte-int-error)   
      - [func (c *conn) Write(b []byte) (int, error)](#func-c-conn-writeb-byte-int-error)   
      - [func (c *conn) Close() error](#func-c-conn-close-error)   
      - [func (c *conn) RemoteAddr() Addr](#func-c-conn-remoteaddr-addr)   
   - [TCP](#tcp)   
      - [func ResolveTCPAddr(network, address string) (*TCPAddr, error)](#func-resolvetcpaddrnetwork-address-string-tcpaddr-error)   
      - [func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)](#func-listentcpnetwork-string-laddr-tcpaddr-tcplistener-error)   
      - [func (l *TCPListener) AcceptTCP() (*TCPConn, error)](#func-l-tcplistener-accepttcp-tcpconn-error)   
      - [func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)](#func-dialtcpnetwork-string-laddr-raddr-tcpaddr-tcpconn-error)   
   - [UDP](#udp)   
      - [func ResolveUDPAddr(network, address string) (*UDPAddr, error)](#func-resolveudpaddrnetwork-address-string-udpaddr-error)   
      - [func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)](#func-listenudpnetwork-string-laddr-udpaddr-udpconn-error)   
      - [func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)](#func-dialudpnetwork-string-laddr-raddr-udpaddr-udpconn-error)   
   - [unix](#unix)   
      - [func ResolveUnixAddr(network, address string) (*UnixAddr, error)](#func-resolveunixaddrnetwork-address-string-unixaddr-error)   
      - [func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)](#func-listenunixnetwork-string-laddr-unixaddr-unixlistener-error)   
      - [func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)](#func-listenunixgramnetwork-string-laddr-unixaddr-unixconn-error)   
      - [func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)](#func-dialunixnetwork-string-laddr-raddr-unixaddr-unixconn-error)   

<!-- /MDTOC -->
# net

## 通用

### func Listen(network, address string) (Listener, error)
### func Dial(network, address string) (Conn, error)
### func (c *conn) Read(b []byte) (int, error)
### func (c *conn) Write(b []byte) (int, error)
### func (c *conn) Close() error




### func (c *conn) RemoteAddr() Addr

* RemoteAddr() 可以获取客户端的地址和端口号
*





## TCP

### func ResolveTCPAddr(network, address string) (*TCPAddr, error)

```
func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", "192.168.1.199:22")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	fmt.Printf("%T -- %#v", conn.RemoteAddr(), conn.RemoteAddr())
	defer conn.Close()
	log.Println("dial ok")
}
```

输出:

```
*net.TCPAddr -- &net.TCPAddr{IP:net.IP{0xc0, 0xa8, 0x1, 0xc7}, Port:22, Zone:""}
```



### func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)

### func (l *TCPListener) AcceptTCP() (*TCPConn, error)

### func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)


## UDP

### func ResolveUDPAddr(network, address string) (*UDPAddr, error)

### func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)

### func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)


## unix

### func ResolveUnixAddr(network, address string) (*UnixAddr, error)

### func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)

### func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)

### func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)



## type IP

```
type IP []byte
```

* IP类型是代表单个IP地址的[]byte切片。
* 本包的函数都可以接受4字节（IPv4）和16字节（IPv6）的切片作为输入。

注意，IP地址是IPv4地址还是IPv6地址是语义上的属性，而不取决于切片的长度：16字节的切片也可以是IPv4地址。

## func ParseIP(s string) IP

* 将s解析为IP类型
* ParseIP将s解析为IP地址，并返回该地址。如果s不是合法的IP地址文本表示，ParseIP会返回nil。
* 字符串可以是小数点分隔的IPv4格式（如"74.125.19.99"）或IPv6格式（如"2001:4860:0:2001::68"）格式。





### func (ip IP) IsLoopback() bool

如果ip是环回地址，则返回真。

示例：

```
b := net.ParseIP("172.0.0.1")
fmt.Println(b.IsLoopback()) // true
```

```
func IPv4(a, b, c, d byte) IP
func ParseIP(s string) IP
func (ip IP) IsGlobalUnicast() bool
func (ip IP) IsLinkLocalUnicast() bool
func (ip IP) IsInterfaceLocalMulticast() bool
func (ip IP) IsLinkLocalMulticast() bool
func (ip IP) IsMulticast() bool
func (ip IP) IsUnspecified() bool
func (ip IP) DefaultMask() IPMask
func (ip IP) Equal(x IP) bool
func (ip IP) To16() IP
func (ip IP) To4() IP
func (ip IP) Mask(mask IPMask) IP
func (ip IP) String() string
func (ip IP) MarshalText() ([]byte, error)
func (ip *IP) UnmarshalText(text []byte) error
```





---
