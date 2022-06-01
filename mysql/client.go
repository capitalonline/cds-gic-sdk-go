package mysql

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
)

const (
	ApiVersion   = "2019-08-08"
	MySqlService = "mysql"
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

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (*DescribeRegionsResponse, error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response := NewDescribeRegionsResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeAvailableDBConfig(request *DescribeAvailableDBConfigRequest) (*DescribeAvailableDBConfigResponse, error) {
	if request == nil {
		request = NewDescribeAvailableDBConfigRequest()
	}
	response := NewDescribeAvailableDBConfigResponse()
	return response, client.Send(request, response)
}

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (*CreateDBInstanceResponse, error) {
	if request == nil {
		request = NewCreateDBInstanceRequest()
	}
	response := NewCreateDBInstanceResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (*DescribeDBInstancesResponse, error) {
	if request == nil {
		request = NewDescribeDBInstancesRequest()
	}
	response := NewDescribeDBInstancesResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeModifiableDBSpec(request *DescribeModifiableDBSpecRequest) (*DescribeModifiableDBSpecResponse, error) {
	if request == nil {
		request = NewDescribeModifiableDBSpecRequest()
	}
	response := NewDescribeModifiableDBSpecResponse()
	return response, client.Send(request, response)
}

func (client *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (*ModifyDBInstanceSpecResponse, error) {
	if request == nil {
		request = NewModifyDBInstanceSpecRequest()
	}
	response := NewModifyDBInstanceSpecResponse()
	return response, client.Send(request, response)
}

func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (*DeleteDBInstanceResponse, error) {
	if request == nil {
		request = NewDeleteDBInstanceRequest()
	}
	response := NewDeleteDBInstanceResponse()
	return response, client.Send(request, response)
}

func (client *Client) CreatePrivilegedAccount(request *CreatePrivilegedAccountRequest) (*CreatePrivilegedAccountResponse, error) {
	if request == nil {
		request = NewCreatePrivilegedAccountRequest()
	}
	response := NewCreatePrivilegedAccountResponse()
	return response, client.Send(request, response)
}

func (client *Client) CreateBackup(request *CreateBackupRequest) (*CreateBackupResponse, error) {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	response := NewCreateBackupResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (*DescribeBackupsResponse, error) {
	if request == nil {
		request = NewDescribeBackupsRequest()
	}
	response := NewDescribeBackupsResponse()
	return response, client.Send(request, response)
}

func (client *Client) DeleteBackup(request *DeleteBackupRequest) (*DeleteBackupResponse, error) {
	if request == nil {
		request = NewDeleteBackupRequest()
	}
	response := NewDeleteBackupResponse()
	return response, client.Send(request, response)
}

func (client *Client) DownloadBackup(request *DownloadBackupRequest) (*DownloadBackupResponse, error) {
	if request == nil {
		request = NewDownloadBackupRequest()
	}
	response := NewDownloadBackupResponse()
	return response, client.Send(request, response)
}

func (client *Client) StartBatchRollback(request *StartBatchRollbackRequest) (*StartBatchRollbackResponse, error) {
	if request == nil {
		request = NewStartBatchRollbackRequest()
	}
	response := NewStartBatchRollbackResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeRollbackRangeTime(request *DescribeRollbackRangeTimeRequest) (*DescribeRollbackRangeTimeResponse, error) {
	if request == nil {
		request = NewDescribeRollbackRangeTimeRequest()
	}
	response := NewDescribeRollbackRangeTimeResponse()
	return response, client.Send(request, response)
}

func (client *Client) StartBatchRollbackToTemporaryDBInstance(request *StartBatchRollbackToTemporaryDBInstanceRequest) (*StartBatchRollbackToTemporaryDBInstanceResponse, error) {
	if request == nil {
		request = NewStartBatchRollbackToTemporaryDBInstanceRequest()
	}
	response := NewStartBatchRollbackToTemporaryDBInstanceResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeTemporaryDBInstances(request *DescribeTemporaryDBInstancesRequest) (*DescribeTemporaryDBInstancesResponse, error) {
	if request == nil {
		request = NewDescribeTemporaryDBInstancesRequest()
	}
	response := NewDescribeTemporaryDBInstancesResponse()
	return response, client.Send(request, response)
}

func (client *Client) RegularizeTemporaryDBInstances(request *RegularizeTemporaryDBInstancesRequest) (*RegularizeTemporaryDBInstancesResponse, error) {
	if request == nil {
		request = NewRegularizeTemporaryDBInstancesRequest()
	}
	response := NewRegularizeTemporaryDBInstancesResponse()
	return response, client.Send(request, response)
}

func (client *Client) DeleteTemporaryDBInstances(request *DeleteTemporaryDBInstancesRequest) (*DeleteTemporaryDBInstancesResponse, error) {
	if request == nil {
		request = NewDeleteTemporaryDBInstancesRequest()
	}
	response := NewDeleteTemporaryDBInstancesResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeAvailableReadOnlyConfig(request *DescribeAvailableReadOnlyConfigRequest) (*DescribeAvailableReadOnlyConfigResponse, error) {
	if request == nil {
		request = NewDescribeAvailableReadOnlyConfigRequest()
	}
	response := NewDescribeAvailableReadOnlyConfigResponse()
	return response, client.Send(request, response)
}

func (client *Client) CreateReadOnlyDBInstance(request *CreateReadOnlyDBInstanceRequest) (*CreateReadOnlyDBInstanceResponse, error) {
	if request == nil {
		request = NewCreateReadOnlyDBInstanceRequest()
	}
	response := NewCreateReadOnlyDBInstanceResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeDBInstancePerformanceRequest(request *DescribeDBInstancePerformanceRequest) (*DescribeDBInstancePerformanceResponse, error) {
	if request == nil {
		request = NewDescribeDBInstancePerformanceRequest()
	}
	response := NewDescribeDBInstancePerformanceResponse()
	return response, client.Send(request, response)
}

func (client *Client) ModifyDBParameter(request *ModifyDBParameterRequest) (*ModifyDBParameterResponse, error) {
	if request == nil {
		request = NewModifyDBParameterRequest()
	}
	response := NewModifyDBParameterResponse()
	return response, client.Send(request, response)
}

func (client *Client) ModifyDbBackupPolicy(request *ModifyDbBackupPolicyRequest) (*ModifyDbBackupPolicyResponse, error) {
	if request == nil {
		request = NewModifyDbBackupPolicyRequest()
	}
	response := NewModifyDbBackupPolicyResponse()
	return response, client.Send(request, response)
}

func (client *Client) ModifyDbPrivilege(request *ModifyDbPrivilegeRequest) (*ModifyDbPrivilegeResponse, error) {
	if request == nil {
		request = NewModifyDbPrivilegeRequest()
	}
	response := NewModifyDbPrivilegeResponse()
	return response, client.Send(request, response)
}

func (client *Client) DescribeDBAccount(request *DescribeDBAccountRequest) (*DescribeDBAccountResponse, error) {
	if request == nil {
		request = NewDescribeDBAccountRequest()
	}
	response := NewDescribeDBAccountResponse()
	return response, client.Send(request, response)
}

func NewDescribeRegionsRequest() *DescribeRegionsRequest {
	request := &DescribeRegionsRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeRegions")
	return request
}

func NewDescribeRegionsResponse() *DescribeRegionsResponse {
	return &DescribeRegionsResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeAvailableDBConfigRequest() *DescribeAvailableDBConfigRequest {
	request := &DescribeAvailableDBConfigRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeAvailableDBConfig")
	return request
}

func NewDescribeAvailableDBConfigResponse() *DescribeAvailableDBConfigResponse {
	return &DescribeAvailableDBConfigResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewCreateDBInstanceRequest() *CreateDBInstanceRequest {
	request := &CreateDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "CreateDBInstance")
	return request
}

func NewCreateDBInstanceResponse() *CreateDBInstanceResponse {
	return &CreateDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeDBInstancesRequest() *DescribeDBInstancesRequest {
	request := &DescribeDBInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeDBInstances")
	return request
}

func NewDescribeDBInstancesResponse() *DescribeDBInstancesResponse {
	return &DescribeDBInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeModifiableDBSpecRequest() *DescribeModifiableDBSpecRequest {
	request := &DescribeModifiableDBSpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeModifiableDBSpec")
	return request
}

func NewDescribeModifiableDBSpecResponse() *DescribeModifiableDBSpecResponse {
	return &DescribeModifiableDBSpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewModifyDBInstanceSpecRequest() *ModifyDBInstanceSpecRequest {
	request := &ModifyDBInstanceSpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "ModifyDBInstanceSpec")
	return request
}

func NewModifyDBInstanceSpecResponse() *ModifyDBInstanceSpecResponse {
	return &ModifyDBInstanceSpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDeleteDBInstanceRequest() *DeleteDBInstanceRequest {
	request := &DeleteDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DeleteDBInstance")
	return request
}

func NewDeleteDBInstanceResponse() *DeleteDBInstanceResponse {
	return &DeleteDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewCreatePrivilegedAccountRequest() *CreatePrivilegedAccountRequest {
	request := &CreatePrivilegedAccountRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "CreatePrivilegedAccount")
	return request
}

func NewCreatePrivilegedAccountResponse() *CreatePrivilegedAccountResponse {
	return &CreatePrivilegedAccountResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewCreateBackupRequest() *CreateBackupRequest {
	request := &CreateBackupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "CreateBackup")
	return request
}

func NewCreateBackupResponse() *CreateBackupResponse {
	return &CreateBackupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeBackupsRequest() *DescribeBackupsRequest {
	request := &DescribeBackupsRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeBackups")
	return request
}

func NewDescribeBackupsResponse() *DescribeBackupsResponse {
	return &DescribeBackupsResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDeleteBackupRequest() *DeleteBackupRequest {
	request := &DeleteBackupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DeleteBackup")
	return request
}

func NewDeleteBackupResponse() *DeleteBackupResponse {
	return &DeleteBackupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDownloadBackupRequest() *DownloadBackupRequest {
	request := &DownloadBackupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DownloadBackup")
	return request
}

func NewDownloadBackupResponse() *DownloadBackupResponse {
	return &DownloadBackupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewStartBatchRollbackRequest() *StartBatchRollbackRequest {
	request := &StartBatchRollbackRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "StartBatchRollback")
	return request
}

func NewStartBatchRollbackResponse() *StartBatchRollbackResponse {
	return &StartBatchRollbackResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeRollbackRangeTimeRequest() *DescribeRollbackRangeTimeRequest {
	request := &DescribeRollbackRangeTimeRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeRollbackRangeTime")
	return request
}

func NewDescribeRollbackRangeTimeResponse() *DescribeRollbackRangeTimeResponse {
	return &DescribeRollbackRangeTimeResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewStartBatchRollbackToTemporaryDBInstanceRequest() *StartBatchRollbackToTemporaryDBInstanceRequest {
	request := &StartBatchRollbackToTemporaryDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "StartBatchRollbackToTemporaryDBInstance")
	return request
}

func NewStartBatchRollbackToTemporaryDBInstanceResponse() *StartBatchRollbackToTemporaryDBInstanceResponse {
	return &StartBatchRollbackToTemporaryDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeTemporaryDBInstancesRequest() *DescribeTemporaryDBInstancesRequest {
	request := &DescribeTemporaryDBInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeTemporaryDBInstances")
	return request
}

func NewDescribeTemporaryDBInstancesResponse() *DescribeTemporaryDBInstancesResponse {
	return &DescribeTemporaryDBInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewRegularizeTemporaryDBInstancesRequest() *RegularizeTemporaryDBInstancesRequest {
	request := &RegularizeTemporaryDBInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "RegularizeTemporaryDBInstances")
	return request
}

func NewRegularizeTemporaryDBInstancesResponse() *RegularizeTemporaryDBInstancesResponse {
	return &RegularizeTemporaryDBInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDeleteTemporaryDBInstancesRequest() *DeleteTemporaryDBInstancesRequest {
	request := &DeleteTemporaryDBInstancesRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DeleteTemporaryDBInstances")
	return request
}

func NewDeleteTemporaryDBInstancesResponse() *DeleteTemporaryDBInstancesResponse {
	return &DeleteTemporaryDBInstancesResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeAvailableReadOnlyConfigRequest() *DescribeAvailableReadOnlyConfigRequest {
	request := &DescribeAvailableReadOnlyConfigRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeAvailableReadOnlyConfig")
	return request
}

func NewDescribeAvailableReadOnlyConfigResponse() *DescribeAvailableReadOnlyConfigResponse {
	return &DescribeAvailableReadOnlyConfigResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewCreateReadOnlyDBInstanceRequest() *CreateReadOnlyDBInstanceRequest {
	request := &CreateReadOnlyDBInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "CreateReadOnlyDBInstance")
	return request
}

func NewCreateReadOnlyDBInstanceResponse() *CreateReadOnlyDBInstanceResponse {
	return &CreateReadOnlyDBInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeDBInstancePerformanceRequest() *DescribeDBInstancePerformanceRequest {
	request := &DescribeDBInstancePerformanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeDBInstancePerformance")
	return request
}

func NewDescribeDBInstancePerformanceResponse() *DescribeDBInstancePerformanceResponse {
	return &DescribeDBInstancePerformanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewModifyDBParameterRequest() *ModifyDBParameterRequest {
	request := &ModifyDBParameterRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "ModifyDBParameter")
	return request
}

func NewModifyDBParameterResponse() *ModifyDBParameterResponse {
	return &ModifyDBParameterResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewModifyDbBackupPolicyRequest() *ModifyDbBackupPolicyRequest {
	request := &ModifyDbBackupPolicyRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "ModifyDbBackupPolicy")
	return request
}

func NewModifyDbBackupPolicyResponse() *ModifyDbBackupPolicyResponse {
	return &ModifyDbBackupPolicyResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewModifyDbPrivilegeRequest() *ModifyDbPrivilegeRequest {
	request := &ModifyDbPrivilegeRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "ModifyDbPrivilege")
	return request
}

func NewModifyDbPrivilegeResponse() *ModifyDbPrivilegeResponse {
	return &ModifyDbPrivilegeResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}

func NewDescribeDBAccountRequest() *DescribeDBAccountRequest {
	request := &DescribeDBAccountRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo(MySqlService, ApiVersion, "DescribeDBAccount")
	return request
}

func NewDescribeDBAccountResponse() *DescribeDBAccountResponse {
	return &DescribeDBAccountResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
}
