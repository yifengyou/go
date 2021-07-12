# zinxV0.1-开发服务器应用

## 服务端

```
package main

import "zinx/znet"

func main() {
	// 创建一个server句柄
	s := znet.NewServer("[zinx V0.1]")
	//启动server
	s.Serve()
}

```


## 客户端

```
package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
*/

func main() {
	fmt.Println("Client start")
	time.Sleep(time.Second)
	// 1. 创建tcp连接，得到conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dail error", err)
		return
	}
	// 2. 链接调用Write写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1..."))
		if err != nil {
			fmt.Println("write failed!", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err")
			return
		}
		fmt.Printf("server call back : %s , cnt = %d\n", string(buf[:cnt]), cnt)
		time.Sleep(1 * time.Second)
	}
}

```
