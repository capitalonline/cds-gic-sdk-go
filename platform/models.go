package platform

import (
	"encoding/json"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
)

type DescribeSubjectsRequest struct {
	*cdshttp.BaseRequest
}

func (request *DescribeSubjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *DescribeSubjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}

type DescribeSubjectsResponse struct {
	*cdshttp.BaseResponse

	Code      *string                      `json:"Code,omitempty"`
	Msg       *string                      `json:"Msg,omitempty"`
	Data      DescribeSubjectsResponseData `json:"Data,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty"`
}

type DescribeSubjectsResponseData struct {
	Total       *int                                   `json:"Total,omitempty"`
	Currency    *string                                `json:"Currency,omitempty"`
	SubjectList []*DescribeSubjectsResponseSubjectList `json:"SubjectList,omitempty"`
}

type DescribeSubjectsResponseSubjectList struct {
	Id                *int    `json:"Id,omitempty"`
	Name              *string `json:"Name,omitempty"`
	Balance           *string `json:"Balance,omitempty"`
	BeginTime         *string `json:"BeginTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty"`
	BillMethod        *string `json:"BillMethod,omitempty"`
	BillMethodDisplay *string `json:"BillMethodDisplay,omitempty"`
	GoodsIds          *string `json:"GoodsIds,omitempty"`
	GoodsNames        *string `json:"GoodsNames,omitempty"`
	SiteIds           *string `json:"SiteIds,omitempty"`
	SiteNames         *string `json:"SiteNames,omitempty"`
}

func (request *DescribeSubjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &request)
}

func (request *DescribeSubjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(request)
	return string(b)
}
