package haproxy

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
)

const (
	ApiVersion     = "2019-08-08"
	HaproxyService = "lb"
)

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

// 获取负载均衡Ha支持的区域
func NewDescribeZonesRequest() *DescribeZonesRequest {
	request := &DescribeZonesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeZones")
	return request
}

func NewDescribeZonesResponse() *DescribeZonesResponse {
	return &DescribeZonesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 获取某个站点支持的Haproxy产品类型以及规格
func NewDescribeLoadBalancersSpecRequest() *DescribeLoadBalancersSpecRequest {
	request := &DescribeLoadBalancersSpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeLoadBalancersSpec")
	return request
}

func NewDescribeLoadBalancersSpecResponse() *DescribeLoadBalancersSpecResponse {
	return &DescribeLoadBalancersSpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 创建Ha实例
func NewCreateLoadBalancer() *CreateLoadBalancerRequest {
	request := &CreateLoadBalancerRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "CreateLoadBalancer")
	return request
}

func NewCreateLoadBalancerResponse() *CreateLoadBalancerResponse {
	return &CreateLoadBalancerResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 获取Ha实例列表
func NewDescribeLoadBalancersRequest() *DescribeLoadBalancersRequest {
	request := &DescribeLoadBalancersRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeLoadBalancers")
	return request
}

func NewDescribeLoadBalancersResponse() *DescribeLoadBalancersResponse {
	return &DescribeLoadBalancersResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 获取实例Ha的配置变更所支持的规格
func NewDescribeLoadBalancersModifySpecRequest() *DescribeLoadBalancersModifySpecRequest {
	request := &DescribeLoadBalancersModifySpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeLoadBalancersModifySpec")
	return request
}

func NewDescribeLoadBalancersModifySpecResponse() *DescribeLoadBalancersModifySpecResponse {
	return &DescribeLoadBalancersModifySpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 修改实例规格
func NewModifyLoadBalancerInstanceSpecRequest() *ModifyLoadBalancerInstanceSpecRequest {
	request := &ModifyLoadBalancerInstanceSpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "ModifyLoadBalancerInstanceSpec")
	return request
}

func NewModifyLoadBalancerInstanceSpecResponse() *ModifyLoadBalancerInstanceSpecResponse {
	return &ModifyLoadBalancerInstanceSpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 删除Ha实例
func NewDeleteLoadBalancerRequest() *DeleteLoadBalancerRequest {
	request := &DeleteLoadBalancerRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DeleteLoadBalancer")
	return request
}

func NewDeleteLoadBalancerResponse() *DeleteLoadBalancerResponse {
	return &DeleteLoadBalancerResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 获取用户的证书列表
func NewDescribeCACertificatesRequest() *DescribeCACertificatesRequest {
	request := DescribeCACertificatesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeCACertificates")
	return &request
}

func NewDescribeCACertificatesResponse() *DescribeCACertificatesResponse {
	return &DescribeCACertificatesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 获取用户的证书详情
func NewDescribeCACertificateRequest() *DescribeCACertificateRequest {
	request := &DescribeCACertificateRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeCACertificate")
	return request
}

func NewDescribeCACertificateResponse() *DescribeCACertificateResponse {
	return &DescribeCACertificateResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 删除证书
func NewDeleteCACertificateRequest() *DeleteCACertificateRequest {
	request := &DeleteCACertificateRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DeleteCACertificate")
	return request
}

func NewDeleteCACertificateResponse() *DeleteCACertificateResponse {
	return &DeleteCACertificateResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 添加证书
func NewUploadCACertificateRequest() *UploadCACertificateRequest {
	request := &UploadCACertificateRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "UploadCACertificate")
	return request
}

func NewUploadCACertificateResponse() *UploadCACertificateResponse {
	return &UploadCACertificateResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 获取Ha实例的当前监听的策略配置列表
func NewDescribeLoadBalancerStrategysRequest() *DescribeLoadBalancerStrategysRequest {
	request := &DescribeLoadBalancerStrategysRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "DescribeLoadBalancerStrategys")
	return request
}

func NewDescribeLoadBalancerStrategysResponse() *DescribeLoadBalancerStrategysResponse {
	return &DescribeLoadBalancerStrategysResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

// 修改（删除、修改、添加）Ha实例的当前监听的策略配置列表
func NewModifyLoadBalancerStrategysRequest() *ModifyLoadBalancerStrategysRequest {
	request := &ModifyLoadBalancerStrategysRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(HaproxyService, ApiVersion, "ModifyLoadBalancerStrategys")
	return request
}

func NewModifyLoadBalancerStrategysResponse() *ModifyLoadBalancerStrategysResponse {
	return &ModifyLoadBalancerStrategysResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}
