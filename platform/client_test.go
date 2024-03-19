package platform

import (
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
	"testing"
)

func TestClient_DescribeSubjects(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeSubjectsRequest()
	response, err := client.DescribeSubjects(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %v", response.ToJsonString(), err)

}
