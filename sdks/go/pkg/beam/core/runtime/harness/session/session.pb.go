// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

/*
Package session is a generated protocol buffer package.

It is generated from these files:
	session.proto

It has these top-level messages:
	Header
	Footer
	EntryHeader
	Entry
*/
package session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import org_apache_beam_model_fn_execution_v1 "github.com/apache/beam/sdks/go/pkg/beam/model/fnexecution_v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Kind int32

const (
	Kind_INVALID              Kind = 0
	Kind_INSTRUCTION_REQUEST  Kind = 1
	Kind_INSTRUCTION_RESPONSE Kind = 2
	Kind_DATA_RECEIVED        Kind = 3
	Kind_DATA_SENT            Kind = 4
	Kind_LOG_ENTRIES          Kind = 5
	Kind_HEADER               Kind = 6
	Kind_FOOTER               Kind = 7
)

var Kind_name = map[int32]string{
	0: "INVALID",
	1: "INSTRUCTION_REQUEST",
	2: "INSTRUCTION_RESPONSE",
	3: "DATA_RECEIVED",
	4: "DATA_SENT",
	5: "LOG_ENTRIES",
	6: "HEADER",
	7: "FOOTER",
}
var Kind_value = map[string]int32{
	"INVALID":              0,
	"INSTRUCTION_REQUEST":  1,
	"INSTRUCTION_RESPONSE": 2,
	"DATA_RECEIVED":        3,
	"DATA_SENT":            4,
	"LOG_ENTRIES":          5,
	"HEADER":               6,
	"FOOTER":               7,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}
func (Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Header struct {
	Version    string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	SdkVersion string `protobuf:"bytes,2,opt,name=sdk_version,json=sdkVersion" json:"sdk_version,omitempty"`
	MaxMsgLen  int64  `protobuf:"varint,3,opt,name=max_msg_len,json=maxMsgLen" json:"max_msg_len,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Header) GetSdkVersion() string {
	if m != nil {
		return m.SdkVersion
	}
	return ""
}

func (m *Header) GetMaxMsgLen() int64 {
	if m != nil {
		return m.MaxMsgLen
	}
	return 0
}

type Footer struct {
}

func (m *Footer) Reset()                    { *m = Footer{} }
func (m *Footer) String() string            { return proto.CompactTextString(m) }
func (*Footer) ProtoMessage()               {}
func (*Footer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EntryHeader struct {
	Len  int64 `protobuf:"varint,1,opt,name=len" json:"len,omitempty"`
	Kind Kind  `protobuf:"varint,2,opt,name=kind,enum=session.Kind" json:"kind,omitempty"`
}

func (m *EntryHeader) Reset()                    { *m = EntryHeader{} }
func (m *EntryHeader) String() string            { return proto.CompactTextString(m) }
func (*EntryHeader) ProtoMessage()               {}
func (*EntryHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EntryHeader) GetLen() int64 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *EntryHeader) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_INVALID
}

type Entry struct {
	Kind Kind `protobuf:"varint,1,opt,name=kind,enum=session.Kind" json:"kind,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*Entry_InstReq
	//	*Entry_InstResp
	//	*Entry_Elems
	//	*Entry_LogEntries
	//	*Entry_Header
	//	*Entry_Footer
	Msg       isEntry_Msg `protobuf_oneof:"msg"`
	Timestamp int64       `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isEntry_Msg interface {
	isEntry_Msg()
}

type Entry_InstReq struct {
	InstReq *org_apache_beam_model_fn_execution_v1.InstructionRequest `protobuf:"bytes,1000,opt,name=inst_req,json=instReq,oneof"`
}
type Entry_InstResp struct {
	InstResp *org_apache_beam_model_fn_execution_v1.InstructionResponse `protobuf:"bytes,1001,opt,name=inst_resp,json=instResp,oneof"`
}
type Entry_Elems struct {
	Elems *org_apache_beam_model_fn_execution_v1.Elements `protobuf:"bytes,1002,opt,name=elems,oneof"`
}
type Entry_LogEntries struct {
	LogEntries *org_apache_beam_model_fn_execution_v1.LogEntry_List `protobuf:"bytes,1003,opt,name=log_entries,json=logEntries,oneof"`
}
type Entry_Header struct {
	Header *Header `protobuf:"bytes,1004,opt,name=header,oneof"`
}
type Entry_Footer struct {
	Footer *Footer `protobuf:"bytes,1005,opt,name=footer,oneof"`
}

func (*Entry_InstReq) isEntry_Msg()    {}
func (*Entry_InstResp) isEntry_Msg()   {}
func (*Entry_Elems) isEntry_Msg()      {}
func (*Entry_LogEntries) isEntry_Msg() {}
func (*Entry_Header) isEntry_Msg()     {}
func (*Entry_Footer) isEntry_Msg()     {}

func (m *Entry) GetMsg() isEntry_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Entry) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_INVALID
}

func (m *Entry) GetInstReq() *org_apache_beam_model_fn_execution_v1.InstructionRequest {
	if x, ok := m.GetMsg().(*Entry_InstReq); ok {
		return x.InstReq
	}
	return nil
}

func (m *Entry) GetInstResp() *org_apache_beam_model_fn_execution_v1.InstructionResponse {
	if x, ok := m.GetMsg().(*Entry_InstResp); ok {
		return x.InstResp
	}
	return nil
}

func (m *Entry) GetElems() *org_apache_beam_model_fn_execution_v1.Elements {
	if x, ok := m.GetMsg().(*Entry_Elems); ok {
		return x.Elems
	}
	return nil
}

func (m *Entry) GetLogEntries() *org_apache_beam_model_fn_execution_v1.LogEntry_List {
	if x, ok := m.GetMsg().(*Entry_LogEntries); ok {
		return x.LogEntries
	}
	return nil
}

func (m *Entry) GetHeader() *Header {
	if x, ok := m.GetMsg().(*Entry_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Entry) GetFooter() *Footer {
	if x, ok := m.GetMsg().(*Entry_Footer); ok {
		return x.Footer
	}
	return nil
}

func (m *Entry) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Entry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Entry_OneofMarshaler, _Entry_OneofUnmarshaler, _Entry_OneofSizer, []interface{}{
		(*Entry_InstReq)(nil),
		(*Entry_InstResp)(nil),
		(*Entry_Elems)(nil),
		(*Entry_LogEntries)(nil),
		(*Entry_Header)(nil),
		(*Entry_Footer)(nil),
	}
}

func _Entry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Entry)
	// msg
	switch x := m.Msg.(type) {
	case *Entry_InstReq:
		b.EncodeVarint(1000<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InstReq); err != nil {
			return err
		}
	case *Entry_InstResp:
		b.EncodeVarint(1001<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InstResp); err != nil {
			return err
		}
	case *Entry_Elems:
		b.EncodeVarint(1002<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Elems); err != nil {
			return err
		}
	case *Entry_LogEntries:
		b.EncodeVarint(1003<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogEntries); err != nil {
			return err
		}
	case *Entry_Header:
		b.EncodeVarint(1004<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case *Entry_Footer:
		b.EncodeVarint(1005<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Footer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Entry.Msg has unexpected type %T", x)
	}
	return nil
}

func _Entry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Entry)
	switch tag {
	case 1000: // msg.inst_req
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(org_apache_beam_model_fn_execution_v1.InstructionRequest)
		err := b.DecodeMessage(msg)
		m.Msg = &Entry_InstReq{msg}
		return true, err
	case 1001: // msg.inst_resp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(org_apache_beam_model_fn_execution_v1.InstructionResponse)
		err := b.DecodeMessage(msg)
		m.Msg = &Entry_InstResp{msg}
		return true, err
	case 1002: // msg.elems
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(org_apache_beam_model_fn_execution_v1.Elements)
		err := b.DecodeMessage(msg)
		m.Msg = &Entry_Elems{msg}
		return true, err
	case 1003: // msg.log_entries
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(org_apache_beam_model_fn_execution_v1.LogEntry_List)
		err := b.DecodeMessage(msg)
		m.Msg = &Entry_LogEntries{msg}
		return true, err
	case 1004: // msg.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Header)
		err := b.DecodeMessage(msg)
		m.Msg = &Entry_Header{msg}
		return true, err
	case 1005: // msg.footer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Footer)
		err := b.DecodeMessage(msg)
		m.Msg = &Entry_Footer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Entry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Entry)
	// msg
	switch x := m.Msg.(type) {
	case *Entry_InstReq:
		s := proto.Size(x.InstReq)
		n += proto.SizeVarint(1000<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_InstResp:
		s := proto.Size(x.InstResp)
		n += proto.SizeVarint(1001<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_Elems:
		s := proto.Size(x.Elems)
		n += proto.SizeVarint(1002<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_LogEntries:
		s := proto.Size(x.LogEntries)
		n += proto.SizeVarint(1003<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_Header:
		s := proto.Size(x.Header)
		n += proto.SizeVarint(1004<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_Footer:
		s := proto.Size(x.Footer)
		n += proto.SizeVarint(1005<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Header)(nil), "session.Header")
	proto.RegisterType((*Footer)(nil), "session.Footer")
	proto.RegisterType((*EntryHeader)(nil), "session.EntryHeader")
	proto.RegisterType((*Entry)(nil), "session.Entry")
	proto.RegisterEnum("session.Kind", Kind_name, Kind_value)
}

func init() { proto.RegisterFile("session.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x80, 0xe3, 0x3a, 0xb1, 0x9b, 0xb1, 0x42, 0xdd, 0x05, 0x09, 0x0b, 0x21, 0x28, 0x39, 0x55,
	0x3d, 0x18, 0x51, 0xb8, 0xc0, 0x2d, 0x6d, 0xb6, 0xd8, 0x22, 0x38, 0xb0, 0x76, 0x0b, 0xe2, 0x62,
	0xb9, 0xf1, 0xd4, 0xb5, 0x62, 0xaf, 0x1d, 0xaf, 0x53, 0x85, 0x1b, 0xaf, 0xc0, 0x5b, 0xf2, 0xfb,
	0x0e, 0xc8, 0xeb, 0xa4, 0x08, 0x24, 0xa4, 0x8a, 0xdb, 0xce, 0xcc, 0xb7, 0xdf, 0x6a, 0x67, 0x76,
	0x61, 0x20, 0x50, 0x88, 0xb4, 0xe0, 0x76, 0x59, 0x15, 0x75, 0x41, 0xf4, 0x75, 0x78, 0x6f, 0xf7,
	0x1c, 0xa3, 0x3c, 0xbc, 0xe0, 0x61, 0x54, 0xa6, 0x6d, 0x6d, 0x38, 0x03, 0xcd, 0xc1, 0x28, 0xc6,
	0x8a, 0x58, 0xa0, 0x5f, 0x61, 0xd5, 0x70, 0x96, 0xb2, 0xa7, 0xec, 0xf7, 0xd9, 0x26, 0x24, 0x0f,
	0xc1, 0x10, 0xf1, 0x3c, 0xdc, 0x54, 0xb7, 0x64, 0x15, 0x44, 0x3c, 0x3f, 0x5b, 0x03, 0x0f, 0xc0,
	0xc8, 0xa3, 0x55, 0x98, 0x8b, 0x24, 0xcc, 0x90, 0x5b, 0xea, 0x9e, 0xb2, 0xaf, 0xb2, 0x7e, 0x1e,
	0xad, 0x5e, 0x8b, 0x64, 0x82, 0x7c, 0xb8, 0x0d, 0xda, 0x49, 0x51, 0xd4, 0x58, 0x0d, 0x8f, 0xc0,
	0xa0, 0xbc, 0xae, 0x3e, 0xae, 0xcf, 0x34, 0x41, 0x6d, 0x36, 0x28, 0x72, 0x43, 0xb3, 0x24, 0x8f,
	0xa0, 0x3b, 0x4f, 0x79, 0x2c, 0x0f, 0xb9, 0x75, 0x38, 0xb0, 0x37, 0x37, 0x79, 0x95, 0xf2, 0x98,
	0xc9, 0xd2, 0xf0, 0x53, 0x17, 0x7a, 0x52, 0x72, 0x0d, 0x2b, 0xff, 0x84, 0xc9, 0x3b, 0xd8, 0x4e,
	0xb9, 0xa8, 0xc3, 0x0a, 0x17, 0xd6, 0x17, 0x7d, 0x4f, 0xd9, 0x37, 0x0e, 0x9f, 0xdb, 0x45, 0x95,
	0xd8, 0x51, 0x19, 0xcd, 0x2e, 0xd1, 0x6e, 0x3a, 0x62, 0xe7, 0x45, 0x8c, 0x99, 0x7d, 0xc1, 0x43,
	0x5c, 0xe1, 0x6c, 0x59, 0x37, 0x8a, 0xab, 0x27, 0xb6, 0xcb, 0x45, 0x5d, 0x2d, 0x67, 0x4d, 0xc8,
	0x70, 0xb1, 0x44, 0x51, 0x3b, 0x1d, 0xa6, 0x37, 0x36, 0x86, 0x0b, 0xf2, 0x01, 0xfa, 0x6b, 0xb1,
	0x28, 0xad, 0xaf, 0xad, 0xf9, 0xc5, 0xff, 0x98, 0x45, 0x59, 0x70, 0x81, 0x4e, 0x87, 0x6d, 0xb7,
	0x6a, 0x51, 0x12, 0x07, 0x7a, 0x98, 0x61, 0x2e, 0xac, 0x6f, 0xad, 0xf7, 0xf1, 0x0d, 0xbd, 0x34,
	0xc3, 0x1c, 0x79, 0x2d, 0x9c, 0x0e, 0x6b, 0x05, 0xe4, 0x3d, 0x18, 0x59, 0x91, 0x84, 0xc8, 0xeb,
	0x2a, 0x45, 0x61, 0x7d, 0x6f, 0x7d, 0xcf, 0x6e, 0xe8, 0x9b, 0x14, 0x89, 0x6c, 0xb4, 0x3d, 0x49,
	0xe5, 0xe5, 0x21, 0x6b, 0x13, 0x29, 0x0a, 0x72, 0x00, 0xda, 0xa5, 0x1c, 0xa2, 0xf5, 0xa3, 0x95,
	0xee, 0x5c, 0xb7, 0xbf, 0x1d, 0xae, 0xd3, 0x61, 0x6b, 0xa2, 0x61, 0x2f, 0xe4, 0xfc, 0xad, 0x9f,
	0x7f, 0xb3, 0xed, 0xbb, 0x68, 0xd8, 0x96, 0x20, 0xf7, 0xa1, 0x5f, 0xa7, 0x39, 0x8a, 0x3a, 0xca,
	0x4b, 0xf9, 0x0a, 0x54, 0xf6, 0x3b, 0x71, 0xd4, 0x03, 0x35, 0x17, 0xc9, 0xc1, 0x67, 0x05, 0xba,
	0xcd, 0x90, 0x89, 0x01, 0xba, 0xeb, 0x9d, 0x8d, 0x26, 0xee, 0xd8, 0xec, 0x90, 0xbb, 0x70, 0xdb,
	0xf5, 0xfc, 0x80, 0x9d, 0x1e, 0x07, 0xee, 0xd4, 0x0b, 0x19, 0x7d, 0x7b, 0x4a, 0xfd, 0xc0, 0x54,
	0x88, 0x05, 0x77, 0xfe, 0x2c, 0xf8, 0x6f, 0xa6, 0x9e, 0x4f, 0xcd, 0x2d, 0xb2, 0x0b, 0x83, 0xf1,
	0x28, 0x18, 0x85, 0x8c, 0x1e, 0x53, 0xf7, 0x8c, 0x8e, 0x4d, 0x95, 0x0c, 0xa0, 0x2f, 0x53, 0x3e,
	0xf5, 0x02, 0xb3, 0x4b, 0x76, 0xc0, 0x98, 0x4c, 0x5f, 0x86, 0xd4, 0x0b, 0x98, 0x4b, 0x7d, 0xb3,
	0x47, 0x00, 0x34, 0x87, 0x8e, 0xc6, 0x94, 0x99, 0x5a, 0xb3, 0x3e, 0x99, 0x4e, 0x03, 0xca, 0x4c,
	0xfd, 0x5c, 0x93, 0x1f, 0xea, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x2f, 0xbd, 0x8f,
	0x7d, 0x03, 0x00, 0x00,
}
