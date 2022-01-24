package mongodb

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
)

const (
	ApiVersion   = "2019-08-08"
	MangoService = "mongodb/v1"
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

func NewDescribeZonesRequest() *DescribeZonesRequest {
	request := &DescribeZonesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MangoService, ApiVersion, "DescribeZones")
	return request
}

func NewDescribeZonesResponse() *DescribeZonesResponse {
	return &DescribeZonesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (*DescribeZonesResponse, error) {
	if request == nil {
		request = NewDescribeZonesRequest()
	}
	response := NewDescribeZonesResponse()
	return response, client.Send(request, response)
}

func NewDescribeSpecInfoRequest() *DescribeSpecInfoRequest {
	request := &DescribeSpecInfoRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MangoService, ApiVersion, "DescribeSpecInfo")
	return request
}

func NewDescribeSpecInfoResponse() *DescribeSpecInfoResponse {
	return &DescribeSpecInfoResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) DescribeSpecInfo(request *DescribeSpecInfoRequest) (*DescribeSpecInfoResponse, error) {
	if request == nil {
		request = NewDescribeSpecInfoRequest()
	}
	response := NewDescribeSpecInfoResponse()

	return response, client.Send(request, response)
}

func NewCreateDBInstanceRequest() *CreateDBInstanceRequest {
	request := &CreateDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MangoService, ApiVersion, "CreateDBInstance")
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

	request.Init().WithApiInfo(MangoService, ApiVersion, "DescribeDBInstances")
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

	request.Init().WithApiInfo(MangoService, ApiVersion, "DeleteDBInstance")
	return request
}

func NewDeleteDBInstanceReponse() *DeleteDBInstanceResponse {
	return &DeleteDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (*DeleteDBInstanceResponse, error) {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	response := NewDeleteDBInstanceReponse()

	return response, client.Send(request, response)
}
