package ft

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

// BatchAuditTest01 invokes the ft.BatchAuditTest01 API synchronously
func (client *Client) BatchAuditTest01(request *BatchAuditTest01Request) (response *BatchAuditTest01Response, err error) {
	response = CreateBatchAuditTest01Response()
	err = client.DoAction(request, response)
	return
}

// BatchAuditTest01WithChan invokes the ft.BatchAuditTest01 API asynchronously
func (client *Client) BatchAuditTest01WithChan(request *BatchAuditTest01Request) (<-chan *BatchAuditTest01Response, <-chan error) {
	responseChan := make(chan *BatchAuditTest01Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchAuditTest01(request)
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

// BatchAuditTest01WithCallback invokes the ft.BatchAuditTest01 API asynchronously
func (client *Client) BatchAuditTest01WithCallback(request *BatchAuditTest01Request, callback func(response *BatchAuditTest01Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchAuditTest01Response
		var err error
		defer close(result)
		response, err = client.BatchAuditTest01(request)
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

// BatchAuditTest01Request is the request struct for api BatchAuditTest01
type BatchAuditTest01Request struct {
	*requests.RpcRequest
	Demo01           string           `position:"Query" name:"Demo01"`
	Test010101       requests.Boolean `position:"Body" name:"Test010101"`
	Name             string           `position:"Query" name:"Name"`
	BatchAuditTest01 string           `position:"Query" name:"BatchAuditTest01"`
}

// BatchAuditTest01Response is the response struct for api BatchAuditTest01
type BatchAuditTest01Response struct {
	*responses.BaseResponse
	Name      string `json:"Name" xml:"Name"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Demo01    Demo01 `json:"Demo01" xml:"Demo01"`
}

// CreateBatchAuditTest01Request creates a request to invoke BatchAuditTest01 API
func CreateBatchAuditTest01Request() (request *BatchAuditTest01Request) {
	request = &BatchAuditTest01Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "BatchAuditTest01", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchAuditTest01Response creates a response to parse from BatchAuditTest01 response
func CreateBatchAuditTest01Response() (response *BatchAuditTest01Response) {
	response = &BatchAuditTest01Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
