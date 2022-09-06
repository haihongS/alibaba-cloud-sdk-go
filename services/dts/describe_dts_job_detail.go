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

// DescribeDtsJobDetail invokes the dts.DescribeDtsJobDetail API synchronously
func (client *Client) DescribeDtsJobDetail(request *DescribeDtsJobDetailRequest) (response *DescribeDtsJobDetailResponse, err error) {
	response = CreateDescribeDtsJobDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDtsJobDetailWithChan invokes the dts.DescribeDtsJobDetail API asynchronously
func (client *Client) DescribeDtsJobDetailWithChan(request *DescribeDtsJobDetailRequest) (<-chan *DescribeDtsJobDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeDtsJobDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDtsJobDetail(request)
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

// DescribeDtsJobDetailWithCallback invokes the dts.DescribeDtsJobDetail API asynchronously
func (client *Client) DescribeDtsJobDetailWithCallback(request *DescribeDtsJobDetailRequest, callback func(response *DescribeDtsJobDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDtsJobDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeDtsJobDetail(request)
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

// DescribeDtsJobDetailRequest is the request struct for api DescribeDtsJobDetail
type DescribeDtsJobDetailRequest struct {
	*requests.RpcRequest
	SyncSubJobHistory        requests.Boolean `position:"Query" name:"SyncSubJobHistory"`
	DtsJobId                 string           `position:"Query" name:"DtsJobId"`
	DtsInstanceID            string           `position:"Query" name:"DtsInstanceID"`
	SynchronizationDirection string           `position:"Query" name:"SynchronizationDirection"`
}

// DescribeDtsJobDetailResponse is the response struct for api DescribeDtsJobDetail
type DescribeDtsJobDetailResponse struct {
	*responses.BaseResponse
	Status                    string                    `json:"Status" xml:"Status"`
	DtsJobName                string                    `json:"DtsJobName" xml:"DtsJobName"`
	FinishTime                string                    `json:"FinishTime" xml:"FinishTime"`
	ErrorMessage              string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	DtsJobId                  string                    `json:"DtsJobId" xml:"DtsJobId"`
	CreateTime                string                    `json:"CreateTime" xml:"CreateTime"`
	PayType                   string                    `json:"PayType" xml:"PayType"`
	Reserved                  string                    `json:"Reserved" xml:"Reserved"`
	DatabaseCount             int                       `json:"DatabaseCount" xml:"DatabaseCount"`
	DtsJobClass               string                    `json:"DtsJobClass" xml:"DtsJobClass"`
	EndTimestamp              string                    `json:"EndTimestamp" xml:"EndTimestamp"`
	AppName                   string                    `json:"AppName" xml:"AppName"`
	DestNetType               string                    `json:"DestNetType" xml:"DestNetType"`
	SubscribeTopic            string                    `json:"SubscribeTopic" xml:"SubscribeTopic"`
	DtsInstanceID             string                    `json:"DtsInstanceID" xml:"DtsInstanceID"`
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	Code                      int                       `json:"Code" xml:"Code"`
	Checkpoint                int64                     `json:"Checkpoint" xml:"Checkpoint"`
	Delay                     int64                     `json:"Delay" xml:"Delay"`
	ExpireTime                string                    `json:"ExpireTime" xml:"ExpireTime"`
	ErrCode                   string                    `json:"ErrCode" xml:"ErrCode"`
	Success                   bool                      `json:"Success" xml:"Success"`
	ErrMessage                string                    `json:"ErrMessage" xml:"ErrMessage"`
	ConsumptionClient         string                    `json:"ConsumptionClient" xml:"ConsumptionClient"`
	DbObject                  string                    `json:"DbObject" xml:"DbObject"`
	DynamicMessage            string                    `json:"DynamicMessage" xml:"DynamicMessage"`
	ConsumptionCheckpoint     string                    `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	EtlCalculator             string                    `json:"EtlCalculator" xml:"EtlCalculator"`
	HttpStatusCode            int                       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	BeginTimestamp            string                    `json:"BeginTimestamp" xml:"BeginTimestamp"`
	GroupId                   string                    `json:"GroupId" xml:"GroupId"`
	SynchronizationDirection  string                    `json:"SynchronizationDirection" xml:"SynchronizationDirection"`
	DtsJobDirection           string                    `json:"DtsJobDirection" xml:"DtsJobDirection"`
	DemoJob                   bool                      `json:"DemoJob" xml:"DemoJob"`
	JobType                   string                    `json:"JobType" xml:"JobType"`
	TaskType                  string                    `json:"TaskType" xml:"TaskType"`
	DedicatedClusterId        string                    `json:"DedicatedClusterId" xml:"DedicatedClusterId"`
	BootTime                  string                    `json:"BootTime" xml:"BootTime"`
	Binlog                    string                    `json:"Binlog" xml:"Binlog"`
	BinlogSite                string                    `json:"BinlogSite" xml:"BinlogSite"`
	BinlogTime                string                    `json:"BinlogTime" xml:"BinlogTime"`
	LastUpdateTime            string                    `json:"LastUpdateTime" xml:"LastUpdateTime"`
	SourceEndpoint            SourceEndpoint            `json:"SourceEndpoint" xml:"SourceEndpoint"`
	DestinationEndpoint       DestinationEndpoint       `json:"DestinationEndpoint" xml:"DestinationEndpoint"`
	MigrationMode             MigrationMode             `json:"MigrationMode" xml:"MigrationMode"`
	SubscriptionHost          SubscriptionHost          `json:"SubscriptionHost" xml:"SubscriptionHost"`
	SubscriptionDataType      SubscriptionDataType      `json:"SubscriptionDataType" xml:"SubscriptionDataType"`
	DataSynchronizationStatus DataSynchronizationStatus `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	SubDistributedJob         []SubDistributedJobItem   `json:"SubDistributedJob" xml:"SubDistributedJob"`
	SubSyncJob                []SubSyncJobItem          `json:"SubSyncJob" xml:"SubSyncJob"`
}

// CreateDescribeDtsJobDetailRequest creates a request to invoke DescribeDtsJobDetail API
func CreateDescribeDtsJobDetailRequest() (request *DescribeDtsJobDetailRequest) {
	request = &DescribeDtsJobDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeDtsJobDetail", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDtsJobDetailResponse creates a response to parse from DescribeDtsJobDetail response
func CreateDescribeDtsJobDetailResponse() (response *DescribeDtsJobDetailResponse) {
	response = &DescribeDtsJobDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
