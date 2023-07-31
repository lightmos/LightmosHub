// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lightmos/restaking/restaking.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CommissionRates defines the initial commission rates to be used for creating
// a validator.
type CommissionRates struct {
	// rate is the commission rate charged to delegators, as a fraction.
	Rate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate"`
	// max_rate defines the maximum commission rate which validator can ever charge, as a fraction.
	MaxRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=max_rate,json=maxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_rate"`
	// max_change_rate defines the maximum daily increase of the validator commission, as a fraction.
	MaxChangeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=max_change_rate,json=maxChangeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_change_rate"`
}

func (m *CommissionRates) Reset()         { *m = CommissionRates{} }
func (m *CommissionRates) String() string { return proto.CompactTextString(m) }
func (*CommissionRates) ProtoMessage()    {}
func (*CommissionRates) Descriptor() ([]byte, []int) {
	return fileDescriptor_7993bec0b80245b1, []int{0}
}
func (m *CommissionRates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommissionRates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommissionRates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommissionRates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommissionRates.Merge(m, src)
}
func (m *CommissionRates) XXX_Size() int {
	return m.Size()
}
func (m *CommissionRates) XXX_DiscardUnknown() {
	xxx_messageInfo_CommissionRates.DiscardUnknown(m)
}

var xxx_messageInfo_CommissionRates proto.InternalMessageInfo

// Commission defines commission parameters for a given validator.
type Commission struct {
	// commission_rates defines the initial commission rates to be used for creating a validator.
	CommissionRates `protobuf:"bytes,1,opt,name=commission_rates,json=commissionRates,proto3,embedded=commission_rates" json:"commission_rates"`
	// update_time is the last time the commission rate was changed.
	UpdateTime time.Time `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3,stdtime" json:"update_time"`
}

func (m *Commission) Reset()         { *m = Commission{} }
func (m *Commission) String() string { return proto.CompactTextString(m) }
func (*Commission) ProtoMessage()    {}
func (*Commission) Descriptor() ([]byte, []int) {
	return fileDescriptor_7993bec0b80245b1, []int{1}
}
func (m *Commission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Commission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Commission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commission.Merge(m, src)
}
func (m *Commission) XXX_Size() int {
	return m.Size()
}
func (m *Commission) XXX_DiscardUnknown() {
	xxx_messageInfo_Commission.DiscardUnknown(m)
}

var xxx_messageInfo_Commission proto.InternalMessageInfo

func (m *Commission) GetUpdateTime() time.Time {
	if m != nil {
		return m.UpdateTime
	}
	return time.Time{}
}

// Description defines a validator description.
type Description struct {
	// moniker defines a human-readable name for the validator.
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	// identity defines an optional identity signature (ex. UPort or Keybase).
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	// website defines an optional website link.
	Website string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	// security_contact defines an optional email for security contact.
	SecurityContact string `protobuf:"bytes,4,opt,name=security_contact,json=securityContact,proto3" json:"security_contact,omitempty"`
	// details define other optional details.
	Details string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *Description) Reset()         { *m = Description{} }
func (m *Description) String() string { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()    {}
func (*Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_7993bec0b80245b1, []int{2}
}
func (m *Description) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Description.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Description.Merge(m, src)
}
func (m *Description) XXX_Size() int {
	return m.Size()
}
func (m *Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Description proto.InternalMessageInfo

func (m *Description) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Description) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Description) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Description) GetSecurityContact() string {
	if m != nil {
		return m.SecurityContact
	}
	return ""
}

func (m *Description) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type RestakerTrace struct {
	Addr        string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	DestChainId string `protobuf:"bytes,2,opt,name=destChainId,proto3" json:"destChainId,omitempty"`
	Amount      string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *RestakerTrace) Reset()         { *m = RestakerTrace{} }
func (m *RestakerTrace) String() string { return proto.CompactTextString(m) }
func (*RestakerTrace) ProtoMessage()    {}
func (*RestakerTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_7993bec0b80245b1, []int{3}
}
func (m *RestakerTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RestakerTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RestakerTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RestakerTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestakerTrace.Merge(m, src)
}
func (m *RestakerTrace) XXX_Size() int {
	return m.Size()
}
func (m *RestakerTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_RestakerTrace.DiscardUnknown(m)
}

var xxx_messageInfo_RestakerTrace proto.InternalMessageInfo

func (m *RestakerTrace) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *RestakerTrace) GetDestChainId() string {
	if m != nil {
		return m.DestChainId
	}
	return ""
}

func (m *RestakerTrace) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*CommissionRates)(nil), "lightmos.restaking.CommissionRates")
	proto.RegisterType((*Commission)(nil), "lightmos.restaking.Commission")
	proto.RegisterType((*Description)(nil), "lightmos.restaking.Description")
	proto.RegisterType((*RestakerTrace)(nil), "lightmos.restaking.RestakerTrace")
}

func init() {
	proto.RegisterFile("lightmos/restaking/restaking.proto", fileDescriptor_7993bec0b80245b1)
}

var fileDescriptor_7993bec0b80245b1 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x8b, 0x13, 0x3f,
	0x18, 0x6e, 0xf6, 0xd7, 0xdf, 0x6e, 0x4d, 0x59, 0xba, 0x04, 0x91, 0x6e, 0x0f, 0x33, 0x4b, 0x05,
	0xd1, 0xc3, 0x4e, 0x61, 0xf5, 0x24, 0x9e, 0xda, 0x7a, 0xf0, 0x26, 0xc3, 0x82, 0x22, 0x48, 0x49,
	0x93, 0x38, 0x0d, 0x6d, 0x92, 0x61, 0x92, 0xc1, 0xf6, 0x5b, 0xec, 0xd1, 0xe3, 0x7e, 0x80, 0x3d,
	0x0a, 0x7e, 0x85, 0x3d, 0x16, 0x4f, 0xe2, 0xa1, 0x4a, 0x7b, 0xf1, 0x63, 0x48, 0xfe, 0x4c, 0xb7,
	0xd4, 0xeb, 0x9e, 0x26, 0xcf, 0xfb, 0xbe, 0x79, 0xf2, 0xcc, 0xf3, 0x24, 0xb0, 0x3b, 0xe3, 0xd9,
	0xc4, 0x08, 0xa5, 0x7b, 0x05, 0xd3, 0x06, 0x4f, 0xb9, 0xcc, 0xee, 0x56, 0x49, 0x5e, 0x28, 0xa3,
	0x10, 0xaa, 0x66, 0x92, 0x6d, 0xa7, 0x73, 0x4a, 0x94, 0x16, 0x4a, 0x8f, 0xdc, 0x44, 0xcf, 0x03,
	0x3f, 0xde, 0x79, 0x98, 0xa9, 0x4c, 0xf9, 0xba, 0x5d, 0x85, 0xea, 0x69, 0xa6, 0x54, 0x36, 0x63,
	0x3d, 0x87, 0xc6, 0xe5, 0xa7, 0x1e, 0x96, 0x8b, 0xd0, 0x8a, 0xf6, 0x5b, 0xb4, 0x2c, 0xb0, 0xe1,
	0x4a, 0x86, 0x7e, 0xbc, 0xdf, 0x37, 0x5c, 0x58, 0x21, 0x22, 0xf7, 0x03, 0xdd, 0x9b, 0x03, 0xd8,
	0x1a, 0x28, 0x21, 0xb8, 0xd6, 0x5c, 0xc9, 0x14, 0x1b, 0xa6, 0xd1, 0x5b, 0x58, 0x2f, 0xb0, 0x61,
	0x6d, 0x70, 0x06, 0x9e, 0x3e, 0xe8, 0xbf, 0xba, 0x5d, 0xc5, 0xb5, 0x9f, 0xab, 0xf8, 0x49, 0xc6,
	0xcd, 0xa4, 0x1c, 0x27, 0x44, 0x89, 0x20, 0x3a, 0x7c, 0xce, 0x35, 0x9d, 0xf6, 0xcc, 0x22, 0x67,
	0x3a, 0x19, 0x32, 0xf2, 0xfd, 0xeb, 0x39, 0x0c, 0xff, 0x34, 0x64, 0x24, 0x75, 0x4c, 0xe8, 0x1d,
	0x6c, 0x08, 0x3c, 0x1f, 0x39, 0xd6, 0x83, 0x7b, 0x60, 0x3d, 0x12, 0x78, 0x6e, 0xb5, 0x22, 0x0a,
	0x5b, 0x96, 0x98, 0x4c, 0xb0, 0xcc, 0x98, 0xe7, 0xff, 0xef, 0x1e, 0xf8, 0x8f, 0x05, 0x9e, 0x0f,
	0x1c, 0xa7, 0x3d, 0xe5, 0x65, 0xe3, 0xcb, 0x75, 0x0c, 0xfe, 0x5c, 0xc7, 0xa0, 0xfb, 0x0d, 0x40,
	0x78, 0x67, 0x17, 0x7a, 0x0f, 0x4f, 0xc8, 0x16, 0xb9, 0xe3, 0xb5, 0x73, 0xad, 0x79, 0xf1, 0x38,
	0xf9, 0x37, 0xf9, 0x64, 0xcf, 0xe8, 0x7e, 0xc3, 0x8a, 0x5c, 0xae, 0x62, 0x90, 0xb6, 0xc8, 0x5e,
	0x06, 0xaf, 0x61, 0xb3, 0xcc, 0x29, 0x36, 0x6c, 0x64, 0x13, 0x73, 0xa6, 0x35, 0x2f, 0x3a, 0x89,
	0x8f, 0x33, 0xa9, 0xe2, 0x4c, 0x2e, 0xab, 0x38, 0x3d, 0xd7, 0xd5, 0xaf, 0x18, 0xa4, 0xd0, 0x6f,
	0xb4, 0xad, 0x1d, 0xe5, 0x37, 0x00, 0x36, 0x87, 0x4c, 0x93, 0x82, 0xe7, 0xf6, 0x7e, 0xa0, 0x36,
	0x3c, 0x12, 0x4a, 0xf2, 0x29, 0x2b, 0x7c, 0xce, 0x69, 0x05, 0x51, 0x07, 0x36, 0x38, 0x65, 0xd2,
	0x70, 0xb3, 0xf0, 0x61, 0xa5, 0x5b, 0x6c, 0x77, 0x7d, 0x66, 0x63, 0xcd, 0x2b, 0x9f, 0xd3, 0x0a,
	0xa2, 0x67, 0xf0, 0x44, 0x33, 0x52, 0x16, 0xdc, 0x2c, 0x46, 0x44, 0x49, 0x83, 0x89, 0x69, 0xd7,
	0xdd, 0x48, 0xab, 0xaa, 0x0f, 0x7c, 0xd9, 0x92, 0x50, 0x66, 0x30, 0x9f, 0xe9, 0xf6, 0xff, 0x9e,
	0x24, 0xc0, 0x1d, 0xb9, 0x1f, 0xe1, 0x71, 0xea, 0x7c, 0x63, 0xc5, 0x65, 0x81, 0x09, 0x43, 0x08,
	0xd6, 0x31, 0xa5, 0x95, 0x58, 0xb7, 0x46, 0x67, 0xb0, 0x49, 0x99, 0x36, 0x83, 0x09, 0xe6, 0xf2,
	0x0d, 0x0d, 0x62, 0x77, 0x4b, 0xe8, 0x11, 0x3c, 0xc4, 0x42, 0x95, 0xd2, 0x04, 0xb9, 0x01, 0xf5,
	0x5f, 0xdc, 0xae, 0x23, 0xb0, 0x5c, 0x47, 0xe0, 0xf7, 0x3a, 0x02, 0x57, 0x9b, 0xa8, 0xb6, 0xdc,
	0x44, 0xb5, 0x1f, 0x9b, 0xa8, 0xf6, 0xa1, 0xb3, 0x7d, 0xd5, 0xf3, 0x9d, 0x77, 0xed, 0x2e, 0xca,
	0xf8, 0xd0, 0xf9, 0xfe, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x3e, 0x8e, 0x28, 0xfa,
	0x03, 0x00, 0x00,
}

func (this *CommissionRates) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CommissionRates)
	if !ok {
		that2, ok := that.(CommissionRates)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Rate.Equal(that1.Rate) {
		return false
	}
	if !this.MaxRate.Equal(that1.MaxRate) {
		return false
	}
	if !this.MaxChangeRate.Equal(that1.MaxChangeRate) {
		return false
	}
	return true
}
func (this *Commission) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Commission)
	if !ok {
		that2, ok := that.(Commission)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CommissionRates.Equal(&that1.CommissionRates) {
		return false
	}
	if !this.UpdateTime.Equal(that1.UpdateTime) {
		return false
	}
	return true
}
func (this *Description) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Description)
	if !ok {
		that2, ok := that.(Description)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if this.Identity != that1.Identity {
		return false
	}
	if this.Website != that1.Website {
		return false
	}
	if this.SecurityContact != that1.SecurityContact {
		return false
	}
	if this.Details != that1.Details {
		return false
	}
	return true
}
func (m *CommissionRates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommissionRates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommissionRates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxChangeRate.Size()
		i -= size
		if _, err := m.MaxChangeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRestaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxRate.Size()
		i -= size
		if _, err := m.MaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRestaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRestaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Commission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.UpdateTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.UpdateTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintRestaking(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommissionRates.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRestaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Description) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Description) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Description) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SecurityContact) > 0 {
		i -= len(m.SecurityContact)
		copy(dAtA[i:], m.SecurityContact)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.SecurityContact)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RestakerTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestakerTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RestakerTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestChainId) > 0 {
		i -= len(m.DestChainId)
		copy(dAtA[i:], m.DestChainId)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.DestChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintRestaking(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRestaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovRestaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommissionRates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Rate.Size()
	n += 1 + l + sovRestaking(uint64(l))
	l = m.MaxRate.Size()
	n += 1 + l + sovRestaking(uint64(l))
	l = m.MaxChangeRate.Size()
	n += 1 + l + sovRestaking(uint64(l))
	return n
}

func (m *Commission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommissionRates.Size()
	n += 1 + l + sovRestaking(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.UpdateTime)
	n += 1 + l + sovRestaking(uint64(l))
	return n
}

func (m *Description) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	l = len(m.SecurityContact)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	return n
}

func (m *RestakerTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	l = len(m.DestChainId)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovRestaking(uint64(l))
	}
	return n
}

func sovRestaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRestaking(x uint64) (n int) {
	return sovRestaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommissionRates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestaking
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
			return fmt.Errorf("proto: CommissionRates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommissionRates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxChangeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxChangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestaking
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
func (m *Commission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestaking
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
			return fmt.Errorf("proto: Commission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommissionRates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.UpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestaking
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
func (m *Description) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestaking
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
			return fmt.Errorf("proto: Description: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Description: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityContact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestaking
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
func (m *RestakerTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRestaking
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
			return fmt.Errorf("proto: RestakerTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestakerTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRestaking
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
				return ErrInvalidLengthRestaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRestaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRestaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRestaking
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
func skipRestaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRestaking
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
					return 0, ErrIntOverflowRestaking
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
					return 0, ErrIntOverflowRestaking
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
				return 0, ErrInvalidLengthRestaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRestaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRestaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRestaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRestaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRestaking = fmt.Errorf("proto: unexpected end of group")
)
