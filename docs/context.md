# context

```
import "context"
```

* go 1.7引入

## type Context

### func Background() Context


```
func Background() Context {
	return background
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

type emptyCtx int
```




### func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

* 传递取消，传递的仅仅是信号
* 传递的是信号，不一定立即结束，也需要善后，实际处理还需要看协程逻辑

实例：

```
func g1(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("g1取消")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go g1(ctx)
	count := 1
	for {
		if count >= 3 {
			cancel()
		}
		time.Sleep(time.Second)
		count++
	}
}

```





### func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

* 传递截止时间信号



### func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

* 传递超时

```
func g1(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("超时取消")
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go g1(ctx)
	count := 1
	for {
		if count >= 3 {
			fmt.Println("手动取消")
			cancel()
			count = 0
		}
		time.Sleep(time.Second)
		if count > 0 {
			count++
		}
	}
}
```

输出：

```
超时取消
手动取消
```



### func WithValue(parent Context, key, val interface{}) Context

* 传递key-value

```
func g2(ctx context.Context) {
	fmt.Println(ctx.Value("begin"))
}

func g(ctx context.Context) {
	fmt.Println(ctx.Value("begin"))
	fmt.Println("你是猪")
	ctx2 := context.WithValue(ctx,"begin", "这电影真好看")
	go g2(ctx2)
}
func main() {
	ct := context.WithValue(context.Background(), "begin", "从台词里看到一句话")
	go g(ct)
	time.Sleep(time.Second)

}
```

输出：

```
从台词里看到一句话
你是猪
这电影真好看
```


### func TODO() Context

















## 参考



















---
