// Code generated by protoc-gen-gogo.
// source: unmarshalmerge.proto
// DO NOT EDIT!

/*
	Package unmarshalmerge is a generated protocol buffer package.

	It is generated from these files:
		unmarshalmerge.proto

	It has these top-level messages:
		Big
		BigUnsafe
		Sub
		IntMerge
*/
package unmarshalmerge

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Big struct {
	Sub              *Sub   `protobuf:"bytes,1,opt,name=Sub" json:"Sub,omitempty"`
	Number           *int64 `protobuf:"varint,2,opt,name=Number" json:"Number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Big) Reset()      { *m = Big{} }
func (*Big) ProtoMessage() {}

func (m *Big) GetSub() *Sub {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *Big) GetNumber() int64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type BigUnsafe struct {
	Sub              *Sub   `protobuf:"bytes,1,opt,name=Sub" json:"Sub,omitempty"`
	Number           *int64 `protobuf:"varint,2,opt,name=Number" json:"Number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BigUnsafe) Reset()      { *m = BigUnsafe{} }
func (*BigUnsafe) ProtoMessage() {}

func (m *BigUnsafe) GetSub() *Sub {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *BigUnsafe) GetNumber() int64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type Sub struct {
	SubNumber        *int64 `protobuf:"varint,1,opt,name=SubNumber" json:"SubNumber,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Sub) Reset()      { *m = Sub{} }
func (*Sub) ProtoMessage() {}

func (m *Sub) GetSubNumber() int64 {
	if m != nil && m.SubNumber != nil {
		return *m.SubNumber
	}
	return 0
}

type IntMerge struct {
	Int64            int64  `protobuf:"varint,1,req,name=Int64" json:"Int64"`
	Int32            int32  `protobuf:"varint,2,opt,name=Int32" json:"Int32"`
	Sint32           int32  `protobuf:"zigzag32,3,req,name=Sint32" json:"Sint32"`
	Sint64           int64  `protobuf:"zigzag64,4,opt,name=Sint64" json:"Sint64"`
	Uint64           uint64 `protobuf:"varint,5,opt,name=Uint64" json:"Uint64"`
	Uint32           uint32 `protobuf:"varint,6,req,name=Uint32" json:"Uint32"`
	Fixed64          uint64 `protobuf:"fixed64,7,opt,name=Fixed64" json:"Fixed64"`
	Fixed32          uint32 `protobuf:"fixed32,8,opt,name=Fixed32" json:"Fixed32"`
	Sfixed32         int32  `protobuf:"fixed32,9,req,name=Sfixed32" json:"Sfixed32"`
	Sfixed64         int64  `protobuf:"fixed64,10,opt,name=Sfixed64" json:"Sfixed64"`
	Bool             bool   `protobuf:"varint,11,opt,name=Bool" json:"Bool"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IntMerge) Reset()      { *m = IntMerge{} }
func (*IntMerge) ProtoMessage() {}

func (m *IntMerge) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *IntMerge) GetInt32() int32 {
	if m != nil {
		return m.Int32
	}
	return 0
}

func (m *IntMerge) GetSint32() int32 {
	if m != nil {
		return m.Sint32
	}
	return 0
}

func (m *IntMerge) GetSint64() int64 {
	if m != nil {
		return m.Sint64
	}
	return 0
}

func (m *IntMerge) GetUint64() uint64 {
	if m != nil {
		return m.Uint64
	}
	return 0
}

func (m *IntMerge) GetUint32() uint32 {
	if m != nil {
		return m.Uint32
	}
	return 0
}

func (m *IntMerge) GetFixed64() uint64 {
	if m != nil {
		return m.Fixed64
	}
	return 0
}

func (m *IntMerge) GetFixed32() uint32 {
	if m != nil {
		return m.Fixed32
	}
	return 0
}

func (m *IntMerge) GetSfixed32() int32 {
	if m != nil {
		return m.Sfixed32
	}
	return 0
}

func (m *IntMerge) GetSfixed64() int64 {
	if m != nil {
		return m.Sfixed64
	}
	return 0
}

func (m *IntMerge) GetBool() bool {
	if m != nil {
		return m.Bool
	}
	return false
}

func init() {
	proto.RegisterType((*Big)(nil), "unmarshalmerge.Big")
	proto.RegisterType((*BigUnsafe)(nil), "unmarshalmerge.BigUnsafe")
	proto.RegisterType((*Sub)(nil), "unmarshalmerge.Sub")
	proto.RegisterType((*IntMerge)(nil), "unmarshalmerge.IntMerge")
}
func (this *Big) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Big)
	if !ok {
		that2, ok := that.(Big)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Big")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Big but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Big but is not nil && this == nil")
	}
	if !this.Sub.Equal(that1.Sub) {
		return fmt.Errorf("Sub this(%v) Not Equal that(%v)", this.Sub, that1.Sub)
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return fmt.Errorf("Number this(%v) Not Equal that(%v)", *this.Number, *that1.Number)
		}
	} else if this.Number != nil {
		return fmt.Errorf("this.Number == nil && that.Number != nil")
	} else if that1.Number != nil {
		return fmt.Errorf("Number this(%v) Not Equal that(%v)", this.Number, that1.Number)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Big) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Big)
	if !ok {
		that2, ok := that.(Big)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Sub.Equal(that1.Sub) {
		return false
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return false
		}
	} else if this.Number != nil {
		return false
	} else if that1.Number != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BigUnsafe) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*BigUnsafe)
	if !ok {
		that2, ok := that.(BigUnsafe)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *BigUnsafe")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *BigUnsafe but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *BigUnsafe but is not nil && this == nil")
	}
	if !this.Sub.Equal(that1.Sub) {
		return fmt.Errorf("Sub this(%v) Not Equal that(%v)", this.Sub, that1.Sub)
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return fmt.Errorf("Number this(%v) Not Equal that(%v)", *this.Number, *that1.Number)
		}
	} else if this.Number != nil {
		return fmt.Errorf("this.Number == nil && that.Number != nil")
	} else if that1.Number != nil {
		return fmt.Errorf("Number this(%v) Not Equal that(%v)", this.Number, that1.Number)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *BigUnsafe) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*BigUnsafe)
	if !ok {
		that2, ok := that.(BigUnsafe)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Sub.Equal(that1.Sub) {
		return false
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return false
		}
	} else if this.Number != nil {
		return false
	} else if that1.Number != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Sub) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Sub)
	if !ok {
		that2, ok := that.(Sub)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Sub")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Sub but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Sub but is not nil && this == nil")
	}
	if this.SubNumber != nil && that1.SubNumber != nil {
		if *this.SubNumber != *that1.SubNumber {
			return fmt.Errorf("SubNumber this(%v) Not Equal that(%v)", *this.SubNumber, *that1.SubNumber)
		}
	} else if this.SubNumber != nil {
		return fmt.Errorf("this.SubNumber == nil && that.SubNumber != nil")
	} else if that1.SubNumber != nil {
		return fmt.Errorf("SubNumber this(%v) Not Equal that(%v)", this.SubNumber, that1.SubNumber)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Sub) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Sub)
	if !ok {
		that2, ok := that.(Sub)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.SubNumber != nil && that1.SubNumber != nil {
		if *this.SubNumber != *that1.SubNumber {
			return false
		}
	} else if this.SubNumber != nil {
		return false
	} else if that1.SubNumber != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IntMerge) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*IntMerge)
	if !ok {
		that2, ok := that.(IntMerge)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *IntMerge")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *IntMerge but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *IntMerge but is not nil && this == nil")
	}
	if this.Int64 != that1.Int64 {
		return fmt.Errorf("Int64 this(%v) Not Equal that(%v)", this.Int64, that1.Int64)
	}
	if this.Int32 != that1.Int32 {
		return fmt.Errorf("Int32 this(%v) Not Equal that(%v)", this.Int32, that1.Int32)
	}
	if this.Sint32 != that1.Sint32 {
		return fmt.Errorf("Sint32 this(%v) Not Equal that(%v)", this.Sint32, that1.Sint32)
	}
	if this.Sint64 != that1.Sint64 {
		return fmt.Errorf("Sint64 this(%v) Not Equal that(%v)", this.Sint64, that1.Sint64)
	}
	if this.Uint64 != that1.Uint64 {
		return fmt.Errorf("Uint64 this(%v) Not Equal that(%v)", this.Uint64, that1.Uint64)
	}
	if this.Uint32 != that1.Uint32 {
		return fmt.Errorf("Uint32 this(%v) Not Equal that(%v)", this.Uint32, that1.Uint32)
	}
	if this.Fixed64 != that1.Fixed64 {
		return fmt.Errorf("Fixed64 this(%v) Not Equal that(%v)", this.Fixed64, that1.Fixed64)
	}
	if this.Fixed32 != that1.Fixed32 {
		return fmt.Errorf("Fixed32 this(%v) Not Equal that(%v)", this.Fixed32, that1.Fixed32)
	}
	if this.Sfixed32 != that1.Sfixed32 {
		return fmt.Errorf("Sfixed32 this(%v) Not Equal that(%v)", this.Sfixed32, that1.Sfixed32)
	}
	if this.Sfixed64 != that1.Sfixed64 {
		return fmt.Errorf("Sfixed64 this(%v) Not Equal that(%v)", this.Sfixed64, that1.Sfixed64)
	}
	if this.Bool != that1.Bool {
		return fmt.Errorf("Bool this(%v) Not Equal that(%v)", this.Bool, that1.Bool)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *IntMerge) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*IntMerge)
	if !ok {
		that2, ok := that.(IntMerge)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Int64 != that1.Int64 {
		return false
	}
	if this.Int32 != that1.Int32 {
		return false
	}
	if this.Sint32 != that1.Sint32 {
		return false
	}
	if this.Sint64 != that1.Sint64 {
		return false
	}
	if this.Uint64 != that1.Uint64 {
		return false
	}
	if this.Uint32 != that1.Uint32 {
		return false
	}
	if this.Fixed64 != that1.Fixed64 {
		return false
	}
	if this.Fixed32 != that1.Fixed32 {
		return false
	}
	if this.Sfixed32 != that1.Sfixed32 {
		return false
	}
	if this.Sfixed64 != that1.Sfixed64 {
		return false
	}
	if this.Bool != that1.Bool {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Big) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&unmarshalmerge.Big{")
	if this.Sub != nil {
		s = append(s, "Sub: "+fmt.Sprintf("%#v", this.Sub)+",\n")
	}
	if this.Number != nil {
		s = append(s, "Number: "+valueToGoStringUnmarshalmerge(this.Number, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BigUnsafe) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&unmarshalmerge.BigUnsafe{")
	if this.Sub != nil {
		s = append(s, "Sub: "+fmt.Sprintf("%#v", this.Sub)+",\n")
	}
	if this.Number != nil {
		s = append(s, "Number: "+valueToGoStringUnmarshalmerge(this.Number, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Sub) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&unmarshalmerge.Sub{")
	if this.SubNumber != nil {
		s = append(s, "SubNumber: "+valueToGoStringUnmarshalmerge(this.SubNumber, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IntMerge) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
	s = append(s, "&unmarshalmerge.IntMerge{")
	s = append(s, "Int64: "+fmt.Sprintf("%#v", this.Int64)+",\n")
	s = append(s, "Int32: "+fmt.Sprintf("%#v", this.Int32)+",\n")
	s = append(s, "Sint32: "+fmt.Sprintf("%#v", this.Sint32)+",\n")
	s = append(s, "Sint64: "+fmt.Sprintf("%#v", this.Sint64)+",\n")
	s = append(s, "Uint64: "+fmt.Sprintf("%#v", this.Uint64)+",\n")
	s = append(s, "Uint32: "+fmt.Sprintf("%#v", this.Uint32)+",\n")
	s = append(s, "Fixed64: "+fmt.Sprintf("%#v", this.Fixed64)+",\n")
	s = append(s, "Fixed32: "+fmt.Sprintf("%#v", this.Fixed32)+",\n")
	s = append(s, "Sfixed32: "+fmt.Sprintf("%#v", this.Sfixed32)+",\n")
	s = append(s, "Sfixed64: "+fmt.Sprintf("%#v", this.Sfixed64)+",\n")
	s = append(s, "Bool: "+fmt.Sprintf("%#v", this.Bool)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUnmarshalmerge(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringUnmarshalmerge(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func NewPopulatedBig(r randyUnmarshalmerge, easy bool) *Big {
	this := &Big{}
	if r.Intn(10) != 0 {
		this.Sub = NewPopulatedSub(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Number = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 3)
	}
	return this
}

func NewPopulatedBigUnsafe(r randyUnmarshalmerge, easy bool) *BigUnsafe {
	this := &BigUnsafe{}
	if r.Intn(10) != 0 {
		this.Sub = NewPopulatedSub(r, easy)
	}
	if r.Intn(10) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Number = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 3)
	}
	return this
}

func NewPopulatedSub(r randyUnmarshalmerge, easy bool) *Sub {
	this := &Sub{}
	if r.Intn(10) != 0 {
		v3 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		this.SubNumber = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 2)
	}
	return this
}

func NewPopulatedIntMerge(r randyUnmarshalmerge, easy bool) *IntMerge {
	this := &IntMerge{}
	this.Int64 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Int64 *= -1
	}
	this.Int32 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Int32 *= -1
	}
	this.Sint32 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Sint32 *= -1
	}
	this.Sint64 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Sint64 *= -1
	}
	this.Uint64 = uint64(uint64(r.Uint32()))
	this.Uint32 = uint32(r.Uint32())
	this.Fixed64 = uint64(uint64(r.Uint32()))
	this.Fixed32 = uint32(r.Uint32())
	this.Sfixed32 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Sfixed32 *= -1
	}
	this.Sfixed64 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Sfixed64 *= -1
	}
	this.Bool = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 12)
	}
	return this
}

type randyUnmarshalmerge interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneUnmarshalmerge(r randyUnmarshalmerge) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringUnmarshalmerge(r randyUnmarshalmerge) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneUnmarshalmerge(r)
	}
	return string(tmps)
}
func randUnrecognizedUnmarshalmerge(r randyUnmarshalmerge, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldUnmarshalmerge(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldUnmarshalmerge(data []byte, r randyUnmarshalmerge, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(v5))
	case 1:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateUnmarshalmerge(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (this *Big) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Big{`,
		`Sub:` + strings.Replace(fmt.Sprintf("%v", this.Sub), "Sub", "Sub", 1) + `,`,
		`Number:` + valueToStringUnmarshalmerge(this.Number) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BigUnsafe) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BigUnsafe{`,
		`Sub:` + strings.Replace(fmt.Sprintf("%v", this.Sub), "Sub", "Sub", 1) + `,`,
		`Number:` + valueToStringUnmarshalmerge(this.Number) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Sub) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Sub{`,
		`SubNumber:` + valueToStringUnmarshalmerge(this.SubNumber) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IntMerge) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IntMerge{`,
		`Int64:` + fmt.Sprintf("%v", this.Int64) + `,`,
		`Int32:` + fmt.Sprintf("%v", this.Int32) + `,`,
		`Sint32:` + fmt.Sprintf("%v", this.Sint32) + `,`,
		`Sint64:` + fmt.Sprintf("%v", this.Sint64) + `,`,
		`Uint64:` + fmt.Sprintf("%v", this.Uint64) + `,`,
		`Uint32:` + fmt.Sprintf("%v", this.Uint32) + `,`,
		`Fixed64:` + fmt.Sprintf("%v", this.Fixed64) + `,`,
		`Fixed32:` + fmt.Sprintf("%v", this.Fixed32) + `,`,
		`Sfixed32:` + fmt.Sprintf("%v", this.Sfixed32) + `,`,
		`Sfixed64:` + fmt.Sprintf("%v", this.Sfixed64) + `,`,
		`Bool:` + fmt.Sprintf("%v", this.Bool) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUnmarshalmerge(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Big) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnmarshalmerge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Big: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Big: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sub", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnmarshalmerge
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sub == nil {
				m.Sub = &Sub{}
			}
			if err := m.Sub.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Number = &v
		default:
			iNdEx = preIndex
			skippy, err := skipUnmarshalmerge(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnmarshalmerge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Sub) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnmarshalmerge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Sub: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sub: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubNumber", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SubNumber = &v
		default:
			iNdEx = preIndex
			skippy, err := skipUnmarshalmerge(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnmarshalmerge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IntMerge) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnmarshalmerge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IntMerge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntMerge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64", wireType)
			}
			m.Int64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Int64 |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32", wireType)
			}
			m.Int32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Int32 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sint32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Sint32 = v
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.Sint64 = int64(v)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint64", wireType)
			}
			m.Uint64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Uint64 |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint32", wireType)
			}
			m.Uint32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Uint32 |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fixed64", wireType)
			}
			m.Fixed64 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Fixed64 = uint64(data[iNdEx-8])
			m.Fixed64 |= uint64(data[iNdEx-7]) << 8
			m.Fixed64 |= uint64(data[iNdEx-6]) << 16
			m.Fixed64 |= uint64(data[iNdEx-5]) << 24
			m.Fixed64 |= uint64(data[iNdEx-4]) << 32
			m.Fixed64 |= uint64(data[iNdEx-3]) << 40
			m.Fixed64 |= uint64(data[iNdEx-2]) << 48
			m.Fixed64 |= uint64(data[iNdEx-1]) << 56
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fixed32", wireType)
			}
			m.Fixed32 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			m.Fixed32 = uint32(data[iNdEx-4])
			m.Fixed32 |= uint32(data[iNdEx-3]) << 8
			m.Fixed32 |= uint32(data[iNdEx-2]) << 16
			m.Fixed32 |= uint32(data[iNdEx-1]) << 24
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sfixed32", wireType)
			}
			m.Sfixed32 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			m.Sfixed32 = int32(data[iNdEx-4])
			m.Sfixed32 |= int32(data[iNdEx-3]) << 8
			m.Sfixed32 |= int32(data[iNdEx-2]) << 16
			m.Sfixed32 |= int32(data[iNdEx-1]) << 24
			hasFields[0] |= uint64(0x00000008)
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sfixed64", wireType)
			}
			m.Sfixed64 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Sfixed64 = int64(data[iNdEx-8])
			m.Sfixed64 |= int64(data[iNdEx-7]) << 8
			m.Sfixed64 |= int64(data[iNdEx-6]) << 16
			m.Sfixed64 |= int64(data[iNdEx-5]) << 24
			m.Sfixed64 |= int64(data[iNdEx-4]) << 32
			m.Sfixed64 |= int64(data[iNdEx-3]) << 40
			m.Sfixed64 |= int64(data[iNdEx-2]) << 48
			m.Sfixed64 |= int64(data[iNdEx-1]) << 56
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Bool = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipUnmarshalmerge(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnmarshalmerge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Int64")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Sint32")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Uint32")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Sfixed32")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUnmarshalmerge(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnmarshalmerge
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnmarshalmerge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUnmarshalmerge
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUnmarshalmerge
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUnmarshalmerge(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUnmarshalmerge = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnmarshalmerge   = fmt.Errorf("proto: integer overflow")
)

func (m *BigUnsafe) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnmarshalmergeUnsafe
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BigUnsafe: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BigUnsafe: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sub", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmergeUnsafe
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUnmarshalmergeUnsafe
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sub == nil {
				m.Sub = &Sub{}
			}
			if err := m.Sub.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnmarshalmergeUnsafe
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Number = &v
		default:
			iNdEx = preIndex
			skippy, err := skipUnmarshalmergeUnsafe(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUnmarshalmergeUnsafe
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUnmarshalmergeUnsafe(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnmarshalmergeUnsafe
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnmarshalmergeUnsafe
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnmarshalmergeUnsafe
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUnmarshalmergeUnsafe
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUnmarshalmergeUnsafe
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUnmarshalmergeUnsafe(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUnmarshalmergeUnsafe = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnmarshalmergeUnsafe   = fmt.Errorf("proto: integer overflow")
)
