package task

import (
	"fmt"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
)

func TestClient_DescribeTask(t *testing.T) {
	credential := common.NewCredential("", "")

	cpftask := profile.NewClientProfile()
	cpftask.HttpProfile.ReqMethod = "GET"
	taskclient, _ := NewClient(credential, regions.Beijing, cpftask)

	taskRequest := NewDescribeTaskRequest()
	taskRequest.TaskId = common.StringPtr("29869364")
	taskResponse, err := taskclient.DescribeTask(taskRequest)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", taskResponse.ToJsonString(), err)
}
