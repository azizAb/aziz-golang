// Code generated by protoc-gen-go. DO NOT EDIT.
// source: report.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Report struct {
	Timestamp            string   `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Userid               string   `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Activities           string   `protobuf:"bytes,3,opt,name=activities,proto3" json:"activities,omitempty"`
	Applicationid        string   `protobuf:"bytes,4,opt,name=applicationid,proto3" json:"applicationid,omitempty"`
	Cif                  string   `protobuf:"bytes,5,opt,name=cif,proto3" json:"cif,omitempty"`
	Accountno            string   `protobuf:"bytes,6,opt,name=accountno,proto3" json:"accountno,omitempty"`
	Ktpid                string   `protobuf:"bytes,7,opt,name=ktpid,proto3" json:"ktpid,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Report) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *Report) GetActivities() string {
	if m != nil {
		return m.Activities
	}
	return ""
}

func (m *Report) GetApplicationid() string {
	if m != nil {
		return m.Applicationid
	}
	return ""
}

func (m *Report) GetCif() string {
	if m != nil {
		return m.Cif
	}
	return ""
}

func (m *Report) GetAccountno() string {
	if m != nil {
		return m.Accountno
	}
	return ""
}

func (m *Report) GetKtpid() string {
	if m != nil {
		return m.Ktpid
	}
	return ""
}

func (m *Report) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ReportList struct {
	List                 []*Report `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReportList) Reset()         { *m = ReportList{} }
func (m *ReportList) String() string { return proto.CompactTextString(m) }
func (*ReportList) ProtoMessage()    {}
func (*ReportList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{1}
}

func (m *ReportList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportList.Unmarshal(m, b)
}
func (m *ReportList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportList.Marshal(b, m, deterministic)
}
func (m *ReportList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportList.Merge(m, src)
}
func (m *ReportList) XXX_Size() int {
	return xxx_messageInfo_ReportList.Size(m)
}
func (m *ReportList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportList.DiscardUnknown(m)
}

var xxx_messageInfo_ReportList proto.InternalMessageInfo

func (m *ReportList) GetList() []*Report {
	if m != nil {
		return m.List
	}
	return nil
}

type ApplicationUserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationUserId) Reset()         { *m = ApplicationUserId{} }
func (m *ApplicationUserId) String() string { return proto.CompactTextString(m) }
func (*ApplicationUserId) ProtoMessage()    {}
func (*ApplicationUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{2}
}

func (m *ApplicationUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationUserId.Unmarshal(m, b)
}
func (m *ApplicationUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationUserId.Marshal(b, m, deterministic)
}
func (m *ApplicationUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationUserId.Merge(m, src)
}
func (m *ApplicationUserId) XXX_Size() int {
	return xxx_messageInfo_ApplicationUserId.Size(m)
}
func (m *ApplicationUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationUserId.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationUserId proto.InternalMessageInfo

func (m *ApplicationUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*Report)(nil), "proto.Report")
	proto.RegisterType((*ReportList)(nil), "proto.ReportList")
	proto.RegisterType((*ApplicationUserId)(nil), "proto.ApplicationUserId")
}

func init() { proto.RegisterFile("report.proto", fileDescriptor_3eedb623aa6ca98c) }

var fileDescriptor_3eedb623aa6ca98c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xa6, 0x02, 0x45, 0x46, 0x49, 0x64, 0x63, 0x70, 0x83, 0xc6, 0x60, 0xe3, 0x81, 0x83, 0x29,
	0x09, 0xea, 0x03, 0x70, 0xf0, 0x40, 0xe2, 0x89, 0xc4, 0xb3, 0x59, 0xba, 0x0b, 0x99, 0x08, 0xdd,
	0x4d, 0x77, 0x4a, 0xe2, 0xeb, 0xfa, 0x24, 0xa6, 0xb3, 0x55, 0x6c, 0x3c, 0xb5, 0xdf, 0xcf, 0xec,
	0x7e, 0xfb, 0x0d, 0x9c, 0x17, 0xc6, 0xd9, 0x82, 0x52, 0x57, 0x58, 0xb2, 0xa2, 0xcb, 0x9f, 0xf1,
	0xf5, 0xd6, 0xda, 0xed, 0xce, 0xcc, 0x18, 0xad, 0xcb, 0xcd, 0xec, 0x65, 0xef, 0xe8, 0x33, 0x78,
	0x92, 0xaf, 0x08, 0xe2, 0x15, 0x0f, 0x89, 0x1b, 0xe8, 0x13, 0xee, 0x8d, 0x27, 0xb5, 0x77, 0x32,
	0x9a, 0x44, 0xd3, 0xfe, 0xea, 0x48, 0x88, 0x11, 0xc4, 0xa5, 0x37, 0x05, 0x6a, 0x79, 0xc2, 0x52,
	0x8d, 0xc4, 0x2d, 0x80, 0xca, 0x08, 0x0f, 0x48, 0x68, 0xbc, 0x6c, 0xb3, 0xf6, 0x87, 0x11, 0xf7,
	0x30, 0x50, 0xce, 0xed, 0x30, 0x53, 0x84, 0x36, 0x47, 0x2d, 0x3b, 0x6c, 0x69, 0x92, 0xe2, 0x02,
	0xda, 0x19, 0x6e, 0x64, 0x97, 0xb5, 0xea, 0xb7, 0x4a, 0xa3, 0xb2, 0xcc, 0x96, 0x39, 0xe5, 0x56,
	0xc6, 0x21, 0xcd, 0x2f, 0x21, 0x2e, 0xa1, 0xfb, 0x41, 0x0e, 0xb5, 0xec, 0xb1, 0x12, 0x40, 0x95,
	0xd1, 0x93, 0xa2, 0xd2, 0xcb, 0xd3, 0x90, 0x31, 0xa0, 0x64, 0x06, 0x10, 0xde, 0xf8, 0x8a, 0x9e,
	0xc4, 0x1d, 0x74, 0x76, 0xe8, 0x49, 0x46, 0x93, 0xf6, 0xf4, 0x6c, 0x3e, 0x08, 0x45, 0xa4, 0xc1,
	0xb0, 0x62, 0x29, 0x79, 0x80, 0xe1, 0xe2, 0x98, 0xef, 0xcd, 0x9b, 0x62, 0xa9, 0xc5, 0x15, 0xf4,
	0xaa, 0x37, 0xbf, 0xa3, 0xae, 0xdb, 0xe1, 0x0a, 0x96, 0x7a, 0x7e, 0x80, 0x5e, 0x98, 0xf6, 0xe2,
	0x09, 0xfa, 0x0b, 0xad, 0xeb, 0x42, 0x9b, 0x47, 0x8f, 0x47, 0x69, 0x58, 0x44, 0xfa, 0xb3, 0x88,
	0x94, 0x17, 0x91, 0xb4, 0xc4, 0x33, 0x74, 0x38, 0x99, 0xac, 0x07, 0xfe, 0xdd, 0x3d, 0x1e, 0x36,
	0x8e, 0xaa, 0xcc, 0x49, 0x6b, 0x1d, 0x33, 0xf7, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xbb,
	0x10, 0x37, 0xf6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReportsClient is the client API for Reports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportsClient interface {
	AddReport(ctx context.Context, in *Report, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ApplicationUserId, opts ...grpc.CallOption) (*ReportList, error)
}

type reportsClient struct {
	cc *grpc.ClientConn
}

func NewReportsClient(cc *grpc.ClientConn) ReportsClient {
	return &reportsClient{cc}
}

func (c *reportsClient) AddReport(ctx context.Context, in *Report, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Reports/AddReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportsClient) List(ctx context.Context, in *ApplicationUserId, opts ...grpc.CallOption) (*ReportList, error) {
	out := new(ReportList)
	err := c.cc.Invoke(ctx, "/proto.Reports/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportsServer is the server API for Reports service.
type ReportsServer interface {
	AddReport(context.Context, *Report) (*empty.Empty, error)
	List(context.Context, *ApplicationUserId) (*ReportList, error)
}

// UnimplementedReportsServer can be embedded to have forward compatible implementations.
type UnimplementedReportsServer struct {
}

func (*UnimplementedReportsServer) AddReport(ctx context.Context, req *Report) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReport not implemented")
}
func (*UnimplementedReportsServer) List(ctx context.Context, req *ApplicationUserId) (*ReportList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterReportsServer(s *grpc.Server, srv ReportsServer) {
	s.RegisterService(&_Reports_serviceDesc, srv)
}

func _Reports_AddReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Report)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServer).AddReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Reports/AddReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServer).AddReport(ctx, req.(*Report))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reports_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Reports/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServer).List(ctx, req.(*ApplicationUserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reports_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Reports",
	HandlerType: (*ReportsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReport",
			Handler:    _Reports_AddReport_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Reports_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report.proto",
}
