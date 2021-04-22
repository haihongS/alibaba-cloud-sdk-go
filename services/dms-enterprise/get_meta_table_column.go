package dms_enterprise

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

// GetMetaTableColumn invokes the dms_enterprise.GetMetaTableColumn API synchronously
func (client *Client) GetMetaTableColumn(request *GetMetaTableColumnRequest) (response *GetMetaTableColumnResponse, err error) {
	response = CreateGetMetaTableColumnResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaTableColumnWithChan invokes the dms_enterprise.GetMetaTableColumn API asynchronously
func (client *Client) GetMetaTableColumnWithChan(request *GetMetaTableColumnRequest) (<-chan *GetMetaTableColumnResponse, <-chan error) {
	responseChan := make(chan *GetMetaTableColumnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaTableColumn(request)
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

// GetMetaTableColumnWithCallback invokes the dms_enterprise.GetMetaTableColumn API asynchronously
func (client *Client) GetMetaTableColumnWithCallback(request *GetMetaTableColumnRequest, callback func(response *GetMetaTableColumnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaTableColumnResponse
		var err error
		defer close(result)
		response, err = client.GetMetaTableColumn(request)
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

// GetMetaTableColumnRequest is the request struct for api GetMetaTableColumn
type GetMetaTableColumnRequest struct {
	*requests.RpcRequest
	TableGuid string           `position:"Query" name:"TableGuid"`
	Tid       requests.Integer `position:"Query" name:"Tid"`
}

// GetMetaTableColumnResponse is the response struct for api GetMetaTableColumn
type GetMetaTableColumnResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	ErrorCode    string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool     `json:"Success" xml:"Success"`
	ColumnList   []Column `json:"ColumnList" xml:"ColumnList"`
}

// CreateGetMetaTableColumnRequest creates a request to invoke GetMetaTableColumn API
func CreateGetMetaTableColumnRequest() (request *GetMetaTableColumnRequest) {
	request = &GetMetaTableColumnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetMetaTableColumn", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMetaTableColumnResponse creates a response to parse from GetMetaTableColumn response
func CreateGetMetaTableColumnResponse() (response *GetMetaTableColumnResponse) {
	response = &GetMetaTableColumnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
