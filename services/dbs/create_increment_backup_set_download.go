package dbs

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

// CreateIncrementBackupSetDownload invokes the dbs.CreateIncrementBackupSetDownload API synchronously
func (client *Client) CreateIncrementBackupSetDownload(request *CreateIncrementBackupSetDownloadRequest) (response *CreateIncrementBackupSetDownloadResponse, err error) {
	response = CreateCreateIncrementBackupSetDownloadResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIncrementBackupSetDownloadWithChan invokes the dbs.CreateIncrementBackupSetDownload API asynchronously
func (client *Client) CreateIncrementBackupSetDownloadWithChan(request *CreateIncrementBackupSetDownloadRequest) (<-chan *CreateIncrementBackupSetDownloadResponse, <-chan error) {
	responseChan := make(chan *CreateIncrementBackupSetDownloadResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIncrementBackupSetDownload(request)
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

// CreateIncrementBackupSetDownloadWithCallback invokes the dbs.CreateIncrementBackupSetDownload API asynchronously
func (client *Client) CreateIncrementBackupSetDownloadWithCallback(request *CreateIncrementBackupSetDownloadRequest, callback func(response *CreateIncrementBackupSetDownloadResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIncrementBackupSetDownloadResponse
		var err error
		defer close(result)
		response, err = client.CreateIncrementBackupSetDownload(request)
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

// CreateIncrementBackupSetDownloadRequest is the request struct for api CreateIncrementBackupSetDownload
type CreateIncrementBackupSetDownloadRequest struct {
	*requests.RpcRequest
	ClientToken         string `position:"Query" name:"ClientToken"`
	BackupSetName       string `position:"Query" name:"BackupSetName"`
	BackupSetId         string `position:"Query" name:"BackupSetId"`
	OwnerId             string `position:"Query" name:"OwnerId"`
	BackupSetDataFormat string `position:"Query" name:"BackupSetDataFormat"`
}

// CreateIncrementBackupSetDownloadResponse is the response struct for api CreateIncrementBackupSetDownload
type CreateIncrementBackupSetDownloadResponse struct {
	*responses.BaseResponse
	HttpStatusCode          int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId               string `json:"RequestId" xml:"RequestId"`
	ErrCode                 string `json:"ErrCode" xml:"ErrCode"`
	Success                 bool   `json:"Success" xml:"Success"`
	ErrMessage              string `json:"ErrMessage" xml:"ErrMessage"`
	BackupSetDownloadTaskId string `json:"BackupSetDownloadTaskId" xml:"BackupSetDownloadTaskId"`
}

// CreateCreateIncrementBackupSetDownloadRequest creates a request to invoke CreateIncrementBackupSetDownload API
func CreateCreateIncrementBackupSetDownloadRequest() (request *CreateIncrementBackupSetDownloadRequest) {
	request = &CreateIncrementBackupSetDownloadRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "CreateIncrementBackupSetDownload", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateIncrementBackupSetDownloadResponse creates a response to parse from CreateIncrementBackupSetDownload response
func CreateCreateIncrementBackupSetDownloadResponse() (response *CreateIncrementBackupSetDownloadResponse) {
	response = &CreateIncrementBackupSetDownloadResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
