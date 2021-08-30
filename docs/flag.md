# flag

```
import "flag"
```

flag包实现了命令行参数的解析。

## func Parse()

* 从os.Args[1:]中解析注册的flag。
* 必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。

```
// Parse parses the command-line flags from os.Args[1:]. Must be called
// after all flags are defined and before flags are accessed by the program.
func Parse() {
	// Ignore errors; CommandLine is set for ExitOnError.
	CommandLine.Parse(os.Args[1:])
}
```




## func StringVar(p *string, name string, value string, usage string)

```
// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.Var(newStringValue(value, p), name, usage)
}
```

* StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。






---
