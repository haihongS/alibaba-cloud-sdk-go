package faas

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

// QueryEcsImages invokes the faas.QueryEcsImages API synchronously
// api document: https://help.aliyun.com/api/faas/queryecsimages.html
func (client *Client) QueryEcsImages(request *QueryEcsImagesRequest) (response *QueryEcsImagesResponse, err error) {
	response = CreateQueryEcsImagesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEcsImagesWithChan invokes the faas.QueryEcsImages API asynchronously
// api document: https://help.aliyun.com/api/faas/queryecsimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEcsImagesWithChan(request *QueryEcsImagesRequest) (<-chan *QueryEcsImagesResponse, <-chan error) {
	responseChan := make(chan *QueryEcsImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEcsImages(request)
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

// QueryEcsImagesWithCallback invokes the faas.QueryEcsImages API asynchronously
// api document: https://help.aliyun.com/api/faas/queryecsimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEcsImagesWithCallback(request *QueryEcsImagesRequest, callback func(response *QueryEcsImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEcsImagesResponse
		var err error
		defer close(result)
		response, err = client.QueryEcsImages(request)
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

// QueryEcsImagesRequest is the request struct for api QueryEcsImages
type QueryEcsImagesRequest struct {
	*requests.RpcRequest
}

// QueryEcsImagesResponse is the response struct for api QueryEcsImages
type QueryEcsImagesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	ECSImages []ECSImage `json:"ECSImages" xml:"ECSImages"`
}

// CreateQueryEcsImagesRequest creates a request to invoke QueryEcsImages API
func CreateQueryEcsImagesRequest() (request *QueryEcsImagesRequest) {
	request = &QueryEcsImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "QueryEcsImages", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEcsImagesResponse creates a response to parse from QueryEcsImages response
func CreateQueryEcsImagesResponse() (response *QueryEcsImagesResponse) {
	response = &QueryEcsImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}