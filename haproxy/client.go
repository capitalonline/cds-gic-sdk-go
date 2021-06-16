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
func (client *Client) DescribeZones(request *DescribeZonesRequest) (*DescribeZonesResponse, error) {
	if request == nil {
		request = NewDescribeZonesRequest()
	}
	response := NewDescribeZonesResponse()
	return response, client.Send(request, response)
}

// 获取某个站点支持的Haproxy产品类型以及规格
func (client *Client) DescribeLoadBalancersSpec(request *DescribeLoadBalancersSpecRequest) (*DescribeLoadBalancersSpecResponse, error) {
	if request == nil {
		request = NewDescribeLoadBalancersSpecRequest()
	}
	response := NewDescribeLoadBalancersSpecResponse()
	return response, client.Send(request, response)
}

// 创建Ha实例
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (*CreateLoadBalancerResponse, error) {
	if request == nil {
		request = NewCreateLoadBalancerRequest()
	}
	response := NewCreateLoadBalancerResponse()
	return response, client.Send(request, response)
}

// 获取Ha实例列表
func (client *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (*DescribeLoadBalancersResponse, error) {
	if request == nil {
		request = NewDescribeLoadBalancersRequest()
	}
	response := NewDescribeLoadBalancersResponse()
	return response, client.Send(request, response)
}

// 获取实例Ha的配置变更所支持的规格
func (client *Client) DescribeLoadBalancersModifySpec(request *DescribeLoadBalancersModifySpecRequest) (*DescribeLoadBalancersModifySpecResponse, error) {
	if request == nil {
		request = NewDescribeLoadBalancersModifySpecRequest()
	}
	response := NewDescribeLoadBalancersModifySpecResponse()
	return response, client.Send(request, response)
}

// 修改实例规格
func (client *Client) ModifyLoadBalancerInstanceSpec(request *ModifyLoadBalancerInstanceSpecRequest) (*ModifyLoadBalancerInstanceSpecResponse, error) {
	if request == nil {
		request = NewModifyLoadBalancerInstanceSpecRequest()
	}
	response := NewModifyLoadBalancerInstanceSpecResponse()
	return response, client.Send(request, response)
}

// 删除Ha实例
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (*DeleteLoadBalancerResponse, error) {
	if request == nil {
		request = NewDeleteLoadBalancerRequest()
	}
	response := NewDeleteLoadBalancerResponse()
	return response, client.Send(request, response)
}

// 获取用户的证书列表
func (client *Client) DescribeCACertificates(request *DescribeCACertificatesRequest) (*DescribeCACertificatesResponse, error) {
	if request == nil {
		request = NewDescribeCACertificatesRequest()
	}
	response := NewDescribeCACertificatesResponse()
	return response, client.Send(request, response)
}

// 获取用户的证书详
func (client *Client) DescribeCACertificate(request *DescribeCACertificateRequest) (*DescribeCACertificateResponse, error) {
	if request == nil {
		request = NewDescribeCACertificateRequest()
	}
	response := NewDescribeCACertificateResponse()
	return response, client.Send(request, response)
}

// 删除证书
func (client *Client) DeleteCACertificate(request *DeleteCACertificateRequest) (*DeleteCACertificateResponse, error) {
	if request == nil {
		request = NewDeleteCACertificateRequest()
	}
	response := NewDeleteCACertificateResponse()
	return response, client.Send(request, response)
}

// 添加证书
func (client *Client) UploadCACertificate(request *UploadCACertificateRequest) (*UploadCACertificateResponse, error) {
	if request == nil {
		request = NewUploadCACertificateRequest()
	}
	response := NewUploadCACertificateResponse()
	return response, client.Send(request, response)
}

// 获取Ha实例的当前监听的策略配置列表
func (client *Client) DescribeLoadBalancerStrategys(request *DescribeLoadBalancerStrategysRequest) (*DescribeLoadBalancerStrategysResponse, error) {
	if request == nil {
		request = NewDescribeLoadBalancerStrategysRequest()
	}
	response := NewDescribeLoadBalancerStrategysResponse()
	return response, client.Send(request, response)
}

// 修改（删除、修改、添加）Ha实例的当前监听的策略配置列表
func (client *Client) ModifyLoadBalancerStrategys(request *ModifyLoadBalancerStrategysRequest) (*ModifyLoadBalancerStrategysResponse, error) {
	if request == nil {
		request = NewModifyLoadBalancerStrategysRequest()
	}
	response := NewModifyLoadBalancerStrategysResponse()
	return response, client.Send(request, response)
}

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

func NewCreateLoadBalancerRequest() *CreateLoadBalancerRequest {
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
