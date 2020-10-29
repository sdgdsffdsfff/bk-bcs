// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/meshmanager/meshmanager.proto

package meshmanager

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// ErrCode error code
type ErrCode int32

const (
	ErrCode_ERROR_NO ErrCode = 0
	// 0 Sunccess
	ErrCode_ERROR_OK ErrCode = 100
	// 404 Not Found
	ErrCode_ERROR_NOT_FOUND ErrCode = 404
	// 91000 mesh cluster failed
	ErrCode_ERROR_MESH_CLUSTER_FAILED ErrCode = 91000
	//91001 MeshCluster already exist
	ErrCode_ERROR_MESH_CLUSTER_EXIST ErrCode = 91001
)

var ErrCode_name = map[int32]string{
	0:     "ERROR_NO",
	100:   "ERROR_OK",
	404:   "ERROR_NOT_FOUND",
	91000: "ERROR_MESH_CLUSTER_FAILED",
	91001: "ERROR_MESH_CLUSTER_EXIST",
}

var ErrCode_value = map[string]int32{
	"ERROR_NO":                  0,
	"ERROR_OK":                  100,
	"ERROR_NOT_FOUND":           404,
	"ERROR_MESH_CLUSTER_FAILED": 91000,
	"ERROR_MESH_CLUSTER_EXIST":  91001,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{0}
}

type CreateMeshClusterReq struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Clusterid            string   `protobuf:"bytes,2,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	Meshtype             string   `protobuf:"bytes,3,opt,name=meshtype,proto3" json:"meshtype,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Configurations       []string `protobuf:"bytes,6,rep,name=configurations,proto3" json:"configurations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMeshClusterReq) Reset()         { *m = CreateMeshClusterReq{} }
func (m *CreateMeshClusterReq) String() string { return proto.CompactTextString(m) }
func (*CreateMeshClusterReq) ProtoMessage()    {}
func (*CreateMeshClusterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{0}
}

func (m *CreateMeshClusterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMeshClusterReq.Unmarshal(m, b)
}
func (m *CreateMeshClusterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMeshClusterReq.Marshal(b, m, deterministic)
}
func (m *CreateMeshClusterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMeshClusterReq.Merge(m, src)
}
func (m *CreateMeshClusterReq) XXX_Size() int {
	return xxx_messageInfo_CreateMeshClusterReq.Size(m)
}
func (m *CreateMeshClusterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMeshClusterReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMeshClusterReq proto.InternalMessageInfo

func (m *CreateMeshClusterReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CreateMeshClusterReq) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

func (m *CreateMeshClusterReq) GetMeshtype() string {
	if m != nil {
		return m.Meshtype
	}
	return ""
}

func (m *CreateMeshClusterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateMeshClusterReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateMeshClusterReq) GetConfigurations() []string {
	if m != nil {
		return m.Configurations
	}
	return nil
}

type CreateMeshClusterResp struct {
	Seq                  uint64   `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	ErrCode              ErrCode  `protobuf:"varint,2,opt,name=errCode,proto3,enum=meshmanager.ErrCode" json:"errCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMeshClusterResp) Reset()         { *m = CreateMeshClusterResp{} }
func (m *CreateMeshClusterResp) String() string { return proto.CompactTextString(m) }
func (*CreateMeshClusterResp) ProtoMessage()    {}
func (*CreateMeshClusterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{1}
}

func (m *CreateMeshClusterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMeshClusterResp.Unmarshal(m, b)
}
func (m *CreateMeshClusterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMeshClusterResp.Marshal(b, m, deterministic)
}
func (m *CreateMeshClusterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMeshClusterResp.Merge(m, src)
}
func (m *CreateMeshClusterResp) XXX_Size() int {
	return xxx_messageInfo_CreateMeshClusterResp.Size(m)
}
func (m *CreateMeshClusterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMeshClusterResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMeshClusterResp proto.InternalMessageInfo

func (m *CreateMeshClusterResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *CreateMeshClusterResp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_ERROR_NO
}

func (m *CreateMeshClusterResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type DeleteMeshClusterReq struct {
	Clusterid            string   `protobuf:"bytes,2,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMeshClusterReq) Reset()         { *m = DeleteMeshClusterReq{} }
func (m *DeleteMeshClusterReq) String() string { return proto.CompactTextString(m) }
func (*DeleteMeshClusterReq) ProtoMessage()    {}
func (*DeleteMeshClusterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{2}
}

func (m *DeleteMeshClusterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMeshClusterReq.Unmarshal(m, b)
}
func (m *DeleteMeshClusterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMeshClusterReq.Marshal(b, m, deterministic)
}
func (m *DeleteMeshClusterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMeshClusterReq.Merge(m, src)
}
func (m *DeleteMeshClusterReq) XXX_Size() int {
	return xxx_messageInfo_DeleteMeshClusterReq.Size(m)
}
func (m *DeleteMeshClusterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMeshClusterReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMeshClusterReq proto.InternalMessageInfo

func (m *DeleteMeshClusterReq) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

type DeleteMeshClusterResp struct {
	Seq                  uint64   `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	ErrCode              ErrCode  `protobuf:"varint,2,opt,name=errCode,proto3,enum=meshmanager.ErrCode" json:"errCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMeshClusterResp) Reset()         { *m = DeleteMeshClusterResp{} }
func (m *DeleteMeshClusterResp) String() string { return proto.CompactTextString(m) }
func (*DeleteMeshClusterResp) ProtoMessage()    {}
func (*DeleteMeshClusterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{3}
}

func (m *DeleteMeshClusterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMeshClusterResp.Unmarshal(m, b)
}
func (m *DeleteMeshClusterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMeshClusterResp.Marshal(b, m, deterministic)
}
func (m *DeleteMeshClusterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMeshClusterResp.Merge(m, src)
}
func (m *DeleteMeshClusterResp) XXX_Size() int {
	return xxx_messageInfo_DeleteMeshClusterResp.Size(m)
}
func (m *DeleteMeshClusterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMeshClusterResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMeshClusterResp proto.InternalMessageInfo

func (m *DeleteMeshClusterResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *DeleteMeshClusterResp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_ERROR_NO
}

func (m *DeleteMeshClusterResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ListMeshClusterReq struct {
	Clusterid            string   `protobuf:"bytes,2,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMeshClusterReq) Reset()         { *m = ListMeshClusterReq{} }
func (m *ListMeshClusterReq) String() string { return proto.CompactTextString(m) }
func (*ListMeshClusterReq) ProtoMessage()    {}
func (*ListMeshClusterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{4}
}

func (m *ListMeshClusterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeshClusterReq.Unmarshal(m, b)
}
func (m *ListMeshClusterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeshClusterReq.Marshal(b, m, deterministic)
}
func (m *ListMeshClusterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeshClusterReq.Merge(m, src)
}
func (m *ListMeshClusterReq) XXX_Size() int {
	return xxx_messageInfo_ListMeshClusterReq.Size(m)
}
func (m *ListMeshClusterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeshClusterReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeshClusterReq proto.InternalMessageInfo

func (m *ListMeshClusterReq) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

type MeshCluster struct {
	Version              string                    `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Clusterid            string                    `protobuf:"bytes,2,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	Deletion             bool                      `protobuf:"varint,3,opt,name=deletion,proto3" json:"deletion,omitempty"`
	Components           map[string]*InstallStatus `protobuf:"bytes,4,rep,name=components,proto3" json:"components,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MeshCluster) Reset()         { *m = MeshCluster{} }
func (m *MeshCluster) String() string { return proto.CompactTextString(m) }
func (*MeshCluster) ProtoMessage()    {}
func (*MeshCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{5}
}

func (m *MeshCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshCluster.Unmarshal(m, b)
}
func (m *MeshCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshCluster.Marshal(b, m, deterministic)
}
func (m *MeshCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshCluster.Merge(m, src)
}
func (m *MeshCluster) XXX_Size() int {
	return xxx_messageInfo_MeshCluster.Size(m)
}
func (m *MeshCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshCluster.DiscardUnknown(m)
}

var xxx_messageInfo_MeshCluster proto.InternalMessageInfo

func (m *MeshCluster) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MeshCluster) GetClusterid() string {
	if m != nil {
		return m.Clusterid
	}
	return ""
}

func (m *MeshCluster) GetDeletion() bool {
	if m != nil {
		return m.Deletion
	}
	return false
}

func (m *MeshCluster) GetComponents() map[string]*InstallStatus {
	if m != nil {
		return m.Components
	}
	return nil
}

type InstallStatus struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallStatus) Reset()         { *m = InstallStatus{} }
func (m *InstallStatus) String() string { return proto.CompactTextString(m) }
func (*InstallStatus) ProtoMessage()    {}
func (*InstallStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{6}
}

func (m *InstallStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallStatus.Unmarshal(m, b)
}
func (m *InstallStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallStatus.Marshal(b, m, deterministic)
}
func (m *InstallStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallStatus.Merge(m, src)
}
func (m *InstallStatus) XXX_Size() int {
	return xxx_messageInfo_InstallStatus.Size(m)
}
func (m *InstallStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InstallStatus proto.InternalMessageInfo

func (m *InstallStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallStatus) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *InstallStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *InstallStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListMeshClusterResp struct {
	Seq                  uint64         `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	ErrCode              ErrCode        `protobuf:"varint,2,opt,name=errCode,proto3,enum=meshmanager.ErrCode" json:"errCode,omitempty"`
	ErrMsg               string         `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	MeshClusters         []*MeshCluster `protobuf:"bytes,4,rep,name=meshClusters,proto3" json:"meshClusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListMeshClusterResp) Reset()         { *m = ListMeshClusterResp{} }
func (m *ListMeshClusterResp) String() string { return proto.CompactTextString(m) }
func (*ListMeshClusterResp) ProtoMessage()    {}
func (*ListMeshClusterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ddc37088f698422, []int{7}
}

func (m *ListMeshClusterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeshClusterResp.Unmarshal(m, b)
}
func (m *ListMeshClusterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeshClusterResp.Marshal(b, m, deterministic)
}
func (m *ListMeshClusterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeshClusterResp.Merge(m, src)
}
func (m *ListMeshClusterResp) XXX_Size() int {
	return xxx_messageInfo_ListMeshClusterResp.Size(m)
}
func (m *ListMeshClusterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeshClusterResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeshClusterResp proto.InternalMessageInfo

func (m *ListMeshClusterResp) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ListMeshClusterResp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_ERROR_NO
}

func (m *ListMeshClusterResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ListMeshClusterResp) GetMeshClusters() []*MeshCluster {
	if m != nil {
		return m.MeshClusters
	}
	return nil
}

func init() {
	proto.RegisterEnum("meshmanager.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*CreateMeshClusterReq)(nil), "meshmanager.CreateMeshClusterReq")
	proto.RegisterType((*CreateMeshClusterResp)(nil), "meshmanager.CreateMeshClusterResp")
	proto.RegisterType((*DeleteMeshClusterReq)(nil), "meshmanager.DeleteMeshClusterReq")
	proto.RegisterType((*DeleteMeshClusterResp)(nil), "meshmanager.DeleteMeshClusterResp")
	proto.RegisterType((*ListMeshClusterReq)(nil), "meshmanager.ListMeshClusterReq")
	proto.RegisterType((*MeshCluster)(nil), "meshmanager.MeshCluster")
	proto.RegisterMapType((map[string]*InstallStatus)(nil), "meshmanager.MeshCluster.ComponentsEntry")
	proto.RegisterType((*InstallStatus)(nil), "meshmanager.InstallStatus")
	proto.RegisterType((*ListMeshClusterResp)(nil), "meshmanager.ListMeshClusterResp")
}

func init() {
	proto.RegisterFile("proto/meshmanager/meshmanager.proto", fileDescriptor_6ddc37088f698422)
}

var fileDescriptor_6ddc37088f698422 = []byte{
	// 1088 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6c, 0xdb, 0x54,
	0x18, 0xc7, 0x4d, 0xda, 0xb5, 0xaf, 0xdd, 0x9a, 0xbd, 0x76, 0xc3, 0x35, 0x1d, 0xf5, 0xdc, 0x4b,
	0x65, 0xd6, 0xb4, 0x75, 0xf9, 0x33, 0x05, 0x71, 0x70, 0xd2, 0x4c, 0xeb, 0xd6, 0x34, 0xe8, 0xa5,
	0x93, 0x00, 0x81, 0x22, 0xe3, 0xbc, 0xa5, 0x66, 0x89, 0xed, 0xf9, 0x39, 0x45, 0xd5, 0x40, 0x1a,
	0x3b, 0x4c, 0xcb, 0xc4, 0x40, 0x32, 0x13, 0xd2, 0xd4, 0x6a, 0x02, 0x34, 0x4d, 0x42, 0xe2, 0xcc,
	0xa6, 0x0e, 0x8d, 0xdb, 0xb8, 0x97, 0x33, 0xb7, 0x25, 0xdd, 0x4e, 0x5c, 0xe1, 0x86, 0xfc, 0x2f,
	0x76, 0x12, 0xa7, 0x45, 0x88, 0xc3, 0x2e, 0xef, 0x7d, 0x7f, 0x9f, 0xbf, 0xef, 0xf7, 0x7d, 0xef,
	0x7b, 0x06, 0xd3, 0xba, 0xa1, 0x99, 0xda, 0x5c, 0x15, 0x93, 0xf5, 0xaa, 0xa4, 0x4a, 0x65, 0x6c,
	0x84, 0xe9, 0xa4, 0xa3, 0x85, 0xc3, 0x21, 0x11, 0x33, 0x59, 0xd6, 0xb4, 0x72, 0x05, 0xcf, 0x49,
	0xba, 0x32, 0x27, 0xa9, 0xaa, 0x66, 0x4a, 0xa6, 0xa2, 0xa9, 0xc4, 0x35, 0x65, 0x4e, 0x39, 0x9b,
	0x3c, 0x5b, 0xc6, 0xea, 0x2c, 0xf9, 0x54, 0x2a, 0xdb, 0x07, 0x6a, 0xba, 0x63, 0xd1, 0x6d, 0xcd,
	0xd5, 0xe3, 0x60, 0x3c, 0x63, 0x60, 0xc9, 0xc4, 0x39, 0x4c, 0xd6, 0x33, 0x95, 0x1a, 0x31, 0xb1,
	0x81, 0xf0, 0x65, 0x38, 0x07, 0x0e, 0x6d, 0x60, 0x83, 0x28, 0x9a, 0x4a, 0x53, 0x2c, 0x35, 0x33,
	0x94, 0x3e, 0x66, 0x89, 0x90, 0xf7, 0x65, 0x82, 0x4f, 0x20, 0x9f, 0x80, 0x6f, 0x81, 0x21, 0xd9,
	0x75, 0x57, 0x4a, 0x74, 0x9f, 0xe3, 0x32, 0x61, 0x89, 0xc7, 0xf9, 0x40, 0x2a, 0x04, 0x24, 0x0a,
	0x48, 0xb8, 0x08, 0x06, 0xed, 0xec, 0xcc, 0x4d, 0x1d, 0xd3, 0x31, 0xc7, 0xef, 0x65, 0x4b, 0x1c,
	0xe7, 0x5b, 0x42, 0xa1, 0x45, 0xa1, 0x16, 0x05, 0xa7, 0x41, 0x5c, 0x95, 0xaa, 0x98, 0x8e, 0x3b,
	0x0e, 0xa3, 0x96, 0x38, 0xc2, 0x3b, 0x02, 0xc1, 0x59, 0x91, 0xb3, 0xda, 0x21, 0xd9, 0x3b, 0xd1,
	0x25, 0x19, 0xd3, 0xfd, 0xa1, 0x90, 0x5a, 0x52, 0x21, 0x20, 0x51, 0x40, 0xc2, 0xf3, 0xe0, 0x88,
	0xac, 0xa9, 0x17, 0x95, 0x72, 0xcd, 0x70, 0xd1, 0xa2, 0x07, 0xd8, 0xd8, 0xcc, 0x50, 0x7a, 0xda,
	0x12, 0x59, 0xbe, 0x43, 0x25, 0x74, 0xf0, 0xa8, 0x83, 0x4f, 0xdd, 0xa7, 0x2c, 0xf1, 0x07, 0x0a,
	0x20, 0x3e, 0x12, 0x68, 0x61, 0xa2, 0x71, 0x63, 0xab, 0xf1, 0x68, 0xe7, 0xd2, 0x69, 0xf2, 0x7c,
	0xeb, 0xfa, 0xde, 0xe3, 0x7b, 0x76, 0x8a, 0xcd, 0xed, 0x5b, 0x8d, 0x9b, 0xdf, 0xef, 0x52, 0x3e,
	0xb8, 0xbb, 0x54, 0x80, 0xd7, 0x2e, 0x15, 0x20, 0xb3, 0x02, 0xcf, 0x5d, 0x61, 0x39, 0xcf, 0x88,
	0x4b, 0xb1, 0xdc, 0x42, 0xf2, 0x8d, 0xe4, 0xeb, 0xdc, 0x29, 0x96, 0x6b, 0x99, 0xdb, 0xd2, 0x74,
	0xa6, 0x30, 0x9b, 0xcb, 0x16, 0xf2, 0x85, 0xd9, 0x85, 0xf9, 0xf9, 0x45, 0xc1, 0xd6, 0xfb, 0x47,
	0xd8, 0x6a, 0x85, 0x98, 0x8a, 0xc6, 0x7d, 0xce, 0xfd, 0x44, 0x81, 0x63, 0x11, 0x21, 0x12, 0x1d,
	0x4e, 0x81, 0x18, 0xc1, 0x97, 0x9d, 0x46, 0x88, 0xa7, 0x0f, 0x5b, 0x22, 0xe0, 0x6d, 0x5e, 0xb0,
	0x17, 0x64, 0x2f, 0x30, 0x0b, 0x0e, 0x61, 0xc3, 0xc8, 0x68, 0x25, 0xec, 0x94, 0xfe, 0x88, 0x30,
	0x9e, 0x0c, 0x37, 0x71, 0xd6, 0xd5, 0x79, 0x3d, 0xe4, 0x59, 0x0a, 0x3e, 0x81, 0x7c, 0x02, 0xbe,
	0x06, 0x06, 0xb0, 0x61, 0xe4, 0x48, 0xd9, 0x6b, 0x84, 0x31, 0x4b, 0x4c, 0xf0, 0x9e, 0x48, 0xf0,
	0x76, 0xe4, 0xed, 0xdc, 0x03, 0x0a, 0x8c, 0x2f, 0xe1, 0x0a, 0xee, 0x6a, 0xdd, 0xff, 0xda, 0x89,
	0xa9, 0x8b, 0x96, 0x28, 0x83, 0x77, 0xf8, 0xc8, 0x53, 0xed, 0x3a, 0x7d, 0xf7, 0xfc, 0xce, 0xbd,
	0xc8, 0x3a, 0x85, 0x4e, 0xe6, 0x20, 0x7b, 0x65, 0x7f, 0xf4, 0x5d, 0xa0, 0x23, 0xbe, 0xf1, 0xa2,
	0x02, 0xfd, 0x07, 0x05, 0xe0, 0x8a, 0x42, 0xcc, 0xff, 0x0b, 0xe6, 0x5b, 0x94, 0x25, 0xd6, 0x29,
	0x20, 0xf3, 0x63, 0x1d, 0xa7, 0x3a, 0x38, 0xaf, 0x3c, 0xfb, 0xea, 0x41, 0xe3, 0xeb, 0x6f, 0x23,
	0x70, 0xde, 0xbb, 0xf9, 0x73, 0xf3, 0xea, 0x17, 0x7f, 0xfe, 0xf2, 0x65, 0xe3, 0xfe, 0xb5, 0xe6,
	0xdd, 0xed, 0xd6, 0x81, 0x4f, 0x1f, 0xee, 0xec, 0xfd, 0xb8, 0xd3, 0xb8, 0x71, 0xe7, 0xd9, 0xaf,
	0xb7, 0x1b, 0x5b, 0x77, 0x9b, 0xf5, 0xab, 0xcd, 0xed, 0xba, 0xeb, 0xfc, 0xaf, 0xca, 0x71, 0xbd,
	0x0f, 0x0c, 0x87, 0xa2, 0x80, 0x74, 0xc7, 0xe8, 0x0b, 0x66, 0xdc, 0x64, 0x57, 0xca, 0xe1, 0x41,
	0xc6, 0x80, 0xc1, 0x92, 0x5d, 0x55, 0xdb, 0xd1, 0x86, 0x75, 0x10, 0xb5, 0x78, 0x78, 0x16, 0x00,
	0x59, 0xab, 0xea, 0x9a, 0x8a, 0x55, 0x93, 0xd0, 0x71, 0x36, 0x36, 0x33, 0x2c, 0xcc, 0xb4, 0x95,
	0x2e, 0x14, 0x41, 0x32, 0xd3, 0x32, 0xcd, 0xaa, 0xa6, 0xb1, 0x89, 0x42, 0xbe, 0xcc, 0xfb, 0x60,
	0xb4, 0x43, 0x0d, 0x13, 0x20, 0x76, 0x09, 0x6f, 0x7a, 0xc1, 0xda, 0x24, 0x9c, 0x07, 0xfd, 0x1b,
	0x52, 0xa5, 0xe6, 0x36, 0xc9, 0xb0, 0xc0, 0xb4, 0x7d, 0x69, 0x59, 0x25, 0xa6, 0x54, 0xa9, 0x14,
	0x4c, 0xc9, 0xac, 0x11, 0xe4, 0x1a, 0xa6, 0xfa, 0x4e, 0x53, 0x1c, 0x01, 0x87, 0xdb, 0x74, 0x10,
	0x7a, 0x53, 0xd6, 0x3d, 0xd9, 0x1d, 0xaa, 0x93, 0xe1, 0xa1, 0xea, 0x61, 0x10, 0x4c, 0xce, 0xe3,
	0x60, 0x80, 0x38, 0xbe, 0x6e, 0x63, 0x21, 0x8f, 0xb3, 0x31, 0xad, 0x62, 0x42, 0xa4, 0xb2, 0x37,
	0xb2, 0x91, 0xcf, 0x72, 0xf5, 0x3e, 0x30, 0xd6, 0xd5, 0x5d, 0x2f, 0xe8, 0x55, 0x80, 0x1f, 0x81,
	0x91, 0x6a, 0x10, 0xa7, 0x5f, 0x48, 0xba, 0x57, 0x21, 0xd3, 0x53, 0x96, 0x38, 0xc9, 0xb7, 0xb9,
	0x08, 0x6d, 0x1c, 0x6a, 0xe3, 0xf8, 0xcf, 0xc0, 0x21, 0x2f, 0x6c, 0x38, 0x02, 0x06, 0xb3, 0x08,
	0xe5, 0x51, 0x71, 0x35, 0x9f, 0x78, 0x29, 0xe0, 0xf2, 0xe7, 0x13, 0x25, 0x38, 0x0e, 0x46, 0x7d,
	0xdd, 0x5a, 0xf1, 0x4c, 0xfe, 0xc2, 0xea, 0x52, 0xe2, 0x76, 0x0c, 0x4e, 0x81, 0x09, 0x57, 0x9a,
	0xcb, 0x16, 0xce, 0x16, 0x33, 0x2b, 0x17, 0x0a, 0x6b, 0x59, 0x54, 0x3c, 0x23, 0x2e, 0xaf, 0x64,
	0x97, 0x12, 0x7f, 0x3d, 0xe9, 0x87, 0xaf, 0x02, 0x3a, 0xc2, 0x20, 0xfb, 0xde, 0x72, 0x61, 0x2d,
	0xf1, 0xf7, 0x93, 0x7e, 0xe1, 0xb7, 0xb8, 0x7b, 0x0f, 0x72, 0x6e, 0x22, 0xf0, 0x11, 0x05, 0x8e,
	0x76, 0xbd, 0x07, 0xf0, 0x64, 0x5b, 0xb2, 0x51, 0x4f, 0x1a, 0xc3, 0x1d, 0x64, 0x42, 0x74, 0xae,
	0x68, 0x89, 0x6f, 0xc2, 0xde, 0x6f, 0x1f, 0xd3, 0x5b, 0x75, 0xed, 0xf7, 0xa7, 0xdf, 0xf4, 0xb1,
	0xdc, 0x2b, 0x6d, 0xbf, 0x4c, 0x1b, 0x0b, 0x0e, 0xeb, 0x5d, 0xc6, 0x14, 0xc5, 0xc3, 0x87, 0x14,
	0x38, 0xda, 0x35, 0x64, 0x3b, 0xa2, 0x8f, 0x1a, 0xf4, 0x1d, 0xd1, 0x47, 0xce, 0x69, 0xee, 0x43,
	0x2f, 0xfa, 0x1e, 0x2f, 0x02, 0xd3, 0x5b, 0xe5, 0x44, 0x7f, 0x82, 0xdf, 0x2f, 0x7a, 0xf8, 0x98,
	0x02, 0xa3, 0x1d, 0x57, 0x02, 0x4e, 0xb5, 0x45, 0xd5, 0x3d, 0x8e, 0x19, 0x76, 0x7f, 0x03, 0xa2,
	0x73, 0xeb, 0x96, 0x98, 0x86, 0x53, 0x07, 0x8c, 0x57, 0xe6, 0x20, 0x03, 0x37, 0x01, 0xb8, 0x5f,
	0x02, 0xe9, 0x8a, 0x25, 0x9e, 0x83, 0x27, 0x01, 0x9d, 0xa9, 0x68, 0xb5, 0x12, 0xbb, 0x8a, 0x4d,
	0x82, 0x8d, 0x0d, 0x45, 0xc6, 0xac, 0xf8, 0xee, 0x32, 0xbb, 0xa4, 0xc9, 0x42, 0xff, 0x7c, 0x72,
	0x3e, 0xb9, 0xc0, 0x53, 0x94, 0x90, 0x90, 0x74, 0xbd, 0xa2, 0xc8, 0xce, 0xbf, 0xd2, 0xdc, 0x27,
	0x44, 0x53, 0x53, 0x5d, 0x92, 0x0f, 0x4e, 0x74, 0xfd, 0x22, 0xbf, 0x1d, 0xa2, 0x3f, 0x1e, 0x70,
	0xd4, 0x8b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x82, 0x0d, 0xa3, 0x47, 0x4a, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeshManagerClient is the client API for MeshManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshManagerClient interface {
	CreateMeshCluster(ctx context.Context, in *CreateMeshClusterReq, opts ...grpc.CallOption) (*CreateMeshClusterResp, error)
	DeleteMeshCluster(ctx context.Context, in *DeleteMeshClusterReq, opts ...grpc.CallOption) (*DeleteMeshClusterResp, error)
	ListMeshCluster(ctx context.Context, in *ListMeshClusterReq, opts ...grpc.CallOption) (*ListMeshClusterResp, error)
}

type meshManagerClient struct {
	cc *grpc.ClientConn
}

func NewMeshManagerClient(cc *grpc.ClientConn) MeshManagerClient {
	return &meshManagerClient{cc}
}

func (c *meshManagerClient) CreateMeshCluster(ctx context.Context, in *CreateMeshClusterReq, opts ...grpc.CallOption) (*CreateMeshClusterResp, error) {
	out := new(CreateMeshClusterResp)
	err := c.cc.Invoke(ctx, "/meshmanager.MeshManager/CreateMeshCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshManagerClient) DeleteMeshCluster(ctx context.Context, in *DeleteMeshClusterReq, opts ...grpc.CallOption) (*DeleteMeshClusterResp, error) {
	out := new(DeleteMeshClusterResp)
	err := c.cc.Invoke(ctx, "/meshmanager.MeshManager/DeleteMeshCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshManagerClient) ListMeshCluster(ctx context.Context, in *ListMeshClusterReq, opts ...grpc.CallOption) (*ListMeshClusterResp, error) {
	out := new(ListMeshClusterResp)
	err := c.cc.Invoke(ctx, "/meshmanager.MeshManager/ListMeshCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshManagerServer is the server API for MeshManager service.
type MeshManagerServer interface {
	CreateMeshCluster(context.Context, *CreateMeshClusterReq) (*CreateMeshClusterResp, error)
	DeleteMeshCluster(context.Context, *DeleteMeshClusterReq) (*DeleteMeshClusterResp, error)
	ListMeshCluster(context.Context, *ListMeshClusterReq) (*ListMeshClusterResp, error)
}

// UnimplementedMeshManagerServer can be embedded to have forward compatible implementations.
type UnimplementedMeshManagerServer struct {
}

func (*UnimplementedMeshManagerServer) CreateMeshCluster(ctx context.Context, req *CreateMeshClusterReq) (*CreateMeshClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeshCluster not implemented")
}
func (*UnimplementedMeshManagerServer) DeleteMeshCluster(ctx context.Context, req *DeleteMeshClusterReq) (*DeleteMeshClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeshCluster not implemented")
}
func (*UnimplementedMeshManagerServer) ListMeshCluster(ctx context.Context, req *ListMeshClusterReq) (*ListMeshClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMeshCluster not implemented")
}

func RegisterMeshManagerServer(s *grpc.Server, srv MeshManagerServer) {
	s.RegisterService(&_MeshManager_serviceDesc, srv)
}

func _MeshManager_CreateMeshCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMeshClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshManagerServer).CreateMeshCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshmanager.MeshManager/CreateMeshCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshManagerServer).CreateMeshCluster(ctx, req.(*CreateMeshClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshManager_DeleteMeshCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMeshClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshManagerServer).DeleteMeshCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshmanager.MeshManager/DeleteMeshCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshManagerServer).DeleteMeshCluster(ctx, req.(*DeleteMeshClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshManager_ListMeshCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMeshClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshManagerServer).ListMeshCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshmanager.MeshManager/ListMeshCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshManagerServer).ListMeshCluster(ctx, req.(*ListMeshClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeshManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meshmanager.MeshManager",
	HandlerType: (*MeshManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeshCluster",
			Handler:    _MeshManager_CreateMeshCluster_Handler,
		},
		{
			MethodName: "DeleteMeshCluster",
			Handler:    _MeshManager_DeleteMeshCluster_Handler,
		},
		{
			MethodName: "ListMeshCluster",
			Handler:    _MeshManager_ListMeshCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/meshmanager/meshmanager.proto",
}