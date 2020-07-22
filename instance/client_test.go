package instance

import (
	"fmt"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
	"github.com/capitalonline/cds-gic-sdk-go/task"
)

func TestClient_CreateInstance(t *testing.T) {
	ak := ""
	sk := ""
	credential := common.NewCredential(ak, sk)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewAddInstanceRequest()
	request.RegionId = common.StringPtr("APAC_Seoul_A")
	request.VdcId = common.StringPtr("")
	request.Password = common.StringPtr("")
	request.InstanceName = common.StringPtr("")
	request.InstanceChargeType = common.StringPtr("PostPaid")
	request.AutoRenew = common.IntPtr(0)
	request.Cpu = common.IntPtr(4)
	request.Ram = common.IntPtr(8)
	request.ImageId = common.StringPtr("Centos_7.6_64")
	request.PublicIp = common.StringPtrs([]string{"auto"})
	request.InstanceType = common.StringPtr("Standard") //7960400
	request.UTC = common.BoolPtr(false)                 //7960400

	dd1 := DataDisk{
		Size: common.IntPtr(200),
		Type: common.StringPtr("high_disk"),
	}
	ip := PrivateIp{
		PrivateID: common.StringPtr("09b2e1a6-6db2-11ea-b6bc-d61f1b70218b"),
		IP:        common.StringPtrs([]string{"auto"}),
	}
	request.DataDisks = []*DataDisk{&dd1}
	request.PrivateIp = []*PrivateIp{&ip}

	request.SystemDisk = &SystemDisk{
		IOPS: common.IntPtr(5),
		Size: common.IntPtr(200),
		Type: common.StringPtr("ssd_system_disk"), // or with type "system_disk"
	}

	response, err := client.CreateInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

	// get instance id
	taskRequest := task.NewDescribeTaskRequest()
	taskRequest.TaskId = response.TaskId
	cpftask := profile.NewClientProfile()
	cpftask.HttpProfile.ReqMethod = "GET"
	taskclient, _ := task.NewClient(credential, regions.Beijing, cpftask)
	taskResponse, err := taskclient.DescribeTask(taskRequest)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", taskResponse.ToJsonString(), err)
}

func TestClient_DescribeInstance(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeInstanceRequest()
	//request.VdcId = common.StringPtr("")
	//request.PageNumber = common.IntPtr(1)
	//request.PageSize = common.IntPtr(1000)
	request.InstanceId = common.StringPtr("vdc id")
	response, err := client.DescribeInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_ModifyInstanceName(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, "CN_Beijing_A", cpf)

	request := NewModifyInstanceNameRequest()
	request.InstanceId = common.StringPtr("instance id")
	request.InstanceName = common.StringPtr("new instance name")
	response, err := client.ModifyInstanceName(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

// wait for test
func TestClient_DeleteInstance(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteInstanceRequest()
	request.InstanceIds = common.StringPtrs([]string{"instance id"})
	response, err := client.DeleteInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}
