package arms

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

// DeleteGrafanaResource invokes the arms.DeleteGrafanaResource API synchronously
func (client *Client) DeleteGrafanaResource(request *DeleteGrafanaResourceRequest) (response *DeleteGrafanaResourceResponse, err error) {
	response = CreateDeleteGrafanaResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGrafanaResourceWithChan invokes the arms.DeleteGrafanaResource API asynchronously
func (client *Client) DeleteGrafanaResourceWithChan(request *DeleteGrafanaResourceRequest) (<-chan *DeleteGrafanaResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteGrafanaResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGrafanaResource(request)
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

// DeleteGrafanaResourceWithCallback invokes the arms.DeleteGrafanaResource API asynchronously
func (client *Client) DeleteGrafanaResourceWithCallback(request *DeleteGrafanaResourceRequest, callback func(response *DeleteGrafanaResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGrafanaResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteGrafanaResource(request)
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

// DeleteGrafanaResourceRequest is the request struct for api DeleteGrafanaResource
type DeleteGrafanaResourceRequest struct {
	*requests.RpcRequest
	ClusterName string `position:"Body" name:"ClusterName"`
	ClusterId   string `position:"Body" name:"ClusterId"`
	UserId      string `position:"Body" name:"UserId"`
}

// DeleteGrafanaResourceResponse is the response struct for api DeleteGrafanaResource
type DeleteGrafanaResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDeleteGrafanaResourceRequest creates a request to invoke DeleteGrafanaResource API
func CreateDeleteGrafanaResourceRequest() (request *DeleteGrafanaResourceRequest) {
	request = &DeleteGrafanaResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteGrafanaResource", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGrafanaResourceResponse creates a response to parse from DeleteGrafanaResource response
func CreateDeleteGrafanaResourceResponse() (response *DeleteGrafanaResourceResponse) {
	response = &DeleteGrafanaResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
