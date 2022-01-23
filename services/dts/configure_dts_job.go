package dts

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

// ConfigureDtsJob invokes the dts.ConfigureDtsJob API synchronously
func (client *Client) ConfigureDtsJob(request *ConfigureDtsJobRequest) (response *ConfigureDtsJobResponse, err error) {
	response = CreateConfigureDtsJobResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureDtsJobWithChan invokes the dts.ConfigureDtsJob API asynchronously
func (client *Client) ConfigureDtsJobWithChan(request *ConfigureDtsJobRequest) (<-chan *ConfigureDtsJobResponse, <-chan error) {
	responseChan := make(chan *ConfigureDtsJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureDtsJob(request)
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

// ConfigureDtsJobWithCallback invokes the dts.ConfigureDtsJob API asynchronously
func (client *Client) ConfigureDtsJobWithCallback(request *ConfigureDtsJobRequest, callback func(response *ConfigureDtsJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureDtsJobResponse
		var err error
		defer close(result)
		response, err = client.ConfigureDtsJob(request)
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

// ConfigureDtsJobRequest is the request struct for api ConfigureDtsJob
type ConfigureDtsJobRequest struct {
	*requests.RpcRequest
	Checkpoint                      string           `position:"Query" name:"Checkpoint"`
	SourceEndpointUserName          string           `position:"Query" name:"SourceEndpointUserName"`
	DelayPhone                      string           `position:"Query" name:"DelayPhone"`
	SourceEndpointIP                string           `position:"Query" name:"SourceEndpointIP"`
	ErrorPhone                      string           `position:"Query" name:"ErrorPhone"`
	DestinationEndpointUserName     string           `position:"Query" name:"DestinationEndpointUserName"`
	DtsJobId                        string           `position:"Query" name:"DtsJobId"`
	DbList                          string           `position:"Body" name:"DbList"`
	DestinationEndpointOracleSID    string           `position:"Query" name:"DestinationEndpointOracleSID"`
	DestinationEndpointPort         string           `position:"Query" name:"DestinationEndpointPort"`
	SourceEndpointPassword          string           `position:"Query" name:"SourceEndpointPassword"`
	OwnerId                         string           `position:"Query" name:"OwnerId"`
	JobType                         string           `position:"Query" name:"JobType"`
	DelayRuleTime                   requests.Integer `position:"Query" name:"DelayRuleTime"`
	DestinationEndpointIP           string           `position:"Query" name:"DestinationEndpointIP"`
	SourceEndpointInstanceType      string           `position:"Query" name:"SourceEndpointInstanceType"`
	DtsJobName                      string           `position:"Query" name:"DtsJobName"`
	DtsInstanceId                   string           `position:"Query" name:"DtsInstanceId"`
	SynchronizationDirection        string           `position:"Query" name:"SynchronizationDirection"`
	SourceEndpointRegion            string           `position:"Query" name:"SourceEndpointRegion"`
	DelayNotice                     requests.Boolean `position:"Query" name:"DelayNotice"`
	DestinationEndpointInstanceType string           `position:"Query" name:"DestinationEndpointInstanceType"`
	DataInitialization              requests.Boolean `position:"Query" name:"DataInitialization"`
	SourceEndpointInstanceID        string           `position:"Query" name:"SourceEndpointInstanceID"`
	StructureInitialization         requests.Boolean `position:"Query" name:"StructureInitialization"`
	SourceEndpointOwnerID           string           `position:"Query" name:"SourceEndpointOwnerID"`
	SourceEndpointDatabaseName      string           `position:"Query" name:"SourceEndpointDatabaseName"`
	DestinationEndpointRegion       string           `position:"Query" name:"DestinationEndpointRegion"`
	Reserve                         string           `position:"Body" name:"Reserve"`
	DataSynchronization             requests.Boolean `position:"Query" name:"DataSynchronization"`
	DestinationEndpointEngineName   string           `position:"Query" name:"DestinationEndpointEngineName"`
	DestinationEndpointInstanceID   string           `position:"Query" name:"DestinationEndpointInstanceID"`
	SourceEndpointPort              string           `position:"Query" name:"SourceEndpointPort"`
	SourceEndpointOracleSID         string           `position:"Query" name:"SourceEndpointOracleSID"`
	DestinationEndpointDataBaseName string           `position:"Query" name:"DestinationEndpointDataBaseName"`
	ErrorNotice                     requests.Boolean `position:"Query" name:"ErrorNotice"`
	SourceEndpointRole              string           `position:"Query" name:"SourceEndpointRole"`
	DestinationEndpointPassword     string           `position:"Query" name:"DestinationEndpointPassword"`
	SourceEndpointEngineName        string           `position:"Query" name:"SourceEndpointEngineName"`
}

// ConfigureDtsJobResponse is the response struct for api ConfigureDtsJob
type ConfigureDtsJobResponse struct {
	*responses.BaseResponse
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	DtsJobId       string `json:"DtsJobId" xml:"DtsJobId"`
	DtsInstanceId  string `json:"DtsInstanceId" xml:"DtsInstanceId"`
	Success        string `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateConfigureDtsJobRequest creates a request to invoke ConfigureDtsJob API
func CreateConfigureDtsJobRequest() (request *ConfigureDtsJobRequest) {
	request = &ConfigureDtsJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ConfigureDtsJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigureDtsJobResponse creates a response to parse from ConfigureDtsJob response
func CreateConfigureDtsJobResponse() (response *ConfigureDtsJobResponse) {
	response = &ConfigureDtsJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
