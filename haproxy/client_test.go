package haproxy

import (
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
	"testing"
)

func TestClient_DescribeDedicatedHostTypes(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeLoadBalancersRequest()
	request.InstanceUuid = common.StringPtr("")
	response, err := client.DescribeLoadBalancers(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %v", response.ToJsonString(), err)

}

func TestClient_DescribeLoadBalancers(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeLoadBalancersRequest()
	request.InstanceUuid = common.StringPtr("")
	response, err := client.DescribeLoadBalancers(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %v", response.ToJsonString(), err)

}
