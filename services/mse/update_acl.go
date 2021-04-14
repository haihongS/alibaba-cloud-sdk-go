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

// UpdateAcl invokes the mse.UpdateAcl API synchronously
func (client *Client) UpdateAcl(request *UpdateAclRequest) (response *UpdateAclResponse, err error) {
	response = CreateUpdateAclResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAclWithChan invokes the mse.UpdateAcl API asynchronously
func (client *Client) UpdateAclWithChan(request *UpdateAclRequest) (<-chan *UpdateAclResponse, <-chan error) {
	responseChan := make(chan *UpdateAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAcl(request)
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

// UpdateAclWithCallback invokes the mse.UpdateAcl API asynchronously
func (client *Client) UpdateAclWithCallback(request *UpdateAclRequest, callback func(response *UpdateAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAclResponse
		var err error
		defer close(result)
		response, err = client.UpdateAcl(request)
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

// UpdateAclRequest is the request struct for api UpdateAcl
type UpdateAclRequest struct {
	*requests.RpcRequest
	ClusterId    string `position:"Query" name:"ClusterId"`
	AclEntryList string `position:"Query" name:"AclEntryList"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

// UpdateAclResponse is the response struct for api UpdateAcl
type UpdateAclResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAclRequest creates a request to invoke UpdateAcl API
func CreateUpdateAclRequest() (request *UpdateAclRequest) {
	request = &UpdateAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateAcl", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAclResponse creates a response to parse from UpdateAcl response
func CreateUpdateAclResponse() (response *UpdateAclResponse) {
	response = &UpdateAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
