// Copyright Istio Authors. All Rights Reserved.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.6
// source: pkg/workloadmetadata/proto/workload_metadata.proto

package istio_telemetry_workloadmetadata_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkloadMetadataResource_WorkloadType int32

const (
	WorkloadMetadataResource_KUBERNETES_DEPLOYMENT WorkloadMetadataResource_WorkloadType = 0
	WorkloadMetadataResource_KUBERNETES_CRONJOB    WorkloadMetadataResource_WorkloadType = 1
	WorkloadMetadataResource_KUBERNETES_POD        WorkloadMetadataResource_WorkloadType = 2
	WorkloadMetadataResource_KUBERNETES_JOB        WorkloadMetadataResource_WorkloadType = 3
)

// Enum value maps for WorkloadMetadataResource_WorkloadType.
var (
	WorkloadMetadataResource_WorkloadType_name = map[int32]string{
		0: "KUBERNETES_DEPLOYMENT",
		1: "KUBERNETES_CRONJOB",
		2: "KUBERNETES_POD",
		3: "KUBERNETES_JOB",
	}
	WorkloadMetadataResource_WorkloadType_value = map[string]int32{
		"KUBERNETES_DEPLOYMENT": 0,
		"KUBERNETES_CRONJOB":    1,
		"KUBERNETES_POD":        2,
		"KUBERNETES_JOB":        3,
	}
)

func (x WorkloadMetadataResource_WorkloadType) Enum() *WorkloadMetadataResource_WorkloadType {
	p := new(WorkloadMetadataResource_WorkloadType)
	*p = x
	return p
}

func (x WorkloadMetadataResource_WorkloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkloadMetadataResource_WorkloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_workloadmetadata_proto_workload_metadata_proto_enumTypes[0].Descriptor()
}

func (WorkloadMetadataResource_WorkloadType) Type() protoreflect.EnumType {
	return &file_pkg_workloadmetadata_proto_workload_metadata_proto_enumTypes[0]
}

func (x WorkloadMetadataResource_WorkloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkloadMetadataResource_WorkloadType.Descriptor instead.
func (WorkloadMetadataResource_WorkloadType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescGZIP(), []int{1, 0}
}

// WorkloadMetadataResources contains a list of workloads and their metadata for
// an Ambient L4 proxy (aka per-Kubernetes-node proxy). This will be sent in
// an xDS response message via ECDS.
type WorkloadMetadataResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provides the xDS node identifier for the ASM proxy that corresponds to the
	// `workload_resources` returned in a response.
	ProxyId string `protobuf:"bytes,1,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	// Contains a set of workload metadata for all workload instances that
	// are currently being proxied by the xDS node.
	WorkloadMetadataResources []*WorkloadMetadataResource `protobuf:"bytes,2,rep,name=workload_metadata_resources,json=workloadMetadataResources,proto3" json:"workload_metadata_resources,omitempty"`
}

func (x *WorkloadMetadataResources) Reset() {
	*x = WorkloadMetadataResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadMetadataResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadMetadataResources) ProtoMessage() {}

func (x *WorkloadMetadataResources) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadMetadataResources.ProtoReflect.Descriptor instead.
func (*WorkloadMetadataResources) Descriptor() ([]byte, []int) {
	return file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *WorkloadMetadataResources) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *WorkloadMetadataResources) GetWorkloadMetadataResources() []*WorkloadMetadataResource {
	if x != nil {
		return x.WorkloadMetadataResources
	}
	return nil
}

// Contains the metadata for a single workload instance.
type WorkloadMetadataResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set of IP addresses associated with an individual workload instance.
	IpAddresses []string `protobuf:"bytes,1,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	// The full name of the workload instance.
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	// The Kubernetes namespace to which the workload instance belongs.
	NamespaceName string `protobuf:"bytes,3,opt,name=namespace_name,json=namespaceName,proto3" json:"namespace_name,omitempty"`
	// The set of containers (if known) that constitute the workload instance.
	Containers []string `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	// The name of the workload provided by the instance. This is typically the
	// Kubernetes deployment name.
	WorkloadName string `protobuf:"bytes,5,opt,name=workload_name,json=workloadName,proto3" json:"workload_name,omitempty"`
	// Type of workload
	WorkloadType WorkloadMetadataResource_WorkloadType `protobuf:"varint,6,opt,name=workload_type,json=workloadType,proto3,enum=istio.telemetry.workloadmetadata.v1.WorkloadMetadataResource_WorkloadType" json:"workload_type,omitempty"`
	// Canonical name of the workload
	CanonicalName string `protobuf:"bytes,7,opt,name=canonical_name,json=canonicalName,proto3" json:"canonical_name,omitempty"`
	// Canonical revision of the workload
	CanonicalRevision string `protobuf:"bytes,8,opt,name=canonical_revision,json=canonicalRevision,proto3" json:"canonical_revision,omitempty"`
}

func (x *WorkloadMetadataResource) Reset() {
	*x = WorkloadMetadataResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadMetadataResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadMetadataResource) ProtoMessage() {}

func (x *WorkloadMetadataResource) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadMetadataResource.ProtoReflect.Descriptor instead.
func (*WorkloadMetadataResource) Descriptor() ([]byte, []int) {
	return file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadMetadataResource) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *WorkloadMetadataResource) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *WorkloadMetadataResource) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

func (x *WorkloadMetadataResource) GetContainers() []string {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *WorkloadMetadataResource) GetWorkloadName() string {
	if x != nil {
		return x.WorkloadName
	}
	return ""
}

func (x *WorkloadMetadataResource) GetWorkloadType() WorkloadMetadataResource_WorkloadType {
	if x != nil {
		return x.WorkloadType
	}
	return WorkloadMetadataResource_KUBERNETES_DEPLOYMENT
}

func (x *WorkloadMetadataResource) GetCanonicalName() string {
	if x != nil {
		return x.CanonicalName
	}
	return ""
}

func (x *WorkloadMetadataResource) GetCanonicalRevision() string {
	if x != nil {
		return x.CanonicalRevision
	}
	return ""
}

var File_pkg_workloadmetadata_proto_workload_metadata_proto protoreflect.FileDescriptor

var file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x22, 0xb5, 0x01, 0x0a, 0x19, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x49, 0x64, 0x12, 0x7d, 0x0a, 0x1b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x22, 0x80, 0x04, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x6f, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e,
	0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x61,
	0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61,
	0x6c, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4b, 0x55, 0x42,
	0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54,
	0x45, 0x53, 0x5f, 0x43, 0x52, 0x4f, 0x4e, 0x4a, 0x4f, 0x42, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x50, 0x4f, 0x44, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x4a,
	0x4f, 0x42, 0x10, 0x03, 0x42, 0x4f, 0x5a, 0x4d, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f,
	0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescOnce sync.Once
	file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescData = file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDesc
)

func file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescGZIP() []byte {
	file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescOnce.Do(func() {
		file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescData)
	})
	return file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDescData
}

var file_pkg_workloadmetadata_proto_workload_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_workloadmetadata_proto_workload_metadata_proto_goTypes = []interface{}{
	(WorkloadMetadataResource_WorkloadType)(0), // 0: istio.telemetry.workloadmetadata.v1.WorkloadMetadataResource.WorkloadType
	(*WorkloadMetadataResources)(nil),          // 1: istio.telemetry.workloadmetadata.v1.WorkloadMetadataResources
	(*WorkloadMetadataResource)(nil),           // 2: istio.telemetry.workloadmetadata.v1.WorkloadMetadataResource
}
var file_pkg_workloadmetadata_proto_workload_metadata_proto_depIdxs = []int32{
	2, // 0: istio.telemetry.workloadmetadata.v1.WorkloadMetadataResources.workload_metadata_resources:type_name -> istio.telemetry.workloadmetadata.v1.WorkloadMetadataResource
	0, // 1: istio.telemetry.workloadmetadata.v1.WorkloadMetadataResource.workload_type:type_name -> istio.telemetry.workloadmetadata.v1.WorkloadMetadataResource.WorkloadType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_workloadmetadata_proto_workload_metadata_proto_init() }
func file_pkg_workloadmetadata_proto_workload_metadata_proto_init() {
	if File_pkg_workloadmetadata_proto_workload_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadMetadataResources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadMetadataResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_workloadmetadata_proto_workload_metadata_proto_goTypes,
		DependencyIndexes: file_pkg_workloadmetadata_proto_workload_metadata_proto_depIdxs,
		EnumInfos:         file_pkg_workloadmetadata_proto_workload_metadata_proto_enumTypes,
		MessageInfos:      file_pkg_workloadmetadata_proto_workload_metadata_proto_msgTypes,
	}.Build()
	File_pkg_workloadmetadata_proto_workload_metadata_proto = out.File
	file_pkg_workloadmetadata_proto_workload_metadata_proto_rawDesc = nil
	file_pkg_workloadmetadata_proto_workload_metadata_proto_goTypes = nil
	file_pkg_workloadmetadata_proto_workload_metadata_proto_depIdxs = nil
}