// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metastore/metadata_request.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/router-protocol/sdk-go/routerchain/multichain/types"
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

type MetadataRequest struct {
	ChainType   types.ChainType `protobuf:"varint,1,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId     string          `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce  uint64          `protobuf:"varint,3,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight uint64          `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	DaapAddress []byte          `protobuf:"bytes,5,opt,name=daapAddress,proto3" json:"daapAddress,omitempty"`
	FeePayer    string          `protobuf:"bytes,6,opt,name=feePayer,proto3" json:"feePayer,omitempty"`
	Status      MetaTxStatus    `protobuf:"varint,7,opt,name=status,proto3,enum=routerprotocol.routerchain.metastore.MetaTxStatus" json:"status,omitempty"`
}

func (m *MetadataRequest) Reset()         { *m = MetadataRequest{} }
func (m *MetadataRequest) String() string { return proto.CompactTextString(m) }
func (*MetadataRequest) ProtoMessage()    {}
func (*MetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bbf9ea4f3996390, []int{0}
}
func (m *MetadataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataRequest.Merge(m, src)
}
func (m *MetadataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataRequest proto.InternalMessageInfo

func (m *MetadataRequest) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *MetadataRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MetadataRequest) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *MetadataRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MetadataRequest) GetDaapAddress() []byte {
	if m != nil {
		return m.DaapAddress
	}
	return nil
}

func (m *MetadataRequest) GetFeePayer() string {
	if m != nil {
		return m.FeePayer
	}
	return ""
}

func (m *MetadataRequest) GetStatus() MetaTxStatus {
	if m != nil {
		return m.Status
	}
	return META_TX_CREATED
}

type MetadataRequestClaimHash struct {
	ChainType   types.ChainType `protobuf:"varint,1,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId     string          `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce  uint64          `protobuf:"varint,3,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight uint64          `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	DaapAddress []byte          `protobuf:"bytes,5,opt,name=daapAddress,proto3" json:"daapAddress,omitempty"`
	FeePayer    string          `protobuf:"bytes,6,opt,name=feePayer,proto3" json:"feePayer,omitempty"`
}

func (m *MetadataRequestClaimHash) Reset()         { *m = MetadataRequestClaimHash{} }
func (m *MetadataRequestClaimHash) String() string { return proto.CompactTextString(m) }
func (*MetadataRequestClaimHash) ProtoMessage()    {}
func (*MetadataRequestClaimHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bbf9ea4f3996390, []int{1}
}
func (m *MetadataRequestClaimHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataRequestClaimHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataRequestClaimHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataRequestClaimHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataRequestClaimHash.Merge(m, src)
}
func (m *MetadataRequestClaimHash) XXX_Size() int {
	return m.Size()
}
func (m *MetadataRequestClaimHash) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataRequestClaimHash.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataRequestClaimHash proto.InternalMessageInfo

func (m *MetadataRequestClaimHash) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *MetadataRequestClaimHash) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MetadataRequestClaimHash) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *MetadataRequestClaimHash) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MetadataRequestClaimHash) GetDaapAddress() []byte {
	if m != nil {
		return m.DaapAddress
	}
	return nil
}

func (m *MetadataRequestClaimHash) GetFeePayer() string {
	if m != nil {
		return m.FeePayer
	}
	return ""
}

func init() {
	proto.RegisterType((*MetadataRequest)(nil), "routerprotocol.routerchain.metastore.MetadataRequest")
	proto.RegisterType((*MetadataRequestClaimHash)(nil), "routerprotocol.routerchain.metastore.MetadataRequestClaimHash")
}

func init() { proto.RegisterFile("metastore/metadata_request.proto", fileDescriptor_0bbf9ea4f3996390) }

var fileDescriptor_0bbf9ea4f3996390 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xad, 0x4b, 0x69, 0xa9, 0x41, 0x20, 0x65, 0xb2, 0x8a, 0x64, 0x45, 0x15, 0x43, 0x16, 0x1c,
	0x54, 0x26, 0x46, 0xe8, 0x52, 0x90, 0xa8, 0x50, 0xda, 0x89, 0x25, 0x72, 0x93, 0xa3, 0x89, 0x48,
	0xea, 0x10, 0x3b, 0xa8, 0xdd, 0xf9, 0x00, 0xfe, 0x0a, 0xc6, 0x8e, 0x8c, 0xa8, 0xfd, 0x11, 0x14,
	0xa7, 0x69, 0x0b, 0x03, 0x62, 0x67, 0xb1, 0xfc, 0x9e, 0xef, 0xde, 0xf3, 0x3d, 0x1d, 0x36, 0x63,
	0x50, 0x5c, 0x2a, 0x91, 0x82, 0x9d, 0xdf, 0x7c, 0xae, 0xb8, 0x9b, 0xc2, 0x53, 0x06, 0x52, 0xb1,
	0x24, 0x15, 0x4a, 0x18, 0x27, 0xa9, 0xc8, 0x14, 0xa4, 0x1a, 0x78, 0x22, 0x62, 0x05, 0xf4, 0x02,
	0x1e, 0x4e, 0xd8, 0xba, 0xb9, 0x75, 0x1c, 0x67, 0x91, 0x0a, 0x35, 0x6b, 0xeb, 0xd3, 0x55, 0xb3,
	0x04, 0x0a, 0x89, 0x16, 0xfd, 0x6e, 0xe2, 0xaa, 0xa9, 0x2b, 0x15, 0x57, 0x99, 0x2c, 0xde, 0xdb,
	0x6f, 0x55, 0x7c, 0x74, 0xbb, 0x72, 0x77, 0x0a, 0x73, 0xa3, 0x8f, 0x9b, 0x5a, 0x67, 0x38, 0x4b,
	0x80, 0x20, 0x13, 0x59, 0x87, 0x9d, 0x33, 0xf6, 0xdb, 0x57, 0xd6, 0xfe, 0xac, 0x5b, 0xf6, 0x39,
	0x1b, 0x09, 0x83, 0xe0, 0x86, 0x06, 0xd7, 0x3e, 0xa9, 0x9a, 0xc8, 0x6a, 0x3a, 0x25, 0x34, 0x28,
	0xc6, 0xf0, 0x0c, 0x13, 0xd5, 0x17, 0x13, 0x0f, 0xc8, 0x8e, 0x89, 0xac, 0x9a, 0xb3, 0xc5, 0x18,
	0x26, 0xde, 0x1f, 0x45, 0xc2, 0x7b, 0xec, 0x41, 0x38, 0x0e, 0x14, 0xa9, 0xe9, 0x82, 0x6d, 0x2a,
	0xaf, 0xf0, 0x39, 0x4f, 0x2e, 0x7d, 0x3f, 0x05, 0x29, 0xc9, 0xae, 0x89, 0xac, 0x03, 0x67, 0x9b,
	0x32, 0x5a, 0x78, 0xef, 0x01, 0xe0, 0x8e, 0xcf, 0x20, 0x25, 0x75, 0x6d, 0xbf, 0xc6, 0xc6, 0x0d,
	0xae, 0x17, 0x69, 0x90, 0x86, 0x1e, 0xb3, 0xc3, 0xfe, 0x92, 0x38, 0xcb, 0x03, 0x1b, 0x4e, 0x07,
	0xba, 0xd3, 0x59, 0x29, 0xb4, 0x5f, 0xaa, 0x98, 0xfc, 0x48, 0xb2, 0x1b, 0xf1, 0x30, 0xee, 0x71,
	0x19, 0xfc, 0x9f, 0x48, 0xaf, 0x06, 0xef, 0x0b, 0x8a, 0xe6, 0x0b, 0x8a, 0x3e, 0x17, 0x14, 0xbd,
	0x2e, 0x69, 0x65, 0xbe, 0xa4, 0x95, 0x8f, 0x25, 0xad, 0xdc, 0x5f, 0x8c, 0x43, 0x15, 0x64, 0x23,
	0xe6, 0x89, 0xd8, 0x2e, 0x66, 0x3d, 0x2d, 0x67, 0x2f, 0x71, 0xb1, 0xc4, 0x53, 0x7b, 0xb3, 0xb4,
	0xf9, 0x2a, 0xcb, 0x51, 0x5d, 0x17, 0x9e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x23, 0x56, 0xa2,
	0x5e, 0x33, 0x03, 0x00, 0x00,
}

func (m *MetadataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if len(m.FeePayer) > 0 {
		i -= len(m.FeePayer)
		copy(dAtA[i:], m.FeePayer)
		i = encodeVarintMetadataRequest(dAtA, i, uint64(len(m.FeePayer)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DaapAddress) > 0 {
		i -= len(m.DaapAddress)
		copy(dAtA[i:], m.DaapAddress)
		i = encodeVarintMetadataRequest(dAtA, i, uint64(len(m.DaapAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.EventNonce != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMetadataRequest(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainType != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MetadataRequestClaimHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataRequestClaimHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataRequestClaimHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeePayer) > 0 {
		i -= len(m.FeePayer)
		copy(dAtA[i:], m.FeePayer)
		i = encodeVarintMetadataRequest(dAtA, i, uint64(len(m.FeePayer)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DaapAddress) > 0 {
		i -= len(m.DaapAddress)
		copy(dAtA[i:], m.DaapAddress)
		i = encodeVarintMetadataRequest(dAtA, i, uint64(len(m.DaapAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.EventNonce != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMetadataRequest(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainType != 0 {
		i = encodeVarintMetadataRequest(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadataRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadataRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetadataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainType != 0 {
		n += 1 + sovMetadataRequest(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMetadataRequest(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovMetadataRequest(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMetadataRequest(uint64(m.BlockHeight))
	}
	l = len(m.DaapAddress)
	if l > 0 {
		n += 1 + l + sovMetadataRequest(uint64(l))
	}
	l = len(m.FeePayer)
	if l > 0 {
		n += 1 + l + sovMetadataRequest(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovMetadataRequest(uint64(m.Status))
	}
	return n
}

func (m *MetadataRequestClaimHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainType != 0 {
		n += 1 + sovMetadataRequest(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMetadataRequest(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovMetadataRequest(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMetadataRequest(uint64(m.BlockHeight))
	}
	l = len(m.DaapAddress)
	if l > 0 {
		n += 1 + l + sovMetadataRequest(uint64(l))
	}
	l = len(m.FeePayer)
	if l > 0 {
		n += 1 + l + sovMetadataRequest(uint64(l))
	}
	return n
}

func sovMetadataRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadataRequest(x uint64) (n int) {
	return sovMetadataRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetadataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadataRequest
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
			return fmt.Errorf("proto: MetadataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
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
				return ErrInvalidLengthMetadataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaapAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
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
				return ErrInvalidLengthMetadataRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaapAddress = append(m.DaapAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DaapAddress == nil {
				m.DaapAddress = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
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
				return ErrInvalidLengthMetadataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeePayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MetaTxStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetadataRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadataRequest
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
func (m *MetadataRequestClaimHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadataRequest
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
			return fmt.Errorf("proto: MetadataRequestClaimHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataRequestClaimHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
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
				return ErrInvalidLengthMetadataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaapAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
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
				return ErrInvalidLengthMetadataRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaapAddress = append(m.DaapAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DaapAddress == nil {
				m.DaapAddress = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadataRequest
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
				return ErrInvalidLengthMetadataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeePayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadataRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadataRequest
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
func skipMetadataRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadataRequest
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
					return 0, ErrIntOverflowMetadataRequest
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
					return 0, ErrIntOverflowMetadataRequest
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
				return 0, ErrInvalidLengthMetadataRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadataRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadataRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadataRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadataRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadataRequest = fmt.Errorf("proto: unexpected end of group")
)
