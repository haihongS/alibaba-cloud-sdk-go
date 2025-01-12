package rds

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

// RequestServiceOfCloudDBA invokes the rds.RequestServiceOfCloudDBA API synchronously
func (client *Client) RequestServiceOfCloudDBA(request *RequestServiceOfCloudDBARequest) (response *RequestServiceOfCloudDBAResponse, err error) {
	response = CreateRequestServiceOfCloudDBAResponse()
	err = client.DoAction(request, response)
	return
}

// RequestServiceOfCloudDBAWithChan invokes the rds.RequestServiceOfCloudDBA API asynchronously
func (client *Client) RequestServiceOfCloudDBAWithChan(request *RequestServiceOfCloudDBARequest) (<-chan *RequestServiceOfCloudDBAResponse, <-chan error) {
	responseChan := make(chan *RequestServiceOfCloudDBAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestServiceOfCloudDBA(request)
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

// RequestServiceOfCloudDBAWithCallback invokes the rds.RequestServiceOfCloudDBA API asynchronously
func (client *Client) RequestServiceOfCloudDBAWithCallback(request *RequestServiceOfCloudDBARequest, callback func(response *RequestServiceOfCloudDBAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestServiceOfCloudDBAResponse
		var err error
		defer close(result)
		response, err = client.RequestServiceOfCloudDBA(request)
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

// RequestServiceOfCloudDBARequest is the request struct for api RequestServiceOfCloudDBA
type RequestServiceOfCloudDBARequest struct {
	*requests.RpcRequest
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
}

// RequestServiceOfCloudDBAResponse is the response struct for api RequestServiceOfCloudDBA
type RequestServiceOfCloudDBAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AttrData  string `json:"AttrData" xml:"AttrData"`
	ListData  string `json:"ListData" xml:"ListData"`
}

// CreateRequestServiceOfCloudDBARequest creates a request to invoke RequestServiceOfCloudDBA API
func CreateRequestServiceOfCloudDBARequest() (request *RequestServiceOfCloudDBARequest) {
	request = &RequestServiceOfCloudDBARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RequestServiceOfCloudDBA", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRequestServiceOfCloudDBAResponse creates a response to parse from RequestServiceOfCloudDBA response
func CreateRequestServiceOfCloudDBAResponse() (response *RequestServiceOfCloudDBAResponse) {
	response = &RequestServiceOfCloudDBAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
