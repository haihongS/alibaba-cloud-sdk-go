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

// ReleaseProduceAuthorization invokes the companyreg.ReleaseProduceAuthorization API synchronously
func (client *Client) ReleaseProduceAuthorization(request *ReleaseProduceAuthorizationRequest) (response *ReleaseProduceAuthorizationResponse, err error) {
	response = CreateReleaseProduceAuthorizationResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseProduceAuthorizationWithChan invokes the companyreg.ReleaseProduceAuthorization API asynchronously
func (client *Client) ReleaseProduceAuthorizationWithChan(request *ReleaseProduceAuthorizationRequest) (<-chan *ReleaseProduceAuthorizationResponse, <-chan error) {
	responseChan := make(chan *ReleaseProduceAuthorizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseProduceAuthorization(request)
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

// ReleaseProduceAuthorizationWithCallback invokes the companyreg.ReleaseProduceAuthorization API asynchronously
func (client *Client) ReleaseProduceAuthorizationWithCallback(request *ReleaseProduceAuthorizationRequest, callback func(response *ReleaseProduceAuthorizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseProduceAuthorizationResponse
		var err error
		defer close(result)
		response, err = client.ReleaseProduceAuthorization(request)
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

// ReleaseProduceAuthorizationRequest is the request struct for api ReleaseProduceAuthorization
type ReleaseProduceAuthorizationRequest struct {
	*requests.RpcRequest
	BizType          string `position:"Body" name:"BizType"`
	AuthorizedUserId string `position:"Body" name:"AuthorizedUserId"`
	BizId            string `position:"Body" name:"BizId"`
}

// ReleaseProduceAuthorizationResponse is the response struct for api ReleaseProduceAuthorization
type ReleaseProduceAuthorizationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateReleaseProduceAuthorizationRequest creates a request to invoke ReleaseProduceAuthorization API
func CreateReleaseProduceAuthorizationRequest() (request *ReleaseProduceAuthorizationRequest) {
	request = &ReleaseProduceAuthorizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "ReleaseProduceAuthorization", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseProduceAuthorizationResponse creates a response to parse from ReleaseProduceAuthorization response
func CreateReleaseProduceAuthorizationResponse() (response *ReleaseProduceAuthorizationResponse) {
	response = &ReleaseProduceAuthorizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
