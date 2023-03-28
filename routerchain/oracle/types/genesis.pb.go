// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params                  Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId                  string                `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	BandTokenPriceStateList []BandTokenPriceState `protobuf:"bytes,3,rep,name=bandTokenPriceStateList,proto3" json:"bandTokenPriceStateList"`
	CallDataRecordList      []CallDataRecord      `protobuf:"bytes,4,rep,name=callDataRecordList,proto3" json:"callDataRecordList"`
	BandOracleRequestList   []BandOracleRequest   `protobuf:"bytes,5,rep,name=bandOracleRequestList,proto3" json:"bandOracleRequestList"`
	GasPriceStateList       []GasPriceState       `protobuf:"bytes,6,rep,name=gasPriceStateList,proto3" json:"gasPriceStateList"`
	TokenPriceList          []TokenPrice          `protobuf:"bytes,7,rep,name=tokenPriceList,proto3" json:"tokenPriceList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce0b3a2b4a184fc3, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetBandTokenPriceStateList() []BandTokenPriceState {
	if m != nil {
		return m.BandTokenPriceStateList
	}
	return nil
}

func (m *GenesisState) GetCallDataRecordList() []CallDataRecord {
	if m != nil {
		return m.CallDataRecordList
	}
	return nil
}

func (m *GenesisState) GetBandOracleRequestList() []BandOracleRequest {
	if m != nil {
		return m.BandOracleRequestList
	}
	return nil
}

func (m *GenesisState) GetGasPriceStateList() []GasPriceState {
	if m != nil {
		return m.GasPriceStateList
	}
	return nil
}

func (m *GenesisState) GetTokenPriceList() []TokenPrice {
	if m != nil {
		return m.TokenPriceList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "routerprotocol.routerchain.oracle.GenesisState")
}

func init() { proto.RegisterFile("oracle/genesis.proto", fileDescriptor_ce0b3a2b4a184fc3) }

var fileDescriptor_ce0b3a2b4a184fc3 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x6d, 0x92, 0x38, 0x62, 0x82, 0x90, 0x18, 0x82, 0x62, 0x45, 0x60, 0x0c, 0x6c, 0xc2,
	0x22, 0x36, 0x04, 0x94, 0x07, 0x08, 0x48, 0x11, 0x12, 0x12, 0x91, 0x61, 0x05, 0x0b, 0x6b, 0x62,
	0x8f, 0x1c, 0x0b, 0xc7, 0xe3, 0xce, 0x4c, 0xaa, 0xf6, 0x05, 0xba, 0xee, 0x63, 0x65, 0x99, 0x65,
	0x57, 0x55, 0x95, 0xbc, 0x48, 0x35, 0x3f, 0x6e, 0xf3, 0x57, 0xd5, 0xbb, 0xb9, 0x9e, 0x7b, 0xce,
	0x77, 0xef, 0x19, 0x83, 0x36, 0xa1, 0x28, 0xca, 0xb0, 0x9f, 0xe0, 0x1c, 0xb3, 0x94, 0x79, 0x05,
	0x25, 0x9c, 0xc0, 0x77, 0x94, 0x2c, 0x38, 0xa6, 0xb2, 0x88, 0x48, 0xe6, 0xa9, 0x32, 0x9a, 0xa1,
	0x34, 0xf7, 0x94, 0xa0, 0xdb, 0x4e, 0x48, 0x42, 0x64, 0x83, 0x2f, 0x4e, 0x4a, 0xd8, 0x7d, 0xa9,
	0xed, 0x0a, 0x44, 0xd1, 0x5c, 0xbb, 0x75, 0x3f, 0xe8, 0x8f, 0x53, 0x94, 0xc7, 0x21, 0x27, 0xff,
	0x71, 0x1e, 0x16, 0x34, 0x8d, 0x70, 0xc8, 0x38, 0xe2, 0x58, 0x37, 0xbd, 0xd1, 0x4d, 0x11, 0xca,
	0xb2, 0x30, 0x46, 0x1c, 0x85, 0x14, 0x47, 0x84, 0xc6, 0xfa, 0xda, 0xdd, 0xf6, 0x50, 0xe7, 0x90,
	0xe2, 0x93, 0x05, 0x66, 0x5c, 0x77, 0xbc, 0x2e, 0x37, 0x41, 0xec, 0x88, 0xbd, 0xad, 0x6f, 0xb7,
	0xf0, 0xea, 0xe6, 0xfd, 0x45, 0x03, 0x3c, 0x1b, 0xab, 0xed, 0x7f, 0x0b, 0x01, 0x1c, 0x03, 0x4b,
	0x8d, 0x6f, 0x9b, 0xae, 0xd9, 0x6b, 0x0d, 0x3e, 0x7a, 0x8f, 0xa6, 0xe1, 0x4d, 0xa4, 0x60, 0x54,
	0x5f, 0x5e, 0xbf, 0x35, 0x02, 0x2d, 0x87, 0x1d, 0xd0, 0x2c, 0x08, 0xe5, 0x61, 0x1a, 0xdb, 0x4f,
	0x5c, 0xb3, 0xf7, 0x34, 0xb0, 0x44, 0xf9, 0x23, 0x86, 0xa7, 0xa0, 0x23, 0xf6, 0xf8, 0x23, 0x66,
	0x99, 0x88, 0x51, 0x24, 0xf8, 0x67, 0xca, 0xb8, 0x5d, 0x73, 0x6b, 0xbd, 0xd6, 0x60, 0x58, 0x01,
	0x39, 0x3a, 0x74, 0xd0, 0xfc, 0x87, 0xcc, 0x61, 0x02, 0xa0, 0x88, 0xf7, 0x3b, 0xe2, 0x28, 0x90,
	0xe1, 0x4a, 0x64, 0x5d, 0x22, 0x3f, 0x57, 0x40, 0x7e, 0xdb, 0x11, 0x6b, 0xda, 0x11, 0x4b, 0x58,
	0x80, 0x57, 0x62, 0x86, 0x5f, 0x52, 0x16, 0xa8, 0x67, 0x92, 0xac, 0x86, 0x64, 0x7d, 0xad, 0xb8,
	0xde, 0x8e, 0x5e, 0xe3, 0x8e, 0x1b, 0xc3, 0x18, 0xbc, 0x48, 0x10, 0xdb, 0x0b, 0xd3, 0x92, 0xb4,
	0x4f, 0x15, 0x68, 0xe3, 0x6d, 0xad, 0x26, 0x1d, 0x1a, 0xc2, 0x7f, 0xe0, 0x39, 0xbf, 0xcb, 0x55,
	0x22, 0x9a, 0x12, 0xd1, 0xaf, 0x80, 0xb8, 0x7f, 0x10, 0xed, 0xbf, 0x67, 0x35, 0x9a, 0x2c, 0xd7,
	0x8e, 0xb9, 0x5a, 0x3b, 0xe6, 0xcd, 0xda, 0x31, 0x2f, 0x37, 0x8e, 0xb1, 0xda, 0x38, 0xc6, 0xd5,
	0xc6, 0x31, 0xfe, 0x0e, 0x93, 0x94, 0xcf, 0x16, 0x53, 0x2f, 0x22, 0x73, 0x5f, 0x39, 0xf7, 0x4b,
	0x52, 0x59, 0x4b, 0x94, 0x7f, 0xe6, 0x97, 0xbf, 0xf9, 0x79, 0x81, 0xd9, 0xd4, 0x92, 0x5d, 0x5f,
	0x6e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x9b, 0xdd, 0x71, 0xe5, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenPriceList) > 0 {
		for iNdEx := len(m.TokenPriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GasPriceStateList) > 0 {
		for iNdEx := len(m.GasPriceStateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPriceStateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.BandOracleRequestList) > 0 {
		for iNdEx := len(m.BandOracleRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BandOracleRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.CallDataRecordList) > 0 {
		for iNdEx := len(m.CallDataRecordList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallDataRecordList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BandTokenPriceStateList) > 0 {
		for iNdEx := len(m.BandTokenPriceStateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BandTokenPriceStateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.BandTokenPriceStateList) > 0 {
		for _, e := range m.BandTokenPriceStateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CallDataRecordList) > 0 {
		for _, e := range m.CallDataRecordList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BandOracleRequestList) > 0 {
		for _, e := range m.BandOracleRequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasPriceStateList) > 0 {
		for _, e := range m.GasPriceStateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokenPriceList) > 0 {
		for _, e := range m.TokenPriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandTokenPriceStateList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BandTokenPriceStateList = append(m.BandTokenPriceStateList, BandTokenPriceState{})
			if err := m.BandTokenPriceStateList[len(m.BandTokenPriceStateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallDataRecordList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallDataRecordList = append(m.CallDataRecordList, CallDataRecord{})
			if err := m.CallDataRecordList[len(m.CallDataRecordList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandOracleRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BandOracleRequestList = append(m.BandOracleRequestList, BandOracleRequest{})
			if err := m.BandOracleRequestList[len(m.BandOracleRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceStateList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPriceStateList = append(m.GasPriceStateList, GasPriceState{})
			if err := m.GasPriceStateList[len(m.GasPriceStateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPriceList = append(m.TokenPriceList, TokenPrice{})
			if err := m.TokenPriceList[len(m.TokenPriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
