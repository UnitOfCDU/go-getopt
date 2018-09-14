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
