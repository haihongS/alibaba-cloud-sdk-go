package dataworks_public

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

// GetNode invokes the dataworks_public.GetNode API synchronously
func (client *Client) GetNode(request *GetNodeRequest) (response *GetNodeResponse, err error) {
	response = CreateGetNodeResponse()
	err = client.DoAction(request, response)
	return
}

// GetNodeWithChan invokes the dataworks_public.GetNode API asynchronously
func (client *Client) GetNodeWithChan(request *GetNodeRequest) (<-chan *GetNodeResponse, <-chan error) {
	responseChan := make(chan *GetNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNode(request)
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

// GetNodeWithCallback invokes the dataworks_public.GetNode API asynchronously
func (client *Client) GetNodeWithCallback(request *GetNodeRequest, callback func(response *GetNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNodeResponse
		var err error
		defer close(result)
		response, err = client.GetNode(request)
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

// GetNodeRequest is the request struct for api GetNode
type GetNodeRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	NodeId     requests.Integer `position:"Body" name:"NodeId"`
}

// GetNodeResponse is the response struct for api GetNode
type GetNodeResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetNodeRequest creates a request to invoke GetNode API
func CreateGetNodeRequest() (request *GetNodeRequest) {
	request = &GetNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetNode", "", "")
	request.Method = requests.POST
	return
}

// CreateGetNodeResponse creates a response to parse from GetNode response
func CreateGetNodeResponse() (response *GetNodeResponse) {
	response = &GetNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}