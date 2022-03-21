package ecd

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

// VerifyCen invokes the ecd.VerifyCen API synchronously
func (client *Client) VerifyCen(request *VerifyCenRequest) (response *VerifyCenResponse, err error) {
	response = CreateVerifyCenResponse()
	err = client.DoAction(request, response)
	return
}

// VerifyCenWithChan invokes the ecd.VerifyCen API asynchronously
func (client *Client) VerifyCenWithChan(request *VerifyCenRequest) (<-chan *VerifyCenResponse, <-chan error) {
	responseChan := make(chan *VerifyCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyCen(request)
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

// VerifyCenWithCallback invokes the ecd.VerifyCen API asynchronously
func (client *Client) VerifyCenWithCallback(request *VerifyCenRequest, callback func(response *VerifyCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyCenResponse
		var err error
		defer close(result)
		response, err = client.VerifyCen(request)
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

// VerifyCenRequest is the request struct for api VerifyCen
type VerifyCenRequest struct {
	*requests.RpcRequest
	CenId      string           `position:"Query" name:"CenId"`
	CenOwnerId requests.Integer `position:"Query" name:"CenOwnerId"`
	VerifyCode string           `position:"Query" name:"VerifyCode"`
	CidrBlock  string           `position:"Query" name:"CidrBlock"`
}

// VerifyCenResponse is the response struct for api VerifyCen
type VerifyCenResponse struct {
	*responses.BaseResponse
	Status       string       `json:"Status" xml:"Status"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	CidrBlocks   []string     `json:"CidrBlocks" xml:"CidrBlocks"`
	RouteEntries []RouteEntry `json:"RouteEntries" xml:"RouteEntries"`
}

// CreateVerifyCenRequest creates a request to invoke VerifyCen API
func CreateVerifyCenRequest() (request *VerifyCenRequest) {
	request = &VerifyCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "VerifyCen", "", "")
	request.Method = requests.POST
	return
}

// CreateVerifyCenResponse creates a response to parse from VerifyCen response
func CreateVerifyCenResponse() (response *VerifyCenResponse) {
	response = &VerifyCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
