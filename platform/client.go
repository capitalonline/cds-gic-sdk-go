package platform

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
	clientProfile.HttpProfile.Endpoint = cdshttp.RootDomainNew
	return
}

func NewDescribeSubjectsRequest() (request *DescribeSubjectsRequest) {
	request = &DescribeSubjectsRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("platform/subject", ApiVersion, "DescribeSubjects")
	request.SetDomain(cdshttp.RootDomainNew + "/platform/subject")
	return
}

func NewDescribeSubjectsResponse() (response *DescribeSubjectsResponse) {
	response = &DescribeSubjectsResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func (c *Client) DescribeSubjects(request *DescribeSubjectsRequest) (response *DescribeSubjectsResponse, err error) {
	if request == nil {
		request = NewDescribeSubjectsRequest()
	}
	response = NewDescribeSubjectsResponse()
	err = c.Send(request, response)
	return
}
