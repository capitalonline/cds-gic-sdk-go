package instance

import (
	"encoding/json"

	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
)

// Create Instance Request
type AddInstanceRequest struct {
	*cdshttp.BaseRequest
	RegionId           *string      `json:"RegionId,omitempty" name:"RegionId"`
	VdcId              *string      `json:"VdcId,omitempty" name:"VdcId"`
	InstanceName       *string      `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceChargeType *string      `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	Password           *string      `json:"Password,omitempty" name:"Password"`
	PublicKey          *string      `json:"PublicKey,omitempty" name:"PublicKey"`
	Cpu                *int         `json:"Cpu,omitempty" name:"Cpu"`
	Ram                *int         `json:"Ram,omitempty" name:"Ram"`
	InstanceType       *string      `json:"InstanceType,omitempty" name:"InstanceType"`
	ImageId            *string      `json:"ImageId,omitempty" name:"ImageId"`
	AssignCCSId        *string      `json:"AssignCCSId,omitempty" name:"AssignCCSId"`
	SystemDisk         *SystemDisk  `json:"SystemDisk,omitempty" name:"SystemDisk"`
	DataDisks          []*DataDisk  `json:"DataDisks,omitempty" name:"DataDisks"`
	AutoRenew          *int         `json:"AutoRenew,omitempty" name:"AutoRenew"`
	PrepaidMonth       *int         `json:"PrepaidMonth,omitempty" name:"PrepaidMonth"`
	PublicIp           []*string    `json:"PublicIp,omitempty" name:"PublicIp"`
	PrivateIp          []*PrivateIp `json:"PrivateIp,omitempty" name:"PrivateIp"`
	Amount             *int         `json:"Amount,omitempty" name:"Amount"`
	UTC                *bool        `json:"UTC,omitempty" name:"UTC"`
	ImagePassword      *string      `json:"ImagePassword,omitempty" name:"ImagePassword"`
	UserData           []*string    `json:"UserData,omitempty" name:"UserData"`
	DescriptionNum     *string      `json:"DescriptionNum,omitempty" name:"DescriptionNum"`
	LabelNames         []*string    `json:"LabelNames,omitempty" name:"LabelNames"`
	DeletionProtection *bool        `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
	HostName           *string      `json:"HostName,omitempty" name:"HostName"`
	SubjectId          *int         `json:"SubjectId,omitempty" name:"SubjectId"`
	DedicatedHostId    *string      `json:"DedicatedHostId,omitempty" name:"DedicatedHostId"`
}

type SystemDisk struct {
	Type *string `json:"Type,omitempty" name:"Type" tf:"type"`
	IOPS *int    `json:"IOPS,omitempty" name:"IOPS" tf:"iops"`
	Size *int    `json:"Size,omitempty" name:"Size" tf:"size"`
}

type DataDisk struct {
	Size *int    `json:"Size,omitempty" name:"Size"`
	Type *string `json:"Type,omitempty" name:"Type"`
	IOPS *int    `json:"IOPS,omitempty" name:"IOPS"`
}

type PrivateIp struct {
	PrivateID *string   `json:"PrivateId,omitempty" name:"PrivateId"`
	IP        []*string `json:"IP,omitempty" name:"IP"`
	Segment   *string   `json:"Segment,omitempty" name:"Segment"`
}

func (instance *AddInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *AddInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Create Instance Reponse
type AddInstanceResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *AddInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *AddInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Describe Instance Request
type DescribeInstanceRequest struct {
	*cdshttp.BaseRequest
	VdcId      *string   `json:"VdcId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty"`
	PublicIp   []*string `json:"PublicIp,omitempty"`
	PageNumber *int      `json:"PageNumber,omitempty"`
	PageSize   *int      `json:"PageSize,omitempty"`
	VdcName    *string   `json:"VdcName,omitempty"`
	TagName    *string   `json:"TagName,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty"`
}

func (instance *DescribeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type DescribeReturnInfo struct {
	InstanceStatus          *string    `json:"InstanceStatus"`
	InstanceName            *string    `json:"InstanceName"`
	InstanceId              *string    `json:"InstanceId"`
	VdcId                   *string    `json:"VdcId"`
	VdcName                 *string    `json:"VdcName"`
	Disks                   *DisksInfo `json:"Disks"`
	Cpu                     *int       `json:"Cpu"`
	Ram                     *int       `json:"Ram"`
	PrivateNetworkInterface []*PrivateNetworkInterfaceInfo
	PublicNetworkInterface  []*PublicNetworkInterfaceInfo `json:"PublicNetworkInterface"`
	InstanceChargeType      *string                       `json:"InstanceChargeType"`
	CreateTime              *string                       `json:"CreateTime" name:"CreateTime"`
}

type DisksInfo struct {
	DataDisks       []*DataDisksInfo `json:"DataDisks"`
	SystemDisk      *SystemDiskInfo  `json:"SystemDisk"`
	LeftDataDiskNum *int             `json:"LeftDataDiskNum"`
}

type DataDisksInfo struct {
	DiskId   *string `json:"DiskId"`
	DiskType *string `json:"DiskType"`
	Size     *int    `json:"Size"`
	IopsSize *int    `json:"IopsSize"`
	Iops     *int    `json:"Iops"`
}

type SystemDiskInfo struct {
	DiskType *string `json:"DiskType"`
	Size     *int    `json:"Size"`
	IopsSize *int    `json:"IopsSize"`
	Iops     *int    `json:"Iops"`
}

type PrivateNetworkInterfaceInfo struct {
	InterfaceId *string `json:"InterfaceId"`
	Name        *string `json:"Name"`
	IP          *string `json:"IP"`
	MAC         *string `json:"MAC"`
	Connected   *int    `json:"Connected"`
	PrivateId   *string `json:"PrivateId"`
}

type PublicNetworkInterfaceInfo struct {
	InterfaceId *string `json:"InterfaceId"`
	Name        *string `json:"Name"`
	IP          *string `json:"IP"`
	MAC         *string `json:"MAC"`
	Connected   *int    `json:"Connected"`
	PublicId    *string `json:"PublicId"`
}

type DescribeInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code"`
	Message *string `json:"Message"`
	Data    struct {
		Instances  []*DescribeReturnInfo
		PageNumber *int `json:"PageNumber"`
		PageCount  *int `json:"PageCount"`
		Total      *int `json:"Total"`
	}
}

func (instance *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Delete Instance Request
type DeleteInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceIds []*string
}

// Delete Instance Reponse
type DeleteInstanceResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Modify Instance Name
type ModifyInstanceNameRequest struct {
	*cdshttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (instance *ModifyInstanceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyInstanceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ModifyInstanceNameResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ModifyInstanceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// ModifyInstanceSpec, cpu, ram
type ModifyInstanceSpecRequest struct {
	*cdshttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Cpu          *int    `json:"Cpu,omitempty" name:"Cpu"`
	Ram          *int    `json:"Ram,omitempty" name:"Ram"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (instance *ModifyInstanceSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyInstanceSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ModifyInstanceSpecResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ModifyInstanceSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Create Disk
type CreateDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string     `json:"InstanceId,omitempty" name:"InstanceId"`
	DataDisks  []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (instance *CreateDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *CreateDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type CreateDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *CreateDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *CreateDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Resize Disk
type ResizeDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DiskId     *string `json:"DiskId,omitempty" name:"DiskId"`
	DataSize   *int    `json:"DataSize,omitempty" name:"DataSize"`
	IOPS       *int    `json:"IOPS,omitempty" name:"IOPS"`
}

func (instance *ResizeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ResizeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ResizeDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ResizeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Resize Disk
type DeleteDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	DiskIds    []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (instance *DeleteDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DeleteDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type DeleteDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *DeleteDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DeleteDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Modify Ip Address
type ModifyIpRequest struct {
	*cdshttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InterfaceId *string `json:"InterfaceId,omitempty" name:"InterfaceId"`
	Address     *string `json:"Address,omitempty" name:"Address"`
	Password    *string `json:"Password,omitempty" name:"Password"`
}

func (instance *ModifyIpRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ModifyIpResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ModifyIpResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ExtendSystemDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Size       *int    `json:"Size,omitempty" name:"Size"`
	IOPS       *int    `json:"IOPS,omitempty" name:"IOPS"`
}

type ExtendSystemDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ExtendSystemDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ExtendSystemDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

func (instance *ExtendSystemDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ExtendSystemDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ResetInstancesPasswordRequest struct {
	*cdshttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	Password    *string `json:"Password,omitempty" name:"Password"`
}

func (request *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *ResetInstancesPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type ResetInstancesPasswordResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *ResetInstancesPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type ResetImageRequest struct {
	*cdshttp.BaseRequest
	InstanceId    *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	ImageId       *string   `json:"ImageId,omitempty" name:"ImageId"`
	ImagePassword *string   `json:"ImagePassword,omitempty" name:"ImagePassword"`
	Password      *string   `json:"Password,omitempty" name:"Password"`
	PublicKey     *string   `json:"PublicKey,omitempty" name:"PublicKey"`
	ProductId     *string   `json:"ProductId,omitempty" name:"ProductId"`
	UserData      []*string `json:"UserData,omitempty" name:"UserData"`
}

func (request *ResetImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *ResetImageRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type ResetImageResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code,omitempty" name:"Code"`
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *ResetImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *ResetImageResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type ModifyInstanceChargeTypeRequest struct {
	*cdshttp.BaseRequest
	InstanceId         *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	AutoRenew          *int    `json:"AutoRenew,omitempty" name:"AutoRenew"`
	PrepaidMonth       *int    `json:"PrepaidMonth,omitempty" name:"PrepaidMonth"`
}

func (request *ModifyInstanceChargeTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *ModifyInstanceChargeTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type ModifyInstanceChargeTypeResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code,omitempty" name:"Code"`
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *ModifyInstanceChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *ModifyInstanceChargeTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type StopInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (request *StopInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *StopInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type StopInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *StopInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *StopInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type RebootInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (request *RebootInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *RebootInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type RebootInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *RebootInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *RebootInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type StartInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (request *StartInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *StartInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type StartInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *StartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *StartInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type StartInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (request *StartInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *StartInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type StartInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *StartInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type StopInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (request *StopInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *StopInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type StopInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *StopInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type RebootInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceIds *string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (request *RebootInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *RebootInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type RebootInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *RebootInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type DescribeInstanceMonitorRequest struct {
	*cdshttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	MetricName  *string `json:"MetricName,omitempty" name:"MetricName"`
	Period      *int    `json:"Period,omitempty" name:"Period"`
	StartTime   *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *string `json:"EndTime,omitempty" name:"EndTime"`
	InterfaceId *string `json:"InterfaceId,omitempty" name:"InterfaceId"`
	DiskId      *string `json:"DiskId,omitempty" name:"DiskId"`
}

func (request *DescribeInstanceMonitorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *DescribeInstanceMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type DescribeInstanceMonitorResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *DescribeInstanceMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *DescribeInstanceMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type GetIpInfoBySegmentRequest struct {
	*cdshttp.BaseRequest
	PrivateId *string `json:"PrivateId,omitempty" name:"PrivateId"`
	Segment   *string `json:"Segment,omitempty" name:"Segment"`
}

func (request *GetIpInfoBySegmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *GetIpInfoBySegmentRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type GetIpInfoBySegmentResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *GetIpInfoBySegmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *GetIpInfoBySegmentResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type ChangeVmDelProtectionRequest struct {
	*cdshttp.BaseRequest
	InstanceIds        []*string
	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
}

func (request *ChangeVmDelProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *ChangeVmDelProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type ChangeVmDelProtectionResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Message *string      `json:"Message,omitempty" name:"Message"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *ChangeVmDelProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *ChangeVmDelProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type AllocateDedicatedHostsRequest struct {
	*cdshttp.BaseRequest
	RegionId            *string `json:"RegionId,omitempty"`
	DedicatedHostType   *string `json:"DedicatedHostType,omitempty"`
	DedicatedHostGoodId *int    `json:"DedicatedHostGoodId,omitempty"`
	DedicatedHostName   *string `json:"DedicatedHostName,omitempty"`
	DedicatedHostCpu    *int    `json:"DedicatedHostCpu,omitempty"`
	DedicatedHostLimit  *int    `json:"DedicatedHostLimit,omitempty"`
	DedicatedHostRam    *int    `json:"DedicatedHostRam,omitempty"`
	PrepaidMonth        *int    `json:"PrepaidMonth,omitempty"`
	AutoRenew           *int    `json:"AutoRenew,omitempty"`
	Amount              *int    `json:"Amount,omitempty"`
	DescriptionNum      *bool   `json:"DescriptionNum,omitempty"`
	SubjectId           *int    `json:"SubjectId,omitempty"`
}

func (request *AllocateDedicatedHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *AllocateDedicatedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type AllocateDedicatedHostsResponse struct {
	*cdshttp.BaseResponse
	Code    *string   `json:"Code,omitempty" name:"Code"`
	Message *string   `json:"Message,omitempty" name:"Message"`
	Data    []*string `json:"Data,omitempty" name:"Data"`
}

func (response *AllocateDedicatedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *AllocateDedicatedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type DescribeDedicatedHostsRequest struct {
	*cdshttp.BaseRequest

	HostId     *string `json:"HostId,omitempty" name:"HostId"`
	PageNumber *int    `json:"PageNumber,omitempty" name:"PageNumber"`
	PageSize   *int    `json:"PageSize,omitempty" name:"PageSize"`
	HostName   *string `json:"HostName,omitempty" name:"HostName"`
}

func (request *DescribeDedicatedHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *DescribeDedicatedHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type DescribeDedicatedHostsResponse struct {
	*cdshttp.BaseResponse
	Code    *string                    `json:"Code,omitempty" name:"Code"`
	Message *string                    `json:"Message,omitempty" name:"Message"`
	Data    DescribeDedicatedHostsData `json:"Data,omitempty" name:"Data"`
}

type DescribeDedicatedHostsData struct {
	HostList []*DescribeDedicatedHostsDataHost `json:"HostList,omitempty"`
	Total    *int                              `json:"Total,omitempty"`
}

type DescribeDedicatedHostsDataHost struct {
	BillMethod    *string `json:"BillMethod,omitempty" name:"BillMethod"`
	CpuRate       *string `json:"CpuRate,omitempty" name:"CpuRate"`
	Duration      *int    `json:"Duration,omitempty" name:"Duration"`
	EndBillTime   *string `json:"EndBillTime,omitempty" name:"EndBillTime"`
	HostId        *string `json:"HostId,omitempty" name:"HostId"`
	HostName      *string `json:"HostName,omitempty" name:"HostName"`
	HostType      *string `json:"HostType,omitempty" name:"HostType"`
	RamRate       *string `json:"RamRate,omitempty" name:"RamRate"`
	Region        *string `json:"Region,omitempty" name:"Region"`
	StartBillTime *string `json:"StartBillTime,omitempty" name:"StartBillTime"`
	Status        *string `json:"Status,omitempty" name:"Status"`
	VmNum         *int    `json:"VmNum,omitempty" name:"VmNum"`
}

func (response *DescribeDedicatedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *DescribeDedicatedHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type DescribeDedicatedHostTypesRequest struct {
	*cdshttp.BaseRequest
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (request *DescribeDedicatedHostTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *DescribeDedicatedHostTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type DescribeDedicatedHostTypesResponse struct {
	*cdshttp.BaseResponse
	Code    *string                                   `json:"Code,omitempty" name:"Code"`
	Message *string                                   `json:"Message,omitempty" name:"Message"`
	Data    []*DescribeDedicatedHostTypesResponseData `json:"Data,omitempty" name:"Data"`
}

type DescribeDedicatedHostTypesResponseData struct {
	BillSchemeId      *string `json:"BillSchemeId"`
	Cpu               *int    `json:"Cpu"`
	GoodsId           *int    `json:"GoodsId"`
	Ram               *int    `json:"Ram"`
	VmFamilyId        *string `json:"VmFamilyId"`
	VmRuleName        *string `json:"VmRuleName"`
	VmSpecId          *string `json:"VmSpecId"`
	VmTypeDescription *string `json:"VmTypeDescription"`
	VmTypeId          *string `json:"VmTypeId"`
	VmTypeName        *string `json:"VmTypeName"`
	VmTypeSort        *int    `json:"VmTypeSort"`
}

func (response *DescribeDedicatedHostTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *DescribeDedicatedHostTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}

type ModifyInstanceHostNameRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	HostName   *string `json:"HostName,omitempty" name:"HostName"`
	Password   *string `json:"Password" name:"Password"`
}

func (request *ModifyInstanceHostNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *ModifyInstanceHostNameRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type ModifyInstanceHostNameResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code,omitempty" name:"Code"`
	Message *string `json:"Message,omitempty" name:"Message"`
	//Data    interface{} `json:"Data,omitempty" name:"Data"`
}

func (response *ModifyInstanceHostNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &response)
}

func (response *ModifyInstanceHostNameResponse) ToJsonString() string {
	b, _ := json.Marshal(response)
	return string(b)
}
