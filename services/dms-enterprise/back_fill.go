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

// BackFill invokes the dms_enterprise.BackFill API synchronously
func (client *Client) BackFill(request *BackFillRequest) (response *BackFillResponse, err error) {
	response = CreateBackFillResponse()
	err = client.DoAction(request, response)
	return
}

// BackFillWithChan invokes the dms_enterprise.BackFill API asynchronously
func (client *Client) BackFillWithChan(request *BackFillRequest) (<-chan *BackFillResponse, <-chan error) {
	responseChan := make(chan *BackFillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BackFill(request)
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

// BackFillWithCallback invokes the dms_enterprise.BackFill API asynchronously
func (client *Client) BackFillWithCallback(request *BackFillRequest, callback func(response *BackFillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BackFillResponse
		var err error
		defer close(result)
		response, err = client.BackFill(request)
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

// BackFillRequest is the request struct for api BackFill
type BackFillRequest struct {
	*requests.RpcRequest
	DagId             requests.Integer `position:"Query" name:"DagId"`
	Tid               requests.Integer `position:"Query" name:"Tid"`
	IsTriggerSubTree  requests.Boolean `position:"Query" name:"IsTriggerSubTree"`
	BackFillDateEnd   string           `position:"Query" name:"BackFillDateEnd"`
	HistoryDagId      requests.Integer `position:"Query" name:"HistoryDagId"`
	StartNodeIds      *[]string        `position:"Query" name:"StartNodeIds"  type:"Json"`
	BackFillDateBegin string           `position:"Query" name:"BackFillDateBegin"`
	BackFillDate      string           `position:"Query" name:"BackFillDate"`
	Asc               requests.Boolean `position:"Query" name:"Asc"`
	Interval          requests.Integer `position:"Query" name:"Interval"`
}

// BackFillResponse is the response struct for api BackFill
type BackFillResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	NodeId       int64  `json:"NodeId" xml:"NodeId"`
}

// CreateBackFillRequest creates a request to invoke BackFill API
func CreateBackFillRequest() (request *BackFillRequest) {
	request = &BackFillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "BackFill", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBackFillResponse creates a response to parse from BackFill response
func CreateBackFillResponse() (response *BackFillResponse) {
	response = &BackFillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
