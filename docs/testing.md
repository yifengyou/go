<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [testing](#testing)   
   - [type T - 单元测试](#type-t-单元测试)   
      - [func TestXxxx(t *testing.T)](#func-testxxxxt-testingt)   
   - [type B - 基准测试](#type-b-基准测试)   
      - [func BenchmarkXxxx(b *testing.B)](#func-benchmarkxxxxb-testingb)   

<!-- /MDTOC -->

# testing

* go test 命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件。
* testing 提供对 Go 包的自动化测试的支持。通过 `go test` 命令，能够自动执行如下形式的任何函数：


```
func TestXxx(*testing.T)
```

* 其中 Xxx 可以是任何字母数字字符串（但开头字母必须大写，不能是 [a-z]，不能是数字），用于识别测试例程。
* 在这些函数中，使用 Error, Fail 或相关方法来发出失败信号。
* 测试返回实例：

```
=== RUN   TestDelayfunc
df.String() =  {DelayFun:, args:[hello zinx!]}
hello   zinx!
--- PASS: TestDelayfunc (0.00s)
PASS
```


* 要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 `TestXxx` 函数，如上所述。
* 将该文件放在与被测试的包相同的包中。该文件将被排除在正常的程序包之外，但在运行 “go test” 命令时将被包含。 有关详细信息，请运行 “go help test” 和 “go help testflag” 了解。

如果有需要，可以调用 *T 和 *B 的 Skip 方法，跳过该测试或基准测试：

```
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

## type T - 单元测试

```
// T is a type passed to Test functions to manage test state and support formatted test logs.
//
// A test ends when its Test function returns or calls any of the methods
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods, as well as
// the Parallel method, must be called only from the goroutine running the
// Test function.
//
// The other reporting methods, such as the variations of Log and Error,
// may be called simultaneously from multiple goroutines.
type T struct {
	common
	isParallel bool
	context    *testContext // For running tests and subtests.
}
```

### func TestXxxx(t *testing.T)

```
func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
```



## type B - 基准测试

```
// B is a type passed to Benchmark functions to manage benchmark
// timing and to specify the number of iterations to run.
//
// A benchmark ends when its Benchmark function returns or calls any of the methods
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods must be called
// only from the goroutine running the Benchmark function.
// The other reporting methods, such as the variations of Log and Error,
// may be called simultaneously from multiple goroutines.
//
// Like in tests, benchmark logs are accumulated during execution
// and dumped to standard output when done. Unlike in tests, benchmark logs
// are always printed, so as not to hide output whose existence may be
// affecting benchmark results.
type B struct {
	common
	importPath       string // import path of the package containing the benchmark
	context          *benchContext
	N                int
	previousN        int           // number of iterations in the previous run
	previousDuration time.Duration // total duration of the previous run
	benchFunc        func(b *B)
	benchTime        benchTimeFlag
	bytes            int64
	missingBytes     bool // one of the subbenchmarks does not have bytes set.
	timerOn          bool
	showAllocResult  bool
	result           BenchmarkResult
	parallelism      int // RunParallel creates parallelism*GOMAXPROCS goroutines
	// The initial states of memStats.Mallocs and memStats.TotalAlloc.
	startAllocs uint64
	startBytes  uint64
	// The net total of this test after being run.
	netAllocs uint64
	netBytes  uint64
	// Extra metrics collected by ReportMetric.
	extra map[string]float64
}
```


### func BenchmarkXxxx(b *testing.B)


* Benchmark 是基准测试，
* 通过 "go test" 命令，加上 -bench flag 来执行

基准测试函数样例看起来如下所示：

```
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```

测试输出：

```
goos: windows
goarch: amd64
pkg: github.com/aceld/zinx/ztimer
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkHello
BenchmarkHello-8   	 9139898	       119.6 ns/op
PASS

Process finished with the exit code 0
```


```
func (c *B) Error(args ...interface{})
func (c *B) Errorf(format string, args ...interface{})
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...interface{})
func (c *B) Fatalf(format string, args ...interface{})
func (c *B) Log(args ...interface{})
func (c *B) Logf(format string, args ...interface{})
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Skip(args ...interface{})
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...interface{})
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()
```







---
