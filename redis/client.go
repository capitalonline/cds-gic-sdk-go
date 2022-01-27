package redis

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
)

const (
	ApiVersion   = "2019-08-08"
	RedisService = "redis"
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

func NewDescribeRegionsRequest() *DescribeRegionsRequest {
	request := &DescribeRegionsRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(RedisService, ApiVersion, "DescribeRegins")
	return request
}

func NewDescribeRegionsResponse() *DescribeRegionsResponse {
	return &DescribeRegionsResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) DescribeRegins(request *DescribeRegionsRequest) (*DescribeRegionsResponse, error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response := NewDescribeRegionsResponse()
	return response, client.Send(request, response)
}

func NewDescribeAvailableDBConfigRequest() *DescribeAvailableDBConfigRequest {
	request := &DescribeAvailableDBConfigRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}

	request.Init().WithApiInfo(RedisService, ApiVersion, "DescribeAvailableDBConfig")
	return request
}

func NewDescribeAvailableDBConfigResponse() *DescribeAvailableDBConfigResponse {
	return &DescribeAvailableDBConfigResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) DescribeAvailableDBConfig(request *DescribeAvailableDBConfigRequest) (*DescribeAvailableDBConfigResponse, error) {
	if request == nil {
		request = NewDescribeAvailableDBConfigRequest()
	}

	response := NewDescribeAvailableDBConfigResponse()
	return response, client.Send(request, response)
}

func NewCreateDBInstanceRequest() *CreateDBInstanceRequest {
	request := &CreateDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}

	request.Init().WithApiInfo(RedisService, ApiVersion, "CreateDBInstance")
	return request
}

func NewCreateDBInstanceResponse() *CreateDBInstanceResponse {
	return &CreateDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (*CreateDBInstanceResponse, error) {
	if request == nil {
		request = NewCreateDBInstanceRequest()
	}

	response := NewCreateDBInstanceResponse()
	return response, client.Send(request, response)
}

func NewDescribeDBInstancesRequest() *DescribeDBInstancesRequest {
	request := &DescribeDBInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}

	request.Init().WithApiInfo(RedisService, ApiVersion, "DescribeDBInstances")
	return request
}

func NewDescribeDBInstancesResponse() *DescribeDBInstancesResponse {
	return &DescribeDBInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (*DescribeDBInstancesResponse, error) {
	if request == nil {
		request = NewDescribeDBInstancesRequest()
	}
	response := NewDescribeDBInstancesResponse()
	return response, client.Send(request, response)
}

func NewDeleteDBInstanceRequest() *DeleteDBInstanceRequest {
	request := &DeleteDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(RedisService, ApiVersion, "DeleteDBInstance")
	return request
}

func NewDeleteDBInstanceResponse() *DeleteDBInstanceResponse {
	return &DeleteDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (clinet *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (*DeleteDBInstanceResponse, error) {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	response := NewDeleteDBInstanceResponse()
	return response, clinet.Send(request, response)
}
