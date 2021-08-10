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
	Code    *string       `json:"Code"`
	Message *string       `json:"Message"`
	Data    []*RegionData `json:"Data"`
}

type RegionData struct {
	CityId     *string `json:"CityId"`
	CityName   *string `json:"CityName"`
	IsSaling   *bool   `json:"IsSaling"`
	RegionId   *string `json:"RegionId"`
	RegionName *string `json:"RegionName"`
	Sitename   *string `json:"SiteName"`
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
	RegionId *string `json:"RegionId"`
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
	Code    *string                   `json:"Code"`
	Message *string                   `json:"Message"`
	Data    []*AvailableDBConfigEntry `json:"Data"`
}

type AvailableDBConfigEntry struct {
	ProductName *string                     `json:"ProductName"`
	RegionId    *string                     `json:"RegionId"`
	Products    []*AvailableDBConfigProduct `json:"Productrs"`
}

type AvailableDBConfigProduct struct {
	Architectures []*AvailableDBConfigArchitecture `json:"Architectures"`
	Version       *string                          `json:"Version"`
}

type AvailableDBConfigArchitecture struct {
	ArchitectureName *string                         `json:"ArchitectureName"`
	ComputeRoles     []*AvailableDBConfigComputeRole `json:"ComputeRole"`
	EnginesType      []interface{}                   `json:"EnginesType"`
	NetworkLinks     []*AvailableDBConfigNetworkLink `json:"NetworkLinks"`
	SubProductName   *string                         `json:"SubProductName"`
}

type AvailableDBConfigComputeRole struct {
	ComputeName *string                       `json:"ComputeName"`
	Standards   []*AvailableDBConfigStandards `json:"Standards"`
}

type AvailableDBConfigStandards struct {
	AttachDisk []*AvailableDBConfigAttchDisk `json:"AttachDisk"`
}

type AvailableDBConfigAttchDisk struct {
	BasicIops     *string `json:"BasicIops"`
	DiskMax       *int    `json:"DiskMax"`
	DiskMaxExpand *int    `json:"DiskMaxExpand"`
	DiskMin       *int    `json:"DiskMin"`
	DiskName      *string `json:"DiskName"`
	DiskStep      *int    `json:"DiskStep"`
	DiskUnit      *string `json:"DiskUnit"`
	DiskValue     *string `json:"DiskValue"`
}

type AvailableDBConfigCpuRam struct {
	CPU         *int    `json:"CPU"`
	Name        *string `json:"Name"`
	PaasGoodsId *int    `json:"PaasGoodsId"`
	RAM         *int    `json:"RAM"`
}

type AvailableDBConfigNetworkLink struct {
	DescDetail *string `json:"DescDetail"`
	LinkType   *string `json:"LinkType"`
	Name       *string `json:"Name"`
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
	RegionId     *string `json:"RegionId"`
	VdcId        *string `json:"VdcId"`
	BasePipeId   *string `json:"BasePipeId"`
	InstanceName *string `json:"InstanceName"`
	PaasGoodsId  *string `json:"PaasGoodsId"`
	DiskType     *string `json:"DiskType"`
	DiskValue    *int    `json:"DiskValue"`
	Password     *string `json:"Password"`
	Amount       *int    `json:"Amount"`
}

func (request *CreateDBInstanceRequest) ToJSonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *CreateDBInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type CreateDBInstanceResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code"`
	Message *string      `json:"Message"`
	Data    *interface{} `json:"Data"`
	TaskId  *string      `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	InstanceName *string `json:"InstanceName"`
	IP           *string `json:"IP"`
}

func (request *DescribeDBInstancesRequest) ToJsonstring() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeDBInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeDBInstancesResponse struct {
	*cdshttp.BaseResponse
	Code    *string                    `json:"Code"`
	Message *string                    `json:"Message"`
	Data    []*DescribeDBInstanceEntry `json:"Data"`
	TaskId  *string                    `json:"TaskId"`
}

type DescribeDBInstanceEntry struct {
	CloneServices   *interface{}         `json:"CloneServices"`
	Cpu             *int                 `json:"Cpu"`
	CreatedTime     *string              `json:"CreatedTime"`
	Disks           *int                 `json:"Disks"`
	DisplayName     *string              `json:"DisplayName"`
	IP              *string              `json:"IP"`
	InstanceName    *string              `json:"InstanceName"`
	InstanceUuid    *string              `json:"InstanceUuid"`
	LinkType        *string              `json:"LinkType"`
	LinkTypeStr     *string              `json:"LinkTypeStr"`
	MasterInfo      *string              `json:"MasterInfo"`
	Port            *int                 `json:"Port"`
	Ram             *int                 `json:"Ram"`
	RegionId        *string              `json:"RegionId"`
	RelationService interface{}          `json:"RelationService"`
	ResourceId      *string              `json:"ResourceId"`
	RoGroups        []*DBInstanceRoGroup `json:"RoGroups"`
	Status          *string              `json:"Status"`
	StatusStr       *string              `json:"StatusStr"`
	SubProductName  *string              `json:"SubProductName"`
	VdcId           *string              `json:"VdcId"`
	VdcName         *string              `json:"VdcName"`
	Version         *string              `json:"Version"`
}

type DBInstanceRoGroup struct {
	Cpu         *int    `json:"Cpu"`
	CreateTime  *string `json:"CreateTime"`
	InpaasId    *string `json:"InpaasId"`
	Ip          *string `json:"Ip"`
	Locked      *string `json:"Locked"`
	Message     *string `json:"Message"`
	Port        *int    `json:"Port"`
	Progress    *int    `json:"Progress"`
	Ram         *int    `json:"Ram"`
	ServiceId   *string `json:"ServiceId"`
	ResourceId  *string `json:"ResourceId"`
	ServiceName *string `json:"ServiceName"`
	Status      *string `json:"Status"`
	SvcType     *string `json:"SvcType"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	AccountName  *string `json:"AccountName"`
	Password     *string `json:"Password"`
	AccountType  *string `json:"AccountType"`
	Description  *string `json:"Description"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
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
	Code    *string           `json:"Code"`
	Message *string           `json:"Message"`
	Data    *ModifiableDBSpec `json:"Data"`
	TaskId  *string           `json:"TaskId"`
}

type ModifiableDBSpec struct {
	AttachDisk  []*ModifiableDBSpecAttachDisk `json:"AttachDisk"`
	CpuRam      []*ModifiableDBSpecCpuRam     `json:"CpuRam"`
	ProductName *string                       `json:"ProductName"`
	RegionId    *string                       `json:"RegionId"`
}

type ModifiableDBSpecAttachDisk struct {
	BasicIops     *string `json:"BasicIops"`
	DiskMax       *int    `json:"DiskMax"`
	DiskMaxExpand *int    `json:"DiskMaxExpand"`
	DiskMin       *int    `json:"DiskMin"`
	DiskName      *string `json:"DiskName"`
	DiskStep      *int    `json:"DiskStep"`
	DiskUnit      *string `json:"DiskUnit"`
	DiskValue     *string `json:"DiskValue"`
}

type ModifiableDBSpecCpuRam struct {
	CPU         *int    `json:"CPU"`
	N           *string `json:"N"`
	PaasGoodsId *int    `json:"PaasGoodsId"`
	RAM         *int    `json:"RAM"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	PaasGoodsId  *int    `json:"PaasGoodsId"`
	DiskType     *string `json:"DiskType"`
	DiskValue    *string `json:"DiskValue"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
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
	Code    *string                `json:"Code"`
	Message *string                `json:"Message"`
	Data    *AvailableDBConfigData `json:"Data"`
	TaskId  *string                `json:"TaskId"`
}

type AvailableDBConfigData struct {
	AttachDisk  []*AvailableDBConfigAttchDisk `json:"AttachDisk"`
	CpuRam      []*AvailableDBConfigCpuRam    `json:"CpuRam"`
	ProductName *string                       `json:"ProductName"`
	RegionId    *string                       `json:"RegionId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	InstanceName *string `json:"InstanceName"`
	PaasGoodsId  *int    `json:"PaasGoodsId"`
	DiskType     *string `json:"DiskType"`
	DiskValue    *int    `json:"DiskValue"`
	TestGroupId  *int    `json:"TestGroupId"`
	Amount       *int    `json:"Amount"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
}

func (response *CreateReadOnlyDBInstanceResponse) ToJSonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *CreateReadOnlyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type CreateBackupRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid"`
	BackupType   *string `json:"BackupType"`
	Desc         *string `json:"Desc"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
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
	Code    *string       `json:"Code"`
	Message *string       `json:"Message"`
	Data    []*BackupData `json:"Data"`
	TaskId  *string       `json:"TaskId"`
}

type BackupData struct {
	BackupId   *string `json:"BackupId"`
	BackupMode *string `json:"BackupMode"`
	BackupSize *int    `json:"BackupSize"`
	BackupType *string `json:"BackupType"`
	Desc       *string `json:"Desc"`
	StartTime  *string `json:"StartTime"`
	EndTime    *string `json:"EndTime"`
	InstanceId *string `json:"InstanceId"`
	Status     *string `json:"Status"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	BackupId     *string `json:"BackupId"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	BackupId     *string `json:"BackupId"`
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
	Code    *string `json:"Code"`
	Message *string `json:"Message"`
	Data    *struct {
		BackupDownloadUrl         *string `json:"BackupDownloadUrl"`
		BackupIntranetDownloadUrl *string `json:"BackupIntranetDownloadUrl"`
		BackupMode                *string `json:"BackupMode"`
		BackupSize                *int    `json:"BackupSize"`
		BackupType                *string `json:"BackupType"`
		Desc                      *string `json:"Desc"`
		StartTime                 *string `json:"StartTime"`
		EndTIme                   *string `json:"EndTime"`
		InstanceId                *string `json:"InstanceId"`
		Status                    *string `json:"Status"`
	} `json:"Data"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	BackupId     *string `json:"BackupId"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
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
	Code    *string `json:"Code"`
	Message *string `json:"Message"`
	Data    *struct {
		LeftBorderTime  *string `json:"LeftBorderTime"`
		RightBorderTime *string `json:"RightBorderTime"`
	} `json:"Data"`
	TaskId *string `json:"TaskId"`
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
	InstanceUuid          *string `json:"InstanceUuid"`
	SelectTime            *string `json:"SelectTime"`
	TemporaryInstanceName *string `json:"TemporaryInstanceName"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
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
	Code    *string `json:"Code"`
	Message *string `json:"Message"`
	Data    []*struct {
		CreateTime            *string `json:"CreateTime"`
		Ip                    *string `json:"Ip"`
		Port                  *int    `json:"Port"`
		SelectRecoveryTime    *string `json:"SelectRecoveryTime"`
		Status                *string `json:"Status"`
		StatusStr             *string `json:"StatusStr"`
		TemporaryInstanceName *string `json:"TemporaryInstanceName"`
		TemporaryInstanceUuid *string `json:"TemporaryInstanceUuid"`
	} `json:"Data"`
	TaskId *string `json:"TaskId"`
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
	InstanceUuid          *string `json:"InstanceUuid"`
	TemporaryInstanceUuid *string `json:"TemporaryInstanceUuid"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid          *string `json:"InstanceUuid"`
	TemporaryInstanceUuid *string `json:"TemporaryInstanceUuid"`
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
	Code    *string     `json:"Code"`
	Message *string     `json:"Message"`
	Data    interface{} `json:"Data"`
	TaskId  *string     `json:"TaskId"`
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
	InstanceUuid *string `json:"InstanceUuid"`
	MetricKey    *string `json:"MetricKey"`
	StartTime    *string `json:"StartTime"`
	EndTime      *string `json:"EndTime"`
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
	Code    *string                    `json:"Code"`
	Message *string                    `json:"Message"`
	Data    *DBInstancePerformanceData `json:"Data"`
	TaskId  *string                    `json:"TaskId"`
}

type DBInstancePerformanceData struct {
	DataPoints   []*DBInstancePerformanceDataPoint `json:"DataPoints"`
	StartTime    *string                           `json:"StartTime"`
	EndTime      *string                           `json:"EndTime"`
	InstanceUuid *string                           `json:"InstanceUuid"`
	MetricKey    *string                           `json:"MetricKey"`
	Period       *int                              `json:"Period"`
	ProductType  *string                           `json:"ProductType"`
}

type DBInstancePerformanceDataPoint struct {
	MetricName  *string                       `json:"MetricName"`
	MetricType  *string                       `json:"MetricType"`
	MonitorType *string                       `json:"MonitorType"`
	Unit        *string                       `json:"Unit"`
	Values      []*DBInstancePerformanceValue `json:"Values"`
}

type DBInstancePerformanceValue struct {
	DateTime *string `json:"DateTime"`
	Value    *int    `json:"Value"`
}

func (response *DescribeDBInstancePerformanceResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeDBInstancePerformanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}
