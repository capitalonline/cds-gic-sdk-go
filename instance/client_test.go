package instance

import (
	"fmt"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
	"github.com/capitalonline/cds-gic-sdk-go/task"
)

var ak string = "xxxxxx"
var sk string = "xxxxxx"

func TestClient_CreateInstance(t *testing.T) {
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
		PrivateID: common.StringPtr("private_id"),
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
	credential := common.NewCredential(ak, sk)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeInstanceRequest()
	request.VdcId = common.StringPtr("")
	request.PageNumber = common.IntPtr(1)
	request.PageSize = common.IntPtr(1000)
	request.InstanceId = common.StringPtr("instance_id")
	var ips = make([]string, 0, 1)
	ips = append(ips, "x.x.x.x")
	request.PublicIp = common.StringPtrs(ips)
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

func TestClient_ModifyInstanceSpec(t *testing.T) {
	credential := common.NewCredential(ak, sk)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewModifyInstanceSpecRequest()
	request.InstanceId = common.StringPtr("")
	request.Ram = common.IntPtr(8)
	request.Cpu = common.IntPtr(4)
	response, err := client.ModifyInstanceSpec(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_CreateDisk(t *testing.T) {
	credential := common.NewCredential(ak, sk)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewCreateDiskRequest()
	request.InstanceId = common.StringPtr("")
	var disks = make([]*DataDisk, 0, 1)
	request.DataDisks = append(disks, &DataDisk{
		Size: common.IntPtr(20),
		Type: common.StringPtr("ssd_disk"),
		IOPS: common.IntPtr(0),
	})
	response, err := client.CreateDisk(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_ResizeDisk(t *testing.T) {
	credential := common.NewCredential(ak, sk)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewResizeDiskRequest()
	request.InstanceId = common.StringPtr("")
	request.DiskId = common.StringPtr("disk_id")
	request.DataSize = common.IntPtr(30)
	response, err := client.ResizeDisk(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_DeleteDisk(t *testing.T) {
	credential := common.NewCredential(ak, sk)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteDiskRequest()
	var ids = make([]*string, 0, 1)
	ids = append(ids, common.StringPtr("disk_id"))
	request.InstanceId = common.StringPtr("instance_id")
	request.DiskIds = ids
	response, err := client.DeleteDisk(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}
