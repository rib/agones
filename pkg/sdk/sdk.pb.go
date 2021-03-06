// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdk.proto

package sdk

import (
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// I am Empty
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// Key, Value entry
type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{1}
}
func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (dst *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(dst, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// time duration, in seconds
type Duration struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{2}
}
func (m *Duration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duration.Unmarshal(m, b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
}
func (dst *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(dst, src)
}
func (m *Duration) XXX_Size() int {
	return xxx_messageInfo_Duration.Size(m)
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

// A GameServer Custom Resource Definition object
// We will only export those resources that make the most
// sense. Can always expand to more as needed.
type GameServer struct {
	ObjectMeta           *GameServer_ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Spec                 *GameServer_Spec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status               *GameServer_Status     `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GameServer) Reset()         { *m = GameServer{} }
func (m *GameServer) String() string { return proto.CompactTextString(m) }
func (*GameServer) ProtoMessage()    {}
func (*GameServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{3}
}
func (m *GameServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer.Unmarshal(m, b)
}
func (m *GameServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer.Marshal(b, m, deterministic)
}
func (dst *GameServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer.Merge(dst, src)
}
func (m *GameServer) XXX_Size() int {
	return xxx_messageInfo_GameServer.Size(m)
}
func (m *GameServer) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer proto.InternalMessageInfo

func (m *GameServer) GetObjectMeta() *GameServer_ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *GameServer) GetSpec() *GameServer_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *GameServer) GetStatus() *GameServer_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// representation of the K8s ObjectMeta resource
type GameServer_ObjectMeta struct {
	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace       string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid             string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	Generation      int64  `protobuf:"varint,5,opt,name=generation,proto3" json:"generation,omitempty"`
	// timestamp is in Epoch format, unit: seconds
	CreationTimestamp int64 `protobuf:"varint,6,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	// optional deletion timestamp in Epoch format, unit: seconds
	DeletionTimestamp    int64             `protobuf:"varint,7,opt,name=deletion_timestamp,json=deletionTimestamp,proto3" json:"deletion_timestamp,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GameServer_ObjectMeta) Reset()         { *m = GameServer_ObjectMeta{} }
func (m *GameServer_ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*GameServer_ObjectMeta) ProtoMessage()    {}
func (*GameServer_ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{3, 0}
}
func (m *GameServer_ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_ObjectMeta.Unmarshal(m, b)
}
func (m *GameServer_ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_ObjectMeta.Marshal(b, m, deterministic)
}
func (dst *GameServer_ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_ObjectMeta.Merge(dst, src)
}
func (m *GameServer_ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_GameServer_ObjectMeta.Size(m)
}
func (m *GameServer_ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_ObjectMeta proto.InternalMessageInfo

func (m *GameServer_ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *GameServer_ObjectMeta) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *GameServer_ObjectMeta) GetCreationTimestamp() int64 {
	if m != nil {
		return m.CreationTimestamp
	}
	return 0
}

func (m *GameServer_ObjectMeta) GetDeletionTimestamp() int64 {
	if m != nil {
		return m.DeletionTimestamp
	}
	return 0
}

func (m *GameServer_ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *GameServer_ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type GameServer_Spec struct {
	Health               *GameServer_Spec_Health `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GameServer_Spec) Reset()         { *m = GameServer_Spec{} }
func (m *GameServer_Spec) String() string { return proto.CompactTextString(m) }
func (*GameServer_Spec) ProtoMessage()    {}
func (*GameServer_Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{3, 1}
}
func (m *GameServer_Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Spec.Unmarshal(m, b)
}
func (m *GameServer_Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Spec.Marshal(b, m, deterministic)
}
func (dst *GameServer_Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Spec.Merge(dst, src)
}
func (m *GameServer_Spec) XXX_Size() int {
	return xxx_messageInfo_GameServer_Spec.Size(m)
}
func (m *GameServer_Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Spec proto.InternalMessageInfo

func (m *GameServer_Spec) GetHealth() *GameServer_Spec_Health {
	if m != nil {
		return m.Health
	}
	return nil
}

type GameServer_Spec_Health struct {
	Disabled             bool     `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	PeriodSeconds        int32    `protobuf:"varint,2,opt,name=period_seconds,json=periodSeconds,proto3" json:"period_seconds,omitempty"`
	FailureThreshold     int32    `protobuf:"varint,3,opt,name=failure_threshold,json=failureThreshold,proto3" json:"failure_threshold,omitempty"`
	InitialDelaySeconds  int32    `protobuf:"varint,4,opt,name=initial_delay_seconds,json=initialDelaySeconds,proto3" json:"initial_delay_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServer_Spec_Health) Reset()         { *m = GameServer_Spec_Health{} }
func (m *GameServer_Spec_Health) String() string { return proto.CompactTextString(m) }
func (*GameServer_Spec_Health) ProtoMessage()    {}
func (*GameServer_Spec_Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{3, 1, 0}
}
func (m *GameServer_Spec_Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Spec_Health.Unmarshal(m, b)
}
func (m *GameServer_Spec_Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Spec_Health.Marshal(b, m, deterministic)
}
func (dst *GameServer_Spec_Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Spec_Health.Merge(dst, src)
}
func (m *GameServer_Spec_Health) XXX_Size() int {
	return xxx_messageInfo_GameServer_Spec_Health.Size(m)
}
func (m *GameServer_Spec_Health) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Spec_Health.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Spec_Health proto.InternalMessageInfo

func (m *GameServer_Spec_Health) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *GameServer_Spec_Health) GetPeriodSeconds() int32 {
	if m != nil {
		return m.PeriodSeconds
	}
	return 0
}

func (m *GameServer_Spec_Health) GetFailureThreshold() int32 {
	if m != nil {
		return m.FailureThreshold
	}
	return 0
}

func (m *GameServer_Spec_Health) GetInitialDelaySeconds() int32 {
	if m != nil {
		return m.InitialDelaySeconds
	}
	return 0
}

type GameServer_Status struct {
	State                string                    `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Address              string                    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Ports                []*GameServer_Status_Port `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GameServer_Status) Reset()         { *m = GameServer_Status{} }
func (m *GameServer_Status) String() string { return proto.CompactTextString(m) }
func (*GameServer_Status) ProtoMessage()    {}
func (*GameServer_Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{3, 2}
}
func (m *GameServer_Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Status.Unmarshal(m, b)
}
func (m *GameServer_Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Status.Marshal(b, m, deterministic)
}
func (dst *GameServer_Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Status.Merge(dst, src)
}
func (m *GameServer_Status) XXX_Size() int {
	return xxx_messageInfo_GameServer_Status.Size(m)
}
func (m *GameServer_Status) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Status.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Status proto.InternalMessageInfo

func (m *GameServer_Status) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *GameServer_Status) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GameServer_Status) GetPorts() []*GameServer_Status_Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type GameServer_Status_Port struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServer_Status_Port) Reset()         { *m = GameServer_Status_Port{} }
func (m *GameServer_Status_Port) String() string { return proto.CompactTextString(m) }
func (*GameServer_Status_Port) ProtoMessage()    {}
func (*GameServer_Status_Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_sdk_a0e878ab4087e6bc, []int{3, 2, 0}
}
func (m *GameServer_Status_Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer_Status_Port.Unmarshal(m, b)
}
func (m *GameServer_Status_Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer_Status_Port.Marshal(b, m, deterministic)
}
func (dst *GameServer_Status_Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer_Status_Port.Merge(dst, src)
}
func (m *GameServer_Status_Port) XXX_Size() int {
	return xxx_messageInfo_GameServer_Status_Port.Size(m)
}
func (m *GameServer_Status_Port) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer_Status_Port.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer_Status_Port proto.InternalMessageInfo

func (m *GameServer_Status_Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameServer_Status_Port) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "agones.dev.sdk.Empty")
	proto.RegisterType((*KeyValue)(nil), "agones.dev.sdk.KeyValue")
	proto.RegisterType((*Duration)(nil), "agones.dev.sdk.Duration")
	proto.RegisterType((*GameServer)(nil), "agones.dev.sdk.GameServer")
	proto.RegisterType((*GameServer_ObjectMeta)(nil), "agones.dev.sdk.GameServer.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "agones.dev.sdk.GameServer.ObjectMeta.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "agones.dev.sdk.GameServer.ObjectMeta.LabelsEntry")
	proto.RegisterType((*GameServer_Spec)(nil), "agones.dev.sdk.GameServer.Spec")
	proto.RegisterType((*GameServer_Spec_Health)(nil), "agones.dev.sdk.GameServer.Spec.Health")
	proto.RegisterType((*GameServer_Status)(nil), "agones.dev.sdk.GameServer.Status")
	proto.RegisterType((*GameServer_Status_Port)(nil), "agones.dev.sdk.GameServer.Status.Port")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SDKClient is the client API for SDK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SDKClient interface {
	// Call when the GameServer is ready
	Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Call to self Allocation the GameServer
	Allocate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Call when the GameServer is shutting down
	Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Send a Empty every d Duration to declare that this GameSever is healthy
	Health(ctx context.Context, opts ...grpc.CallOption) (SDK_HealthClient, error)
	// Retrieve the current GameServer data
	GetGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameServer, error)
	// Send GameServer details whenever the GameServer is updated
	WatchGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SDK_WatchGameServerClient, error)
	// Apply a Label to the backing GameServer metadata
	SetLabel(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error)
	// Apply a Annotation to the backing GameServer metadata
	SetAnnotation(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error)
	// Marks the GameServer as the Reserved state for Duration
	Reserve(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*Empty, error)
}

type sDKClient struct {
	cc *grpc.ClientConn
}

func NewSDKClient(cc *grpc.ClientConn) SDKClient {
	return &sDKClient{cc}
}

func (c *sDKClient) Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Allocate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Health(ctx context.Context, opts ...grpc.CallOption) (SDK_HealthClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SDK_serviceDesc.Streams[0], "/agones.dev.sdk.SDK/Health", opts...)
	if err != nil {
		return nil, err
	}
	x := &sDKHealthClient{stream}
	return x, nil
}

type SDK_HealthClient interface {
	Send(*Empty) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type sDKHealthClient struct {
	grpc.ClientStream
}

func (x *sDKHealthClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sDKHealthClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sDKClient) GetGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameServer, error) {
	out := new(GameServer)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/GetGameServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) WatchGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SDK_WatchGameServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SDK_serviceDesc.Streams[1], "/agones.dev.sdk.SDK/WatchGameServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &sDKWatchGameServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SDK_WatchGameServerClient interface {
	Recv() (*GameServer, error)
	grpc.ClientStream
}

type sDKWatchGameServerClient struct {
	grpc.ClientStream
}

func (x *sDKWatchGameServerClient) Recv() (*GameServer, error) {
	m := new(GameServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sDKClient) SetLabel(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/SetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) SetAnnotation(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/SetAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Reserve(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDKServer is the server API for SDK service.
type SDKServer interface {
	// Call when the GameServer is ready
	Ready(context.Context, *Empty) (*Empty, error)
	// Call to self Allocation the GameServer
	Allocate(context.Context, *Empty) (*Empty, error)
	// Call when the GameServer is shutting down
	Shutdown(context.Context, *Empty) (*Empty, error)
	// Send a Empty every d Duration to declare that this GameSever is healthy
	Health(SDK_HealthServer) error
	// Retrieve the current GameServer data
	GetGameServer(context.Context, *Empty) (*GameServer, error)
	// Send GameServer details whenever the GameServer is updated
	WatchGameServer(*Empty, SDK_WatchGameServerServer) error
	// Apply a Label to the backing GameServer metadata
	SetLabel(context.Context, *KeyValue) (*Empty, error)
	// Apply a Annotation to the backing GameServer metadata
	SetAnnotation(context.Context, *KeyValue) (*Empty, error)
	// Marks the GameServer as the Reserved state for Duration
	Reserve(context.Context, *Duration) (*Empty, error)
}

func RegisterSDKServer(s *grpc.Server, srv SDKServer) {
	s.RegisterService(&_SDK_serviceDesc, srv)
}

func _SDK_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Ready(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Allocate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Shutdown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Health_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SDKServer).Health(&sDKHealthServer{stream})
}

type SDK_HealthServer interface {
	SendAndClose(*Empty) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type sDKHealthServer struct {
	grpc.ServerStream
}

func (x *sDKHealthServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sDKHealthServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SDK_GetGameServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).GetGameServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/GetGameServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).GetGameServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_WatchGameServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SDKServer).WatchGameServer(m, &sDKWatchGameServerServer{stream})
}

type SDK_WatchGameServerServer interface {
	Send(*GameServer) error
	grpc.ServerStream
}

type sDKWatchGameServerServer struct {
	grpc.ServerStream
}

func (x *sDKWatchGameServerServer) Send(m *GameServer) error {
	return x.ServerStream.SendMsg(m)
}

func _SDK_SetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).SetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/SetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).SetLabel(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_SetAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).SetAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/SetAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).SetAnnotation(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Duration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Reserve(ctx, req.(*Duration))
	}
	return interceptor(ctx, in, info, handler)
}

var _SDK_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agones.dev.sdk.SDK",
	HandlerType: (*SDKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ready",
			Handler:    _SDK_Ready_Handler,
		},
		{
			MethodName: "Allocate",
			Handler:    _SDK_Allocate_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _SDK_Shutdown_Handler,
		},
		{
			MethodName: "GetGameServer",
			Handler:    _SDK_GetGameServer_Handler,
		},
		{
			MethodName: "SetLabel",
			Handler:    _SDK_SetLabel_Handler,
		},
		{
			MethodName: "SetAnnotation",
			Handler:    _SDK_SetAnnotation_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _SDK_Reserve_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Health",
			Handler:       _SDK_Health_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "WatchGameServer",
			Handler:       _SDK_WatchGameServer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk.proto",
}

func init() { proto.RegisterFile("sdk.proto", fileDescriptor_sdk_a0e878ab4087e6bc) }

var fileDescriptor_sdk_a0e878ab4087e6bc = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0xc0, 0xe5, 0x4b, 0xe2, 0x24, 0x13, 0xee, 0xdf, 0xde, 0x55, 0x72, 0xad, 0x8a, 0x16, 0x8b,
	0xa2, 0xe3, 0x10, 0x36, 0xa4, 0x12, 0xa2, 0x27, 0x54, 0xa9, 0xe8, 0x4a, 0x8b, 0x5a, 0x28, 0x72,
	0xaa, 0x16, 0xf1, 0x12, 0xed, 0x79, 0x87, 0xc4, 0x9c, 0xe3, 0xb5, 0x76, 0x37, 0x57, 0xe5, 0x11,
	0x1e, 0xf8, 0x02, 0x3c, 0xf1, 0x05, 0x78, 0xe2, 0xdb, 0xf0, 0x15, 0xf8, 0x0a, 0xbc, 0xa3, 0x9d,
	0xb5, 0x2f, 0xe1, 0xe8, 0xb5, 0x17, 0x78, 0xca, 0xce, 0xbf, 0xdf, 0x6c, 0xc6, 0x33, 0xb3, 0xd0,
	0xd7, 0xe2, 0x34, 0xae, 0x94, 0x34, 0x92, 0x6d, 0xf1, 0x89, 0x2c, 0x51, 0xc7, 0x02, 0xcf, 0x62,
	0x2d, 0x4e, 0xc3, 0x1b, 0x13, 0x29, 0x27, 0x05, 0x26, 0xbc, 0xca, 0x13, 0x5e, 0x96, 0xd2, 0x70,
	0x93, 0xcb, 0x52, 0x3b, 0xef, 0xa8, 0x0b, 0x9d, 0x07, 0xb3, 0xca, 0x2c, 0xa2, 0x21, 0xf4, 0x1e,
	0xe3, 0xe2, 0x39, 0x2f, 0xe6, 0xc8, 0x76, 0xa0, 0x75, 0x8a, 0x8b, 0xc0, 0xbb, 0xe5, 0x1d, 0xf4,
	0x53, 0x7b, 0x64, 0xfb, 0xd0, 0x39, 0xb3, 0xa6, 0x60, 0x83, 0x74, 0x4e, 0x88, 0xde, 0x85, 0xde,
	0xf1, 0x5c, 0x11, 0x8f, 0x05, 0xd0, 0xd5, 0x98, 0xc9, 0x52, 0x68, 0x8a, 0x6b, 0xa5, 0x8d, 0x18,
	0xfd, 0xd8, 0x07, 0x78, 0xc8, 0x67, 0x38, 0x42, 0x75, 0x86, 0x8a, 0x7d, 0x01, 0x03, 0x79, 0xf2,
	0x03, 0x66, 0x66, 0x3c, 0x43, 0xc3, 0xc9, 0x79, 0x30, 0xbc, 0x1d, 0xff, 0xf3, 0xd6, 0xf1, 0x32,
	0x20, 0x7e, 0x4a, 0xde, 0x5f, 0xa1, 0xe1, 0x29, 0xc8, 0xf3, 0x33, 0xbb, 0x03, 0x6d, 0x5d, 0x61,
	0x46, 0x37, 0x1a, 0x0c, 0x6f, 0xbe, 0x06, 0x30, 0xaa, 0x30, 0x4b, 0xc9, 0x99, 0xdd, 0x05, 0x5f,
	0x1b, 0x6e, 0xe6, 0x3a, 0x68, 0x51, 0xd8, 0x3b, 0xaf, 0x0b, 0x23, 0xc7, 0xb4, 0x0e, 0x08, 0x7f,
	0x6d, 0x03, 0x2c, 0xaf, 0xc2, 0x18, 0xb4, 0x4b, 0x3e, 0xc3, 0xba, 0x48, 0x74, 0x66, 0x37, 0xa0,
	0x6f, 0x7f, 0x75, 0xc5, 0xb3, 0xa6, 0x52, 0x4b, 0x85, 0xad, 0xea, 0x3c, 0x17, 0x94, 0xb8, 0x9f,
	0xda, 0x23, 0x7b, 0x1f, 0x76, 0x14, 0x6a, 0x39, 0x57, 0x19, 0x8e, 0xcf, 0x50, 0xe9, 0x5c, 0x96,
	0x41, 0x9b, 0xcc, 0xdb, 0x8d, 0xfe, 0xb9, 0x53, 0xb3, 0xb7, 0x01, 0x26, 0x58, 0xa2, 0x2b, 0x76,
	0xd0, 0xa1, 0x0a, 0xaf, 0x68, 0xd8, 0x87, 0xc0, 0x32, 0x85, 0x74, 0x1e, 0x9b, 0x7c, 0x86, 0xda,
	0xf0, 0x59, 0x15, 0xf8, 0xe4, 0xb7, 0xdb, 0x58, 0x9e, 0x35, 0x06, 0xeb, 0x2e, 0xb0, 0xc0, 0x0b,
	0xee, 0x5d, 0xe7, 0xde, 0x58, 0x96, 0xee, 0xdf, 0xc2, 0x60, 0xa5, 0x75, 0x82, 0xde, 0xad, 0xd6,
	0xc1, 0x60, 0xf8, 0xc9, 0x95, 0xbe, 0x59, 0x7c, 0x7f, 0x19, 0xf8, 0xa0, 0x34, 0x6a, 0x91, 0xae,
	0xa2, 0xd8, 0x97, 0xe0, 0x17, 0xfc, 0x04, 0x0b, 0x1d, 0xf4, 0x09, 0xfa, 0xf1, 0xd5, 0xa0, 0x4f,
	0x28, 0xc6, 0xf1, 0x6a, 0x40, 0x78, 0x0f, 0x76, 0x2e, 0xe6, 0xba, 0x6a, 0x27, 0x1f, 0x6d, 0x7c,
	0xea, 0x85, 0x77, 0x61, 0xb0, 0x82, 0x5d, 0x2b, 0xf4, 0x2f, 0x0f, 0xda, 0xb6, 0xcb, 0xd8, 0x3d,
	0xf0, 0xa7, 0xc8, 0x0b, 0x33, 0xad, 0xfb, 0xfa, 0xbd, 0x37, 0xb4, 0x65, 0xfc, 0x88, 0xbc, 0xd3,
	0x3a, 0x2a, 0xfc, 0xdd, 0x03, 0xdf, 0xa9, 0x58, 0x08, 0x3d, 0x91, 0x6b, 0x7e, 0x52, 0xa0, 0x20,
	0x58, 0x2f, 0x3d, 0x97, 0xd9, 0x6d, 0xd8, 0xaa, 0x50, 0xe5, 0x52, 0x8c, 0x9b, 0x99, 0xb3, 0x57,
	0xea, 0xa4, 0x9b, 0x4e, 0x3b, 0x72, 0x4a, 0xf6, 0x01, 0xec, 0x7e, 0xcf, 0xf3, 0x62, 0xae, 0x70,
	0x6c, 0xa6, 0x0a, 0xf5, 0x54, 0x16, 0xae, 0xff, 0x3a, 0xe9, 0x4e, 0x6d, 0x78, 0xd6, 0xe8, 0xd9,
	0x10, 0xae, 0xe5, 0x65, 0x6e, 0x72, 0x5e, 0x8c, 0x05, 0x16, 0x7c, 0x71, 0x8e, 0x6e, 0x53, 0xc0,
	0x5e, 0x6d, 0x3c, 0xb6, 0xb6, 0x3a, 0x41, 0xf8, 0x9b, 0x07, 0xbe, 0x1b, 0x13, 0x5b, 0x1c, 0x3b,
	0x28, 0xcd, 0x40, 0x38, 0xc1, 0x6e, 0x05, 0x2e, 0x84, 0x42, 0xad, 0xeb, 0xa2, 0x35, 0x22, 0xfb,
	0x0c, 0x3a, 0x95, 0x54, 0xc6, 0x0e, 0x62, 0xeb, 0x4d, 0x85, 0xa2, 0x0c, 0xf1, 0x37, 0x52, 0x99,
	0xd4, 0x05, 0x85, 0x31, 0xb4, 0xad, 0xf8, 0xca, 0x29, 0x64, 0xd0, 0xb6, 0x4e, 0x75, 0x49, 0xe8,
	0x3c, 0xfc, 0xd9, 0x87, 0xd6, 0xe8, 0xf8, 0x31, 0x7b, 0x04, 0x9d, 0x14, 0xb9, 0x58, 0xb0, 0x6b,
	0x17, 0xf3, 0xd1, 0x16, 0x0c, 0x5f, 0xad, 0x8e, 0x76, 0x7f, 0xfa, 0xe3, 0xcf, 0x5f, 0x36, 0x06,
	0x91, 0x9f, 0x28, 0x1b, 0x7d, 0xe4, 0x1d, 0xb2, 0xaf, 0xa1, 0x77, 0xbf, 0x28, 0x64, 0x66, 0xff,
	0xe5, 0x7a, 0xb0, 0x7d, 0x82, 0x6d, 0x45, 0xfd, 0x84, 0xd7, 0x80, 0x9a, 0x37, 0x9a, 0xce, 0x8d,
	0x90, 0x2f, 0xcb, 0xff, 0xcc, 0xd3, 0x35, 0xc0, 0xf2, 0x9e, 0x9c, 0x37, 0xd2, 0x7a, 0x34, 0x46,
	0xb4, 0xb7, 0xa2, 0x6e, 0xe2, 0x5a, 0xf2, 0xc8, 0x3b, 0x3c, 0xf0, 0xd8, 0x0b, 0xd8, 0x7c, 0x88,
	0x66, 0x65, 0x8b, 0x5f, 0x02, 0x0d, 0x2f, 0xff, 0x8c, 0xd1, 0x1e, 0x91, 0x37, 0xd9, 0x20, 0x99,
	0xd8, 0x9d, 0xe8, 0x38, 0x1c, 0xb6, 0x5f, 0x70, 0x93, 0x4d, 0xff, 0x1f, 0xfa, 0x3a, 0xa1, 0xf7,
	0xd8, 0x6e, 0xf2, 0xd2, 0xc2, 0x56, 0x12, 0x7c, 0x64, 0xef, 0xde, 0x1b, 0xa1, 0xa1, 0xd1, 0x66,
	0xc1, 0x45, 0x48, 0xf3, 0xe6, 0x5d, 0x56, 0x8e, 0x90, 0xc8, 0xfb, 0xe1, 0x76, 0x62, 0x5f, 0x2b,
	0xc1, 0x0d, 0x4f, 0x68, 0xdd, 0xd8, 0x12, 0x73, 0xd8, 0x1c, 0xa1, 0x59, 0xee, 0x9c, 0xf5, 0xe9,
	0x37, 0x89, 0x7e, 0x3d, 0xdc, 0x5f, 0xd2, 0x97, 0xcb, 0xd1, 0xa6, 0x78, 0x0a, 0xdd, 0xd4, 0xfd,
	0x93, 0x7f, 0xc3, 0x9b, 0xa7, 0xf7, 0x32, 0x78, 0x5d, 0xef, 0xa8, 0x97, 0x28, 0x87, 0x38, 0xf2,
	0x0e, 0x3f, 0xef, 0x7c, 0xd7, 0xd2, 0xe2, 0xf4, 0xc4, 0xa7, 0xd7, 0xff, 0xce, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x42, 0xa7, 0x46, 0xe0, 0x38, 0x08, 0x00, 0x00,
}
