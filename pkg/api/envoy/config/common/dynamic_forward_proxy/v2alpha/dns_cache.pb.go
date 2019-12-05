// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto

package envoy_config_common_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	v2 "github.com/datawire/ambassador/pkg/api/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Configuration for the dynamic forward proxy DNS cache. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#next-free-field: 6]
type DnsCacheConfig struct {
	// The name of the cache. Multiple named caches allow independent dynamic forward proxy
	// configurations to operate within a single Envoy process using different configurations. All
	// configurations with the same name *must* otherwise have the same settings when referenced
	// from different configuration components. Configuration will fail to load if this is not
	// the case.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The DNS lookup family to use during resolution.
	//
	// [#comment:TODO(mattklein123): Figure out how to support IPv4/IPv6 "happy eyeballs" mode. The
	// way this might work is a new lookup family which returns both IPv4 and IPv6 addresses, and
	// then configures a host to have a primary and fall back address. With this, we could very
	// likely build a "happy eyeballs" connection pool which would race the primary / fall back
	// address and return the one that wins. This same method could potentially also be used for
	// QUIC to TCP fall back.]
	DnsLookupFamily v2.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.api.v2.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	// The DNS refresh rate for currently cached DNS hosts. If not specified defaults to 60s.
	//
	// .. note:
	//
	//  The returned DNS TTL is not currently used to alter the refresh rate. This feature will be
	//  added in a future change.
	DnsRefreshRate *types.Duration `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	// The TTL for hosts that are unused. Hosts that have not been used in the configured time
	// interval will be purged. If not specified defaults to 5m.
	//
	// .. note:
	//
	//   The TTL is only checked at the time of DNS refresh, as specified by *dns_refresh_rate*. This
	//   means that if the configured TTL is shorter than the refresh rate the host may not be removed
	//   immediately.
	//
	//  .. note:
	//
	//   The TTL has no relation to DNS TTL and is only used to control Envoy's resource usage.
	HostTtl *types.Duration `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	// The maximum number of hosts that the cache will hold. If not specified defaults to 1024.
	//
	// .. note:
	//
	//   The implementation is approximate and enforced independently on each worker thread, thus
	//   it is possible for the maximum hosts in the cache to go slightly above the configured
	//   value depending on timing. This is similar to how other circuit breakers work.
	MaxHosts             *types.UInt32Value `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2d9297e0c94cb56, []int{0}
}
func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return m.Size()
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v2.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v2.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *types.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *types.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *types.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto", fileDescriptor_d2d9297e0c94cb56)
}

var fileDescriptor_d2d9297e0c94cb56 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd7, 0xd9, 0x2e, 0x6c, 0x8d, 0x28, 0x4b, 0x0e, 0x10, 0x56, 0x50, 0x45, 0x48, 0x48,
	0xd5, 0x1e, 0x6c, 0x91, 0x7d, 0x00, 0xa0, 0xad, 0x10, 0x48, 0x48, 0xac, 0x22, 0xe0, 0x06, 0xd1,
	0x6c, 0xe2, 0xb4, 0x11, 0x8e, 0x6d, 0xd9, 0x4e, 0xb6, 0xbd, 0xf2, 0x38, 0x3c, 0x02, 0x27, 0x8e,
	0x1c, 0x79, 0x04, 0xd4, 0x0b, 0xe2, 0x2d, 0x90, 0xed, 0xf6, 0x80, 0x80, 0x43, 0x6f, 0x99, 0xcc,
	0xfc, 0x5f, 0xfe, 0x7f, 0x26, 0xf8, 0x19, 0x13, 0xbd, 0x5c, 0xd3, 0x52, 0x8a, 0xba, 0x59, 0xd0,
	0x52, 0xb6, 0xad, 0x14, 0xb4, 0x5a, 0x0b, 0x68, 0x9b, 0xb2, 0xa8, 0xa5, 0xbe, 0x02, 0x5d, 0x15,
	0x4a, 0xcb, 0xd5, 0x9a, 0xf6, 0x19, 0x70, 0xb5, 0x04, 0x5a, 0x09, 0x53, 0x94, 0x50, 0x2e, 0x19,
	0x51, 0x5a, 0x5a, 0x19, 0x3f, 0xf6, 0x08, 0x12, 0x10, 0x24, 0x20, 0xc8, 0x3f, 0x11, 0x64, 0x8b,
	0x38, 0xbd, 0x13, 0xbe, 0x0a, 0xaa, 0xa1, 0x7d, 0x46, 0xcb, 0xca, 0x04, 0xd4, 0xe9, 0x78, 0x21,
	0xe5, 0x82, 0x33, 0xea, 0xab, 0xcb, 0xae, 0xa6, 0x55, 0xa7, 0xc1, 0x36, 0x52, 0xfc, 0xaf, 0x7f,
	0xa5, 0x41, 0x29, 0xa6, 0x77, 0xfa, 0xbb, 0x3d, 0xf0, 0xa6, 0x02, 0xcb, 0xe8, 0xee, 0x21, 0x34,
	0x1e, 0xfe, 0x8c, 0xf0, 0x68, 0x2e, 0xcc, 0xcc, 0xd9, 0x9e, 0x79, 0xa3, 0xf1, 0x03, 0x3c, 0x10,
	0xd0, 0xb2, 0x04, 0xa5, 0x68, 0x32, 0x9c, 0x0e, 0xbf, 0xfc, 0xfa, 0x7a, 0x38, 0xd0, 0x51, 0x8a,
	0x72, 0xff, 0x3a, 0x7e, 0x8f, 0x6f, 0xbb, 0xa0, 0x5c, 0xca, 0x8f, 0x9d, 0x2a, 0x6a, 0x68, 0x1b,
	0xbe, 0x4e, 0xa2, 0x14, 0x4d, 0x46, 0xd9, 0x23, 0x12, 0x12, 0x83, 0x6a, 0x48, 0x9f, 0x91, 0x19,
	0xef, 0x8c, 0x65, 0x9a, 0xcc, 0x85, 0x79, 0xe5, 0xa7, 0x9f, 0xfb, 0xe1, 0x29, 0x76, 0xc8, 0xa3,
	0x4f, 0x28, 0x3a, 0x41, 0xf9, 0xad, 0xea, 0xcf, 0x66, 0xfc, 0x1a, 0x9f, 0x38, 0xbc, 0x66, 0xb5,
	0x66, 0x66, 0x59, 0x68, 0xb0, 0x2c, 0x39, 0x4c, 0xd1, 0xe4, 0x46, 0x76, 0x8f, 0x84, 0x90, 0x64,
	0x17, 0x92, 0xcc, 0xb7, 0x4b, 0xd8, 0x12, 0x3f, 0xa3, 0xe8, 0xec, 0x20, 0x1f, 0x55, 0xc2, 0xe4,
	0x41, 0x9d, 0x83, 0x65, 0xf1, 0x53, 0x7c, 0xbc, 0x94, 0xc6, 0x16, 0xd6, 0xf2, 0x64, 0xb0, 0x0f,
	0xe8, 0xba, 0x93, 0xbd, 0xb1, 0x3c, 0x9e, 0xe3, 0x61, 0x0b, 0xab, 0xc2, 0x95, 0x26, 0x39, 0xf2,
	0x88, 0xfb, 0x7f, 0x21, 0xde, 0xbe, 0x14, 0xf6, 0x3c, 0x7b, 0x07, 0xbc, 0x63, 0xdb, 0x9d, 0x9d,
	0x45, 0xe9, 0x41, 0x7e, 0xdc, 0xc2, 0xea, 0x85, 0x13, 0x4e, 0x3f, 0x7c, 0xdb, 0x8c, 0xd1, 0xf7,
	0xcd, 0x18, 0xfd, 0xd8, 0x8c, 0x11, 0x7e, 0xd2, 0xc8, 0xb0, 0xac, 0x70, 0xff, 0xbd, 0xff, 0x94,
	0xe9, 0xcd, 0xdd, 0xd5, 0x2e, 0x9c, 0x83, 0x0b, 0x74, 0x79, 0xcd, 0x5b, 0x39, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x37, 0x01, 0xb8, 0xa6, 0xb9, 0x02, 0x00, 0x00,
}

func (m *DnsCacheConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DnsCacheConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DnsCacheConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxHosts != nil {
		{
			size, err := m.MaxHosts.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDnsCache(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.HostTtl != nil {
		{
			size, err := m.HostTtl.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDnsCache(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.DnsRefreshRate != nil {
		{
			size, err := m.DnsRefreshRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDnsCache(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.DnsLookupFamily != 0 {
		i = encodeVarintDnsCache(dAtA, i, uint64(m.DnsLookupFamily))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDnsCache(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDnsCache(dAtA []byte, offset int, v uint64) int {
	offset -= sovDnsCache(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DnsCacheConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDnsCache(uint64(l))
	}
	if m.DnsLookupFamily != 0 {
		n += 1 + sovDnsCache(uint64(m.DnsLookupFamily))
	}
	if m.DnsRefreshRate != nil {
		l = m.DnsRefreshRate.Size()
		n += 1 + l + sovDnsCache(uint64(l))
	}
	if m.HostTtl != nil {
		l = m.HostTtl.Size()
		n += 1 + l + sovDnsCache(uint64(l))
	}
	if m.MaxHosts != nil {
		l = m.MaxHosts.Size()
		n += 1 + l + sovDnsCache(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDnsCache(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDnsCache(x uint64) (n int) {
	return sovDnsCache(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DnsCacheConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDnsCache
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
			return fmt.Errorf("proto: DnsCacheConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DnsCacheConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsCache
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
				return ErrInvalidLengthDnsCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDnsCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsLookupFamily", wireType)
			}
			m.DnsLookupFamily = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DnsLookupFamily |= v2.Cluster_DnsLookupFamily(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsRefreshRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsCache
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
				return ErrInvalidLengthDnsCache
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDnsCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsRefreshRate == nil {
				m.DnsRefreshRate = &types.Duration{}
			}
			if err := m.DnsRefreshRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostTtl", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsCache
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
				return ErrInvalidLengthDnsCache
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDnsCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HostTtl == nil {
				m.HostTtl = &types.Duration{}
			}
			if err := m.HostTtl.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHosts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsCache
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
				return ErrInvalidLengthDnsCache
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDnsCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxHosts == nil {
				m.MaxHosts = &types.UInt32Value{}
			}
			if err := m.MaxHosts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDnsCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDnsCache
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDnsCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDnsCache(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDnsCache
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
					return 0, ErrIntOverflowDnsCache
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
					return 0, ErrIntOverflowDnsCache
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
				return 0, ErrInvalidLengthDnsCache
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDnsCache
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDnsCache
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
				next, err := skipDnsCache(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDnsCache
				}
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
	ErrInvalidLengthDnsCache = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDnsCache   = fmt.Errorf("proto: integer overflow")
)
