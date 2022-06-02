package amptest

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

// Huichengetest invokes the amptest.Huichengetest API synchronously
func (client *Client) Huichengetest(request *HuichengetestRequest) (response *HuichengetestResponse, err error) {
	response = CreateHuichengetestResponse()
	err = client.DoAction(request, response)
	return
}

// HuichengetestWithChan invokes the amptest.Huichengetest API asynchronously
func (client *Client) HuichengetestWithChan(request *HuichengetestRequest) (<-chan *HuichengetestResponse, <-chan error) {
	responseChan := make(chan *HuichengetestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Huichengetest(request)
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

// HuichengetestWithCallback invokes the amptest.Huichengetest API asynchronously
func (client *Client) HuichengetestWithCallback(request *HuichengetestRequest, callback func(response *HuichengetestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HuichengetestResponse
		var err error
		defer close(result)
		response, err = client.Huichengetest(request)
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

// HuichengetestRequest is the request struct for api Huichengetest
type HuichengetestRequest struct {
	*requests.RpcRequest
}

// HuichengetestResponse is the response struct for api Huichengetest
type HuichengetestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateHuichengetestRequest creates a request to invoke Huichengetest API
func CreateHuichengetestRequest() (request *HuichengetestRequest) {
	request = &HuichengetestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AmpTest", "2020-12-30", "Huichengetest", "", "")
	request.Method = requests.POST
	return
}

// CreateHuichengetestResponse creates a response to parse from Huichengetest response
func CreateHuichengetestResponse() (response *HuichengetestResponse) {
	response = &HuichengetestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
