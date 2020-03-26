package main

import (
	"encoding/xml"
	"github.com/wudiliujie/common/rpath"
	"github.com/wudiliujie/common/writer"
	"io/ioutil"
	"os"
	"strings"
)

type Attrs struct {
	As    []A `xml:"a"`
	write *writer.Writer
}

func (a *Attrs) MakeGo(fileName string) {
	a.write = &writer.Writer{}
	a.write.AddLine("package consts")
	a.write.AddLine("type AttrType int32")
	a.write.AddLine("const (")
	for _, v := range a.As {
		a.write.AddLineFmt("	%v         AttrType = %v //%v", v.N, v.I, v.D)
	}
	a.write.AddLine(")")

	//var SendPlayerInt = []AttrType{
	a.write.AddLine("var SendPlayerInt = []AttrType{")
	for _, v := range a.As {
		if v.T == 0 && v.C {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")

	a.write.AddLine("var SendPlayerStr = []AttrType{")
	for _, v := range a.As {
		if v.T == 1 && v.C {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")

	a.write.AddLine("var SendOtherInt = []AttrType{")
	for _, v := range a.As {
		if v.T == 0 && v.OC {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	a.write.AddLine("var SendOtherStr = []AttrType{")
	for _, v := range a.As {
		if v.T == 1 && v.OC {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")

	a.write.AddLine("var SaveInt = []AttrType{")
	for _, v := range a.As {
		if v.T == 0 && v.S {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	a.write.AddLine("var SaveStr = []AttrType{")
	for _, v := range a.As {
		if v.T == 1 && v.S {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	a.write.AddLine("var FightAttr = []AttrType{")
	for _, v := range a.As {
		if v.F {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")

	a.write.AddLine("var ResetAttr = []AttrType{")
	for _, v := range a.As {
		if v.R && v.T < 1 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")

	a.write.AddLine("var StrResetAttr = []AttrType{")
	for _, v := range a.As {
		if v.R && v.T == 1 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	//保存基础信息
	a.write.AddLine("var IntBaseAttr = []AttrType{")
	for _, v := range a.As {
		if v.B && v.T == 0 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	a.write.AddLine("var StrBaseAttr = []AttrType{")
	for _, v := range a.As {
		if v.B && v.T == 1 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	//跨服属性
	a.write.AddLine("var IntCorssAttr = []AttrType{")
	for _, v := range a.As {
		if v.K && v.T == 0 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	a.write.AddLine("var StrCrossAttr = []AttrType{")
	for _, v := range a.As {
		if v.B && v.T == 1 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	a.write.AddLine("var BaseFightAttr = []AttrType{")
	for _, v := range a.As {
		if v.BF {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	//排行榜属性
	a.write.AddLine("var IntRankAttr = []AttrType{")
	for _, v := range a.As {
		if v.Rank && v.T == 0 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")
	//计数器属性
	a.write.AddLine("var CounterAttr = []AttrType{")
	for _, v := range a.As {
		if v.M && v.T == 0 {
			a.write.AddLineFmt("	%v,", v.N)
		}
	}
	a.write.AddLine("}")

	a.write.AddLine("var AttrName = map[AttrType]string{")
	for _, v := range a.As {
		a.write.AddLineFmt("%v:    \"%v\",", v.N, v.D)
	}
	a.write.AddLine("}")
	rpath.SaveFile(fileName, a.write.Content.String())
}
func (a *Attrs) MakeTs(fileName string) {
	a.write = &writer.Writer{}
	a.write.AddLine("export class RoleModelData {")
	a.write.AddLine("static AttrType = {")
	for _, v := range a.As {
		if v.C {
			a.write.AddLineFmt(`%v: "%v", //%v`, v.I, ReplaceClientName(v.N), v.D)
		}

	}
	a.write.AddLine("}")
	a.write.AddLine("")
	for _, v := range a.As {
		if v.C {
			a.write.AddLine("/*")
			a.write.AddLineFmt("* %v", v.D)
			a.write.AddLine("*/")
			if v.T == 1 {
				a.write.AddLineFmt("%v = \"\"; ", ReplaceClientName(v.N))
			} else {
				a.write.AddLineFmt("%v = 0; ", ReplaceClientName(v.N))
			}
		}
	}
	a.write.AddLine("}")

	rpath.SaveFile(fileName, a.write.Content.String())
}
func ReplaceClientName(n string) string {
	return strings.Replace(n, "AttrType_", "", 1)
}

type A struct {
	I    int32  `xml:"i,attr"`
	N    string `xml:"n,attr"`
	D    string `xml:"d,attr"`
	F    bool   `xml:"f,attr"`
	C    bool   `xml:"c,attr"`
	OC   bool   `xml:"oc,attr"`
	S    bool   `xml:"s,attr"`
	T    int32  `xml:"t,attr"`
	R    bool   `xml:"r,attr"`    //是否重置
	B    bool   `xml:"b,attr"`    //是不是基础信息
	K    bool   `xml:"k,attr"`    //属性发生变化，是否同步到跨服
	M    bool   `xml:"m,attr"`    //属性发生变化，是否记录到计数器里面
	Rank bool   `xml:"rank,attr"` //排行榜属性
	BF   bool   `xml:"bf,attr"`
}

func main() {
	f, err := os.Open(`attr.xml`)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	// v := make(map[string]interface{})
	var v Attrs
	err = xml.Unmarshal(data, &v)
	if err != nil {
		panic(err)
	}
	v.MakeGo(`AttrType.go`)
	v.MakeTs("RoleModelData.ts")
}
