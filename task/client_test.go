package task

import (
	"fmt"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
)

func TestClient_DescribeTask(t *testing.T) {
	credential := common.NewCredential("eb4796589e1d11ee92ba0a18980946d7", "a1833d5911612ea9b5e3e5c6dc59eef1")

	cpftask := profile.NewClientProfile()
	cpftask.HttpProfile.ReqMethod = "GET"
	//cpftask.HttpProfile.Endpoint = "gateway.gic.test/openapi/ccs"
	taskclient, _ := NewClient(credential, regions.Beijing, cpftask)

	taskRequest := NewDescribeTaskRequest()
	taskRequest.TaskId = common.StringPtr("29869364")
	taskResponse, err := taskclient.DescribeTask(taskRequest)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", taskResponse.ToJsonString(), err)
}
