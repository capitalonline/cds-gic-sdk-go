package main

import (
	"fmt"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/instance"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"
)

const (
	// 北京
	Beijing = "CN_Beijing_A"
	// 上海
	Shanghai = "CN_Shanghai_A"
)

func main() {

	// Init a credential with Access Key Id and Secret Access Key
	// You can apply them from the CDS web portal
	credential := common.NewCredential(
		os.Getenv("CDS_SECRET_ID"),
		os.Getenv("CDS_SECRET_KEY"),
	)

	// init a client profile with method type
	cpf := profile.NewClientProfile()

	// Example: Get vdc
	cpf.HttpProfile.ReqMethod = "GET"
	vdcClient, _ := vdc.NewClient(credential, "", cpf)
	descVdcRequest := vdc.DescribeVdcRequest()

	// comment out the line below, you can get all the vdc data
	descVdcRequest.RegionId = common.StringPtr(Shanghai)
	descVdcResponse, err := vdcClient.DescribeVdc(descVdcRequest)
	if err != nil {
		fmt.Println("API request fail:", err.Error())
	} else {
		fmt.Println(descVdcResponse.ToJsonString())
	}

	// Example: Create vdc
	cpf.HttpProfile.ReqMethod = "POST"
	vdcClient, _ = vdc.NewClient(credential, "", cpf)
	createVdcRequest := vdc.NewAddVdcRequest()
	createVdcRequest.RegionId = common.StringPtr(Beijing)
	createVdcRequest.VdcName = common.StringPtr("beijing-vdc")
	createVdcRequest.PublicNetwork = &vdc.PublicNetwork{
		// the account of public IP
		IPNum: common.IntPtr(8),
		// the bandwidth of public network
		Qos:           common.IntPtr(10),
		Name:          common.StringPtr("pubnet"),
		BillingMethod: common.StringPtr("Bandwidth"),
		Type:          common.StringPtr("Bandwidth_BGP"),
	}
	createVdcResponse, err := vdcClient.CreateVdc(createVdcRequest)
	if err != nil {
		fmt.Println("API request fail:", err.Error())
	} else {
		fmt.Printf("Task: %v, code: %v", *createVdcResponse.TaskId, *createVdcResponse.Code)
	}

	// Example: Get vm
	cpf.HttpProfile.ReqMethod = "POST"
	instanceClient, _ := instance.NewClient(credential, "", cpf)
	descInstanceRequest := instance.NewDescribeInstanceRequest()
	// comment out the two line below, you can get all the vm info of all the vdc
	descInstanceRequest.VdcId = common.StringPtr("vdc_id")
	descInstanceRequest.InstanceId = common.StringPtr("instance_id")
	descInstanceResponse, err := instanceClient.DescribeInstance(descInstanceRequest)
	if err != nil {
		fmt.Println("API request fail:", err.Error())
	} else {
		fmt.Println(descInstanceResponse.ToJsonString())
	}

	// Example: Create a vm
	cpf.HttpProfile.ReqMethod = "POST"
	createInstanceRequest := instance.NewAddInstanceRequest()
	createInstanceRequest.RegionId = common.StringPtr(Shanghai)
	createInstanceRequest.VdcId = common.StringPtr("vdc_id")
	createInstanceRequest.Password = common.StringPtr("password")
	createInstanceRequest.InstanceName = common.StringPtr("hostname001")
	createInstanceRequest.InstanceChargeType = common.StringPtr("PostPaid")
	createInstanceRequest.AutoRenew = common.IntPtr(0)
	createInstanceRequest.Cpu = common.IntPtr(1)
	createInstanceRequest.Ram = common.IntPtr(1)
	createInstanceRequest.PrepaidMonth = common.IntPtr(1)
	createInstanceRequest.Amount = common.IntPtr(1)
	createInstanceRequest.ImageId = common.StringPtr("Ubuntu_16.04_64")
	createInstanceRequest.PublicIp = common.StringPtrs([]string{"auto"})
	createInstanceRequest.InstanceType = common.StringPtr("Standard")
	dd1 := instance.DataDisk{
		Size: common.IntPtr(100),
		Type: common.StringPtr("high_disk"),
	}
	ip := instance.PrivateIp{
		PrivateID: common.StringPtr("available-private-network-id-under-the-target-vdc"),
		IP:        common.StringPtrs([]string{"auto"}),
	}
	createInstanceRequest.DataDisks = []*instance.DataDisk{&dd1}
	createInstanceRequest.PrivateIp = []*instance.PrivateIp{&ip}

	createInstanceResponse, err := instanceClient.CreateInstance(createInstanceRequest)
	if err != nil {
		fmt.Println("API request fail:", err.Error())
	} else {
		fmt.Printf("Task: %v, code: %v", *createInstanceResponse.TaskId, *createInstanceResponse.Code)
	}
}