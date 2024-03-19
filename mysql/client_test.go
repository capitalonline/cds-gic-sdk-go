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
	credential := common.NewCredential("abc07d56422e11eea9798e96c407823e", "b9042f4c207d6b59fdcb9bc99f655928")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	//cpf.HttpProfile.Endpoint = "gateway.gic.test/openapi/ccs"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeDBInstancesRequest()
	//request.InstanceUuid = common.StringPtr("26de01dc-e1c8-11ee-96b3-ea38d02cc52b")
	response, err := client.DescribeDBInstances(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %v", response.ToJsonString(), err)

}
