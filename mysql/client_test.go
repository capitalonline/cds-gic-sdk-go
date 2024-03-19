package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common/regions"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"
)

func clientProfile(method string) *profile.ClientProfile {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = method
	return cpf
}

func TestClient(t *testing.T) {

	client, _ := vdc.NewClient(common.NewCredential("bd27cbbc72f711eab98c824ee1207173", "422cd7dc78a99c77092c687c39fe1dac"), "CN_Beijing_E", clientProfile("GET"))

	request := NewDescribeModifiableDBSpecRequest()
	request.InstanceUuid = common.StringPtr("e58f0f1a-f96b-42eb-b78d-a17a4dd400cd")

	response := NewDescribeModifiableDBSpecResponse()

	if err := client.Send(request, response); err != nil {
		t.Fatal(err)
	}

	bytes, _ := json.Marshal(response)

	fmt.Println(string(bytes))
}

func TestClient_DescribeDBInstances(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeDBInstancesRequest()
	request.InstanceUuid = common.StringPtr("")
	response, err := client.DescribeDBInstances(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %v", response.ToJsonString(), err)

}
