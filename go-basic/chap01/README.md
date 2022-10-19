# Go入门

## c01_helloWorld.go

运行： `go run c01_helloWorld.go`
编译： `go build c01_helloWorld.go`

## c02_osArgs.go

获取命令行参数：`os.Args`。其中`Args[0]`固定是当前文件的位置参数

## c03_printf.go

`fmt.Printf`的占位符类型：

```
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```