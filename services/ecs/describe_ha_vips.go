package ecs

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

// DescribeHaVips invokes the ecs.DescribeHaVips API synchronously
func (client *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
	response = CreateDescribeHaVipsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHaVipsWithChan invokes the ecs.DescribeHaVips API asynchronously
func (client *Client) DescribeHaVipsWithChan(request *DescribeHaVipsRequest) (<-chan *DescribeHaVipsResponse, <-chan error) {
	responseChan := make(chan *DescribeHaVipsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHaVips(request)
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

// DescribeHaVipsWithCallback invokes the ecs.DescribeHaVips API asynchronously
func (client *Client) DescribeHaVipsWithCallback(request *DescribeHaVipsRequest, callback func(response *DescribeHaVipsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHaVipsResponse
		var err error
		defer close(result)
		response, err = client.DescribeHaVips(request)
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

// DescribeHaVipsRequest is the request struct for api DescribeHaVips
type DescribeHaVipsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer        `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer        `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer        `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer        `position:"Query" name:"OwnerId"`
	Filter               *[]DescribeHaVipsFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

// DescribeHaVipsFilter is a repeated param struct in DescribeHaVipsRequest
type DescribeHaVipsFilter struct {
	Value *[]string `name:"Value" type:"Repeated"`
	Key   string    `name:"Key"`
}

// DescribeHaVipsResponse is the response struct for api DescribeHaVips
type DescribeHaVipsResponse struct {
	*responses.BaseResponse
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	HaVips     HaVips `json:"HaVips" xml:"HaVips"`
}

// CreateDescribeHaVipsRequest creates a request to invoke DescribeHaVips API
func CreateDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
	request = &DescribeHaVipsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeHaVips", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeHaVipsResponse creates a response to parse from DescribeHaVips response
func CreateDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
	response = &DescribeHaVipsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
