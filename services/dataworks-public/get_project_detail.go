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

// GetProjectDetail invokes the dataworks_public.GetProjectDetail API synchronously
func (client *Client) GetProjectDetail(request *GetProjectDetailRequest) (response *GetProjectDetailResponse, err error) {
	response = CreateGetProjectDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectDetailWithChan invokes the dataworks_public.GetProjectDetail API asynchronously
func (client *Client) GetProjectDetailWithChan(request *GetProjectDetailRequest) (<-chan *GetProjectDetailResponse, <-chan error) {
	responseChan := make(chan *GetProjectDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProjectDetail(request)
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

// GetProjectDetailWithCallback invokes the dataworks_public.GetProjectDetail API asynchronously
func (client *Client) GetProjectDetailWithCallback(request *GetProjectDetailRequest, callback func(response *GetProjectDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectDetailResponse
		var err error
		defer close(result)
		response, err = client.GetProjectDetail(request)
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

// GetProjectDetailRequest is the request struct for api GetProjectDetail
type GetProjectDetailRequest struct {
	*requests.RpcRequest
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
}

// GetProjectDetailResponse is the response struct for api GetProjectDetail
type GetProjectDetailResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Success        bool                   `json:"Success" xml:"Success"`
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	Data           DataInGetProjectDetail `json:"Data" xml:"Data"`
}

// CreateGetProjectDetailRequest creates a request to invoke GetProjectDetail API
func CreateGetProjectDetailRequest() (request *GetProjectDetailRequest) {
	request = &GetProjectDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetProjectDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateGetProjectDetailResponse creates a response to parse from GetProjectDetail response
func CreateGetProjectDetailResponse() (response *GetProjectDetailResponse) {
	response = &GetProjectDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}