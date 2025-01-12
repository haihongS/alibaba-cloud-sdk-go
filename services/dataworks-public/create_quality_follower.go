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

// CreateQualityFollower invokes the dataworks_public.CreateQualityFollower API synchronously
func (client *Client) CreateQualityFollower(request *CreateQualityFollowerRequest) (response *CreateQualityFollowerResponse, err error) {
	response = CreateCreateQualityFollowerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQualityFollowerWithChan invokes the dataworks_public.CreateQualityFollower API asynchronously
func (client *Client) CreateQualityFollowerWithChan(request *CreateQualityFollowerRequest) (<-chan *CreateQualityFollowerResponse, <-chan error) {
	responseChan := make(chan *CreateQualityFollowerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQualityFollower(request)
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

// CreateQualityFollowerWithCallback invokes the dataworks_public.CreateQualityFollower API asynchronously
func (client *Client) CreateQualityFollowerWithCallback(request *CreateQualityFollowerRequest, callback func(response *CreateQualityFollowerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQualityFollowerResponse
		var err error
		defer close(result)
		response, err = client.CreateQualityFollower(request)
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

// CreateQualityFollowerRequest is the request struct for api CreateQualityFollower
type CreateQualityFollowerRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Body" name:"ProjectName"`
	Follower    string           `position:"Body" name:"Follower"`
	EntityId    requests.Integer `position:"Body" name:"EntityId"`
	AlarmMode   requests.Integer `position:"Body" name:"AlarmMode"`
}

// CreateQualityFollowerResponse is the response struct for api CreateQualityFollower
type CreateQualityFollowerResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int    `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateCreateQualityFollowerRequest creates a request to invoke CreateQualityFollower API
func CreateCreateQualityFollowerRequest() (request *CreateQualityFollowerRequest) {
	request = &CreateQualityFollowerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateQualityFollower", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateQualityFollowerResponse creates a response to parse from CreateQualityFollower response
func CreateCreateQualityFollowerResponse() (response *CreateQualityFollowerResponse) {
	response = &CreateQualityFollowerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
