<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [time](#time)   
   - [func Sleep(d Duration)](#func-sleepd-duration)   
   - [type Duration](#type-duration)   
   - [type Time](#type-time)   
      - [func Now() Time](#func-now-time)   
      - [func (t Time) Unix() int64](#func-t-time-unix-int64)   
      - [func (t Time) UnixMilli() int64](#func-t-time-unixmilli-int64)   
      - [func (t Time) UnixMicro() int64](#func-t-time-unixmicro-int64)   
      - [func (t Time) UnixNano() int64](#func-t-time-unixnano-int64)   

<!-- /MDTOC -->

# time

## 时间度量单位


```
time.Hour
time.Minute
time.Second
time.Millisecond  //1000
time.Microsecond  //1000
time.Nanosecond   //1000
```


## func Sleep(d Duration)

* 参数需要提供时间单位的整数倍数

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




## type Duration

```
// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
type Duration int64
```

* 根据上面定义，毫无疑问，time.Duration()是类型转换
* 常规time.Sleep都是要求Duration类型，所以一般都要进行类型转换，除非常熟

```
//tw.interval是int64类型，但是编译器只要Duration，强类型要求必须转换，因此
time.Sleep(time.Duration(tw.interval) * time.Millisecond)

//但如果是常数字面量就不需要
time.Sleep(123 * time.Millisecond)
```







```
func ParseDuration(s string) (Duration, error)
func Since(t Time) Duration
func (d Duration) Hours() float64
func (d Duration) Minutes() float64
func (d Duration) Seconds() float64
func (d Duration) Nanoseconds() int64
func (d Duration) String() string
```


## type Time


```
type Time struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	wall uint64
	ext  int64

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// The nil location means UTC.
	// All UTC times are represented with loc==nil, never loc==&utcLoc.
	loc *Location
}
```

### func Now() Time

// Now returns the current local time.
Now返回当前本地时间。






### func (t Time) Unix() int64
### func (t Time) UnixMilli() int64
### func (t Time) UnixMicro() int64
### func (t Time) UnixNano() int64

* 秒、毫秒、微妙、纳秒关系：

1s = 1000 ms
1ms = 1000 微妙
1微妙 = 1000 纳秒

1s = 1000,000,0000 纳秒
1s = 1000,000 微妙


* UnixMilli、UnixMicro在Go 1.17中有，不是所有版本都有


### func (t Time) Add(d Duration) Time

### func (t Time) AddDate(years int, months int, days int) Time

### func (t Time) Sub(u Time) Duration

```
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
func Parse(layout, value string) (Time, error)
func ParseInLocation(layout, value string, loc *Location) (Time, error)

func Unix(sec int64, nsec int64) Time
func (t Time) Location() *Location
func (t Time) Zone() (name string, offset int)
func (t Time) IsZero() bool
func (t Time) Local() Time
func (t Time) UTC() Time
func (t Time) In(loc *Location) Time
func (t Time) Equal(u Time) bool
func (t Time) Before(u Time) bool
func (t Time) After(u Time) bool
func (t Time) Date() (year int, month Month, day int)
func (t Time) Clock() (hour, min, sec int)
func (t Time) Year() int
func (t Time) Month() Month
func (t Time) ISOWeek() (year, week int)
func (t Time) YearDay() int
func (t Time) Day() int
func (t Time) Weekday() Weekday
func (t Time) Hour() int
func (t Time) Minute() int
func (t Time) Second() int
func (t Time) Nanosecond() int
func (t Time) Round(d Duration) Time
func (t Time) Truncate(d Duration) Time
func (t Time) Format(layout string) string
func (t Time) String() string
func (t Time) GobEncode() ([]byte, error)
func (t *Time) GobDecode(data []byte) error
func (t Time) MarshalBinary() ([]byte, error)
func (t *Time) UnmarshalBinary(data []byte) error
func (t Time) MarshalJSON() ([]byte, error)
func (t *Time) UnmarshalJSON(data []byte) error
func (t Time) MarshalText() ([]byte, error)
func (t *Time) UnmarshalText(data []byte) error
```





示例：

```
	fmt.Printf("时间戳（秒钟）：%v\n", time.Now().Unix())
	//时间戳（秒钟）：1629640901
	fmt.Printf("时间戳（毫秒）：%v\n",time.Now().UnixMilli())
	//时间戳（毫秒）：1629640901326
	fmt.Printf("时间戳（微妙）：%v\n",time.Now().UnixMicro())
	//时间戳（微妙）：1629640901326004
	fmt.Printf("时间戳（纳秒）：%v\n",time.Now().UnixNano())
	//时间戳（纳秒）：1629640901326004800

	fmt.Printf("时间戳（毫秒）：%v\n",time.Now().UnixNano() / 1e6)
	//时间戳（毫秒）：1629640901326
	fmt.Printf("时间戳（秒钟）：%v\n",time.Now().UnixNano() / 1e9)
	//时间戳（秒钟）：1629640901

```








---
