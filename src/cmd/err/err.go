package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"cmd/common/writer"
	"cmd/common"
)
type err struct {
	N string
	I   int
	M string
}

type Errslice struct {
	Errs []err
}

func main(){
	args := os.Args //获取用户输入的所有参数
	errpath :=""
	if args==nil || len(args)<2{
		fmt.Println("请输入err路径?")
		errpath=`D:\work\yxlserver\src\common\errmsg\`

	}else{
		errpath =args[1];
		fmt.Println(args[1])
	}


	var s Errslice
	content, err := ioutil.ReadFile(errpath+"err.json")
	if err != nil {
		fmt.Println(err);
	}
	json.Unmarshal(content, &s)
	fmt.Println(s)
	var writer writer.Writer
	writer.AddLine("package errmsg");
	writer.AddLine("const (")
	for _,v :=range  s.Errs{
		writer.AddLineFmt("\t%v=%v\t//%v",v.N,v.I,v.M)
	}
	writer.AddLine(")")
	fmt.Println(writer.Content)
	common.SaveFile(errpath+"err.go",writer.Content)
}

