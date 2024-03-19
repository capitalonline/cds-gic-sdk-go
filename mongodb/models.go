package mongodb

import (
	"encoding/json"

	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
)

type DescribeZonesRequest struct {
	*cdshttp.BaseRequest
}

func (request *DescribeZonesRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeZonesResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code" name:"Code"`
	Message *string      `json:"Message" name:"Message"`
	Data    []*ZonesData `json:"Data" name:"Data"`
}

type ZonesData struct {
	CityId     *string `json:"CityId" name:"CityId"`
	CityName   *string `json:"CityName" name:"CityName"`
	IsSaling   *int    `json:"IsSaling" name:"IsSaling"`
	RegionId   *string `json:"RegionId" name:"RegionId"`
	RegionName *string `json:"RegionName" name:"RegionName"`
	SiteName   *string `json:"SiteName" name:"SiteName"`
}

func (response *DescribeZonesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeSpecInfoRequest struct {
	*cdshttp.BaseRequest
	RegionId *string `json:"RegionId" name:"RegionId"`
}

func (request *DescribeSpecInfoRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeSpecInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeSpecInfoResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *AvailableDBConfigEntry `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

type AvailableDBConfigEntry struct {
	ProductName *string                     `json:"ProductName" name:"ProductName"`
	Products    []*AvailableDBConfigProduct `json:"Products" name:"Products"`
	RegionId    *string                     `json:"RegionId" name:"RegionId"`
}

type AvailableDBConfigProduct struct {
	Architectures []*AvailableDBConfigArchitecture `json:"Architectures" name:"Architectures"`
	Version       *string                          `json:"Version" name:"Version"`
}

type AvailableDBConfigArchitecture struct {
	ArchitectureName *string                          `json:"ArchitectureName" name:"ArchitectureName"`
	ComputeRoles     []*AvailableDBConfigComputeRoles `json:"ComputeRoles" name:"ComputeRoles"`
	EnginesType      []*string                        `json:"EnginesType" name:"EnginesType"`
	NetworkLinks     []*AvailableDBConfigNetworkLink  `json:"NetworkLinks" name:"NetworkLinks"`
	SubProductName   *string                          `json:"SubProductName" name:"SubProductName"`
}

type AvailableDBConfigComputeRoles struct {
	ComputeName *string                     `json:"ComputeName" name:"ComputeName"`
	Standards   *AvailableDBConfigStandards `json:"Standards" name:"Standards"`
}

type AvailableDBConfigStandards struct {
	AttachDisk []*AvailableDBConfigAttachDisk `json:"AttachDisk" name:"AttachDisk"`
	CpuRam     []*AvailableDBConfigCpuRam     `json:"CpuRam" name:"CpuRam"`
}

type AvailableDBConfigAttachDisk struct {
	BasicIops     *string `json:"BasicIops" name:"BasicIops"`
	DiskMax       *int    `json:"DiskMax" name:"DiskMax"`
	DiskMaxExpand *int    `json:"DiskMaxExpand" name:"DiskMaxExpand"`
	DiskMin       *int    `json:"DiskMin" name:"DiskMin"`
	DiskName      *string `json:"DiskName" name:"DiskName"`
	DiskStep      *int    `json:"DiskStep" name:"DiskStep"`
	DiskUnit      *string `json:"DiskUnit" name:"DiskUnit"`
	DiskValue     *string `json:"DiskValue" name:"DiskValue"`
}

type AvailableDBConfigCpuRam struct {
	CPU         *int    `json:"CPU" name:"CPU"`
	Name        *string `json:"Name" name:"Name"`
	PaasGoodsId *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	RAM         *int    `json:"RAM" name:"RAM"`
}

type AvailableDBConfigNetworkLink struct {
	DescDetail    *string `json:"DescDetail" name:"DescDetail"`
	LinkType      *string `json:"LinkType" name:"LinkType"`
	Name          *string `json:"Name" name:"Name"`
	SplitRwSwitch *int    `json:"SplitRwSwitch" name:"SplitRwSwitch"`
}

func (response *DescribeSpecInfoResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeSpecInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type CreateDBInstanceRequest struct {
	*cdshttp.BaseRequest
	RegionId     *string `json:"RegionId" name:"RegionId"`
	VdcId        *string `json:"VdcId" name:"VdcId"`
	BasePipeId   *string `json:"BasePipeId" name:"BasePipeId"`
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	PaasGoodsId  *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	DiskType     *string `json:"DiskType" name:"DiskType"`
	DiskValue    *int    `json:"DiskValue" name:"DiskValue"`
	Password     *string `json:"Password" name:"Password"`
	Version      *string `json:"Version" name:"Version"`
	SubjectId    *int    `json:"SubjectId,omitempty" name:"SubjectId"`
}

func (request *CreateDBInstanceRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *CreateDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type CreateDBInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string        `json:"Code" name:"Code"`
	Message *string        `json:"Message" name:"Message"`
	Data    *CreateDBEntry `json:"Data" name:"Data"`
	TaskId  *string        `json:"TaskId" name:"TaskId"`
}

type CreateDBEntry struct {
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (response *CreateDBInstanceResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeDBInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	IP           *string `json:"IP,omitempty" name:"IP"`
}

func (request *DescribeDBInstancesRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeDBInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string        `json:"Code" name:"Code"`
	Message *string        `json:"Message" name:"Message"`
	Data    []*DBInstances `json:"Data" name:"Data"`
	TaskId  *string        `json:"TaskId" name:"TaskId"`
}

type DBInstances struct {
	CloneServices   []*string   `json:"CloneServices" name:"CloneServices"`
	DisplayName     *string     `json:"DisplayName" name:"DisplayName"`
	IP              *string     `json:"IP" name:"IP"`
	InstanceName    *string     `json:"InstanceName" name:"InstanceName"`
	InstanceUuid    *string     `json:"InstanceUuid" name:"InstanceUuid"`
	LinkType        *string     `json:"LinkType" name:"LinkType"`
	LinkTypeStr     *string     `json:"LinkTypeStr" name:"LinkTypeStr"`
	MasterInfo      interface{} `json:"MasterInfo" name:"MasterInfo"`
	Port            *int        `json:"Port" name:"Port"`
	Ram             *int        `json:"Ram" name:"Ram"`
	RegionId        *string     `json:"RegionId" name:"RegionId"`
	RelationService interface{} `json:"RelationService" name:"RelationService"`
	Status          *string     `json:"Status" name:"Status"`
	StatusStr       *string     `json:"StatusStr" name:"StatusStr"`
	SubProductName  *string     `json:"SubProductName" name:"SubProductName"`
	VdcId           *string     `json:"VdcId" name:"VdcId"`
	VdcName         *string     `json:"VdcName" name:"VdcName"`
	Version         *string     `json:"Version" name:"Version"`
	Message         *string     `json:"Message" name:"Message"`
}

func (response *DescribeDBInstancesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DeleteDBInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (response *DeleteDBInstanceRequest) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DeleteDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DeleteDBInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *DeleteDBInstanceResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DeleteDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}
