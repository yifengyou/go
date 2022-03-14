<!-- TOC -->

- [go命令](#go命令)
  - [go generate](#go-generate)

<!-- /TOC -->
# go命令


```
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

```

## go build

## go clean

## go doc

## go install

## go env



## go run

```
compile and run Go program
```

举栗子:

```
go run ./dir    # dir目录下包含main包main函数
go run -exec 'sudo ' ./kprobe  # 使用sudo指定kprobe目录下的main包main函数
```



## go generate

```
generate Go files by processing source
```

* go generate命令是go 1.4版本里面新添加的一个命令
* 当运行go generate时，它将扫描与当前包相关的源代码文件，找出所有包含"//go:generate"的特殊注释，提取并执行该特殊注释后面的命令，命令为可执行程序，形同shell下面执行。

有几点需要注意：

1. 该特殊注释必须在.go源码文件中。
2. 每个源码文件可以包含多个generate特殊注释时。
3. 显示运行go generate命令时，才会执行特殊注释后面的命令。
4. 命令串行执行的，如果出错，就终止后面的执行。
5. 特殊注释必须以"//go:generate"开头，双斜线后面没有空格。

举栗子：

```
//go:generate echo hello
//go:generate go run main.go
//go:generate echo file=$GOFILE pkg=$GOPACKAGE
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf kprobe.c -- -I../headers
```



---
