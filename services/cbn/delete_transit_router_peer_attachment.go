package cbn

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

// DeleteTransitRouterPeerAttachment invokes the cbn.DeleteTransitRouterPeerAttachment API synchronously
func (client *Client) DeleteTransitRouterPeerAttachment(request *DeleteTransitRouterPeerAttachmentRequest) (response *DeleteTransitRouterPeerAttachmentResponse, err error) {
	response = CreateDeleteTransitRouterPeerAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTransitRouterPeerAttachmentWithChan invokes the cbn.DeleteTransitRouterPeerAttachment API asynchronously
func (client *Client) DeleteTransitRouterPeerAttachmentWithChan(request *DeleteTransitRouterPeerAttachmentRequest) (<-chan *DeleteTransitRouterPeerAttachmentResponse, <-chan error) {
	responseChan := make(chan *DeleteTransitRouterPeerAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTransitRouterPeerAttachment(request)
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

// DeleteTransitRouterPeerAttachmentWithCallback invokes the cbn.DeleteTransitRouterPeerAttachment API asynchronously
func (client *Client) DeleteTransitRouterPeerAttachmentWithCallback(request *DeleteTransitRouterPeerAttachmentRequest, callback func(response *DeleteTransitRouterPeerAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTransitRouterPeerAttachmentResponse
		var err error
		defer close(result)
		response, err = client.DeleteTransitRouterPeerAttachment(request)
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

// DeleteTransitRouterPeerAttachmentRequest is the request struct for api DeleteTransitRouterPeerAttachment
type DeleteTransitRouterPeerAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType              string           `position:"Query" name:"ResourceType"`
	Version                   string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId string           `position:"Query" name:"TransitRouterAttachmentId"`
}

// DeleteTransitRouterPeerAttachmentResponse is the response struct for api DeleteTransitRouterPeerAttachment
type DeleteTransitRouterPeerAttachmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTransitRouterPeerAttachmentRequest creates a request to invoke DeleteTransitRouterPeerAttachment API
func CreateDeleteTransitRouterPeerAttachmentRequest() (request *DeleteTransitRouterPeerAttachmentRequest) {
	request = &DeleteTransitRouterPeerAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteTransitRouterPeerAttachment", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTransitRouterPeerAttachmentResponse creates a response to parse from DeleteTransitRouterPeerAttachment response
func CreateDeleteTransitRouterPeerAttachmentResponse() (response *DeleteTransitRouterPeerAttachmentResponse) {
	response = &DeleteTransitRouterPeerAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
