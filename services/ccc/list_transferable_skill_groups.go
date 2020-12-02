package ccc

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

// ListTransferableSkillGroups invokes the ccc.ListTransferableSkillGroups API synchronously
func (client *Client) ListTransferableSkillGroups(request *ListTransferableSkillGroupsRequest) (response *ListTransferableSkillGroupsResponse, err error) {
	response = CreateListTransferableSkillGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransferableSkillGroupsWithChan invokes the ccc.ListTransferableSkillGroups API asynchronously
func (client *Client) ListTransferableSkillGroupsWithChan(request *ListTransferableSkillGroupsRequest) (<-chan *ListTransferableSkillGroupsResponse, <-chan error) {
	responseChan := make(chan *ListTransferableSkillGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransferableSkillGroups(request)
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

// ListTransferableSkillGroupsWithCallback invokes the ccc.ListTransferableSkillGroups API asynchronously
func (client *Client) ListTransferableSkillGroupsWithCallback(request *ListTransferableSkillGroupsRequest, callback func(response *ListTransferableSkillGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransferableSkillGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListTransferableSkillGroups(request)
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

// ListTransferableSkillGroupsRequest is the request struct for api ListTransferableSkillGroups
type ListTransferableSkillGroupsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListTransferableSkillGroupsResponse is the response struct for api ListTransferableSkillGroups
type ListTransferableSkillGroupsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SkillGroups    []SkillGroup `json:"SkillGroups" xml:"SkillGroups"`
}

// CreateListTransferableSkillGroupsRequest creates a request to invoke ListTransferableSkillGroups API
func CreateListTransferableSkillGroupsRequest() (request *ListTransferableSkillGroupsRequest) {
	request = &ListTransferableSkillGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListTransferableSkillGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateListTransferableSkillGroupsResponse creates a response to parse from ListTransferableSkillGroups response
func CreateListTransferableSkillGroupsResponse() (response *ListTransferableSkillGroupsResponse) {
	response = &ListTransferableSkillGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}