package mse

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

// ListEurekaServices invokes the mse.ListEurekaServices API synchronously
func (client *Client) ListEurekaServices(request *ListEurekaServicesRequest) (response *ListEurekaServicesResponse, err error) {
	response = CreateListEurekaServicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListEurekaServicesWithChan invokes the mse.ListEurekaServices API asynchronously
func (client *Client) ListEurekaServicesWithChan(request *ListEurekaServicesRequest) (<-chan *ListEurekaServicesResponse, <-chan error) {
	responseChan := make(chan *ListEurekaServicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEurekaServices(request)
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

// ListEurekaServicesWithCallback invokes the mse.ListEurekaServices API asynchronously
func (client *Client) ListEurekaServicesWithCallback(request *ListEurekaServicesRequest, callback func(response *ListEurekaServicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEurekaServicesResponse
		var err error
		defer close(result)
		response, err = client.ListEurekaServices(request)
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

// ListEurekaServicesRequest is the request struct for api ListEurekaServices
type ListEurekaServicesRequest struct {
	*requests.RpcRequest
	ClusterId   string           `position:"Query" name:"ClusterId"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	RequestPars string           `position:"Query" name:"RequestPars"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// ListEurekaServicesResponse is the response struct for api ListEurekaServices
type ListEurekaServicesResponse struct {
	*responses.BaseResponse
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	Success    bool            `json:"Success" xml:"Success"`
	Message    string          `json:"Message" xml:"Message"`
	ErrorCode  string          `json:"ErrorCode" xml:"ErrorCode"`
	PageNumber int             `json:"PageNumber" xml:"PageNumber"`
	PageSize   int             `json:"PageSize" xml:"PageSize"`
	TotalCount int             `json:"TotalCount" xml:"TotalCount"`
	HttpCode   string          `json:"HttpCode" xml:"HttpCode"`
	Data       []SimpleService `json:"Data" xml:"Data"`
}

// CreateListEurekaServicesRequest creates a request to invoke ListEurekaServices API
func CreateListEurekaServicesRequest() (request *ListEurekaServicesRequest) {
	request = &ListEurekaServicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListEurekaServices", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListEurekaServicesResponse creates a response to parse from ListEurekaServices response
func CreateListEurekaServicesResponse() (response *ListEurekaServicesResponse) {
	response = &ListEurekaServicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
