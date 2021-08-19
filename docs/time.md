# time

## func Sleep(d Duration)

```
type Duration int64
```

```
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
```

```
const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)
```


















---
