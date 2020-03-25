package rtc

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

// DescribeMAURule invokes the rtc.DescribeMAURule API synchronously
// api document: https://help.aliyun.com/api/rtc/describemaurule.html
func (client *Client) DescribeMAURule(request *DescribeMAURuleRequest) (response *DescribeMAURuleResponse, err error) {
	response = CreateDescribeMAURuleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMAURuleWithChan invokes the rtc.DescribeMAURule API asynchronously
// api document: https://help.aliyun.com/api/rtc/describemaurule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMAURuleWithChan(request *DescribeMAURuleRequest) (<-chan *DescribeMAURuleResponse, <-chan error) {
	responseChan := make(chan *DescribeMAURuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMAURule(request)
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

// DescribeMAURuleWithCallback invokes the rtc.DescribeMAURule API asynchronously
// api document: https://help.aliyun.com/api/rtc/describemaurule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMAURuleWithCallback(request *DescribeMAURuleRequest, callback func(response *DescribeMAURuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMAURuleResponse
		var err error
		defer close(result)
		response, err = client.DescribeMAURule(request)
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

// DescribeMAURuleRequest is the request struct for api DescribeMAURule
type DescribeMAURuleRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// DescribeMAURuleResponse is the response struct for api DescribeMAURule
type DescribeMAURuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Rules     []Rule `json:"Rules" xml:"Rules"`
}

// CreateDescribeMAURuleRequest creates a request to invoke DescribeMAURule API
func CreateDescribeMAURuleRequest() (request *DescribeMAURuleRequest) {
	request = &DescribeMAURuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeMAURule", "rtc", "openAPI")
	return
}

// CreateDescribeMAURuleResponse creates a response to parse from DescribeMAURule response
func CreateDescribeMAURuleResponse() (response *DescribeMAURuleResponse) {
	response = &DescribeMAURuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
