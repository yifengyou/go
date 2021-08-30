# context

```
import "context"
```

* go 1.7引入

## type Context

### func Background() Context

### func TODO() Context

### func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

### func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

### func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

### func WithValue(parent Context, key, val interface{}) Context




## 参考



















---
