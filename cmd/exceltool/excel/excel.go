package excel

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/wudiliujie/common/convert"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/writer"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const Cfg_Err = "Cfg_Err"

type CfgProp struct {
	Name     string
	Type     string
	Desc     string
	IsClient bool
}
type CfgStruct struct {
	Name     string
	IsClient bool
	Desc     string
	Props    []CfgProp
}
type OutData struct {
	Ver   string
	Items map[string]interface{}
}
type OutServerData struct {
	Ver   string
	Items []map[string]interface{}
}
type OutClient struct {
	Items map[string]interface{}
}

var Pf = make(map[int32]string)

func init() {
	Pf[1] = "ios"
	Pf[2] = "android"
	Pf[4] = "h5"
	Pf[8] = "xyx_android"
	Pf[16] = "xyx_ios"
}

type Table struct {
	Struct      *CfgStruct
	ServerItems []map[string]interface{}
	ClientItems []map[string]interface{}
	//客户端生成5个平台
	PfData map[int32][]map[string]interface{}
}

var OnHandleData func(table *Table)

func CreateTable() *Table {
	table := &Table{}
	table.Struct = &CfgStruct{}
	table.Struct.Props = make([]CfgProp, 0)
	return table
}

var gVer = time.Now().Format("200601021504")

type Excel struct {
	Tables     []*Table
	maping     []string //对应关系
	waitGroutp *sync.WaitGroup
}

func (e *Excel) InitTable(sourcePath string) {
	e.maping = make([]string, 0)
	e.waitGroutp = &sync.WaitGroup{}
	files, err := WalkDir(sourcePath, ".xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	//这里执行多线程
	for _, file := range files {
		e.ExcelGetCfgStructByFile(file)
	}
	e.waitGroutp.Wait()
	log.Debug("处理完成")
}
func (e *Excel) ExcelGetCfgStructByFile(file string) {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	sheets := xlsx.GetSheetMap()
	for _, v := range sheets {
		if strings.HasPrefix(v, "!") {
			e.waitGroutp.Add(1)
			go func(x *excelize.File, name string) {
				e.ExcelGetCfgStructBySheet(x, name)
				e.waitGroutp.Done()
			}(xlsx, v)
		}
	}
}

func (e *Excel) ExcelGetCfgStructBySheet(xlsx *excelize.File, sheetName string) {
	// Get all the rows in the Sheet1.
	log.Debug("开始 %v>>%v\r\n", xlsx.Path, sheetName)
	rows := xlsx.GetRows(sheetName)

	if len(rows) < 3 {
		return
	}
	//读取数据
	firstRow := rows[0]
	secondRow := rows[1]
	thridRow := rows[2]
	columnCount := len(firstRow)

	str := strings.Split(sheetName, "!")
	table := CreateTable()
	table.Struct.Name = str[1]
	if len(str) >= 3 {
		table.Struct.Desc = str[2]
	}
	for i := 0; i < columnCount; i++ {
		var p CfgProp
		p.Desc = thridRow[i]
		p.Name = firstRow[i]
		if p.Name == "" {
			continue
		}

		types := strings.Split(secondRow[i], "_")
		p.Type = types[0]
		if len(types) >= 2 && types[1] == "C" {
			table.Struct.IsClient = true
			p.IsClient = true

		} else {
			p.IsClient = false
		}
		if len(types) == 1 || len(types) == 2 && types[1] != "D" {
			table.Struct.Props = append(table.Struct.Props, p)
		}
	}
	table.ServerItems = make([]map[string]interface{}, 0)
	table.ClientItems = make([]map[string]interface{}, 0)

	for i := 3; i < len(rows); i++ {
		serverItem := make(map[string]interface{})
		clientItem := make(map[string]interface{})
		if rows[i][0] == "" {
			continue
		}

		for j := 0; j < len(table.Struct.Props); j++ {
			if firstRow[j] == "" {
				continue
			}
			client := table.Struct.Props[j].IsClient
			if secondRow[j] == "INT" || secondRow[j] == "INT_C" {
				v, _ := strconv.Atoi(rows[i][j])
				if v != 0 {
					serverItem[firstRow[j]] = v
					if client {
						clientItem[firstRow[j]] = v
					}

				}
			} else if secondRow[j] == "INT64" || secondRow[j] == "INT64_C" {
				v, _ := strconv.ParseInt(rows[i][j], 10, 64)
				if v != 0 {
					serverItem[firstRow[j]] = v
					if client {
						clientItem[firstRow[j]] = v
					}

				}
			} else {
				if rows[i][j] != "" {
					serverItem[firstRow[j]] = rows[i][j]
					if client {
						clientItem[firstRow[j]] = rows[i][j]

					}

				}
			}
		}
		table.ServerItems = append(table.ServerItems, serverItem)
		if table.Struct.IsClient {
			table.ClientItems = append(table.ClientItems, clientItem)
		}
	}
	if OnHandleData != nil {
		OnHandleData(table)
	}
	e.Tables = append(e.Tables, table)
	e.maping = append(e.maping, fmt.Sprintf("%v\t%v", xlsx.Path, table.Struct.Name))
	log.Debug("结束  %v>>%v    %v,%v\r\n", xlsx.Path, sheetName, len(rows), columnCount)
}

//生成go文件
func (e *Excel) ExcelToGo(fileName string) {
	var writer writer.Writer
	writer.AddLine("package tb")
	writer.AddLine("import (")
	writer.AddLine(`	"encoding/json"`)
	writer.AddLine(")")
	for _, table := range e.Tables {
		writer.AddLineFmt("type %v struct {", table.Struct.Name)
		for _, p := range table.Struct.Props {
			stype := "string"
			if p.Type == "INT" {
				stype = "int32"
			}
			if p.Type == "INT64" {
				stype = "int64"
			}
			if p.Type == "STRING" {
				stype = "string"
			}
			writer.AddLineFmt("%v		%v", p.Name, stype)
		}
		writer.AddLine("}")

		//writer.AddLineFmt("var %v_Pool struct {", table.Struct.Name)
		//writer.AddLineFmt("Items			[]*%v", table.Struct.Name)
		//writer.AddLine("}")
	}
	writer.AddLine("type  CfgData struct{")
	writer.AddLine("Ver string")
	writer.AddLine("Items *CfgDataAll")
	writer.AddLine("}")
	//
	writer.AddLineFmt("type CfgDataAll struct {")
	for _, table := range e.Tables {
		writer.AddLineFmt(`%v []*%v`, table.Struct.Name, table.Struct.Name)
	}
	writer.AddLineFmt("}")

	writer.AddLineFmt("var Data *CfgDataAll")
	writer.AddLineFmt("var Ver string")

	//写入init
	writer.AddLine("func InitCfgData(data []byte)error{")
	writer.AddLine("var errcode error")
	writer.AddLine("var allData CfgData")
	writer.AddLine("errcode = json.Unmarshal(data, &allData)")
	writer.AddLine("if errcode != nil {")
	writer.AddLine("return errcode")
	writer.AddLine("}")
	writer.AddLine("Ver = allData.Ver")
	writer.AddLine("Data = allData.Items")
	writer.AddLine(`return nil`)
	writer.AddLine("}")

	Save(fileName, writer.Content.String())
}

//生成go文件
func (e *Excel) ExcelToCSharpClient(fileName string) {
	var writer writer.Writer
	for _, table := range e.Tables {
		if table.Struct.IsClient == false {
			continue
		}
		writer.AddLineFmt("public class %v", table.Struct.Name)
		writer.AddLineFmt("{")
		for _, p := range table.Struct.Props {
			if p.IsClient {
				stype := "string"
				value := `""`
				if p.Type == "INT" {
					stype = "int"
					value = "0"
				}
				if p.Type == "INT64" {
					stype = "long"
					value = "0"
				}
				if p.Type == "STRING" {
					stype = "string"
				}
				writer.AddLineFmt("public %v %v = %v;", stype, p.Name, value)
			}
		}
		writer.AddLine("}")
	}
	Save(fileName, writer.Content.String())
}

//根据消息提示文件，生成err.go
func (e *Excel) ExcelToErrGo(fileName string) {
	for _, table := range e.Tables {
		if table.Struct.Name == Cfg_Err {
			write := writer.Writer{}
			write.AddLine("package errmsg")
			write.AddLine("type ErrCode int32")
			write.AddLine("const (")

			for _, row := range table.ServerItems {
				l := int((50 - len(convert.ToString(row["Name"]))) / 4)
				if row["Id"] == nil {
					write.AddLineFmt("\t%v"+strings.Repeat("\t", l)+"= %v   //%v", row["Name"], 0, row["MSG"])
				} else {
					write.AddLineFmt("\t%v"+strings.Repeat("\t", l)+"= %v   //%v", row["Name"], row["Id"], row["MSG"])
				}

			}

			write.AddLine(")")
			Save(fileName, write.Content.String())
		}
	}
}

func (e *Excel) ExcelToClient(targetPath string) {
	if !strings.HasSuffix(targetPath, "\\") {
		targetPath += "\\"
	}
	ok, _ := PathExists(targetPath)
	if !ok {
		os.MkdirAll(targetPath, 0777)
	}
	ok1, _ := PathExists(targetPath + "tmp\\")
	if !ok1 {
		os.MkdirAll(targetPath+"tmp\\", 0777)
	}

	outData := &OutData{}
	outData.Items = make(map[string]interface{})
	for _, table := range e.Tables {
		if len(table.ClientItems) > 0 {
			outData.Items[table.Struct.Name] = table.ClientItems
			//直接写成文件
			strjson1, _ := json.MarshalIndent(table.ClientItems, "", "\t")
			Save(targetPath+"tmp\\"+table.Struct.Name+".json", string(strjson1))
		}
	}
	outData.Ver = gVer
	strjson, _ := json.MarshalIndent(outData, "", "\t")
	Save(targetPath+"cfg_data.json", string(strjson))
	for k := range Pf {
		e.ExcelToClientPf(targetPath, k)
	}
}

func (e *Excel) ExcelToClientPf(targetPath string, pf int32) {
	if !strings.HasSuffix(targetPath, "\\") {
		targetPath += "\\"
	}
	ok, _ := PathExists(targetPath)
	if !ok {
		os.MkdirAll(targetPath, 0777)
	}
	tmp := targetPath + "tmp" + convert.ToString(pf) + "\\"
	ok1, _ := PathExists(tmp)
	if !ok1 {
		os.MkdirAll(tmp, 0777)
	}

	outData := &OutData{}
	outData.Items = make(map[string]interface{})
	for _, table := range e.Tables {
		if len(table.PfData[pf]) > 0 {
			outData.Items[table.Struct.Name] = table.PfData[pf]
			//直接写成文件
			strjson1, _ := json.MarshalIndent(table.PfData[pf], "", "\t")
			Save(tmp+table.Struct.Name+".json", string(strjson1))
		}
	}
	outData.Ver = gVer
	strjson, _ := json.MarshalIndent(outData, "", "\t")
	Save(targetPath+"cfg_data_"+Pf[pf]+".json", string(strjson))
}

func (e *Excel) ExcelToServer(targetPath string) {

	ok, _ := PathExists(targetPath)
	if !ok {
		os.MkdirAll(targetPath, 0777)
	}

	//for _, table := range e.Tables {
	//	outData := &OutServerData{}
	//	outData.Items = table.ServerItems
	//	outData.Ver = "0"
	//	strjson, _ := json.MarshalIndent(outData, "", "\t")
	//	Save(targetPath+table.Struct.Name+".json", string(strjson))
	//}
	//
	tmp := targetPath + "tmp" + "\\"
	ok1, _ := PathExists(tmp)
	if !ok1 {
		os.MkdirAll(tmp, 0777)
	}
	outData := &OutData{}
	outData.Items = make(map[string]interface{})
	for _, table := range e.Tables {
		outData.Items[table.Struct.Name] = table.ServerItems
		//直接写成文件
		strjson1, _ := json.MarshalIndent(table.ServerItems, "", "\t")
		Save(targetPath+"tmp\\"+table.Struct.Name+".json", string(strjson1))
	}
	outData.Ver = gVer
	strjson, _ := json.MarshalIndent(outData, "", "\t")
	Save(targetPath+"cfg_data.json", string(strjson))
}

//保存映射文件
func (e *Excel) SaveMaping(filename string) {
	Save(filename, strings.Join(e.maping, "\r\n"))
}

func (e *Excel) SaveCfgDataTs(filename string) {
	writer := &writer.Writer{}
	writer.AddLine("export class CfgData{")
	for _, v := range e.Tables {
		if v.Struct.IsClient {
			writer.AddLineFmt("static %vs : Map<number, %v>;", v.Struct.Name, v.Struct.Name)
		}
	}
	writer.AddLine("static Init(){")
	for _, v := range e.Tables {
		if v.Struct.IsClient {
			writer.AddLineFmt("this.%vs= new Map<number,%v>();", v.Struct.Name, v.Struct.Name)
			//for _,data:=range v.ClientItems{
			//	writer.AddLineFmt("this.%vs.set(%v,new %v(%v))",v.Struct.Name,data[v.Struct.Props[0].Name],v.Struct.Name,GetFuncParam(v.Struct.Props,data))
			//}
		}
	}

	writer.AddLine("}")

	writer.AddLine("}")
	e.WriteCfgDataClass(writer)
	Save(filename, writer.Content.String())
}
func (e *Excel) WriteCfgDataClass(writer *writer.Writer) {
	for _, v := range e.Tables {
		if v.Struct.IsClient == true {
			writer.AddLineFmt("export class %v{", v.Struct.Name)
			writer.AddLineFmt("constructor(%v){", GetConstructorParam(v.Struct.Props))
			idx := 0
			for _, p := range v.Struct.Props {
				if p.IsClient {
					writer.AddLineFmt("this.%v= p%v", p.Name, idx)
					idx++
				}
			}
			writer.AddLine("}")
			for _, p := range v.Struct.Props {
				if p.IsClient {
					writer.AddLineFmt("%v:%v", p.Name, GetClientType(p.Type))
					idx++
				}
			}
			writer.AddLine("}")
		}
	}
}
func (e *Excel) WriteCfgDataText(filename string) {
	writer := &writer.Writer{}
	for _, v := range e.Tables {
		if v.Struct.IsClient {
			writer.AddLineFmt("[Table:%v]", v.Struct.Name)
			for _, data := range v.ClientItems {
				writer.AddLineFmt("%v", GetFuncParam(v.Struct.Props, data))
			}
		}
	}
	Save(filename, writer.Content.String())
}

func (e *Excel) WriteCfgProto(filename string) {
	writer := &writer.Writer{}
	writer.AddLineFmt(`syntax = "proto3";`)
	writer.AddLineFmt("package msg;")
	for _, v := range e.Tables {
		if v.Struct.IsClient {
			writer.AddLineFmt("message %v {", v.Struct.Name)
			idx := int32(1)
			for _, p := range v.Struct.Props {
				writer.AddLineFmt(" %v %v = %v; //%v", GetPropType(p.Type), p.Name, idx, strings.Replace(p.Desc, "\n", "\t", -1))
				idx++
			}
			writer.AddLineFmt("}")
		}
	}
	//生成总结构体
	writer.AddLineFmt("message AllData {")
	idx := int32(1)
	for _, v := range e.Tables {
		if v.Struct.IsClient {
			writer.AddLineFmt("repeated %v %v = %v;", v.Struct.Name, v.Struct.Name, idx)
			idx++
		}
	}
	writer.AddLineFmt("}")
	writer.AddLineFmt("message Data{")
	writer.AddLineFmt("AllData Items =1;")
	writer.AddLineFmt("string Ver =2;")
	writer.AddLineFmt("}")

	Save(filename, writer.Content.String())
}

func GetConstructorParam(prop []CfgProp) string {
	count := 0
	for _, v := range prop {
		if v.IsClient {
			count++
		}
	}
	p := "p0"
	for i := 1; i < count; i++ {
		p += ",p" + strconv.Itoa(i)
	}
	return p
}
func GetFuncParam(prop []CfgProp, data map[string]interface{}) string {
	ret := ""
	for _, v := range prop {
		if v.Type == "INT" {
			ret += convert.ToString(data[v.Name]) + ","
		} else {
			ret += "\"" + convert.ToString(data[v.Name]) + "\","
		}
	}
	return strings.TrimRight(ret, ",")
}
func GetClientType(t string) string {
	if t == "INT" {
		return "number"
	}
	return "string"
}
func GetPropType(t string) string {
	stype := "string"
	if t == "INT" {
		stype = "int32"
	}
	if t == "INT64" {
		stype = "int64"
	}
	if t == "STRING" {
		stype = "string"
	}
	return stype
}
