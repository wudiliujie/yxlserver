package generate

import (
	"fmt"
	"github.com/wudiliujie/common/rpath"
	"github.com/wudiliujie/common/writer"
)

type P struct {
	I     int32  `xml:"i,attr"`
	N     string `xml:"n,attr"`
	T     string `xml:"t,attr"`
	Array bool   `xml:"a,attr"` //是否是数组
	D     string `xml:"d,attr"` //描述
}
type Package struct {
	Id int32  `xml:"id,attr"`
	N  string `xml:"n,attr"`
	D  string `xml:"d,attr"` //描述
	C  bool   `xml:"c,attr"` //是否客户端需要

	Prop []*P `xml:"p"`
}
type Packages struct {
	Package    []*Package `xml:"package"`
	write      *writer.Writer
	customType map[string]int32
}

func (a *Packages) Init() {
	a.customType = make(map[string]int32)
	for _, v := range a.Package {
		a.customType[v.N] = 1
	}
}
func (a *Packages) existType(t string) bool {
	_, ok := a.customType[t]
	return ok
}

func (a *Packages) MakeGo(fileName string) {
	a.write = &writer.Writer{}
	a.write.AddLine("package msg")
	a.write.AddLine("import (")
	a.write.AddLine(`	"github.com/golang/protobuf/proto"`)
	a.write.AddLine(`	"github.com/wudiliujie/common/network"`)
	a.write.AddLine(`	"github.com/wudiliujie/common/network/protobuf"`)
	a.write.AddLineFmt(")")
	a.write.AddLineFmt("var Processor = protobuf.NewProcessor()")
	a.write.AddLineFmt("func init() {")
	for _, pck := range a.Package {
		if pck.Id > 0 {
			a.write.AddLineFmt("Processor.Register(PCKID_%v, func() network.IMessage { return new(%v) })", pck.N, pck.N)
		}
	}
	a.write.AddLine("}")

	for _, pck := range a.Package {
		if pck.Id > 0 {
			a.write.AddLineFmt("const PCKID_%v  = %v  //%v", pck.N, pck.Id, pck.D)
		}
		a.write.AddLineFmt("//%v", pck.D)
		a.write.AddLineFmt("type %v struct {", pck.N)
		for _, p := range pck.Prop {
			a.write.AddLineFmt("//%v", p.D)
			a.write.AddLineFmt("%v %v `protobuf:\"%v,%v,%v,name=%v,proto3\" json:\"%v,omitempty\"`", p.N, a.GetType(p), a.GetProtoBufType(p.T), p.I, a.GetOpt(p.Array), p.N, p.N)
		}
		a.write.AddLine("}")
		//
		a.write.AddLineFmt("func (m *%v) Reset()         { *m = %v{} }", pck.N, pck.N)
		a.write.AddLineFmt("func (m *%v) String() string { return proto.CompactTextString(m) }", pck.N)
		a.write.AddLineFmt("func (*%v) ProtoMessage()    {}", pck.N)
		if pck.Id > 0 {
			a.write.AddLineFmt("func (m *%v) GetId() uint16   { return uint16(PCKID_%v) }", pck.N, pck.N)
		}
	}
	rpath.SaveFile(fileName, a.write.Content)
}
func (a *Packages) GetProtoBufType(t string) string {
	switch t {
	case "string":
		return "bytes"
	case "int32":
		return "varint"
	case "int64":
		return "varint"
	default:
		//判断是否是自定义类型， 如果不是报错
		if a.existType(t) {
			return "bytes"
		} else {
			panic(fmt.Sprintf("类型不存在：%", t))
		}
	}
	return ""

}
func (a *Packages) GetOpt(array bool) string {
	if array {
		return "rep"
	} else {
		return "opt"
	}
}
func (a *Packages) GetType(p *P) string {
	if p.Array {
		if a.existType(p.T) {
			return "[]*" + p.T
		} else {
			return "[]" + p.T
		}
	}
	return p.T
}
