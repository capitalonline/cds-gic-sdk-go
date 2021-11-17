package mysql

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
	ProductName *string                     `json:"ProductName" name:"ProductName"`
	RegionId    *string                     `json:"RegionId" name:"RegionId"`
	Products    []*AvailableDBConfigProduct `json:"Products" name:"Products"`
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
	DiskType     *string `json:"DiskType" name:"DiskType"`
	DiskValue    *int    `json:"DiskValue" name:"DiskValue"`
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
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	IP           *string `json:"IP" name:"IP"`
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
	CloneServices   *interface{}         `json:"CloneServices" name:"CloneServices"`
	Cpu             *int                 `json:"Cpu" name:"Cpu"`
	CreatedTime     *string              `json:"CreatedTime" name:"CreatedTime"`
	Disks           *int                 `json:"Disks" name:"Disks"`
	DisplayName     *string              `json:"DisplayName" name:"DisplayName"`
	IP              *string              `json:"IP" name:"IP"`
	InstanceName    *string              `json:"InstanceName" name:"InstanceName"`
	InstanceUuid    *string              `json:"InstanceUuid" name:"InstanceUuid"`
	LinkType        *string              `json:"LinkType" name:"LinkType"`
	LinkTypeStr     *string              `json:"LinkTypeStr" name:"LinkTypeStr"`
	MasterInfo      *string              `json:"MasterInfo" name:"MasterInfo"`
	Port            *int                 `json:"Port" name:"Port"`
	Ram             *int                 `json:"Ram" name:"Ram"`
	RegionId        *string              `json:"RegionId" name:"RegionId"`
	RelationService interface{}          `json:"RelationService" name:"RelationService"`
	ResourceId      *string              `json:"ResourceId" name:"ResourceId"`
	RoGroups        []*DBInstanceRoGroup `json:"RoGroups" name:"RoGroups"`
	Status          *string              `json:"Status" name:"Status"`
	StatusStr       *string              `json:"StatusStr" name:"StatusStr"`
	SubProductName  *string              `json:"SubProductName" name:"SubProductName"`
	VdcId           *string              `json:"VdcId" name:"VdcId"`
	VdcName         *string              `json:"VdcName" name:"VdcName"`
	Version         *string              `json:"Version" name:"Version"`
}

type DBInstanceRoGroup struct {
	Cpu         *int    `json:"Cpu" name:"Cpu"`
	CreateTime  *string `json:"CreateTime" name:"CreateTime"`
	InpaasId    *string `json:"InpaasId" name:"InpaasId"`
	Ip          *string `json:"Ip" name:"Ip"`
	Locked      *string `json:"Locked" name:"Locked"`
	Message     *string `json:"Message" name:"Message"`
	Port        *int    `json:"Port" name:"Port"`
	Progress    *int    `json:"Progress" name:"Progress"`
	Ram         *int    `json:"Ram" name:"Ram"`
	ServiceId   *string `json:"ServiceId" name:"ServiceId"`
	ResourceId  *string `json:"ResourceId" name:"ResourceId"`
	ServiceName *string `json:"ServiceName" name:"ServiceName"`
	Status      *string `json:"Status" name:"Status"`
	SvcType     *string `json:"SvcType" name:"SvcType"`
}

func (response *DescribeDBInstancesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type CreatePrivilegedAccountRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	AccountName  *string `json:"AccountName" name:"AccountName"`
	Password     *string `json:"Password" name:"Password"`
	AccountType  *string `json:"AccountType" name:"AccountType"`
	Description  *string `json:"Description" name:"Description"`
}

func (request *CreatePrivilegedAccountRequest) ToJsonSting() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *CreatePrivilegedAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type CreatePrivilegedAccountResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *CreatePrivilegedAccountResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *CreatePrivilegedAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeModifiableDBSpecRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (request *DescribeModifiableDBSpecRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeModifiableDBSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeModifiableDBSpecResponse struct {
	*cdshttp.BaseResponse
	Code    *string           `json:"Code" name:"Code"`
	Message *string           `json:"Message" name:"Message"`
	Data    *ModifiableDBSpec `json:"Data" name:"Data"`
	TaskId  *string           `json:"TaskId" name:"TaskId"`
}

type ModifiableDBSpec struct {
	AttachDisk  []*ModifiableDBSpecAttachDisk `json:"AttachDisk" name:"AttachDisk"`
	CpuRam      []*ModifiableDBSpecCpuRam     `json:"CpuRam" name:"CpuRam"`
	ProductName *string                       `json:"ProductName" name:"ProductName"`
	RegionId    *string                       `json:"RegionId" name:"RegionId"`
}

type ModifiableDBSpecAttachDisk struct {
	BasicIops     *string `json:"BasicIops" name:"BasicIops"`
	DiskMax       *int    `json:"DiskMax" name:"DiskMax"`
	DiskMaxExpand *int    `json:"DiskMaxExpand" name:"DiskMaxExpand"`
	DiskMin       *int    `json:"DiskMin" name:"DiskMin"`
	DiskName      *string `json:"DiskName" name:"DiskName"`
	DiskStep      *int    `json:"DiskStep" name:"DiskStep"`
	DiskUnit      *string `json:"DiskUnit" name:"DiskUnit"`
	DiskValue     *string `json:"DiskValue" name:"DiskValue"`
}

type ModifiableDBSpecCpuRam struct {
	CPU         *int    `json:"CPU" name:"CPU"`
	N           *string `json:"N" name:"N"`
	PaasGoodsId *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	RAM         *int    `json:"RAM" name:"RAM"`
}

func (response *DescribeModifiableDBSpecResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeModifiableDBSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type ModifyDBInstanceSpecRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	PaasGoodsId  *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	DiskType     *string `json:"DiskType" name:"DiskType"`
	DiskValue    *int    `json:"DiskValue" name:"DiskValue"`
}

func (request *ModifyDBInstanceSpecRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type ModifyDBInstanceSpecResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *ModifyDBInstanceSpecResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
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

type DescribeAvailableReadOnlyConfigRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (request *DescribeAvailableReadOnlyConfigRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeAvailableReadOnlyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeAvailableReadOnlyConfigResponse struct {
	*cdshttp.BaseResponse
	Code    *string                `json:"Code" name:"Code"`
	Message *string                `json:"Message" name:"Message"`
	Data    *AvailableDBConfigData `json:"Data" name:"Data"`
	TaskId  *string                `json:"TaskId" name:"TaskId"`
}

type AvailableDBConfigData struct {
	AttachDisk  []*AvailableDBConfigAttchDisk `json:"AttachDisk" name:"AttachDisk"`
	CpuRam      []*AvailableDBConfigCpuRam    `json:"CpuRam" name:"CpuRam"`
	ProductName *string                       `json:"ProductName" name:"ProductName"`
	RegionId    *string                       `json:"RegionId" name:"RegionId"`
}

func (response *DescribeAvailableReadOnlyConfigResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeAvailableReadOnlyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type CreateReadOnlyDBInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	PaasGoodsId  *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	DiskType     *string `json:"DiskType" name:"DiskType"`
	DiskValue    *int    `json:"DiskValue" name:"DiskValue"`
	TestGroupId  *int    `json:"TestGroupId" name:"TestGroupId"`
	Amount       *int    `json:"Amount" name:"Amount"`
}

func (request *CreateReadOnlyDBInstanceRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *CreateReadOnlyDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type CreateReadOnlyDBInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *CreateReadOnlyDBInstanceResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *CreateReadOnlyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type CreateBackupRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	BackupType   *string `json:"BackupType" name:"BackupType"`
	Desc         *string `json:"Desc" name:"Desc"`
}

func (request *CreateBackupRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *CreateBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type CreateBackupResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *CreateBackupResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeBackupsRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (request *DescribeBackupsRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeBackupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeBackupsResponse struct {
	*cdshttp.BaseResponse
	Code    *string       `json:"Code" name:"Code"`
	Message *string       `json:"Message" name:"Message"`
	Data    []*BackupData `json:"Data" name:"Data"`
	TaskId  *string       `json:"TaskId" name:"TaskId"`
}

type BackupData struct {
	BackupId   *string `json:"BackupId" name:"BackupId"`
	BackupMode *string `json:"BackupMode" name:"BackupMode"`
	BackupSize *int    `json:"BackupSize" name:"BackupSize"`
	BackupType *string `json:"BackupType" name:"BackupType"`
	Desc       *string `json:"Desc" name:"Desc"`
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	Status     *string `json:"Status" name:"Status"`
}

func (response *DescribeBackupsResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DeleteBackupRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	BackupId     *string `json:"BackupId" name:"BackupId"`
}

func (request *DeleteBackupRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DeleteBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DeleteBackupResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *DeleteBackupResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DeleteBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DownloadBackupRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	BackupId     *string `json:"BackupId" name:"BackupId"`
}

func (request *DownloadBackupRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DownloadBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DownloadBackupResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	Data    *struct {
		BackupDownloadUrl         *string `json:"BackupDownloadUrl" name:"BackupDownloadUrl"`
		BackupIntranetDownloadUrl *string `json:"BackupIntranetDownloadUrl" name:"BackupIntranetDownloadUrl"`
		BackupMode                *string `json:"BackupMode" name:"BackupMode"`
		BackupSize                *int    `json:"BackupSize" name:"BackupSize"`
		BackupType                *string `json:"BackupType" name:"BackupType"`
		Desc                      *string `json:"Desc" name:"Desc"`
		StartTime                 *string `json:"StartTime" name:"StartTime"`
		EndTIme                   *string `json:"EndTime" name:"EndTime"`
		InstanceId                *string `json:"InstanceId" name:"InstanceId"`
		Status                    *string `json:"Status" name:"Status"`
	} `json:"Data" name:"Data"`
}

func (response *DownloadBackupResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DownloadBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type StartBatchRollbackRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	BackupId     *string `json:"BackupId" name:"BackupId"`
}

func (request *StartBatchRollbackRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *StartBatchRollbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type StartBatchRollbackResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name::"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *StartBatchRollbackResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *StartBatchRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeRollbackRangeTimeRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (request *DescribeRollbackRangeTimeRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeRollbackRangeTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeRollbackRangeTimeResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	Data    *struct {
		LeftBorderTime  *string `json:"LeftBorderTime" name:"LeftBorderTime"`
		RightBorderTime *string `json:"RightBorderTime" name:"RightBorderTime"`
	} `json:"Data" name:"Data"`
	TaskId *string `json:"TaskId" name:"TaskId"`
}

func (response *DescribeRollbackRangeTimeResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeRollbackRangeTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type StartBatchRollbackToTemporaryDBInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid          *string `json:"InstanceUuid" name:"InstanceUuid"`
	SelectTime            *string `json:"SelectTime" name:"SelectTime"`
	TemporaryInstanceName *string `json:"TemporaryInstanceName" name:"TemporaryInstanceName"`
}

func (request *StartBatchRollbackToTemporaryDBInstanceRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *StartBatchRollbackToTemporaryDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type StartBatchRollbackToTemporaryDBInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *StartBatchRollbackToTemporaryDBInstanceResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *StartBatchRollbackToTemporaryDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeTemporaryDBInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
}

func (request *DescribeTemporaryDBInstancesRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeTemporaryDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeTemporaryDBInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	Data    []*struct {
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		Ip                    *string `json:"Ip" name:"Ip"`
		Port                  *int    `json:"Port" name:"Port"`
		SelectRecoveryTime    *string `json:"SelectRecoveryTime" name:"SelectRecoveryTime"`
		Status                *string `json:"Status" name:"Status"`
		StatusStr             *string `json:"StatusStr" name:"StatusStr"`
		TemporaryInstanceName *string `json:"TemporaryInstanceName" name:"TemporaryInstanceName"`
		TemporaryInstanceUuid *string `json:"TemporaryInstanceUuid" name:"TemporaryInstanceUuid"`
	} `json:"Data" name:"Data"`
	TaskId *string `json:"TaskId" name:"TaskId"`
}

func (response *DescribeTemporaryDBInstancesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeTemporaryDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type RegularizeTemporaryDBInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid          *string `json:"InstanceUuid" name:"InstanceUuid"`
	TemporaryInstanceUuid *string `json:"TemporaryInstanceUuid" name:"TemporaryInstanceUuid"`
}

func (request *RegularizeTemporaryDBInstancesRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *RegularizeTemporaryDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type RegularizeTemporaryDBInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *RegularizeTemporaryDBInstancesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *RegularizeTemporaryDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DeleteTemporaryDBInstancesRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid          *string `json:"InstanceUuid" name:"InstanceUuid"`
	TemporaryInstanceUuid *string `json:"TemporaryInstanceUuid" name:"TemporaryInstanceUuid"`
}

func (request *DeleteTemporaryDBInstancesRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DeleteTemporaryDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DeleteTemporaryDBInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string     `json:"Code" name:"Code"`
	Message *string     `json:"Message" name:"Message"`
	Data    interface{} `json:"Data" name:"Data"`
	TaskId  *string     `json:"TaskId" name:"TaskId"`
}

func (response *DeleteTemporaryDBInstancesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DeleteTemporaryDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeDBInstancePerformanceRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	MetricKey    *string `json:"MetricKey" name:"MetricKey"`
	StartTime    *string `json:"StartTime" name:"StartTime"`
	EndTime      *string `json:"EndTime" name:"EndTime"`
}

func (request *DescribeDBInstancePerformanceRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeDBInstancePerformanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeDBInstancePerformanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string                    `json:"Code" name:"Code"`
	Message *string                    `json:"Message" name:"Message"`
	Data    *DBInstancePerformanceData `json:"Data" name:"Data"`
	TaskId  *string                    `json:"TaskId" name:"TaskId"`
}

type DBInstancePerformanceData struct {
	DataPoints   []*DBInstancePerformanceDataPoint `json:"DataPoints" name:"DataPoints"`
	StartTime    *string                           `json:"StartTime" name:"StartTime"`
	EndTime      *string                           `json:"EndTime" name:"EndTime"`
	InstanceUuid *string                           `json:"InstanceUuid" name:"InstanceUuid"`
	MetricKey    *string                           `json:"MetricKey" name:"MetricKey"`
	Period       *int                              `json:"Period" name:"Period"`
	ProductType  *string                           `json:"ProductType" name:"ProductType"`
}

type DBInstancePerformanceDataPoint struct {
	MetricName  *string                       `json:"MetricName" name:"MetricName"`
	MetricType  *string                       `json:"MetricType" name:"MetricType"`
	MonitorType *string                       `json:"MonitorType" name:"MonitorType"`
	Unit        *string                       `json:"Unit" name:"Unit"`
	Values      []*DBInstancePerformanceValue `json:"Values" name:"Values"`
}

type DBInstancePerformanceValue struct {
	DateTime *string `json:"DateTime" name:"DateTime"`
	Value    *int    `json:"Value" name:"Value"`
}

func (response *DescribeDBInstancePerformanceResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeDBInstancePerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}
