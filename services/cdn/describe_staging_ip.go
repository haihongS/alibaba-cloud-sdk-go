package cdn

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

// DescribeStagingIp invokes the cdn.DescribeStagingIp API synchronously
func (client *Client) DescribeStagingIp(request *DescribeStagingIpRequest) (response *DescribeStagingIpResponse, err error) {
	response = CreateDescribeStagingIpResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStagingIpWithChan invokes the cdn.DescribeStagingIp API asynchronously
func (client *Client) DescribeStagingIpWithChan(request *DescribeStagingIpRequest) (<-chan *DescribeStagingIpResponse, <-chan error) {
	responseChan := make(chan *DescribeStagingIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStagingIp(request)
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

// DescribeStagingIpWithCallback invokes the cdn.DescribeStagingIp API asynchronously
func (client *Client) DescribeStagingIpWithCallback(request *DescribeStagingIpRequest, callback func(response *DescribeStagingIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStagingIpResponse
		var err error
		defer close(result)
		response, err = client.DescribeStagingIp(request)
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

// DescribeStagingIpRequest is the request struct for api DescribeStagingIp
type DescribeStagingIpRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeStagingIpResponse is the response struct for api DescribeStagingIp
type DescribeStagingIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IPV4s     IPV4s  `json:"IPV4s" xml:"IPV4s"`
}

// CreateDescribeStagingIpRequest creates a request to invoke DescribeStagingIp API
func CreateDescribeStagingIpRequest() (request *DescribeStagingIpRequest) {
	request = &DescribeStagingIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeStagingIp", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStagingIpResponse creates a response to parse from DescribeStagingIp response
func CreateDescribeStagingIpResponse() (response *DescribeStagingIpResponse) {
	response = &DescribeStagingIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
