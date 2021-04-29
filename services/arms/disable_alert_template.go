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

// DisableAlertTemplate invokes the arms.DisableAlertTemplate API synchronously
func (client *Client) DisableAlertTemplate(request *DisableAlertTemplateRequest) (response *DisableAlertTemplateResponse, err error) {
	response = CreateDisableAlertTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DisableAlertTemplateWithChan invokes the arms.DisableAlertTemplate API asynchronously
func (client *Client) DisableAlertTemplateWithChan(request *DisableAlertTemplateRequest) (<-chan *DisableAlertTemplateResponse, <-chan error) {
	responseChan := make(chan *DisableAlertTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableAlertTemplate(request)
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

// DisableAlertTemplateWithCallback invokes the arms.DisableAlertTemplate API asynchronously
func (client *Client) DisableAlertTemplateWithCallback(request *DisableAlertTemplateRequest, callback func(response *DisableAlertTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableAlertTemplateResponse
		var err error
		defer close(result)
		response, err = client.DisableAlertTemplate(request)
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

// DisableAlertTemplateRequest is the request struct for api DisableAlertTemplate
type DisableAlertTemplateRequest struct {
	*requests.RpcRequest
	Id          requests.Integer `position:"Query" name:"Id"`
	ProxyUserId string           `position:"Query" name:"ProxyUserId"`
}

// DisableAlertTemplateResponse is the response struct for api DisableAlertTemplate
type DisableAlertTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDisableAlertTemplateRequest creates a request to invoke DisableAlertTemplate API
func CreateDisableAlertTemplateRequest() (request *DisableAlertTemplateRequest) {
	request = &DisableAlertTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DisableAlertTemplate", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableAlertTemplateResponse creates a response to parse from DisableAlertTemplate response
func CreateDisableAlertTemplateResponse() (response *DisableAlertTemplateResponse) {
	response = &DisableAlertTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
