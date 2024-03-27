package redis

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

func TestDescribeRegions(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("GET")

	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewDescribeRegionsRequest()
	response, err := client.DescribeRegins(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestRedisDescribeAvailableDBConfig(t *testing.T) {
	credential := common.NewCredential("", "")
	cpf := clientProfile("GET")
	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewDescribeAvailableDBConfigRequest()
	request.RegionId = common.StringPtr(regions.Beijing)
	response, err := client.DescribeAvailableDBConfig(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestRedisCreateDBInstance(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("POST")
	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewCreateDBInstanceRequest()
	request.RegionId = common.StringPtr(regions.Beijing)
	request.VdcId = common.StringPtr("c26529b9-f455-47d7-b10c-7eb1f1f72bd0")
	request.BasePipeId = common.StringPtr("9fd9bf3e-540a-11ec-9d8e-96e971c86150")
	request.InstanceName = common.StringPtr("my_redis_test")
	request.PaasGoodsId = common.IntPtr(12467)
	request.Password = common.StringPtr("test_redis123")
	request.Amount = common.IntPtr(1)
	request.SubjectId = common.IntPtr(23423)
	response, err := client.CreateDBInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestRedisDescribeDBInstances(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("GET")
	client, _ := NewClient(credential, regions.Beijing, cpf)
	request := NewDescribeDBInstancesRequest()

	// request.InstanceUuid = common.StringPtr("be84174f-2ef6-4934-a800-9ddbdbf43183")
	// request.InstanceName = common.StringPtr("my_redis_test")
	request.InstanceUuid = common.StringPtr("ebd33210-546c-4f96-896e-f0d716dc9666")

	response, err := client.DescribeDBInstances(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestRedisDeleteDBInstance(t *testing.T) {
	credential := common.NewCredential(s_id, s_key)
	cpf := clientProfile("POST")
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteDBInstanceRequest()
	request.InstanceUuid = common.StringPtr("be84174f-2ef6-4934-a800-9ddbdbf43183")
	response, err := client.DeleteDBInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}
