package mongodb

import (
	"fmt"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
)

func clientProfile(method string) *profile.ClientProfile {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = method
	return cpf
}

var (
	s_id  = "379c6346eca011e9abc00242ac110007"
	s_key = "a8ea19d012ee13faeecf62ff6656b6ae"
)

func TestDescribeZones(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("GET")
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeZonesRequest()
	response, err := client.DescribeZones(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)
}

func TestDescribeSpecInfo(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("GET")

	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewDescribeSpecInfoRequest()
	request.RegionId = common.StringPtr("CN_Shanghai_A")
	response, err := client.DescribeSpecInfo(request)

	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)
}

func TestCreateDBInstance(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("POST")
	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewCreateDBInstanceRequest()
	request.RegionId = common.StringPtr(regions.Beijing)
	request.VdcId = common.StringPtr("c26529b9-f455-47d7-b10c-7eb1f1f72bd0")
	request.BasePipeId = common.StringPtr("9fd9bf3e-540a-11ec-9d8e-96e971c86150")

	request.InstanceName = common.StringPtr("my_mongodb_test")
	request.PaasGoodsId = common.IntPtr(7295)
	request.DiskType = common.StringPtr("ssd_disk")
	request.DiskValue = common.IntPtr(100)
	request.Password = common.StringPtr("test_sdfaferdfd12312@213")
	request.Version = common.StringPtr("3.2.21")
	request.SubjectId = common.IntPtr(2342)
	response, err := client.CreateDBInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)
}

func TestDescribeDBInstances(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("GET")

	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewDescribeDBInstancesRequest()
	// request.InstanceUuid = common.StringPtr("")
	request.InstanceName = common.StringPtr("MongoDB-2022-01-18-FHfYp")
	response, err := client.DescribeDBInstances(request)

	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestDeleteDBInstance(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("POST")

	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewDeleteDBInstanceRequest()
	request.InstanceUuid = common.StringPtr("982189b8-c277-4e4e-ae12-3ac5fd0f37cf")

	response, err := client.DeleteDBInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)
}
