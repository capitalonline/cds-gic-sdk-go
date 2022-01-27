package redis

import (
	"encoding/json"

	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
)

type DescribeRegionsRequest struct {
	*cdshttp.BaseRequest
}

func (request *DescribeRegionsRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeRegionsResponse struct {
	*cdshttp.BaseResponse
	Code    *string       `json:"Code" name:"Code"`
	Message *string       `json:"Message" name:"Message"`
	Data    []*RegionData `json:"Data" name:"Data"`
}

type RegionData struct {
	CityId     *string `json:"CityId" name:"CityId"`
	CityName   *string `json:"CityName" name:"CityName"`
	IsSaling   *int    `json:"IsSaling" name:"IsSaling"`
	RegionId   *string `json:"RegionId" name:"RegionId"`
	RegionName *string `json:"RegionName" name:"RegionName"`
	SiteName   *string `json:"SiteName" name:"SiteName"`
}

func (response *DescribeRegionsResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeAvailableDBConfigRequest struct {
	*cdshttp.BaseRequest
	RegionId *string `json:"RegionId" name:"RegionId"`
}

func (request *DescribeAvailableDBConfigRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeAvailableDBConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeAvailableDBConfigResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *AvailableDBConfigEntry `json:"Data" name:"Data"`
}

type AvailableDBConfigEntry struct {
	ProductName *string `json:"ProductName" name:"ProductName"`
	Products    *[]AvailableDBConfigProduct
	RegionId    *string `json:"RegionId" name:"RegionId"`
}

type AvailableDBConfigProduct struct {
	Architectures []*AvailableDBConfigArchitecture `json:"Architectures" name:"Architectures"`
	Version       *string                          `json:"Version" name:"Version"`
}

type AvailableDBConfigArchitecture struct {
	ArchitectureName *string                         `json:"ArchitectureName" name:"ArchitectureName"`
	ArchitectureType *int                            `json:"ArchitectureType" name:"ArchitectureType"`
	ComputeRoles     []*AvailableDBConfigComputeRole `json:"ComputeRoles" name:"ComputeRoles"`
	EnginesType      []interface{}                   `json:"EnginesType" name:"EnginesType"`
	NetworkLinks     []*AvailableDBConfigNetworkLink `json:"NetworkLinks" name:"NetworkLinks"`
	SubProductName   *string                         `json:"SubProductName" name:"SubProductName"`
}

type AvailableDBConfigComputeRole struct {
	ComputeName *string                     `json:"ComputeName" name:"ComputeName"`
	ComputeType *int                        `json:"ComputeType" name:"ComputeType"`
	Standards   *AvailableDBConfigStandards `json:"Standards" name:"Standards"`
}

type AvailableDBConfigStandards struct {
	AttachDisk []*AvailableDBConfigAttchDisk `json:"AttachDisk" name:"AttachDisk"`
	CpuRam     []*AvailableDBConfigCpuRam    `json:"CpuRam" name:"CpuRam"`
}

type AvailableDBConfigAttchDisk struct {
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
	DescDetail *string `json:"DescDetail" name:"DescDetail"`
	LinkType   *string `json:"LinkType" name:"LinkType"`
	Name       *string `json:"Name" name:"Name"`
}

func (response *DescribeAvailableDBConfigResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeAvailableDBConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type CreateDBInstanceRequest struct {
	*cdshttp.BaseRequest
	RegionId     *string `json:"RegionId" name:"RegionId"`
	VdcId        *string `json:"VdcId" name:"VdcId"`
	BasePipeId   *string `json:"BasePipeId" name:"BasePipeId"`
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	PaasGoodsId  *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	Password     *string `json:"Password" name:"Password"`
	Amount       *int    `json:"Amount" name:"Amount"`
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
	Code    *string               `json:"Code" name:"Code"`
	Message *string               `json:"Message" name:"Message"`
	Data    *CreateDBInstanceData `json:"Data" name:"Data"`
	TaskId  *string               `json:"TaskId" name:"TaskId"`
}

type CreateDBInstanceData struct {
	InstancesUuid []string `json:"InstancesUuid" name:"InstancesUuid"`
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
	Code    *string                    `json:"Code" name:"Code"`
	Message *string                    `json:"Message" name:"Message"`
	Data    []*DescribeDBInstanceEntry `json:"Data" name:"Data"`
	TaskId  *string                    `json:"TaskId" name:"TaskId"`
}

type DescribeDBInstanceEntry struct {
	RelationService interface{}  `json:"RelationService" name:"RelationService"`
	CloneServices   *interface{} `json:"CloneServices" name:"CloneServices"`
	AppName         *string      `json:"AppName" name:"AppName"`
	Cpu             *int         `json:"Cpu" name:"Cpu"`
	IP              *string      `json:"IP" name:"IP"`
	Port            *int         `json:"Port" name:"Port"`
	Ram             *int         `json:"Ram" name:"Ram"`
	RegionId        *string      `json:"RegionId" name:"RegionId"`
	VdcId           *string      `json:"VdcId" name:"VdcId"`
	MasterInfo      *string      `json:"MasterInfo" name:"MasterInfo"`
	LinkType        *string      `json:"LinkType" name:"LinkType"`
	LinkTypeStr     *string      `json:"LinkTypeStr" name:"LinkTypeStr"`
	SubProductName  *string      `json:"SubProductName" name:"SubProductName"`
	Version         *string      `json:"Version" name:"Version"`
	Status          *string      `json:"Status" name:"Status"`
	StatusStr       *string      `json:"StatusStr" name:"StatusStr"`
	InstanceName    *string      `json:"InstanceName" name:"InstanceName"`
	InstanceUuid    *string      `json:"InstanceUuid" name:"InstanceUuid"`
	CreatedTime     *string      `json:"CreatedTime" name:"CreatedTime"`
	VdcName         *string      `json:"VdcName" name:"VdcName"`
	DisplayName     *string      `json:"DisplayName" name:"DisplayName"`

	ResourceId *string       `json:"ResourceId" name:"ResourceId"`
	RoGroups   []interface{} `json:"RoGroups" name:"RoGroups"`
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

func (request *DeleteDBInstanceRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DeleteDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
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
