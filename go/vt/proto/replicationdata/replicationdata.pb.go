// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replicationdata.proto

package replicationdata

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Status is the replication status for MySQL (returned by 'show slave status'
// and parsed into a Position and fields).
type Status struct {
	Position string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// RelayLogPosition will be empty for flavors that do not support returning the full GTIDSet from the relay log, such as MariaDB.
	RelayLogPosition     string   `protobuf:"bytes,8,opt,name=relay_log_position,json=relayLogPosition,proto3" json:"relay_log_position,omitempty"`
	FilePosition         string   `protobuf:"bytes,9,opt,name=file_position,json=filePosition,proto3" json:"file_position,omitempty"`
	FileRelayLogPosition string   `protobuf:"bytes,10,opt,name=file_relay_log_position,json=fileRelayLogPosition,proto3" json:"file_relay_log_position,omitempty"`
	MasterServerId       uint32   `protobuf:"varint,11,opt,name=master_server_id,json=masterServerId,proto3" json:"master_server_id,omitempty"`
	SlaveIoRunning       bool     `protobuf:"varint,2,opt,name=slave_io_running,json=slaveIoRunning,proto3" json:"slave_io_running,omitempty"`
	SlaveSqlRunning      bool     `protobuf:"varint,3,opt,name=slave_sql_running,json=slaveSqlRunning,proto3" json:"slave_sql_running,omitempty"`
	SecondsBehindMaster  uint32   `protobuf:"varint,4,opt,name=seconds_behind_master,json=secondsBehindMaster,proto3" json:"seconds_behind_master,omitempty"`
	MasterHost           string   `protobuf:"bytes,5,opt,name=master_host,json=masterHost,proto3" json:"master_host,omitempty"`
	MasterPort           int32    `protobuf:"varint,6,opt,name=master_port,json=masterPort,proto3" json:"master_port,omitempty"`
	MasterConnectRetry   int32    `protobuf:"varint,7,opt,name=master_connect_retry,json=masterConnectRetry,proto3" json:"master_connect_retry,omitempty"`
	MasterUuid           string   `protobuf:"bytes,12,opt,name=master_uuid,json=masterUuid,proto3" json:"master_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8ee22b8c4b9d06, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Status) GetRelayLogPosition() string {
	if m != nil {
		return m.RelayLogPosition
	}
	return ""
}

func (m *Status) GetFilePosition() string {
	if m != nil {
		return m.FilePosition
	}
	return ""
}

func (m *Status) GetFileRelayLogPosition() string {
	if m != nil {
		return m.FileRelayLogPosition
	}
	return ""
}

func (m *Status) GetMasterServerId() uint32 {
	if m != nil {
		return m.MasterServerId
	}
	return 0
}

func (m *Status) GetSlaveIoRunning() bool {
	if m != nil {
		return m.SlaveIoRunning
	}
	return false
}

func (m *Status) GetSlaveSqlRunning() bool {
	if m != nil {
		return m.SlaveSqlRunning
	}
	return false
}

func (m *Status) GetSecondsBehindMaster() uint32 {
	if m != nil {
		return m.SecondsBehindMaster
	}
	return 0
}

func (m *Status) GetMasterHost() string {
	if m != nil {
		return m.MasterHost
	}
	return ""
}

func (m *Status) GetMasterPort() int32 {
	if m != nil {
		return m.MasterPort
	}
	return 0
}

func (m *Status) GetMasterConnectRetry() int32 {
	if m != nil {
		return m.MasterConnectRetry
	}
	return 0
}

func (m *Status) GetMasterUuid() string {
	if m != nil {
		return m.MasterUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "replicationdata.Status")
}

func init() { proto.RegisterFile("replicationdata.proto", fileDescriptor_ee8ee22b8c4b9d06) }

var fileDescriptor_ee8ee22b8c4b9d06 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0x80, 0xf1, 0xb2, 0x64, 0x89, 0xf2, 0x73, 0x5a, 0xc2, 0xc4, 0x2e, 0x33, 0xdb, 0xc5, 0x94,
	0x10, 0x97, 0x96, 0xfe, 0x03, 0xe9, 0xa5, 0x81, 0x16, 0x82, 0x43, 0x2f, 0xbd, 0x08, 0xc7, 0x56,
	0x1d, 0x81, 0xab, 0xe7, 0x48, 0xb2, 0x21, 0x7f, 0x79, 0xaf, 0xc5, 0x4f, 0x89, 0x1b, 0x42, 0x6f,
	0xf6, 0xf7, 0x7d, 0x3c, 0x49, 0xf0, 0xc8, 0x4c, 0x8b, 0x22, 0x97, 0x49, 0x6c, 0x25, 0xa8, 0x34,
	0xb6, 0xf1, 0xa2, 0xd0, 0x60, 0x81, 0x8e, 0x2f, 0xf0, 0xbf, 0xf7, 0x16, 0xe9, 0x6c, 0x6c, 0x6c,
	0x4b, 0x43, 0xff, 0x90, 0x6e, 0x01, 0x46, 0xd6, 0x8a, 0x79, 0xbe, 0x17, 0xf4, 0xa2, 0xe6, 0x9f,
	0xce, 0x09, 0xd5, 0x22, 0x8f, 0x0f, 0x3c, 0x87, 0x8c, 0x37, 0x55, 0x17, 0xab, 0x09, 0x9a, 0x47,
	0xc8, 0xd6, 0xa7, 0xfa, 0x3f, 0x19, 0xbe, 0xca, 0x5c, 0x7c, 0x86, 0x3d, 0x0c, 0x07, 0x35, 0x6c,
	0xa2, 0x3b, 0xf2, 0x1b, 0xa3, 0x2f, 0xe6, 0x12, 0xcc, 0xa7, 0xb5, 0x8e, 0x2e, 0x67, 0x07, 0x64,
	0xf2, 0x16, 0x1b, 0x2b, 0x34, 0x37, 0x42, 0x57, 0x42, 0x73, 0x99, 0xb2, 0xbe, 0xef, 0x05, 0xc3,
	0x68, 0xe4, 0xf8, 0x06, 0xf1, 0x2a, 0xad, 0x4b, 0x93, 0xc7, 0x95, 0xe0, 0x12, 0xb8, 0x2e, 0x95,
	0x92, 0x2a, 0x63, 0xdf, 0x7c, 0x2f, 0xe8, 0x46, 0x23, 0xe4, 0x2b, 0x88, 0x1c, 0xa5, 0x57, 0xe4,
	0xa7, 0x2b, 0xcd, 0x3e, 0x6f, 0xd2, 0x16, 0xa6, 0x63, 0x14, 0x9b, 0x7d, 0x7e, 0x6a, 0x6f, 0xc8,
	0xcc, 0x88, 0x04, 0x54, 0x6a, 0xf8, 0x56, 0xec, 0xa4, 0x4a, 0xb9, 0x3b, 0x96, 0x7d, 0xc7, 0x4b,
	0xfc, 0x3a, 0xca, 0x25, 0xba, 0x27, 0x54, 0xf4, 0x2f, 0xe9, 0x1f, 0xef, 0xbc, 0x03, 0x63, 0x59,
	0x1b, 0x9f, 0x47, 0x1c, 0x7a, 0x00, 0x63, 0xcf, 0x82, 0x02, 0xb4, 0x65, 0x1d, 0xdf, 0x0b, 0xda,
	0xa7, 0x60, 0x0d, 0xda, 0xd2, 0x6b, 0x32, 0x3d, 0x06, 0x09, 0x28, 0x25, 0x12, 0xcb, 0xb5, 0xb0,
	0xfa, 0xc0, 0x7e, 0x60, 0x49, 0x9d, 0xbb, 0x77, 0x2a, 0xaa, 0xcd, 0xd9, 0xc8, 0xb2, 0x94, 0x29,
	0x1b, 0x9c, 0x9f, 0xf9, 0x5c, 0xca, 0x74, 0xb9, 0x78, 0x99, 0x57, 0xd2, 0x0a, 0x63, 0x16, 0x12,
	0x42, 0xf7, 0x15, 0x66, 0x10, 0x56, 0x36, 0xc4, 0x55, 0x09, 0x2f, 0x36, 0x65, 0xdb, 0x41, 0x7c,
	0xfb, 0x11, 0x00, 0x00, 0xff, 0xff, 0x42, 0xb5, 0xb6, 0xf7, 0x5a, 0x02, 0x00, 0x00,
}
