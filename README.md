# getopt for go
### Installation
```
go get ...
```
### Base Usage
```go
package main

import (
	"fmt"
	"go-getopt"
)

func main() {
	opts := []getopt.Option{
		getopt.Option{'n',true,"zxj","name","user name"},
		getopt.Option{'a',true,"12","age","user age"},
		getopt.Option{'w',false,"code","work","user work"},
		getopt.Option{'c',true,"","cc","user cc"},
	}
	res,err :=getopt.GetOpt(opts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```
传参：
```console
xxx -n=xxx -c=xxx -w -a=12
xxx -n=xxx -c=xxx -wa=12
xxx -n -c=xxx -wa=12
xxx -nwa -c=xxx
xxx -na -c=xxx
xxx --name --age -wc=cccc
```
### 结构
```go
type Option struct {
	Short int32 //短名字
	Required bool //是否必须
	Default string //默认值
	Long string //长名字
	Useage string	//提示信息
}
```
字段 | 组合
--- | ---
Short | 短参数名
Required | 是否为必传参数
Default | 默认值。如果是必传参数，将此字段设置为""(空字符串)，那么命令行传参数时就必须写成 -x=xx
Long | 长参数名
Useage | 参数提示信息

### 结果
- 如果是非必传参数，如果不传，即使设置了默认值，结果中不会有该字段的值
- 如果是必传字段，如果不传，则会返回错误，并显示useage。如果传了，但是没有给值，那么就会被赋予默认值。如果默认值为""空字符串，那么会返回错误，告知某字段必须为 -x=xx的形式
- 如果传递了没有注册的字段，则结果会忽略该字段。

<div style="text-indent:2em;">结果为一个map[string]string类型的map<br />
其中 键 为 短参数名，值为命令行传入的参数值或者默认值
</div>
