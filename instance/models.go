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

//Describe Instance Request
type DescribeInstanceRequest struct {
	*cdshttp.BaseRequest
	VdcId      *string   `json:"VdcId,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty"`
	PublicIp   []*string `json:"PublicIp,omitempty"`
	PageNumber *int      `json:"PageNumber,omitempty"`
	PageSize   *int      `json:"PageSize,omitempty"`
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

type DescribeInstanceReponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code"`
	Message *string `json:"Message"`
	Data    struct {
		Instances  []*DescribeReturnInfo
		PageNumber *int `json:"PageNumber"`
		PageCount  *int `json:"PageCount"`
	}
}

func (instance *DescribeInstanceReponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeInstanceReponse) FromJsonString(s string) error {
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
