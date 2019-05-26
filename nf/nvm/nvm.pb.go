// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nvm.proto

package nvm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NVMConfigBundle struct {
	ScriptSrc                       string   `protobuf:"bytes,1,opt,name=script_src,json=scriptSrc,proto3" json:"script_src,omitempty"`
	ScriptType                      string   `protobuf:"bytes,2,opt,name=script_type,json=scriptType,proto3" json:"script_type,omitempty"`
	ScriptHash                      string   `protobuf:"bytes,3,opt,name=script_hash,json=scriptHash,proto3" json:"script_hash,omitempty"`
	RunnableSrc                     string   `protobuf:"bytes,4,opt,name=runnable_src,json=runnableSrc,proto3" json:"runnable_src,omitempty"`
	BlockJson                       string   `protobuf:"bytes,5,opt,name=block_json,json=blockJson,proto3" json:"block_json,omitempty"`
	TxJson                          string   `protobuf:"bytes,6,opt,name=tx_json,json=txJson,proto3" json:"tx_json,omitempty"`
	ModuleId                        string   `protobuf:"bytes,7,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	BlockHeight                     uint64   `protobuf:"varint,8,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	ChainId                         uint32   `protobuf:"varint,9,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	MetaVersion                     string   `protobuf:"bytes,10,opt,name=meta_version,json=metaVersion,proto3" json:"meta_version,omitempty"`
	EnableLimits                    bool     `protobuf:"varint,11,opt,name=enable_limits,json=enableLimits,proto3" json:"enable_limits,omitempty"`
	LimitsExeInstruction            uint64   `protobuf:"varint,12,opt,name=limits_exe_instruction,json=limitsExeInstruction,proto3" json:"limits_exe_instruction,omitempty"`
	LimitsTotalMemSize              uint64   `protobuf:"varint,13,opt,name=limits_total_mem_size,json=limitsTotalMemSize,proto3" json:"limits_total_mem_size,omitempty"`
	ExecutionTimeout                uint64   `protobuf:"varint,14,opt,name=execution_timeout,json=executionTimeout,proto3" json:"execution_timeout,omitempty"`
	TimeoutGasLimitCost             uint64   `protobuf:"varint,15,opt,name=timeout_gas_limit_cost,json=timeoutGasLimitCost,proto3" json:"timeout_gas_limit_cost,omitempty"`
	MaxLimitsOfExecutionInstruction uint64   `protobuf:"varint,16,opt,name=max_limits_of_execution_instruction,json=maxLimitsOfExecutionInstruction,proto3" json:"max_limits_of_execution_instruction,omitempty"`
	DefaultLimitsOfTotalMemSize     uint64   `protobuf:"varint,17,opt,name=default_limits_of_total_mem_size,json=defaultLimitsOfTotalMemSize,proto3" json:"default_limits_of_total_mem_size,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *NVMConfigBundle) Reset()         { *m = NVMConfigBundle{} }
func (m *NVMConfigBundle) String() string { return proto.CompactTextString(m) }
func (*NVMConfigBundle) ProtoMessage()    {}
func (*NVMConfigBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{0}
}

func (m *NVMConfigBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMConfigBundle.Unmarshal(m, b)
}
func (m *NVMConfigBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMConfigBundle.Marshal(b, m, deterministic)
}
func (m *NVMConfigBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMConfigBundle.Merge(m, src)
}
func (m *NVMConfigBundle) XXX_Size() int {
	return xxx_messageInfo_NVMConfigBundle.Size(m)
}
func (m *NVMConfigBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMConfigBundle.DiscardUnknown(m)
}

var xxx_messageInfo_NVMConfigBundle proto.InternalMessageInfo

func (m *NVMConfigBundle) GetScriptSrc() string {
	if m != nil {
		return m.ScriptSrc
	}
	return ""
}

func (m *NVMConfigBundle) GetScriptType() string {
	if m != nil {
		return m.ScriptType
	}
	return ""
}

func (m *NVMConfigBundle) GetScriptHash() string {
	if m != nil {
		return m.ScriptHash
	}
	return ""
}

func (m *NVMConfigBundle) GetRunnableSrc() string {
	if m != nil {
		return m.RunnableSrc
	}
	return ""
}

func (m *NVMConfigBundle) GetBlockJson() string {
	if m != nil {
		return m.BlockJson
	}
	return ""
}

func (m *NVMConfigBundle) GetTxJson() string {
	if m != nil {
		return m.TxJson
	}
	return ""
}

func (m *NVMConfigBundle) GetModuleId() string {
	if m != nil {
		return m.ModuleId
	}
	return ""
}

func (m *NVMConfigBundle) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *NVMConfigBundle) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *NVMConfigBundle) GetMetaVersion() string {
	if m != nil {
		return m.MetaVersion
	}
	return ""
}

func (m *NVMConfigBundle) GetEnableLimits() bool {
	if m != nil {
		return m.EnableLimits
	}
	return false
}

func (m *NVMConfigBundle) GetLimitsExeInstruction() uint64 {
	if m != nil {
		return m.LimitsExeInstruction
	}
	return 0
}

func (m *NVMConfigBundle) GetLimitsTotalMemSize() uint64 {
	if m != nil {
		return m.LimitsTotalMemSize
	}
	return 0
}

func (m *NVMConfigBundle) GetExecutionTimeout() uint64 {
	if m != nil {
		return m.ExecutionTimeout
	}
	return 0
}

func (m *NVMConfigBundle) GetTimeoutGasLimitCost() uint64 {
	if m != nil {
		return m.TimeoutGasLimitCost
	}
	return 0
}

func (m *NVMConfigBundle) GetMaxLimitsOfExecutionInstruction() uint64 {
	if m != nil {
		return m.MaxLimitsOfExecutionInstruction
	}
	return 0
}

func (m *NVMConfigBundle) GetDefaultLimitsOfTotalMemSize() uint64 {
	if m != nil {
		return m.DefaultLimitsOfTotalMemSize
	}
	return 0
}

type NVMCallbackResult struct {
	FuncName             string   `protobuf:"bytes,1,opt,name=func_name,json=funcName,proto3" json:"func_name,omitempty"`
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	NotNull              bool     `protobuf:"varint,3,opt,name=not_null,json=notNull,proto3" json:"not_null,omitempty"`
	Extra                []string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NVMCallbackResult) Reset()         { *m = NVMCallbackResult{} }
func (m *NVMCallbackResult) String() string { return proto.CompactTextString(m) }
func (*NVMCallbackResult) ProtoMessage()    {}
func (*NVMCallbackResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{1}
}

func (m *NVMCallbackResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMCallbackResult.Unmarshal(m, b)
}
func (m *NVMCallbackResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMCallbackResult.Marshal(b, m, deterministic)
}
func (m *NVMCallbackResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMCallbackResult.Merge(m, src)
}
func (m *NVMCallbackResult) XXX_Size() int {
	return xxx_messageInfo_NVMCallbackResult.Size(m)
}
func (m *NVMCallbackResult) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMCallbackResult.DiscardUnknown(m)
}

var xxx_messageInfo_NVMCallbackResult proto.InternalMessageInfo

func (m *NVMCallbackResult) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *NVMCallbackResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *NVMCallbackResult) GetNotNull() bool {
	if m != nil {
		return m.NotNull
	}
	return false
}

func (m *NVMCallbackResult) GetExtra() []string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type NVMDataRequest struct {
	RequestType          string             `protobuf:"bytes,1,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	RequestIndx          uint32             `protobuf:"varint,2,opt,name=request_indx,json=requestIndx,proto3" json:"request_indx,omitempty"`
	LcsHandler           uint64             `protobuf:"varint,3,opt,name=lcs_handler,json=lcsHandler,proto3" json:"lcs_handler,omitempty"`
	GcsHandler           uint64             `protobuf:"varint,4,opt,name=gcs_handler,json=gcsHandler,proto3" json:"gcs_handler,omitempty"`
	CallbackResult       *NVMCallbackResult `protobuf:"bytes,5,opt,name=callback_result,json=callbackResult,proto3" json:"callback_result,omitempty"`
	ConfigBundle         *NVMConfigBundle   `protobuf:"bytes,6,opt,name=config_bundle,json=configBundle,proto3" json:"config_bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NVMDataRequest) Reset()         { *m = NVMDataRequest{} }
func (m *NVMDataRequest) String() string { return proto.CompactTextString(m) }
func (*NVMDataRequest) ProtoMessage()    {}
func (*NVMDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{2}
}

func (m *NVMDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMDataRequest.Unmarshal(m, b)
}
func (m *NVMDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMDataRequest.Marshal(b, m, deterministic)
}
func (m *NVMDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMDataRequest.Merge(m, src)
}
func (m *NVMDataRequest) XXX_Size() int {
	return xxx_messageInfo_NVMDataRequest.Size(m)
}
func (m *NVMDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NVMDataRequest proto.InternalMessageInfo

func (m *NVMDataRequest) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *NVMDataRequest) GetRequestIndx() uint32 {
	if m != nil {
		return m.RequestIndx
	}
	return 0
}

func (m *NVMDataRequest) GetLcsHandler() uint64 {
	if m != nil {
		return m.LcsHandler
	}
	return 0
}

func (m *NVMDataRequest) GetGcsHandler() uint64 {
	if m != nil {
		return m.GcsHandler
	}
	return 0
}

func (m *NVMDataRequest) GetCallbackResult() *NVMCallbackResult {
	if m != nil {
		return m.CallbackResult
	}
	return nil
}

func (m *NVMDataRequest) GetConfigBundle() *NVMConfigBundle {
	if m != nil {
		return m.ConfigBundle
	}
	return nil
}

// server side response
type NVMStatsBundle struct {
	ActualCountOfExecutionInstruction uint64   `protobuf:"varint,1,opt,name=actual_count_of_execution_instruction,json=actualCountOfExecutionInstruction,proto3" json:"actual_count_of_execution_instruction,omitempty"`
	ActualUsedMemSize                 uint64   `protobuf:"varint,2,opt,name=actual_used_mem_size,json=actualUsedMemSize,proto3" json:"actual_used_mem_size,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *NVMStatsBundle) Reset()         { *m = NVMStatsBundle{} }
func (m *NVMStatsBundle) String() string { return proto.CompactTextString(m) }
func (*NVMStatsBundle) ProtoMessage()    {}
func (*NVMStatsBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{3}
}

func (m *NVMStatsBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMStatsBundle.Unmarshal(m, b)
}
func (m *NVMStatsBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMStatsBundle.Marshal(b, m, deterministic)
}
func (m *NVMStatsBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMStatsBundle.Merge(m, src)
}
func (m *NVMStatsBundle) XXX_Size() int {
	return xxx_messageInfo_NVMStatsBundle.Size(m)
}
func (m *NVMStatsBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMStatsBundle.DiscardUnknown(m)
}

var xxx_messageInfo_NVMStatsBundle proto.InternalMessageInfo

func (m *NVMStatsBundle) GetActualCountOfExecutionInstruction() uint64 {
	if m != nil {
		return m.ActualCountOfExecutionInstruction
	}
	return 0
}

func (m *NVMStatsBundle) GetActualUsedMemSize() uint64 {
	if m != nil {
		return m.ActualUsedMemSize
	}
	return 0
}

type NVMFinalResponse struct {
	Result               int32           `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg                  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	NotNull              bool            `protobuf:"varint,3,opt,name=not_null,json=notNull,proto3" json:"not_null,omitempty"`
	StatsBundle          *NVMStatsBundle `protobuf:"bytes,4,opt,name=stats_bundle,json=statsBundle,proto3" json:"stats_bundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NVMFinalResponse) Reset()         { *m = NVMFinalResponse{} }
func (m *NVMFinalResponse) String() string { return proto.CompactTextString(m) }
func (*NVMFinalResponse) ProtoMessage()    {}
func (*NVMFinalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{4}
}

func (m *NVMFinalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMFinalResponse.Unmarshal(m, b)
}
func (m *NVMFinalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMFinalResponse.Marshal(b, m, deterministic)
}
func (m *NVMFinalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMFinalResponse.Merge(m, src)
}
func (m *NVMFinalResponse) XXX_Size() int {
	return xxx_messageInfo_NVMFinalResponse.Size(m)
}
func (m *NVMFinalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMFinalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NVMFinalResponse proto.InternalMessageInfo

func (m *NVMFinalResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *NVMFinalResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *NVMFinalResponse) GetNotNull() bool {
	if m != nil {
		return m.NotNull
	}
	return false
}

func (m *NVMFinalResponse) GetStatsBundle() *NVMStatsBundle {
	if m != nil {
		return m.StatsBundle
	}
	return nil
}

type NVMCallbackResponse struct {
	FuncName             string   `protobuf:"bytes,1,opt,name=func_name,json=funcName,proto3" json:"func_name,omitempty"`
	FuncParams           []string `protobuf:"bytes,2,rep,name=func_params,json=funcParams,proto3" json:"func_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NVMCallbackResponse) Reset()         { *m = NVMCallbackResponse{} }
func (m *NVMCallbackResponse) String() string { return proto.CompactTextString(m) }
func (*NVMCallbackResponse) ProtoMessage()    {}
func (*NVMCallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{5}
}

func (m *NVMCallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMCallbackResponse.Unmarshal(m, b)
}
func (m *NVMCallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMCallbackResponse.Marshal(b, m, deterministic)
}
func (m *NVMCallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMCallbackResponse.Merge(m, src)
}
func (m *NVMCallbackResponse) XXX_Size() int {
	return xxx_messageInfo_NVMCallbackResponse.Size(m)
}
func (m *NVMCallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMCallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NVMCallbackResponse proto.InternalMessageInfo

func (m *NVMCallbackResponse) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *NVMCallbackResponse) GetFuncParams() []string {
	if m != nil {
		return m.FuncParams
	}
	return nil
}

type NVMDataResponse struct {
	ResponseType         string               `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	ResponseIndx         uint32               `protobuf:"varint,2,opt,name=response_indx,json=responseIndx,proto3" json:"response_indx,omitempty"`
	LcsHandler           uint64               `protobuf:"varint,3,opt,name=lcs_handler,json=lcsHandler,proto3" json:"lcs_handler,omitempty"`
	GcsHandler           uint64               `protobuf:"varint,4,opt,name=gcs_handler,json=gcsHandler,proto3" json:"gcs_handler,omitempty"`
	FinalResponse        *NVMFinalResponse    `protobuf:"bytes,5,opt,name=final_response,json=finalResponse,proto3" json:"final_response,omitempty"`
	CallbackResponse     *NVMCallbackResponse `protobuf:"bytes,6,opt,name=callback_response,json=callbackResponse,proto3" json:"callback_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NVMDataResponse) Reset()         { *m = NVMDataResponse{} }
func (m *NVMDataResponse) String() string { return proto.CompactTextString(m) }
func (*NVMDataResponse) ProtoMessage()    {}
func (*NVMDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a3db176960b780, []int{6}
}

func (m *NVMDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMDataResponse.Unmarshal(m, b)
}
func (m *NVMDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMDataResponse.Marshal(b, m, deterministic)
}
func (m *NVMDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMDataResponse.Merge(m, src)
}
func (m *NVMDataResponse) XXX_Size() int {
	return xxx_messageInfo_NVMDataResponse.Size(m)
}
func (m *NVMDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NVMDataResponse proto.InternalMessageInfo

func (m *NVMDataResponse) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *NVMDataResponse) GetResponseIndx() uint32 {
	if m != nil {
		return m.ResponseIndx
	}
	return 0
}

func (m *NVMDataResponse) GetLcsHandler() uint64 {
	if m != nil {
		return m.LcsHandler
	}
	return 0
}

func (m *NVMDataResponse) GetGcsHandler() uint64 {
	if m != nil {
		return m.GcsHandler
	}
	return 0
}

func (m *NVMDataResponse) GetFinalResponse() *NVMFinalResponse {
	if m != nil {
		return m.FinalResponse
	}
	return nil
}

func (m *NVMDataResponse) GetCallbackResponse() *NVMCallbackResponse {
	if m != nil {
		return m.CallbackResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*NVMConfigBundle)(nil), "NVMConfigBundle")
	proto.RegisterType((*NVMCallbackResult)(nil), "NVMCallbackResult")
	proto.RegisterType((*NVMDataRequest)(nil), "NVMDataRequest")
	proto.RegisterType((*NVMStatsBundle)(nil), "NVMStatsBundle")
	proto.RegisterType((*NVMFinalResponse)(nil), "NVMFinalResponse")
	proto.RegisterType((*NVMCallbackResponse)(nil), "NVMCallbackResponse")
	proto.RegisterType((*NVMDataResponse)(nil), "NVMDataResponse")
}

func init() { proto.RegisterFile("nvm.proto", fileDescriptor_83a3db176960b780) }

var fileDescriptor_83a3db176960b780 = []byte{
	// 877 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0x66, 0xd3, 0xb4, 0x4d, 0x26, 0x49, 0x9b, 0xec, 0x95, 0xb2, 0x70, 0x3a, 0x35, 0x97, 0x0a,
	0x29, 0x12, 0x52, 0x80, 0x1e, 0x48, 0x48, 0xf0, 0x02, 0xa5, 0xd0, 0x9e, 0xae, 0xe1, 0xb4, 0x29,
	0x7d, 0xb5, 0x1c, 0xaf, 0x93, 0x2c, 0xe7, 0xb5, 0xc3, 0xda, 0x5b, 0x6d, 0xef, 0x2f, 0xc0, 0x1b,
	0x6f, 0xfc, 0x2a, 0x7e, 0x12, 0xf2, 0xd8, 0x69, 0x36, 0x81, 0xeb, 0xd3, 0xbd, 0xad, 0xbf, 0xef,
	0x9b, 0xb1, 0x67, 0xfc, 0x8d, 0x17, 0x9a, 0xf2, 0x2e, 0x1b, 0x2d, 0x73, 0x65, 0xd4, 0xe0, 0x9f,
	0x5d, 0x38, 0x1c, 0xdf, 0x5e, 0x9f, 0x2b, 0x39, 0x4b, 0xe7, 0x3f, 0x14, 0x32, 0x11, 0x3c, 0x7c,
	0x06, 0xa0, 0x59, 0x9e, 0x2e, 0x0d, 0xd1, 0x39, 0x8b, 0x82, 0x7e, 0x30, 0x6c, 0xc6, 0x4d, 0x87,
	0x4c, 0x72, 0x16, 0x9e, 0x40, 0xcb, 0xd3, 0xe6, 0x7e, 0xc9, 0xa3, 0x1a, 0xf2, 0x3e, 0xe2, 0xe6,
	0x7e, 0xc9, 0x2b, 0x82, 0x05, 0xd5, 0x8b, 0x68, 0xa7, 0x2a, 0xb8, 0xa4, 0x7a, 0x11, 0x3e, 0x87,
	0x76, 0x5e, 0x48, 0x49, 0xa7, 0x82, 0xe3, 0x16, 0x75, 0x54, 0xb4, 0x56, 0x98, 0xdd, 0xe4, 0x19,
	0xc0, 0x54, 0x28, 0xf6, 0x86, 0xfc, 0xa6, 0x95, 0x8c, 0x76, 0xdd, 0x19, 0x10, 0x79, 0xa9, 0x95,
	0x0c, 0x3f, 0x82, 0x7d, 0x53, 0x3a, 0x6e, 0x0f, 0xb9, 0x3d, 0x53, 0x22, 0xf1, 0x14, 0x9a, 0x99,
	0x4a, 0x0a, 0xc1, 0x49, 0x9a, 0x44, 0xfb, 0x48, 0x35, 0x1c, 0x70, 0x95, 0xd8, 0x7d, 0x5d, 0xd2,
	0x05, 0x4f, 0xe7, 0x0b, 0x13, 0x35, 0xfa, 0xc1, 0xb0, 0x1e, 0xb7, 0x10, 0xbb, 0x44, 0x28, 0xfc,
	0x18, 0x1a, 0x6c, 0x41, 0x53, 0x69, 0xc3, 0x9b, 0xfd, 0x60, 0xd8, 0x89, 0xf7, 0x71, 0xed, 0xa2,
	0x33, 0x6e, 0x28, 0xb9, 0xe3, 0xb9, 0x4e, 0x95, 0x8c, 0xc0, 0x9d, 0xda, 0x62, 0xb7, 0x0e, 0x0a,
	0x4f, 0xa1, 0xc3, 0x5d, 0x59, 0x22, 0xcd, 0x52, 0xa3, 0xa3, 0x56, 0x3f, 0x18, 0x36, 0xe2, 0xb6,
	0x03, 0x5f, 0x21, 0x16, 0x7e, 0x05, 0xc7, 0x8e, 0x25, 0xbc, 0xe4, 0x24, 0x95, 0xda, 0xe4, 0x05,
	0x33, 0x36, 0x63, 0x1b, 0xcf, 0x73, 0xe4, 0xd8, 0x8b, 0x92, 0x5f, 0xad, 0xb9, 0xf0, 0x4b, 0xf8,
	0xd0, 0x47, 0x19, 0x65, 0xa8, 0x20, 0x19, 0xcf, 0x88, 0x4e, 0xdf, 0xf2, 0xa8, 0x83, 0x41, 0xa1,
	0x23, 0x6f, 0x2c, 0x77, 0xcd, 0xb3, 0x49, 0xfa, 0x96, 0x87, 0x9f, 0x41, 0x8f, 0x97, 0x9c, 0x15,
	0x36, 0x9e, 0x98, 0x34, 0xe3, 0xaa, 0x30, 0xd1, 0x01, 0xca, 0xbb, 0x0f, 0xc4, 0x8d, 0xc3, 0xc3,
	0x17, 0x70, 0xec, 0x25, 0x64, 0x4e, 0xb5, 0x3b, 0x3f, 0x61, 0x4a, 0x9b, 0xe8, 0x10, 0x23, 0x9e,
	0x78, 0xf6, 0x67, 0xaa, 0xb1, 0x8e, 0x73, 0xa5, 0x4d, 0xf8, 0x0a, 0x4e, 0x33, 0x5a, 0xfa, 0x62,
	0x89, 0x9a, 0x91, 0xf5, 0x7e, 0xd5, 0xba, 0xba, 0x98, 0xe1, 0x24, 0xa3, 0xa5, 0x6b, 0xc1, 0x2f,
	0xb3, 0x8b, 0x95, 0xae, 0x5a, 0xe2, 0x05, 0xf4, 0x13, 0x3e, 0xa3, 0x85, 0x30, 0x95, 0x8c, 0x5b,
	0xd5, 0xf6, 0x30, 0xd5, 0x53, 0xaf, 0x5b, 0xa5, 0xab, 0x96, 0x3d, 0xb8, 0x87, 0x9e, 0x75, 0x34,
	0x15, 0x62, 0x4a, 0xd9, 0x9b, 0x98, 0xeb, 0x42, 0x18, 0xeb, 0x8b, 0x59, 0x21, 0x19, 0x91, 0x34,
	0xe3, 0xde, 0xd2, 0x0d, 0x0b, 0x8c, 0x69, 0xc6, 0xc3, 0x63, 0xd8, 0xcb, 0x51, 0xe6, 0xcd, 0xec,
	0x57, 0xd6, 0x0c, 0x52, 0x19, 0x22, 0x0b, 0x21, 0xd0, 0xc5, 0x8d, 0x78, 0x5f, 0x2a, 0x33, 0x2e,
	0x84, 0x08, 0x8f, 0x60, 0x97, 0x97, 0x26, 0xa7, 0x51, 0xbd, 0xbf, 0x33, 0x6c, 0xc6, 0x6e, 0x31,
	0xf8, 0xb3, 0x06, 0x07, 0xe3, 0xdb, 0xeb, 0x1f, 0xa9, 0xa1, 0x31, 0xff, 0xbd, 0xe0, 0xda, 0xa0,
	0xd7, 0xdd, 0xa7, 0x1b, 0x97, 0xc0, 0x7b, 0xdd, 0x61, 0x38, 0x2f, 0x15, 0x49, 0x2a, 0x93, 0x12,
	0x0f, 0xd1, 0x79, 0x90, 0x5c, 0xc9, 0xa4, 0xb4, 0x23, 0x25, 0x98, 0x26, 0x0b, 0x6a, 0x07, 0x34,
	0xc7, 0xc3, 0xd4, 0x63, 0x10, 0x4c, 0x5f, 0x3a, 0xc4, 0x0a, 0xe6, 0x15, 0x41, 0xdd, 0x09, 0xe6,
	0x6b, 0xc1, 0xb7, 0x70, 0xc8, 0x7c, 0x4b, 0x88, 0x2f, 0xd6, 0x4e, 0x55, 0xeb, 0x2c, 0x1c, 0xfd,
	0xa7, 0x5b, 0xf1, 0x01, 0xdb, 0xec, 0xde, 0xd7, 0xd0, 0x61, 0xf8, 0x42, 0x90, 0x29, 0x3e, 0x11,
	0x38, 0x74, 0xad, 0xb3, 0xee, 0x68, 0xeb, 0xe9, 0x88, 0xdb, 0xac, 0xb2, 0x1a, 0xfc, 0x15, 0x60,
	0x3b, 0x26, 0x86, 0x1a, 0xed, 0xdf, 0x96, 0xd7, 0xf0, 0x29, 0x65, 0xa6, 0xa0, 0x82, 0x30, 0x55,
	0x48, 0xf3, 0x6e, 0xcf, 0x04, 0x58, 0xc1, 0x73, 0x27, 0x3e, 0xb7, 0xda, 0x77, 0xb8, 0xe6, 0x73,
	0x38, 0xf2, 0x19, 0x0b, 0xcd, 0x93, 0xb5, 0x53, 0x6a, 0x98, 0xa0, 0xe7, 0xb8, 0x5f, 0x35, 0x4f,
	0x56, 0xfe, 0xf8, 0x23, 0x80, 0xee, 0xf8, 0xf6, 0xfa, 0xa7, 0x54, 0x52, 0x11, 0x73, 0xbd, 0x54,
	0x52, 0x57, 0x2d, 0x60, 0x37, 0xde, 0x7d, 0xb0, 0x40, 0x17, 0x76, 0x32, 0x3d, 0xf7, 0xbe, 0xb0,
	0x9f, 0x8f, 0x99, 0xe2, 0x0c, 0xda, 0xda, 0xd6, 0xba, 0xea, 0x52, 0x1d, 0xbb, 0x74, 0x38, 0xda,
	0xec, 0x41, 0xdc, 0xd2, 0xeb, 0xc5, 0x60, 0x02, 0x4f, 0x36, 0xfb, 0xef, 0xce, 0xf3, 0xa8, 0x5f,
	0x4f, 0xa0, 0x85, 0xe4, 0x92, 0xe6, 0x34, 0xd3, 0x51, 0x0d, 0x2d, 0x08, 0x16, 0x7a, 0x8d, 0xc8,
	0xe0, 0xef, 0x1a, 0xbe, 0xea, 0xce, 0x87, 0x3e, 0xe3, 0x29, 0x74, 0x72, 0xff, 0x5d, 0x75, 0x62,
	0x7b, 0x05, 0xa2, 0x15, 0xab, 0xa2, 0x8a, 0x17, 0x1f, 0x44, 0xef, 0xc9, 0x8c, 0xdf, 0xc0, 0xc1,
	0xcc, 0xb6, 0x9f, 0xac, 0xf2, 0x7a, 0x2f, 0xf6, 0x46, 0xdb, 0x17, 0x13, 0x77, 0x66, 0x1b, 0xf7,
	0xf4, 0x3d, 0xf4, 0xaa, 0x36, 0x76, 0xc1, 0xce, 0x8d, 0x47, 0xa3, 0xff, 0x69, 0x64, 0xdc, 0x65,
	0x5b, 0xc8, 0xd9, 0x4b, 0x00, 0x7b, 0x21, 0x3c, 0xbf, 0x4b, 0x19, 0x0f, 0xbf, 0x83, 0xde, 0x24,
	0xa3, 0xb9, 0x39, 0x57, 0xd2, 0xe4, 0x94, 0x19, 0x9b, 0x20, 0xc4, 0x2b, 0xab, 0x4c, 0xf1, 0x27,
	0xdd, 0xd1, 0x56, 0x3b, 0x07, 0x1f, 0x0c, 0x83, 0x2f, 0x82, 0xe9, 0x1e, 0xfe, 0x45, 0x5f, 0xfc,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x90, 0xbd, 0x59, 0x84, 0x52, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NVMServiceClient is the client API for NVMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NVMServiceClient interface {
	SmartContractCall(ctx context.Context, opts ...grpc.CallOption) (NVMService_SmartContractCallClient, error)
}

type nVMServiceClient struct {
	cc *grpc.ClientConn
}

func NewNVMServiceClient(cc *grpc.ClientConn) NVMServiceClient {
	return &nVMServiceClient{cc}
}

func (c *nVMServiceClient) SmartContractCall(ctx context.Context, opts ...grpc.CallOption) (NVMService_SmartContractCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NVMService_serviceDesc.Streams[0], "/NVMService/SmartContractCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &nVMServiceSmartContractCallClient{stream}
	return x, nil
}

type NVMService_SmartContractCallClient interface {
	Send(*NVMDataRequest) error
	Recv() (*NVMDataResponse, error)
	grpc.ClientStream
}

type nVMServiceSmartContractCallClient struct {
	grpc.ClientStream
}

func (x *nVMServiceSmartContractCallClient) Send(m *NVMDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nVMServiceSmartContractCallClient) Recv() (*NVMDataResponse, error) {
	m := new(NVMDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NVMServiceServer is the server API for NVMService service.
type NVMServiceServer interface {
	SmartContractCall(NVMService_SmartContractCallServer) error
}

// UnimplementedNVMServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNVMServiceServer struct {
}

func (*UnimplementedNVMServiceServer) SmartContractCall(srv NVMService_SmartContractCallServer) error {
	return status.Errorf(codes.Unimplemented, "method SmartContractCall not implemented")
}

func RegisterNVMServiceServer(s *grpc.Server, srv NVMServiceServer) {
	s.RegisterService(&_NVMService_serviceDesc, srv)
}

func _NVMService_SmartContractCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NVMServiceServer).SmartContractCall(&nVMServiceSmartContractCallServer{stream})
}

type NVMService_SmartContractCallServer interface {
	Send(*NVMDataResponse) error
	Recv() (*NVMDataRequest, error)
	grpc.ServerStream
}

type nVMServiceSmartContractCallServer struct {
	grpc.ServerStream
}

func (x *nVMServiceSmartContractCallServer) Send(m *NVMDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nVMServiceSmartContractCallServer) Recv() (*NVMDataRequest, error) {
	m := new(NVMDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NVMService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NVMService",
	HandlerType: (*NVMServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SmartContractCall",
			Handler:       _NVMService_SmartContractCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "nvm.proto",
}
