package linkwan

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

// RemoveNodeFromGroup invokes the linkwan.RemoveNodeFromGroup API synchronously
func (client *Client) RemoveNodeFromGroup(request *RemoveNodeFromGroupRequest) (response *RemoveNodeFromGroupResponse, err error) {
	response = CreateRemoveNodeFromGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveNodeFromGroupWithChan invokes the linkwan.RemoveNodeFromGroup API asynchronously
func (client *Client) RemoveNodeFromGroupWithChan(request *RemoveNodeFromGroupRequest) (<-chan *RemoveNodeFromGroupResponse, <-chan error) {
	responseChan := make(chan *RemoveNodeFromGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveNodeFromGroup(request)
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

// RemoveNodeFromGroupWithCallback invokes the linkwan.RemoveNodeFromGroup API asynchronously
func (client *Client) RemoveNodeFromGroupWithCallback(request *RemoveNodeFromGroupRequest, callback func(response *RemoveNodeFromGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveNodeFromGroupResponse
		var err error
		defer close(result)
		response, err = client.RemoveNodeFromGroup(request)
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

// RemoveNodeFromGroupRequest is the request struct for api RemoveNodeFromGroup
type RemoveNodeFromGroupRequest struct {
	*requests.RpcRequest
	DevEui      string `position:"Query" name:"DevEui"`
	NodeGroupId string `position:"Query" name:"NodeGroupId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// RemoveNodeFromGroupResponse is the response struct for api RemoveNodeFromGroup
type RemoveNodeFromGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRemoveNodeFromGroupRequest creates a request to invoke RemoveNodeFromGroup API
func CreateRemoveNodeFromGroupRequest() (request *RemoveNodeFromGroupRequest) {
	request = &RemoveNodeFromGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "RemoveNodeFromGroup", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveNodeFromGroupResponse creates a response to parse from RemoveNodeFromGroup response
func CreateRemoveNodeFromGroupResponse() (response *RemoveNodeFromGroupResponse) {
	response = &RemoveNodeFromGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}