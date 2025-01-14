// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store/storepb/types.proto

package storepb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/thanos-io/thanos/pkg/store/labelpb"
	github_com_thanos_io_thanos_pkg_store_labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// / PartialResponseStrategy controls partial response handling.
type PartialResponseStrategy int32

const (
	/// WARN strategy tells server to treat any error that will related to single StoreAPI (e.g missing chunk series because of underlying
	/// storeAPI is temporarily not available) as warning which will not fail the whole query (still OK response).
	/// Server should produce those as a warnings field in response.
	PartialResponseStrategy_WARN PartialResponseStrategy = 0
	/// ABORT strategy tells server to treat any error that will related to single StoreAPI (e.g missing chunk series because of underlying
	/// storeAPI is temporarily not available) as the gRPC error that aborts the query.
	///
	/// This is especially useful for any rule/alert evaluations on top of StoreAPI which usually does not tolerate partial
	/// errors.
	PartialResponseStrategy_ABORT PartialResponseStrategy = 1
)

var PartialResponseStrategy_name = map[int32]string{
	0: "WARN",
	1: "ABORT",
}

var PartialResponseStrategy_value = map[string]int32{
	"WARN":  0,
	"ABORT": 1,
}

func (x PartialResponseStrategy) String() string {
	return proto.EnumName(PartialResponseStrategy_name, int32(x))
}

func (PartialResponseStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{0}
}

type Chunk_Encoding int32

const (
	Chunk_XOR       Chunk_Encoding = 0
	Chunk_HISTOGRAM Chunk_Encoding = 1
)

var Chunk_Encoding_name = map[int32]string{
	0: "XOR",
	1: "HISTOGRAM",
}

var Chunk_Encoding_value = map[string]int32{
	"XOR":       0,
	"HISTOGRAM": 1,
}

func (x Chunk_Encoding) String() string {
	return proto.EnumName(Chunk_Encoding_name, int32(x))
}

func (Chunk_Encoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{0, 0}
}

type LabelMatcher_Type int32

const (
	LabelMatcher_EQ  LabelMatcher_Type = 0
	LabelMatcher_NEQ LabelMatcher_Type = 1
	LabelMatcher_RE  LabelMatcher_Type = 2
	LabelMatcher_NRE LabelMatcher_Type = 3
)

var LabelMatcher_Type_name = map[int32]string{
	0: "EQ",
	1: "NEQ",
	2: "RE",
	3: "NRE",
}

var LabelMatcher_Type_value = map[string]int32{
	"EQ":  0,
	"NEQ": 1,
	"RE":  2,
	"NRE": 3,
}

func (x LabelMatcher_Type) String() string {
	return proto.EnumName(LabelMatcher_Type_name, int32(x))
}

func (LabelMatcher_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{3, 0}
}

type Chunk struct {
	Type Chunk_Encoding `protobuf:"varint,1,opt,name=type,proto3,enum=thanos.Chunk_Encoding" json:"type,omitempty"`
	Data []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Hash uint64         `protobuf:"varint,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return m.Size()
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

type Series struct {
	Labels []github_com_thanos_io_thanos_pkg_store_labelpb.ZLabel `protobuf:"bytes,1,rep,name=labels,proto3,customtype=github.com/thanos-io/thanos/pkg/store/labelpb.ZLabel" json:"labels"`
	Chunks []AggrChunk                                            `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks"`
}

func (m *Series) Reset()         { *m = Series{} }
func (m *Series) String() string { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()    {}
func (*Series) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{1}
}
func (m *Series) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Series) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Series.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Series) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Series.Merge(m, src)
}
func (m *Series) XXX_Size() int {
	return m.Size()
}
func (m *Series) XXX_DiscardUnknown() {
	xxx_messageInfo_Series.DiscardUnknown(m)
}

var xxx_messageInfo_Series proto.InternalMessageInfo

type AggrChunk struct {
	MinTime int64  `protobuf:"varint,1,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64  `protobuf:"varint,2,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
	Raw     *Chunk `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	Count   *Chunk `protobuf:"bytes,4,opt,name=count,proto3" json:"count,omitempty"`
	Sum     *Chunk `protobuf:"bytes,5,opt,name=sum,proto3" json:"sum,omitempty"`
	Min     *Chunk `protobuf:"bytes,6,opt,name=min,proto3" json:"min,omitempty"`
	Max     *Chunk `protobuf:"bytes,7,opt,name=max,proto3" json:"max,omitempty"`
	Counter *Chunk `protobuf:"bytes,8,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (m *AggrChunk) Reset()         { *m = AggrChunk{} }
func (m *AggrChunk) String() string { return proto.CompactTextString(m) }
func (*AggrChunk) ProtoMessage()    {}
func (*AggrChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{2}
}
func (m *AggrChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggrChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggrChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggrChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggrChunk.Merge(m, src)
}
func (m *AggrChunk) XXX_Size() int {
	return m.Size()
}
func (m *AggrChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_AggrChunk.DiscardUnknown(m)
}

var xxx_messageInfo_AggrChunk proto.InternalMessageInfo

// Matcher specifies a rule, which can match or set of labels or not.
type LabelMatcher struct {
	Type  LabelMatcher_Type `protobuf:"varint,1,opt,name=type,proto3,enum=thanos.LabelMatcher_Type" json:"type,omitempty"`
	Name  string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value string            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()         { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()    {}
func (*LabelMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{3}
}
func (m *LabelMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelMatcher.Merge(m, src)
}
func (m *LabelMatcher) XXX_Size() int {
	return m.Size()
}
func (m *LabelMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_LabelMatcher proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("thanos.PartialResponseStrategy", PartialResponseStrategy_name, PartialResponseStrategy_value)
	proto.RegisterEnum("thanos.Chunk_Encoding", Chunk_Encoding_name, Chunk_Encoding_value)
	proto.RegisterEnum("thanos.LabelMatcher_Type", LabelMatcher_Type_name, LabelMatcher_Type_value)
	proto.RegisterType((*Chunk)(nil), "thanos.Chunk")
	proto.RegisterType((*Series)(nil), "thanos.Series")
	proto.RegisterType((*AggrChunk)(nil), "thanos.AggrChunk")
	proto.RegisterType((*LabelMatcher)(nil), "thanos.LabelMatcher")
}

func init() { proto.RegisterFile("store/storepb/types.proto", fileDescriptor_121fba57de02d8e0) }

var fileDescriptor_121fba57de02d8e0 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xbd, 0xb6, 0xe3, 0x24, 0xf3, 0x6b, 0x7f, 0x32, 0xab, 0x0a, 0xdc, 0x1e, 0x9c, 0xc8,
	0x08, 0x11, 0x55, 0xaa, 0x2d, 0x15, 0x8e, 0x5c, 0x12, 0x14, 0x01, 0x12, 0x6d, 0xe8, 0x26, 0x12,
	0xa8, 0x17, 0xb4, 0x71, 0x57, 0xb6, 0xd5, 0xf8, 0x8f, 0xbc, 0x6b, 0x48, 0x1e, 0x80, 0x3b, 0x88,
	0x3b, 0xcf, 0x93, 0x63, 0x8f, 0x88, 0x43, 0x04, 0xc9, 0x8b, 0x20, 0xaf, 0x1d, 0x20, 0x52, 0x2e,
	0xd6, 0x78, 0x3e, 0xdf, 0x99, 0xd9, 0x99, 0x9d, 0x85, 0x63, 0x2e, 0xd2, 0x9c, 0x79, 0xf2, 0x9b,
	0x4d, 0x3d, 0xb1, 0xc8, 0x18, 0x77, 0xb3, 0x3c, 0x15, 0x29, 0x36, 0x44, 0x48, 0x93, 0x94, 0x9f,
	0x1c, 0x05, 0x69, 0x90, 0x4a, 0x97, 0x57, 0x5a, 0x15, 0x3d, 0xa9, 0x03, 0x67, 0x74, 0xca, 0x66,
	0xbb, 0x81, 0xce, 0x27, 0x04, 0x8d, 0xe7, 0x61, 0x91, 0xdc, 0xe2, 0x53, 0xd0, 0x4b, 0x60, 0xa1,
	0x2e, 0xea, 0xfd, 0x7f, 0x7e, 0xdf, 0xad, 0x32, 0xba, 0x12, 0xba, 0xc3, 0xc4, 0x4f, 0x6f, 0xa2,
	0x24, 0x20, 0x52, 0x83, 0x31, 0xe8, 0x37, 0x54, 0x50, 0x4b, 0xed, 0xa2, 0xde, 0x01, 0x91, 0x36,
	0xb6, 0x40, 0x0f, 0x29, 0x0f, 0x2d, 0xad, 0x8b, 0x7a, 0xfa, 0x40, 0x5f, 0xae, 0x3a, 0x88, 0x48,
	0x8f, 0xe3, 0x40, 0x6b, 0x1b, 0x8f, 0x9b, 0xa0, 0xbd, 0x1b, 0x11, 0x53, 0xc1, 0x87, 0xd0, 0x7e,
	0xf9, 0x6a, 0x3c, 0x19, 0xbd, 0x20, 0xfd, 0x0b, 0x13, 0x39, 0xdf, 0x10, 0x18, 0x63, 0x96, 0x47,
	0x8c, 0x63, 0x1f, 0x0c, 0x79, 0x52, 0x6e, 0xa1, 0xae, 0xd6, 0xfb, 0xef, 0xfc, 0x70, 0x7b, 0x94,
	0xd7, 0xa5, 0x77, 0xf0, 0x6c, 0xb9, 0xea, 0x28, 0x3f, 0x56, 0x9d, 0xa7, 0x41, 0x24, 0xc2, 0x62,
	0xea, 0xfa, 0x69, 0xec, 0x55, 0x82, 0xb3, 0x28, 0xad, 0x2d, 0x2f, 0xbb, 0x0d, 0xbc, 0x9d, 0xa6,
	0xdd, 0x6b, 0x19, 0x4d, 0xea, 0xd4, 0xd8, 0x03, 0xc3, 0x2f, 0x3b, 0xe3, 0x96, 0x2a, 0x8b, 0xdc,
	0xdb, 0x16, 0xe9, 0x07, 0x41, 0x2e, 0x7b, 0x96, 0x2d, 0x28, 0xa4, 0x96, 0x39, 0x5f, 0x55, 0x68,
	0xff, 0x61, 0xf8, 0x18, 0x5a, 0x71, 0x94, 0xbc, 0x17, 0x51, 0x5c, 0x0d, 0x4c, 0x23, 0xcd, 0x38,
	0x4a, 0x26, 0x51, 0xcc, 0x24, 0xa2, 0xf3, 0x0a, 0xa9, 0x35, 0xa2, 0x73, 0x89, 0x3a, 0xa0, 0xe5,
	0xf4, 0xa3, 0x9c, 0xd0, 0x3f, 0x6d, 0xc9, 0x8c, 0xa4, 0x24, 0xf8, 0x21, 0x34, 0xfc, 0xb4, 0x48,
	0x84, 0xa5, 0xef, 0x93, 0x54, 0xac, 0xcc, 0xc2, 0x8b, 0xd8, 0x6a, 0xec, 0xcd, 0xc2, 0x8b, 0xb8,
	0x14, 0xc4, 0x51, 0x62, 0x19, 0x7b, 0x05, 0x71, 0x94, 0x48, 0x01, 0x9d, 0x5b, 0xcd, 0xfd, 0x02,
	0x3a, 0xc7, 0x8f, 0xa1, 0x29, 0x6b, 0xb1, 0xdc, 0x6a, 0xed, 0x13, 0x6d, 0xa9, 0xf3, 0x05, 0xc1,
	0x81, 0x1c, 0xec, 0x05, 0x15, 0x7e, 0xc8, 0x72, 0x7c, 0xb6, 0xb3, 0x45, 0xc7, 0x3b, 0x57, 0x57,
	0x6b, 0xdc, 0xc9, 0x22, 0x63, 0x7f, 0x17, 0x29, 0xa1, 0xf5, 0xa0, 0xda, 0x44, 0xda, 0xf8, 0x08,
	0x1a, 0x1f, 0xe8, 0xac, 0x60, 0x72, 0x4e, 0x6d, 0x52, 0xfd, 0x38, 0x3d, 0xd0, 0xcb, 0x38, 0x6c,
	0x80, 0x3a, 0xbc, 0x32, 0x95, 0x72, 0x91, 0x2e, 0x87, 0x57, 0x26, 0x2a, 0x1d, 0x64, 0x68, 0xaa,
	0xd2, 0x41, 0x86, 0xa6, 0x76, 0xea, 0xc2, 0x83, 0x37, 0x34, 0x17, 0x11, 0x9d, 0x11, 0xc6, 0xb3,
	0x34, 0xe1, 0x6c, 0x2c, 0x72, 0x2a, 0x58, 0xb0, 0xc0, 0x2d, 0xd0, 0xdf, 0xf6, 0xc9, 0xa5, 0xa9,
	0xe0, 0x36, 0x34, 0xfa, 0x83, 0x11, 0x99, 0x98, 0x68, 0xf0, 0x68, 0xf9, 0xcb, 0x56, 0x96, 0x6b,
	0x1b, 0xdd, 0xad, 0x6d, 0xf4, 0x73, 0x6d, 0xa3, 0xcf, 0x1b, 0x5b, 0xb9, 0xdb, 0xd8, 0xca, 0xf7,
	0x8d, 0xad, 0x5c, 0x37, 0xeb, 0xe7, 0x36, 0x35, 0xe4, 0x83, 0x79, 0xf2, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x44, 0x9d, 0x95, 0xa2, 0x86, 0x03, 0x00, 0x00,
}

func (m *Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hash != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Hash))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Series) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Series) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Series) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		for iNdEx := len(m.Chunks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chunks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Labels[iNdEx].Size()
				i -= size
				if _, err := m.Labels[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AggrChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggrChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggrChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Counter != nil {
		{
			size, err := m.Counter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Max != nil {
		{
			size, err := m.Max.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Min != nil {
		{
			size, err := m.Min.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Sum != nil {
		{
			size, err := m.Sum.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Count != nil {
		{
			size, err := m.Count.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Raw != nil {
		{
			size, err := m.Raw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxTime != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxTime))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTime != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LabelMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LabelMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovTypes(uint64(m.Hash))
	}
	return n
}

func (m *Series) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *AggrChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTime != 0 {
		n += 1 + sovTypes(uint64(m.MinTime))
	}
	if m.MaxTime != 0 {
		n += 1 + sovTypes(uint64(m.MaxTime))
	}
	if m.Raw != nil {
		l = m.Raw.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Count != nil {
		l = m.Count.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Sum != nil {
		l = m.Sum.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Min != nil {
		l = m.Min.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Max != nil {
		l = m.Max.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Counter != nil {
		l = m.Counter.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *LabelMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Chunk_Encoding(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Series) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Series: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Series: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, github_com_thanos_io_thanos_pkg_store_labelpb.ZLabel{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, AggrChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AggrChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AggrChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggrChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			m.MinTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			m.MaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Raw == nil {
				m.Raw = &Chunk{}
			}
			if err := m.Raw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Count == nil {
				m.Count = &Chunk{}
			}
			if err := m.Count.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sum == nil {
				m.Sum = &Chunk{}
			}
			if err := m.Sum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Min == nil {
				m.Min = &Chunk{}
			}
			if err := m.Min.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Max == nil {
				m.Max = &Chunk{}
			}
			if err := m.Max.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Counter == nil {
				m.Counter = &Chunk{}
			}
			if err := m.Counter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LabelMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= LabelMatcher_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
