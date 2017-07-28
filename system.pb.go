// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: system.proto

/*
	Package systeminfo is a generated protocol buffer package.

	It is generated from these files:
		system.proto

	It has these top-level messages:
		System
		CPU
*/
package systeminfo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type System struct {
	KernelOS          string   `protobuf:"bytes,1,opt,name=KernelOS,proto3" json:"KernelOS,omitempty"`
	KernelVersion     string   `protobuf:"bytes,2,opt,name=KernelVersion,proto3" json:"KernelVersion,omitempty"`
	KernelArch        string   `protobuf:"bytes,3,opt,name=KernelArch,proto3" json:"KernelArch,omitempty"`
	KernelType        string   `protobuf:"bytes,4,opt,name=KernelType,proto3" json:"KernelType,omitempty"`
	KernelCompileDate string   `protobuf:"bytes,5,opt,name=KernelCompileDate,proto3" json:"KernelCompileDate,omitempty"`
	OSName            string   `protobuf:"bytes,6,opt,name=OSName,proto3" json:"OSName,omitempty"`
	OSID              string   `protobuf:"bytes,7,opt,name=OSID,proto3" json:"OSID,omitempty"`
	OSIDLike          string   `protobuf:"bytes,8,opt,name=OSIDLike,proto3" json:"OSIDLike,omitempty"`
	OSVersion         string   `protobuf:"bytes,9,opt,name=OSVersion,proto3" json:"OSVersion,omitempty"`
	MemTotal          uint64   `protobuf:"varint,10,opt,name=MemTotal,proto3" json:"MemTotal,omitempty"`
	SwapTotal         uint64   `protobuf:"varint,11,opt,name=SwapTotal,proto3" json:"SwapTotal,omitempty"`
	NetDev            []string `protobuf:"bytes,12,rep,name=NetDev" json:"NetDev,omitempty"`
	Sockets           int32    `protobuf:"varint,13,opt,name=Sockets,proto3" json:"Sockets,omitempty"`
	CPU               []*CPU   `protobuf:"bytes,14,rep,name=CPU" json:"CPU,omitempty"`
}

func (m *System) Reset()                    { *m = System{} }
func (m *System) String() string            { return proto.CompactTextString(m) }
func (*System) ProtoMessage()               {}
func (*System) Descriptor() ([]byte, []int) { return fileDescriptorSystem, []int{0} }

func (m *System) GetKernelOS() string {
	if m != nil {
		return m.KernelOS
	}
	return ""
}

func (m *System) GetKernelVersion() string {
	if m != nil {
		return m.KernelVersion
	}
	return ""
}

func (m *System) GetKernelArch() string {
	if m != nil {
		return m.KernelArch
	}
	return ""
}

func (m *System) GetKernelType() string {
	if m != nil {
		return m.KernelType
	}
	return ""
}

func (m *System) GetKernelCompileDate() string {
	if m != nil {
		return m.KernelCompileDate
	}
	return ""
}

func (m *System) GetOSName() string {
	if m != nil {
		return m.OSName
	}
	return ""
}

func (m *System) GetOSID() string {
	if m != nil {
		return m.OSID
	}
	return ""
}

func (m *System) GetOSIDLike() string {
	if m != nil {
		return m.OSIDLike
	}
	return ""
}

func (m *System) GetOSVersion() string {
	if m != nil {
		return m.OSVersion
	}
	return ""
}

func (m *System) GetMemTotal() uint64 {
	if m != nil {
		return m.MemTotal
	}
	return 0
}

func (m *System) GetSwapTotal() uint64 {
	if m != nil {
		return m.SwapTotal
	}
	return 0
}

func (m *System) GetNetDev() []string {
	if m != nil {
		return m.NetDev
	}
	return nil
}

func (m *System) GetSockets() int32 {
	if m != nil {
		return m.Sockets
	}
	return 0
}

func (m *System) GetCPU() []*CPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

type CPU struct {
	PhysicalID int32    `protobuf:"varint,1,opt,name=PhysicalID,proto3" json:"PhysicalID,omitempty"`
	CoreID     int32    `protobuf:"varint,2,opt,name=CoreID,proto3" json:"CoreID,omitempty"`
	Siblings   int32    `protobuf:"varint,3,opt,name=Siblings,proto3" json:"Siblings,omitempty"`
	VendorID   string   `protobuf:"bytes,4,opt,name=VendorID,proto3" json:"VendorID,omitempty"`
	CPUFamily  string   `protobuf:"bytes,5,opt,name=CPUFamily,proto3" json:"CPUFamily,omitempty"`
	Model      string   `protobuf:"bytes,6,opt,name=Model,proto3" json:"Model,omitempty"`
	ModelName  string   `protobuf:"bytes,7,opt,name=ModelName,proto3" json:"ModelName,omitempty"`
	Stepping   string   `protobuf:"bytes,8,opt,name=Stepping,proto3" json:"Stepping,omitempty"`
	Microcode  string   `protobuf:"bytes,9,opt,name=Microcode,proto3" json:"Microcode,omitempty"`
	CPUMHz     float32  `protobuf:"fixed32,10,opt,name=CPUMHz,proto3" json:"CPUMHz,omitempty"`
	BogoMIPS   float32  `protobuf:"fixed32,11,opt,name=BogoMIPS,proto3" json:"BogoMIPS,omitempty"`
	CacheSize  string   `protobuf:"bytes,12,opt,name=CacheSize,proto3" json:"CacheSize,omitempty"`
	CPUCores   int32    `protobuf:"varint,13,opt,name=CPUCores,proto3" json:"CPUCores,omitempty"`
	Flags      []string `protobuf:"bytes,14,rep,name=Flags" json:"Flags,omitempty"`
	Bugs       []string `protobuf:"bytes,15,rep,name=Bugs" json:"Bugs,omitempty"`
}

func (m *CPU) Reset()                    { *m = CPU{} }
func (m *CPU) String() string            { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()               {}
func (*CPU) Descriptor() ([]byte, []int) { return fileDescriptorSystem, []int{1} }

func (m *CPU) GetPhysicalID() int32 {
	if m != nil {
		return m.PhysicalID
	}
	return 0
}

func (m *CPU) GetCoreID() int32 {
	if m != nil {
		return m.CoreID
	}
	return 0
}

func (m *CPU) GetSiblings() int32 {
	if m != nil {
		return m.Siblings
	}
	return 0
}

func (m *CPU) GetVendorID() string {
	if m != nil {
		return m.VendorID
	}
	return ""
}

func (m *CPU) GetCPUFamily() string {
	if m != nil {
		return m.CPUFamily
	}
	return ""
}

func (m *CPU) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *CPU) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *CPU) GetStepping() string {
	if m != nil {
		return m.Stepping
	}
	return ""
}

func (m *CPU) GetMicrocode() string {
	if m != nil {
		return m.Microcode
	}
	return ""
}

func (m *CPU) GetCPUMHz() float32 {
	if m != nil {
		return m.CPUMHz
	}
	return 0
}

func (m *CPU) GetBogoMIPS() float32 {
	if m != nil {
		return m.BogoMIPS
	}
	return 0
}

func (m *CPU) GetCacheSize() string {
	if m != nil {
		return m.CacheSize
	}
	return ""
}

func (m *CPU) GetCPUCores() int32 {
	if m != nil {
		return m.CPUCores
	}
	return 0
}

func (m *CPU) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *CPU) GetBugs() []string {
	if m != nil {
		return m.Bugs
	}
	return nil
}

func init() {
	proto.RegisterType((*System)(nil), "protobuf.System")
	proto.RegisterType((*CPU)(nil), "protobuf.CPU")
}
func (m *System) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *System) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.KernelOS) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.KernelOS)))
		i += copy(dAtA[i:], m.KernelOS)
	}
	if len(m.KernelVersion) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.KernelVersion)))
		i += copy(dAtA[i:], m.KernelVersion)
	}
	if len(m.KernelArch) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.KernelArch)))
		i += copy(dAtA[i:], m.KernelArch)
	}
	if len(m.KernelType) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.KernelType)))
		i += copy(dAtA[i:], m.KernelType)
	}
	if len(m.KernelCompileDate) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.KernelCompileDate)))
		i += copy(dAtA[i:], m.KernelCompileDate)
	}
	if len(m.OSName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.OSName)))
		i += copy(dAtA[i:], m.OSName)
	}
	if len(m.OSID) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.OSID)))
		i += copy(dAtA[i:], m.OSID)
	}
	if len(m.OSIDLike) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.OSIDLike)))
		i += copy(dAtA[i:], m.OSIDLike)
	}
	if len(m.OSVersion) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.OSVersion)))
		i += copy(dAtA[i:], m.OSVersion)
	}
	if m.MemTotal != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.MemTotal))
	}
	if m.SwapTotal != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.SwapTotal))
	}
	if len(m.NetDev) > 0 {
		for _, s := range m.NetDev {
			dAtA[i] = 0x62
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Sockets != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.Sockets))
	}
	if len(m.CPU) > 0 {
		for _, msg := range m.CPU {
			dAtA[i] = 0x72
			i++
			i = encodeVarintSystem(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CPU) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CPU) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PhysicalID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.PhysicalID))
	}
	if m.CoreID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.CoreID))
	}
	if m.Siblings != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.Siblings))
	}
	if len(m.VendorID) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.VendorID)))
		i += copy(dAtA[i:], m.VendorID)
	}
	if len(m.CPUFamily) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.CPUFamily)))
		i += copy(dAtA[i:], m.CPUFamily)
	}
	if len(m.Model) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.Model)))
		i += copy(dAtA[i:], m.Model)
	}
	if len(m.ModelName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.ModelName)))
		i += copy(dAtA[i:], m.ModelName)
	}
	if len(m.Stepping) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.Stepping)))
		i += copy(dAtA[i:], m.Stepping)
	}
	if len(m.Microcode) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.Microcode)))
		i += copy(dAtA[i:], m.Microcode)
	}
	if m.CPUMHz != 0 {
		dAtA[i] = 0x55
		i++
		i = encodeFixed32System(dAtA, i, uint32(math.Float32bits(float32(m.CPUMHz))))
	}
	if m.BogoMIPS != 0 {
		dAtA[i] = 0x5d
		i++
		i = encodeFixed32System(dAtA, i, uint32(math.Float32bits(float32(m.BogoMIPS))))
	}
	if len(m.CacheSize) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintSystem(dAtA, i, uint64(len(m.CacheSize)))
		i += copy(dAtA[i:], m.CacheSize)
	}
	if m.CPUCores != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintSystem(dAtA, i, uint64(m.CPUCores))
	}
	if len(m.Flags) > 0 {
		for _, s := range m.Flags {
			dAtA[i] = 0x72
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Bugs) > 0 {
		for _, s := range m.Bugs {
			dAtA[i] = 0x7a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64System(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32System(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSystem(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *System) Size() (n int) {
	var l int
	_ = l
	l = len(m.KernelOS)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.KernelVersion)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.KernelArch)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.KernelType)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.KernelCompileDate)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.OSName)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.OSID)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.OSIDLike)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.OSVersion)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	if m.MemTotal != 0 {
		n += 1 + sovSystem(uint64(m.MemTotal))
	}
	if m.SwapTotal != 0 {
		n += 1 + sovSystem(uint64(m.SwapTotal))
	}
	if len(m.NetDev) > 0 {
		for _, s := range m.NetDev {
			l = len(s)
			n += 1 + l + sovSystem(uint64(l))
		}
	}
	if m.Sockets != 0 {
		n += 1 + sovSystem(uint64(m.Sockets))
	}
	if len(m.CPU) > 0 {
		for _, e := range m.CPU {
			l = e.Size()
			n += 1 + l + sovSystem(uint64(l))
		}
	}
	return n
}

func (m *CPU) Size() (n int) {
	var l int
	_ = l
	if m.PhysicalID != 0 {
		n += 1 + sovSystem(uint64(m.PhysicalID))
	}
	if m.CoreID != 0 {
		n += 1 + sovSystem(uint64(m.CoreID))
	}
	if m.Siblings != 0 {
		n += 1 + sovSystem(uint64(m.Siblings))
	}
	l = len(m.VendorID)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.CPUFamily)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.Model)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.ModelName)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.Stepping)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	l = len(m.Microcode)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	if m.CPUMHz != 0 {
		n += 5
	}
	if m.BogoMIPS != 0 {
		n += 5
	}
	l = len(m.CacheSize)
	if l > 0 {
		n += 1 + l + sovSystem(uint64(l))
	}
	if m.CPUCores != 0 {
		n += 1 + sovSystem(uint64(m.CPUCores))
	}
	if len(m.Flags) > 0 {
		for _, s := range m.Flags {
			l = len(s)
			n += 1 + l + sovSystem(uint64(l))
		}
	}
	if len(m.Bugs) > 0 {
		for _, s := range m.Bugs {
			l = len(s)
			n += 1 + l + sovSystem(uint64(l))
		}
	}
	return n
}

func sovSystem(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSystem(x uint64) (n int) {
	return sovSystem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *System) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystem
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: System: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: System: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelOS", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelOS = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelArch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelArch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelCompileDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KernelCompileDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OSName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OSID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSIDLike", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OSIDLike = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OSVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemTotal", wireType)
			}
			m.MemTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemTotal |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapTotal", wireType)
			}
			m.SwapTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SwapTotal |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetDev", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetDev = append(m.NetDev, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sockets", wireType)
			}
			m.Sockets = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sockets |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPU", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPU = append(m.CPU, &CPU{})
			if err := m.CPU[len(m.CPU)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSystem
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
func (m *CPU) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSystem
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CPU: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CPU: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhysicalID", wireType)
			}
			m.PhysicalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhysicalID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreID", wireType)
			}
			m.CoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoreID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Siblings", wireType)
			}
			m.Siblings = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Siblings |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPUFamily", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CPUFamily = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Model = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stepping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stepping = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Microcode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Microcode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPUMHz", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.CPUMHz = float32(math.Float32frombits(v))
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field BogoMIPS", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.BogoMIPS = float32(math.Float32frombits(v))
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheSize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CacheSize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPUCores", wireType)
			}
			m.CPUCores = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CPUCores |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Flags = append(m.Flags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bugs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSystem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bugs = append(m.Bugs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSystem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSystem
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
func skipSystem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSystem
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
					return 0, ErrIntOverflowSystem
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSystem
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSystem
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSystem
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipSystem(dAtA[start:])
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
	ErrInvalidLengthSystem = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSystem   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("system.proto", fileDescriptorSystem) }

var fileDescriptorSystem = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0x4f, 0x8e, 0xd3, 0x4c,
	0x10, 0xc5, 0x3f, 0xc7, 0x71, 0x26, 0xe9, 0x49, 0x3e, 0x44, 0x0b, 0xa1, 0x16, 0x1a, 0x85, 0x68,
	0xc4, 0x22, 0x0b, 0x94, 0x05, 0x9c, 0x80, 0xd8, 0x1a, 0x11, 0x41, 0x26, 0x56, 0xf7, 0x78, 0x16,
	0xec, 0x1c, 0xa7, 0xc6, 0x69, 0x8d, 0xed, 0xb6, 0x6c, 0x0f, 0x28, 0x73, 0x12, 0x38, 0x03, 0x17,
	0x61, 0xc9, 0x11, 0x50, 0xb8, 0x08, 0xea, 0x6a, 0xff, 0x09, 0x62, 0xe5, 0x7a, 0xef, 0x55, 0x25,
	0x55, 0xfe, 0x25, 0x64, 0x5c, 0x1e, 0xca, 0x0a, 0xd2, 0x45, 0x5e, 0xa8, 0x4a, 0xd1, 0x21, 0x3e,
	0xb6, 0x0f, 0x77, 0x97, 0xdf, 0x6d, 0x32, 0x10, 0x18, 0xd1, 0x17, 0x64, 0xf8, 0x01, 0x8a, 0x0c,
	0x92, 0x8d, 0x60, 0xd6, 0xcc, 0x9a, 0x8f, 0x78, 0xab, 0xe9, 0x2b, 0x32, 0x31, 0xf5, 0x2d, 0x14,
	0xa5, 0x54, 0x19, 0xeb, 0x61, 0xc3, 0xdf, 0x26, 0x9d, 0x12, 0x62, 0x8c, 0x77, 0x45, 0xb4, 0x67,
	0x36, 0xb6, 0x9c, 0x38, 0x5d, 0x7e, 0x73, 0xc8, 0x81, 0xf5, 0x4f, 0x73, 0xed, 0xd0, 0xd7, 0xe4,
	0xa9, 0x51, 0xae, 0x4a, 0x73, 0x99, 0x80, 0x17, 0x56, 0xc0, 0x1c, 0x6c, 0xfb, 0x37, 0xa0, 0xcf,
	0xc9, 0x60, 0x23, 0xae, 0xc3, 0x14, 0xd8, 0x00, 0x5b, 0x6a, 0x45, 0x29, 0xe9, 0x6f, 0xc4, 0xca,
	0x63, 0x67, 0xe8, 0x62, 0xad, 0x6f, 0xd3, 0xcf, 0x8f, 0xf2, 0x1e, 0xd8, 0xd0, 0xdc, 0xd6, 0x68,
	0x7a, 0x41, 0x46, 0x1b, 0xd1, 0xdc, 0x35, 0xc2, 0xb0, 0x33, 0xf4, 0xe4, 0x1a, 0xd2, 0x1b, 0x55,
	0x85, 0x09, 0x23, 0x33, 0x6b, 0xde, 0xe7, 0xad, 0xd6, 0x93, 0xe2, 0x4b, 0x98, 0x9b, 0xf0, 0x1c,
	0xc3, 0xce, 0xd0, 0xfb, 0x5d, 0x43, 0xe5, 0xc1, 0x67, 0x36, 0x9e, 0xd9, 0x7a, 0x3f, 0xa3, 0x28,
	0x23, 0x67, 0x42, 0x45, 0xf7, 0x50, 0x95, 0x6c, 0x32, 0xb3, 0xe6, 0x0e, 0x6f, 0x24, 0x7d, 0x49,
	0x6c, 0xd7, 0x0f, 0xd8, 0xff, 0x33, 0x7b, 0x7e, 0xfe, 0x66, 0xb2, 0x68, 0x20, 0x2d, 0x5c, 0x3f,
	0xe0, 0x3a, 0xb9, 0xfc, 0x66, 0x63, 0x87, 0x7e, 0x91, 0xfe, 0xfe, 0x50, 0xca, 0x28, 0x4c, 0x56,
	0x1e, 0xc2, 0x72, 0xf8, 0x89, 0xa3, 0xbf, 0xda, 0x55, 0x05, 0xac, 0x3c, 0xe4, 0xe4, 0xf0, 0x5a,
	0xe9, 0x63, 0x84, 0xdc, 0x26, 0x32, 0x8b, 0x4b, 0xc4, 0xe3, 0xf0, 0x56, 0xeb, 0xec, 0x16, 0xb2,
	0x9d, 0x2a, 0x56, 0x5e, 0x8d, 0xa6, 0xd5, 0xfa, 0x50, 0xd7, 0x0f, 0xae, 0xc2, 0x54, 0x26, 0x87,
	0x1a, 0x48, 0x67, 0xd0, 0x67, 0xc4, 0x59, 0xab, 0x1d, 0x24, 0x35, 0x07, 0x23, 0xf4, 0x0c, 0x16,
	0x48, 0xc8, 0xb0, 0xe8, 0x0c, 0xdc, 0xa4, 0x82, 0x3c, 0x97, 0x59, 0xdc, 0x00, 0x69, 0x34, 0x4e,
	0xca, 0xa8, 0x50, 0x91, 0xda, 0x41, 0x03, 0xa4, 0x35, 0xf0, 0x36, 0x3f, 0x58, 0xbf, 0x7f, 0x44,
	0x1c, 0x3d, 0x5e, 0x2b, 0xfd, 0x89, 0x4b, 0x15, 0xab, 0xf5, 0xca, 0x17, 0xc8, 0xa2, 0xc7, 0x5b,
	0x8d, 0xfb, 0x87, 0xd1, 0x1e, 0x84, 0x7c, 0x04, 0x36, 0xae, 0xf7, 0x6f, 0x0c, 0x3d, 0xe9, 0xfa,
	0x81, 0x7e, 0x45, 0x0d, 0x91, 0x56, 0xeb, 0xdb, 0xae, 0x92, 0x30, 0x2e, 0x11, 0xca, 0x88, 0x1b,
	0xa1, 0x7f, 0x62, 0xcb, 0x87, 0xb8, 0x64, 0x4f, 0xd0, 0xc4, 0x7a, 0x79, 0xf1, 0xe3, 0x38, 0xb5,
	0x7e, 0x1e, 0xa7, 0xd6, 0xaf, 0xe3, 0xd4, 0xfa, 0xfa, 0x7b, 0xfa, 0xdf, 0x27, 0x62, 0xfe, 0x73,
	0x32, 0xbb, 0x53, 0xdb, 0x01, 0xc2, 0x7c, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xd0, 0xe1,
	0x59, 0x88, 0x03, 0x00, 0x00,
}
