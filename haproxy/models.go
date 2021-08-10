package haproxy

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
	Code    *string              `json:"Code,omitempty" name:"Code"`
	Data    []*DescribeZonesData `json:"Data,omitempty" name:"Data"`
	Message *string              `json:"Message,omitempty" name:"Message"`
}

type DescribeZonesData struct {
	CityId     *string `json:"CityId,omitempty" name:"CityId"`
	CityName   *string `json:"CityName,omitempty" name:"CityName"`
	IsSaling   *int    `json:"IsSaling,omitempty" name:"IsSaling"`
	RegionId   *string `json:"RegionId,omitempty" name:"RegionId"`
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	SiteName   *string `json:"SiteName,omitempty" name:"SiteName"`
}

func (response *DescribeZonesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancersSpecRequest struct {
	*cdshttp.BaseRequest
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (request *DescribeLoadBalancersSpecRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeLoadBalancersSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeLoadBalancersSpecResponse struct {
	*cdshttp.BaseResponse
	Code    *string                        `json:"Code,omitempty" name:"Code"`
	Data    *DescribeLoadBalancersSpecData `json:"data,omitempty" name:"data"`
	Message *string                        `json:"Message,omitempty" name:"Message"`
	TaskId  *string                        `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *DescribeLoadBalancersSpecResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeLoadBalancersSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancersSpecData struct {
	ProductName *string                              `json:"ProductName,omitempty" name:"ProductName"`
	Products    []*DescribeLoadBalancersSpecProducts `json:"Products,omitempty" name:"Products"`
	RegionId    *string                              `json:"RegionId,omitempty" name:"RegionId"`
}

type DescribeLoadBalancersSpecProducts struct {
	Architectures []*DescribeLoadBalancersSpecArchitectures `json:"Architectures,omitempty" name:"Architectures"`
	Version       *string                                   `json:"Version,omitempty" name:"Version"`
}

type DescribeLoadBalancersSpecArchitectures struct {
	ArchitectureName *string                                  `json:"Architectures,omitempty" name:"Architectures"`
	ComputeRoles     []*DescribeLoadBalancersSpecComputeRoles `json:"ComputeRoles,omitempty" name:"ComputeRoles"`
	EnginesType      []*interface{}                           `json:"EnginesType,omitempty" name:"EnginesType"`   // UNKNOW TYPE
	NetworkLinks     *interface{}                             `json:"NetworkLinks,omitempty" name:"NetworkLinks"` // UNKNOW TYPE
	SubProductName   *string                                  `json:"SubProductName,omitempty" name:"SubProductName"`
}

type DescribeLoadBalancersSpecComputeRoles struct {
	ComputeName *string                             `json:"ComputeName,omitempty" name:"ComputeName"`
	Standards   *DescribeLoadBalancersSpecStandards `json:"Standards,omitempty" name:"Standards"`
}

type DescribeLoadBalancersSpecStandards struct {
	AttachDisk []*DescribeLoadBalancersSpecAttachDisk `json:"AttachDisk,omitempty" name:"AttachDisk"`
	CpuRam     []*DescribeLoadBalancersSpecCpuRam     `json:"CpuRam,omitempty" name:"CpuRam"`
}

type DescribeLoadBalancersSpecAttachDisk struct {
	BasicIops     *string `json:"BasicIops,omitempty" name:"BasicIops"`
	DiskMax       *int    `json:"DiskMax,omitempty" name:"DiskMax"`
	DiskMaxExpand *int    `json:"DiskMaxExpand,omitempty" name:"DiskMaxExpand"`
	DiskMin       *int    `json:"DiskMin,omitempty" name:"DiskMin"`
	DiskName      *string `json:"DiskName,omitempty" name:"DiskName"`
	DiskStep      *int    `json:"DiskStep,omitempty" name:"DiskStep"`
	DiskUnit      *string `json:"DiskUnit,omitempty" name:"DiskUnit"`
	DiskValue     *string `json:"DiskValue,omitempty" name:"DiskValue"`
}

type DescribeLoadBalancersSpecCpuRam struct {
	CPU         *int    `json:"CPU,omitempty" name:"CPU"`
	Name        *string `json:"Name,omitempty" name:"Name"`
	PaasGoodsId *int    `json:"PaasGoodsId,omitempty" name:"PaasGoodsId"`
	RAM         *int    `json:"RAM,omitempty" name:"RAM"`
}

type CreateLoadBalancerRequest struct {
	*cdshttp.BaseRequest
	RegionId     *string                  `json:"RegionId,omitempty" name:"RegionId"`
	VdcId        *string                  `json:"VdcId,omitempty" name:"VdcId"`
	BasePipeId   *string                  `json:"BasePipeId,omitempty" name:"BasePipeId"`
	InstanceName *string                  `json:"InstanceName,omitempty" name:"InstanceName"`
	PaasGoodsId  *int                     `json:"PaasGoodsId,omitempty" name:"PaasGoodsId"`
	Ips          []*CreateLoadBalancerIps `json:"Ips,omitempty" name:"Ips"`
	Amount       *int                     `json:"Amount,omitempty" name:"Amount"`
}

func (request *CreateLoadBalancerRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *CreateLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type CreateLoadBalancerIps struct {
	PipeType  *string `json:"PipeType,omitempty" name:"PipeType"` //UNKNOW TYPE
	PipeId    *string `json:"PipeId,omitempty" name:"PipeId"`
	SegmentId *string `json:"SegmentId,omitempty" name:"SegmentId"`
}

type CreateLoadBalancerResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code,omitempty" name:"Code"`
	Data    *map[string]interface{} `json:"Data,omitempty" name:"Data"` //UNKNOW TYPE
	Message *string                 `json:"Message,omitempty" name:"Message"`
	TaskId  *string                 `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *CreateLoadBalancerResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancersRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid" name:"InstanceUuid"`
	IP           *string `json:"IP,omitempty" name:"IP"`
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (request *DescribeLoadBalancersRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeLoadBalancersResponse struct {
	*cdshttp.BaseResponse
	Code    *string                      `json:"Code,omitempty" name:"Code"`
	Data    []*DescribeLoadBalancersData `json:"Data,omitempty" name:"Data"`
	Message *string                      `json:"Message,omitempty" name:"Message"`
}

func (response *DescribeLoadBalancersResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancersData struct {
	CloneServices   *interface{}                 `json:"CloneServices,omitempty" name:"CloneServices"` //UNKNOW TYPE
	Cpu             *int                         `json:"Cpu,omitempty" name:"Cpu"`
	CreatedTime     *string                      `json:"CreatedTime,omitempty" name:"CreatedTime"`
	DisplayName     *string                      `json:"DisplayName,omitempty" name:"DisplayName"`
	IP              *string                      `json:"IP,omitempty" name:"IP"`
	InstanceName    *string                      `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceUuid    *string                      `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	LinkType        *string                      `json:"LinkType,omitempty" name:"LinkType"`
	LinkTypeStr     *string                      `json:"LinkTypeStr,omitempty" name:"LinkTypeStr"`
	MasterInfo      *string                      `json:"MasterInfo,omitempty" name:"MasterInfo"`
	Port            *int                         `json:"Port,omitempty" name:"Port"`
	Ram             *int                         `json:"Ram,omitempty" name:"Ram"`
	RegionId        *string                      `json:"RegionId,omitempty" name:"RegionId"`
	RelationService *interface{}                 `json:"RelationService,omitempty" name:"RelationService"` //UNKNOW TYPE
	ResourceId      *string                      `json:"ResourceId,omitempty" name:"ResourceId"`
	RoGroups        *interface{}                 `json:"RoGroups,omitempty" name:"RoGroups"` //UNKNOW TYPE
	Status          *string                      `json:"Status,omitempty" name:"Status"`
	StatusStr       *string                      `json:"StatusStr,omitempty" name:"StatusStr"`
	SubProductName  *string                      `json:"SubProductName,omitempty" name:"SubProductName"`
	VdcId           *string                      `json:"VdcId,omitempty" name:"VdcId"`
	VdcName         *string                      `json:"VdcName,omitempty" name:"VdcName"`
	Version         *string                      `json:"Version,omitempty" name:"Version"`
	ProjectName     *string                      `json:"ProjectName" name:"ProjectName"`
	Vips            []*DescribeLoadBalancersVips `json:"Vips,omitempty" name:"Vips"`
}

type DescribeLoadBalancersVips struct {
	IP   string `json:"IP,omitempty" name:"IP"`
	Type string `json:"Type,omitempty" name:"Type"`
}

type DescribeLoadBalancersModifySpecRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

func (request *DescribeLoadBalancersModifySpecRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeLoadBalancersModifySpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeLoadBalancersModifySpecResponse struct {
	*cdshttp.BaseResponse
	Code    *string                              `json:"Code,omitempty" name:"Code"`
	Data    *DescribeLoadBalancersModifySpecData `json:"Data,omitempty" name:"Data"`
	Message *string                              `json:"Message,omitempty" name:"Message"`
	TaskId  *string                              `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *DescribeLoadBalancersModifySpecResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeLoadBalancersModifySpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancersModifySpecData struct {
	AttachDisk  []*DescribeLoadBalancersModifySpecAttachDisk `json:"AttachDisk,omitempty" name:"AttachDisk"`
	CpuRam      []*DescribeLoadBalancersModifySpecCpuRam     `json:"CpuRam,omitempty" name:"CpuRam"`
	ProductName *string                                      `json:"ProductName,omitempty" name:"ProductName"`
	RegionId    *string                                      `json:"RegionId,omitempty" name:"RegionId"`
}

type DescribeLoadBalancersModifySpecAttachDisk struct {
	BasicIops     *string `json:"BasicIops,omitempty" name:"BasicIops"`
	DiskMax       *int    `json:"DiskMax,omitempty" name:"DiskMax"`
	DiskMaxExpand *int    `json:"DiskMaxExpand,omitempty" name:"DiskMaxExpand"`
	DiskMin       *int    `json:"DiskMin,omitempty" name:"DiskMin"`
	DiskName      *string `json:"DiskName,omitempty" name:"DiskName"`
	DiskStep      *int    `json:"DiskStep,omitempty" name:"DiskStep"`
	DiskUnit      *string `json:"DiskUnit,omitempty" name:"DiskUnit"`
	DiskValue     *string `json:"DiskValue,omitempty" name:"DiskValue"`
}

type DescribeLoadBalancersModifySpecCpuRam struct {
	CPU         *int    `json:"CPU,omitempty" name:"CPU"`
	N           *string `json:"N,omitempty" name:"N"`
	PaasGoodsId *int    `json:"PaasGoodsId" name:"PaasGoodsId"`
	RAM         *int    `json:"RAM,omitempty" name:"RAM"`
}

type ModifyLoadBalancerInstanceSpecRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	PaasGoodsId  *int    `json:"PaasGoodsId,omitempty" name:"PaasGoodsId"`
}

func (request *ModifyLoadBalancerInstanceSpecRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *ModifyLoadBalancerInstanceSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type ModifyLoadBalancerInstanceSpecResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"` // UNKNOW TYPE
	Message *string      `json:"Message,omitempty" name:"Message"`
	TaskId  *string      `json:"TaksId,omitempty" name:"TaksId"`
}

func (response *ModifyLoadBalancerInstanceSpecResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *ModifyLoadBalancerInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DeleteLoadBalancerRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

func (request *DeleteLoadBalancerRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DeleteLoadBalancerResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"` // UNKNOW TYPE
	Message *string      `json:"Message,omitempty" name:"Message"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *DeleteLoadBalancerResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeCACertificatesRequest struct {
	*cdshttp.BaseRequest
}

func (request *DescribeCACertificatesRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeCACertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeCACertificatesResponse struct {
	*cdshttp.BaseResponse
	Code    *string                       `json:"Code,omitempty" name:"Code"`
	Message *string                       `json:"Message,omitempty" name:"Message"`
	Data    []*DescribeCACertificatesData `json:"Data,omitempty" name:"Data"`
}

func (response *DescribeCACertificatesResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeCACertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeCACertificatesData struct {
	Brand           *string `json:"Brand,omitempty" name:"Brand"`
	CertificateId   *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`
	Domain          *string `json:"Domain,omitempty" name:"Domain"`
	StartTime       *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime         *string `json:"EndTime,omitempty" name:"EndTime"`
	Organization    *string `json:"Organization,omitempty" name:"Organization"`
	Valid           *int    `json:"Valid,omitempty" name:"Valid"`
}

type DescribeCACertificateRequest struct {
	*cdshttp.BaseRequest
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (request *DescribeCACertificateRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeCACertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeCACertificateResponse struct {
	*cdshttp.BaseResponse
	Code    *string                    `json:"Code,omitempty" name:"Code"`
	Data    *DescribeCACertificateData `json:"Data,omitempty" name:"Data"`
	Message *string                    `json:"Message,omitempty" name:"Message"`
	TaskId  *string                    `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *DescribeCACertificateResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeCACertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeCACertificateData struct {
	Brand           *string `json:"Brand,omitempty" name:"Brand"`
	Certificate     *string `json:"Certificate,omitempty" name:"Certificate"`
	CertificateId   *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`
	Domain          *string `json:"Domain,omitempty" name:"Domain"`
	StartTime       *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime         *string `json:"EndTime,omitempty" name:"EndTime"`
	Organization    *string `json:"Organization,omitempty" name:"Organization"`
	PrivateKey      *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	PublicKey       *string `json:"PublicKey,omitempty" name:"PublicKey"`
	Valid           *int    `json:"Valid,omitempty" name:"Valid"`
}

type DeleteCACertificateRequest struct {
	*cdshttp.BaseRequest
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (request *DeleteCACertificateRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DeleteCACertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DeleteCACertificateResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"` // UNKNOW TYPE
	Message *string      `json:"Message,omitempty" name:"Message"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *DeleteCACertificateResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DeleteCACertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type UploadCACertificateRequest struct {
	*cdshttp.BaseRequest
	Certificate     *string `json:"Certificate,omitempty" name:"Certificate"`
	PrivateKey      *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`
}

func (request *UploadCACertificateRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *UploadCACertificateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type UploadCACertificateResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"` // UNKNOW TYPE
	Message *string      `json:"Message,omitempty" name:"Message"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *UploadCACertificateResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *UploadCACertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancerStrategysRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
}

func (request *DescribeLoadBalancerStrategysRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *DescribeLoadBalancerStrategysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type DescribeLoadBalancerStrategysResponse struct {
	*cdshttp.BaseResponse
	Code    *string                            `json:"Code,omitempty" name:"Code"`
	Data    *DescribeLoadBalancerStrategysData `json:"Data,omitempty" name:"Data"`
	Message *string                            `json:"Message,omitempty" name:"Message"`
	TaskId  *string                            `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *DescribeLoadBalancerStrategysResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *DescribeLoadBalancerStrategysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}

type DescribeLoadBalancerStrategysData struct {
	HttpListeners []*DescribeLoadBalancerStrategysHttpListeners `json:"HttpListeners,omitempty" name:"HttpListeners"`
	TcpListeners  []*DescribeLoadBalancerStrategysTcpListeners  `json:"TcpListeners,omitempty" name:"TcpListeners"`
}

type DescribeLoadBalancerStrategysHttpListeners struct {
	AclWhiteList       []*string                                      `json:"AclWhiteList,omitempty" name:"AclWhiteList"`
	BackendServer      []*DescribeLoadBalancerStrategysBackendServer  `json:"BackendServer,omitempty" name:"BackendServer"`
	CertificateIds     []*DescribeLoadBalancerStrategysCertificateIds `json:"CertificateIds,omitempty" name:"CertificateIds"`
	ClientTimeout      *string                                        `json:"ClientTimeout,omitempty" name:"ClientTimeout"`
	ClientTimeoutUnit  *string                                        `json:"ClientTimeoutUnit,omitempty" name:"ClientTimeoutUnit"`
	ConnectTimeout     *string                                        `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`
	ConnectTimeoutUnit *string                                        `json:"ConnectTimeoutUnit,omitempty" name:"ConnectTimeoutUnit"`
	ListenerMode       *string                                        `json:"ListenerMode,omitempty" name:"ListenerMode"`
	ListenerName       *string                                        `json:"ListenerName,omitempty" name:"ListenerName"`
	ListenerPort       *int                                           `json:"ListenerPort,omitempty" name:"ListenerPort"`
	MaxConn            *int                                           `json:"MaxConn,omitempty" name:"MaxConn"`
	Scheduler          *string                                        `json:"Scheduler,omitempty" name:"Scheduler"`
	ServerTimeout      *string                                        `json:"ServerTimeout,omitempty" name:"ServerTimeout"`
	ServerTimeoutUnit  *string                                        `json:"ServerTimeoutUnit,omitempty" name:"ServerTimeoutUnit"`
	StickySession      *string                                        `json:"StickySession"`
}

type DescribeLoadBalancerStrategysTcpListeners struct {
	AclWhiteList       []*string                                     `json:"AclWhiteList,omitempty" name:"AclWhiteList"`
	BackendServer      []*DescribeLoadBalancerStrategysBackendServer `json:"BackendServer,omitempty" name:"BackendServer"`
	ClientTimeout      *string                                       `json:"ClientTimeout,omitempty" name:"ClientTimeout"`
	ClientTimeoutUnit  *string                                       `json:"ClientTimeoutUnit,omitempty" name:"ClientTimeoutUnit"`
	ConnectTimeout     *string                                       `json:"ConnectTimeout,omitempty" name:"ConnectTimeout"`
	ConnectTimeoutUnit *string                                       `json:"ConnectTimeoutUnit,omitempty" name:"ConnectTimeoutUnit"`
	ListenerMode       *string                                       `json:"ListenerMode,omitempty" name:"ListenerMode"`
	ListenerName       *string                                       `json:"ListenerName,omitempty" name:"ListenerName"`
	ListenerPort       *int                                          `json:"ListenerPort,omitempty" name:"ListenerPort"`
	MaxConn            *int                                          `json:"MaxConn,omitempty" name:"MaxConn"`
	Scheduler          *string                                       `json:"Scheduler,omitempty" name:"Scheduler"`
	ServerTimeout      *string                                       `json:"ServerTimeout,omitempty" name:"ServerTimeout"`
	ServerTimeoutUnit  *string                                       `json:"ServerTimeoutUnit,omitempty" name:"ServerTimeoutUnit"`
}

type DescribeLoadBalancerStrategysBackendServer struct {
	IP      *string `json:"IP,omitempty" name:"IP"`
	MaxConn *int    `json:"MaxConn,omitempty" name:"MaxConn"`
	Port    *int    `json:"Port,omitempty" name:"Port"`
	Weight  *string `json:"Weight,omitempty" name:"Weight"`
}

type DescribeLoadBalancerStrategysCertificateIds struct {
	CertificateId   *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateName *string `json:"CertificateName,omitempty" name:"CertificateName"`
}

type ModifyLoadBalancerStrategysRequest struct {
	*cdshttp.BaseRequest
	InstanceUuid  *string                                       `json:"InstanceUuid,omitempty" name:"InstanceUuid"`
	HttpListeners []*DescribeLoadBalancerStrategysHttpListeners `json:"HttpListeners,omitempty" name:"HttpListeners"`
	TcpListeners  []*DescribeLoadBalancerStrategysTcpListeners  `json:"TcpListeners,omitempty" name:"TcpListeners"`
}

func (request *ModifyLoadBalancerStrategysRequest) ToJsonString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}

func (request *ModifyLoadBalancerStrategysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), request)
}

type ModifyLoadBalancerStrategysResponse struct {
	*cdshttp.BaseResponse
	Code    *string      `json:"Code,omitempty" name:"Code"`
	Data    *interface{} `json:"Data,omitempty" name:"Data"` // UNKNOW TYPE
	Message *string      `json:"Message,omitempty" name:"Message"`
	TaskId  *string      `json:"TaskId,omitempty" name:"TaskId"`
}

func (response *ModifyLoadBalancerStrategysResponse) ToJsonString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response *ModifyLoadBalancerStrategysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), response)
}
