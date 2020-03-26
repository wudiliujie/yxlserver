package main

import (
	"github.com/wudiliujie/common/log"
	"os"
	"strings"
	"yxlserver/cmd/exceltool/excel"
)

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("os args of len(%v) less than 2", argsLen)
	}
	sourcePath := os.Args[1]
	targetPath := os.Args[2]
	if !strings.HasSuffix(targetPath, "\\") {
		targetPath += "\\"
	}
	os.MkdirAll(targetPath, 0777)
	os.MkdirAll(targetPath+`server\`, 0777)
	os.MkdirAll(targetPath+`client\`, 0777)
	e := excel.Excel{}
	e.InitTable(sourcePath)
	e.ExcelToGo(targetPath + "\\tbdata.go")
	//e.WriteCfgDataText(targetPath + "\\confdata.txt")
	//e.ExcelToCSharpClient(targetPath + "\\ConfigClass.cs")
	e.ExcelToClient(targetPath + `client\`)

	e.ExcelToServer(targetPath + `server\`)
	e.ExcelToErrGo(targetPath + "\\errcode.go")
	//e.SaveMaping(targetPath + "\\maping.txt")
	//e.SaveCfgDataTs(targetPath + "\\cfgdata.ts")

	//e.WriteCfgProto(targetPath + "\\data.proto")
	//e.WriteCfgDataText(targetPath + `client\cfgdata.txt`)

}
