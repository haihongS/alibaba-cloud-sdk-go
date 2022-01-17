package companyreg

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetCompanyRegOrder invokes the companyreg.GetCompanyRegOrder API synchronously
func (client *Client) GetCompanyRegOrder(request *GetCompanyRegOrderRequest) (response *GetCompanyRegOrderResponse, err error) {
	response = CreateGetCompanyRegOrderResponse()
	err = client.DoAction(request, response)
	return
}

// GetCompanyRegOrderWithChan invokes the companyreg.GetCompanyRegOrder API asynchronously
func (client *Client) GetCompanyRegOrderWithChan(request *GetCompanyRegOrderRequest) (<-chan *GetCompanyRegOrderResponse, <-chan error) {
	responseChan := make(chan *GetCompanyRegOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCompanyRegOrder(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetCompanyRegOrderWithCallback invokes the companyreg.GetCompanyRegOrder API asynchronously
func (client *Client) GetCompanyRegOrderWithCallback(request *GetCompanyRegOrderRequest, callback func(response *GetCompanyRegOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCompanyRegOrderResponse
		var err error
		defer close(result)
		response, err = client.GetCompanyRegOrder(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetCompanyRegOrderRequest is the request struct for api GetCompanyRegOrder
type GetCompanyRegOrderRequest struct {
	*requests.RpcRequest
	ActionTypes      string           `position:"Query" name:"ActionTypes"`
	BizCode          string           `position:"Query" name:"BizCode"`
	BizId            string           `position:"Query" name:"BizId"`
	MaxOperationSize requests.Integer `position:"Query" name:"MaxOperationSize"`
	BizSubCode       string           `position:"Query" name:"BizSubCode"`
}

// GetCompanyRegOrderResponse is the response struct for api GetCompanyRegOrder
type GetCompanyRegOrderResponse struct {
	*responses.BaseResponse
	RequestId            string     `json:"RequestId" xml:"RequestId"`
	BizId                string     `json:"BizId" xml:"BizId"`
	CompanyName          string     `json:"CompanyName" xml:"CompanyName"`
	BizStatus            string     `json:"BizStatus" xml:"BizStatus"`
	BizInfo              string     `json:"BizInfo" xml:"BizInfo"`
	SupplementBizInfo    string     `json:"SupplementBizInfo" xml:"SupplementBizInfo"`
	AliyunOrderId        string     `json:"AliyunOrderId" xml:"AliyunOrderId"`
	GmtModified          int64      `json:"GmtModified" xml:"GmtModified"`
	OrderAmount          float64    `json:"OrderAmount" xml:"OrderAmount"`
	YunMarketOrderAmount float64    `json:"YunMarketOrderAmount" xml:"YunMarketOrderAmount"`
	GmtPaid              int64      `json:"GmtPaid" xml:"GmtPaid"`
	Extend               string     `json:"Extend" xml:"Extend"`
	BizStatusStage       string     `json:"BizStatusStage" xml:"BizStatusStage"`
	PlatformName         string     `json:"PlatformName" xml:"PlatformName"`
	InboundPhone         string     `json:"InboundPhone" xml:"InboundPhone"`
	OutboundPhone        string     `json:"OutboundPhone" xml:"OutboundPhone"`
	BizSubCode           string     `json:"BizSubCode" xml:"BizSubCode"`
	Operations           Operations `json:"Operations" xml:"Operations"`
}

// CreateGetCompanyRegOrderRequest creates a request to invoke GetCompanyRegOrder API
func CreateGetCompanyRegOrderRequest() (request *GetCompanyRegOrderRequest) {
	request = &GetCompanyRegOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "GetCompanyRegOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCompanyRegOrderResponse creates a response to parse from GetCompanyRegOrder response
func CreateGetCompanyRegOrderResponse() (response *GetCompanyRegOrderResponse) {
	response = &GetCompanyRegOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
