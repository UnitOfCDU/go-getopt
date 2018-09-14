package getopt

import (
	"os"
	"errors"
	"fmt"
	"strings"
	"expvar"
	"path/filepath"
)

type Option struct {
	Short int32 //短名字
	Required bool //是否必须
	Default string //默认值
	Long string //长名字
	Useage string	//提示信息
}
var(
	useage string
)

func GetOpt(opts []Option)(map[string]interface{},error){
	args := getArgs()
	createUseage(opts)
	res := make(map[string]interface{})
	err := parseArgs(args,opts,res)
	if err != nil {
		return nil,err
	}
	//检查是否所有的参数都传值了。
	for _,opt := range opts{
		if opt.Required {
			_,ok := res[string(opt.Short)]
			if !ok {
				return nil,errors.New(opt.Long + " is required")
			}
		}
	}
	return res, nil
}
func parseArgs(args []string,opt []Option,res map[string]interface{})error{
	var err error

	for _,arg := range args {
		if len(arg) < 2{
			//参数格式错误
			err = errors.New("参数格式错误")
			break
		}
		if arg[0:2] == "--" {
			err = parseLongArgs(arg[2:],opt,res)
		}else if arg[0] == '-' {
			err = parseShortArgs(arg[1:],opt,res)
		}else{
			//参数格式错误
			err = errors.New("参数格式错误")
			break
		}
	}
	if err != nil {
		//显示提示信息
		PrintUseage()
		return err
	}
	return nil
}
func PrintUseage(){
	fmt.Println(useage)
}
func createUseage(opts []Option){
	cmd := expvar.Get("cmdline").String()
	cmd = strings.Split(cmd,",")[0]
	cmd = strings.Trim(cmd,`["`)
	cmd = filepath.Base(cmd)
	res := ""
	header := cmd + " "
	for _,opt := range opts {
		str := "\t"
		h := ""
		h += "[-"+string(opt.Short)+"=xxx]"
		str += "-" + string(opt.Short) + "\t"
		str += "--" + string(opt.Long) + "\t\t"
		if opt.Required {
			str += " required \t"
			h = strings.Trim(h,"[]")
		}else{
			str += " optional \t"
		}
		str += opt.Useage+"\n"
		res += str
		header += h + "  "
	}
	useage = header+"\n"+res
}
func getArgs()[]string{
	return os.Args[1:]
}

func parseShortArgs(sarg string,opts []Option,res map[string]interface{}) error{
	l := len(sarg)
	for i,c := range sarg {
		if c != '=' && i+1<l && sarg[i+1] == '='{
			k := sarg[i]
			v := sarg[i+2:]
			res[string(k)] = v
			goto END
		}else{//单个字符
			for _,opt := range opts{
				if opt.Short == c {
					if opt.Required && opt.Default == "" {
						return errors.New(opt.Long+" is required")
					}else{
						res[string(c)] = opt.Default
					}
				}
			}
		}

	}
	END:
	return nil
}
func parseLongArgs(larg string,opts []Option,res map[string]interface{}) error{
	if len(larg) == 0 {
		return nil
	}
	kv := strings.Split(larg,"=")
	var k,v string
	k = kv[0]
	if len(kv) > 1 {//没有等号
		v = kv[1]
	}
	if k == "" {
		PrintUseage()
		return errors.New("参数格式不正确")
	}
	for _,opt := range opts{
		if opt.Long == k {
			if opt.Required && v == "" && opt.Default == ""{
				return errors.New(opt.Long+" is required")
			}else if v == "" && len(kv) <= 1 {
				res[string(opt.Short)] = opt.Default
			}else{
				res[string(opt.Short)] = v
			}
		}
	}
	return nil
}


