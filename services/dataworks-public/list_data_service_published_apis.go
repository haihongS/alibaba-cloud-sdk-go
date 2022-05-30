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

// ListDataServicePublishedApis invokes the dataworks_public.ListDataServicePublishedApis API synchronously
func (client *Client) ListDataServicePublishedApis(request *ListDataServicePublishedApisRequest) (response *ListDataServicePublishedApisResponse, err error) {
	response = CreateListDataServicePublishedApisResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataServicePublishedApisWithChan invokes the dataworks_public.ListDataServicePublishedApis API asynchronously
func (client *Client) ListDataServicePublishedApisWithChan(request *ListDataServicePublishedApisRequest) (<-chan *ListDataServicePublishedApisResponse, <-chan error) {
	responseChan := make(chan *ListDataServicePublishedApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataServicePublishedApis(request)
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

// ListDataServicePublishedApisWithCallback invokes the dataworks_public.ListDataServicePublishedApis API asynchronously
func (client *Client) ListDataServicePublishedApisWithCallback(request *ListDataServicePublishedApisRequest, callback func(response *ListDataServicePublishedApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataServicePublishedApisResponse
		var err error
		defer close(result)
		response, err = client.ListDataServicePublishedApis(request)
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

// ListDataServicePublishedApisRequest is the request struct for api ListDataServicePublishedApis
type ListDataServicePublishedApisRequest struct {
	*requests.RpcRequest
	ApiNameKeyword string           `position:"Body" name:"ApiNameKeyword"`
	ApiPathKeyword string           `position:"Body" name:"ApiPathKeyword"`
	CreatorId      string           `position:"Body" name:"CreatorId"`
	PageNumber     requests.Integer `position:"Body" name:"PageNumber"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	TenantId       requests.Integer `position:"Body" name:"TenantId"`
	ProjectId      requests.Integer `position:"Body" name:"ProjectId"`
}

// ListDataServicePublishedApisResponse is the response struct for api ListDataServicePublishedApis
type ListDataServicePublishedApisResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListDataServicePublishedApisRequest creates a request to invoke ListDataServicePublishedApis API
func CreateListDataServicePublishedApisRequest() (request *ListDataServicePublishedApisRequest) {
	request = &ListDataServicePublishedApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListDataServicePublishedApis", "", "")
	request.Method = requests.POST
	return
}

// CreateListDataServicePublishedApisResponse creates a response to parse from ListDataServicePublishedApis response
func CreateListDataServicePublishedApisResponse() (response *ListDataServicePublishedApisResponse) {
	response = &ListDataServicePublishedApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
