package instance

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
)

const ApiVersion = "2019-08-08"

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
func NewAddInstanceRequest() (request *AddInstanceRequest) {
	request = &AddInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", ApiVersion, "CreateInstance")
	return
}

func NewAddInstanceResponse() (response *AddInstanceResponse) {
	response = &AddInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DescribeInstances")
	return
}
func NewDescribeInstanceResponse() (response *DescribeInstanceReponse) {
	response = &DescribeInstanceReponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DeleteInstance")
	return
}
func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
	request = &ModifyInstanceNameRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyInstanceName")
	return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
	response = &ModifyInstanceNameResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyInstanceSpecRequest() (request *ModifyInstanceSpecRequest) {
	request = &ModifyInstanceSpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyInstanceSpec")
	return
}

func NewModifyInstanceSpecResponse() (response *ModifyInstanceSpecResponse) {
	response = &ModifyInstanceSpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewCreateDiskRequest() (request *CreateDiskRequest) {
	request = &CreateDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "CreateDisk")
	return
}
func NewCreateDiskResponse() (response *CreateDiskResponse) {
	response = &CreateDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ResizeDisk")
	return
}
func NewResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteDiskRequest() (request *DeleteDiskRequest) {
	request = &DeleteDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DeleteDisk")
	return
}
func NewDeleteDiskResponse() (response *DeleteDiskResponse) {
	response = &DeleteDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyIpRequest() (request *ModifyIpRequest) {
	request = &ModifyIpRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyIpAddress")
	return
}
func NewModifyIpResponse() (response *ModifyIpResponse) {
	response = &ModifyIpResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

// Create Instance
func (c *Client) CreateInstance(request *AddInstanceRequest) (response *AddInstanceResponse, err error) {
	if request == nil {
		request = NewAddInstanceRequest()
	}
	response = NewAddInstanceResponse()
	err = c.Send(request, response)
	return
}

// Describe Instance
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceReponse, err error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	response = NewDescribeInstanceResponse()
	err = c.Send(request, response)
	return
}

// Delete Instance
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	response = NewDeleteInstanceResponse()
	err = c.Send(request, response)
	return
}

// Modify Instance Name
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNameRequest()
	}
	response = NewModifyInstanceNameResponse()
	err = c.Send(request, response)
	return
}

// ModifyInstanceSpec cpu, ram
func (c *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (response *ModifyInstanceSpecResponse, err error) {
	if request == nil {
		request = NewModifyInstanceSpecRequest()
	}
	response = NewModifyInstanceSpecResponse()
	err = c.Send(request, response)
	return
}

// Create Disk
func (c *Client) CreateDisk(request *CreateDiskRequest) (response *CreateDiskResponse, err error) {
	if request == nil {
		request = NewCreateDiskRequest()
	}
	response = NewCreateDiskResponse()
	err = c.Send(request, response)
	return
}

// Resize Disk
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	if request == nil {
		request = NewResizeDiskRequest()
	}
	response = NewResizeDiskResponse()
	err = c.Send(request, response)
	return
}

// Delete Disk
func (c *Client) DeleteDisk(request *DeleteDiskRequest) (response *DeleteDiskResponse, err error) {
	if request == nil {
		request = NewDeleteDiskRequest()
	}
	response = NewDeleteDiskResponse()
	err = c.Send(request, response)
	return
}

// Delete Disk
func (c *Client) ModifyIpAddress(request *ModifyIpRequest) (response *ModifyIpResponse, err error) {
	if request == nil {
		request = NewModifyIpRequest()
	}
	response = NewModifyIpResponse()
	err = c.Send(request, response)
	return
}

func NewExtendSystemDiskRequest() (request *ExtendSystemDiskRequest) {
	request = &ExtendSystemDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ExtendSystemDisk")
	return
}

func NewExtendSystemDiskResponse() (response *ExtendSystemDiskResponse) {
	response = &ExtendSystemDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func (c *Client) ExtendSystemDisk(request *ExtendSystemDiskRequest) (response *ExtendSystemDiskResponse, err error) {
	if request == nil {
		request = NewExtendSystemDiskRequest()
	}
	response = NewExtendSystemDiskResponse()
	err = c.Send(request, response)
	return
}

func NewResetInstancesPasswordRequest() *ResetInstancesPasswordRequest {
	request := &ResetInstancesPasswordRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ResetInstancesPassword")
	return request
}

func NewResetInstancesPasswordResponse() *ResetInstancesPasswordResponse {
	response := &ResetInstancesPasswordResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
	if request == nil {
		request = NewResetInstancesPasswordRequest()
	}
	response = NewResetInstancesPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewResetImageRequest() *ResetImageRequest {
	request := &ResetImageRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ResetImage")
	return request
}

func NewResetImageResponse() *ResetImageResponse {
	response := &ResetImageResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) ResetImage(request *ResetImageRequest) (response *ResetImageResponse, err error) {
	if request == nil {
		request = NewResetImageRequest()
	}
	response = NewResetImageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceChargeTypeRequest() *ModifyInstanceChargeTypeRequest {
	request := &ModifyInstanceChargeTypeRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyInstanceChargeType")
	return request
}

func NewModifyInstanceChargeTypeResponse() *ModifyInstanceChargeTypeResponse {
	response := &ModifyInstanceChargeTypeResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) ModifyInstanceChargeType(request *ModifyInstanceChargeTypeRequest) (response *ModifyInstanceChargeTypeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceChargeTypeRequest()
	}
	response = NewModifyInstanceChargeTypeResponse()
	err = c.Send(request, response)
	return
}

func NewStopInstanceRequest() *StopInstanceRequest {
	request := &StopInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "StopInstance")
	return request
}

func NewStopInstanceResponse() *StopInstanceResponse {
	response := &StopInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
	if request == nil {
		request = NewStopInstanceRequest()
	}
	response = NewStopInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewRebootInstanceRequest() *RebootInstanceRequest {
	request := &RebootInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "RebootInstance")
	return request
}

func NewRebootInstanceResponse() *RebootInstanceResponse {
	response := &RebootInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) RebootInstance(request *RebootInstanceRequest) (response *RebootInstanceResponse, err error) {
	if request == nil {
		request = NewRebootInstanceRequest()
	}
	response = NewRebootInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstanceRequest() *StartInstanceRequest {
	request := &StartInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "StartInstance")
	return request
}

func NewStartInstanceResponse() *StartInstanceResponse {
	response := &StartInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) StartInstance(request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
	if request == nil {
		request = NewStartInstanceRequest()
	}
	response = NewStartInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstancesRequest() *StartInstancesRequest {
	request := &StartInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "StartInstances")
	return request
}

func NewStartInstancesResponse() *StartInstancesResponse {
	response := &StartInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	if request == nil {
		request = NewStartInstancesRequest()
	}
	response = NewStartInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewStopInstancesRequest() *StopInstancesRequest {
	request := &StopInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "StopInstances")
	return request
}

func NewStopInstancesResponse() *StopInstancesResponse {
	response := &StopInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	if request == nil {
		request = NewStopInstancesRequest()
	}
	response = NewStopInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewRebootInstancesRequest() *RebootInstancesRequest {
	request := &RebootInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "RebootInstances")
	return request
}

func NewRebootInstancesResponse() *RebootInstancesResponse {
	response := &RebootInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	if request == nil {
		request = NewRebootInstancesRequest()
	}
	response = NewRebootInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceMonitorRequest() *DescribeInstanceMonitorRequest {
	request := &DescribeInstanceMonitorRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DescribeInstanceMonitor")
	return request
}

func NewDescribeInstanceMonitorResponse() *DescribeInstanceMonitorResponse {
	response := &DescribeInstanceMonitorResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return response
}

func (c *Client) DescribeInstanceMonitor(request *DescribeInstanceMonitorRequest) (response *DescribeInstanceMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceMonitorRequest()
	}
	response = NewDescribeInstanceMonitorResponse()
	err = c.Send(request, response)
	return
}
