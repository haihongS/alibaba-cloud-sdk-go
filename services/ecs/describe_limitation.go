package ecs

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

// DescribeLimitation invokes the ecs.DescribeLimitation API synchronously
func (client *Client) DescribeLimitation(request *DescribeLimitationRequest) (response *DescribeLimitationResponse, err error) {
	response = CreateDescribeLimitationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLimitationWithChan invokes the ecs.DescribeLimitation API asynchronously
func (client *Client) DescribeLimitationWithChan(request *DescribeLimitationRequest) (<-chan *DescribeLimitationResponse, <-chan error) {
	responseChan := make(chan *DescribeLimitationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLimitation(request)
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

// DescribeLimitationWithCallback invokes the ecs.DescribeLimitation API asynchronously
func (client *Client) DescribeLimitationWithCallback(request *DescribeLimitationRequest, callback func(response *DescribeLimitationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLimitationResponse
		var err error
		defer close(result)
		response, err = client.DescribeLimitation(request)
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

// DescribeLimitationRequest is the request struct for api DescribeLimitation
type DescribeLimitationRequest struct {
	*requests.RpcRequest
	Limitation           string           `position:"Query" name:"Limitation"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLimitationResponse is the response struct for api DescribeLimitation
type DescribeLimitationResponse struct {
	*responses.BaseResponse
	Limitation string `json:"Limitation" xml:"Limitation"`
	Value      string `json:"Value" xml:"Value"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeLimitationRequest creates a request to invoke DescribeLimitation API
func CreateDescribeLimitationRequest() (request *DescribeLimitationRequest) {
	request = &DescribeLimitationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeLimitation", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLimitationResponse creates a response to parse from DescribeLimitation response
func CreateDescribeLimitationResponse() (response *DescribeLimitationResponse) {
	response = &DescribeLimitationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
