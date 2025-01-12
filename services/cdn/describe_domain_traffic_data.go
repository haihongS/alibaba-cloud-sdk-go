package cdn

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

// DescribeDomainTrafficData invokes the cdn.DescribeDomainTrafficData API synchronously
func (client *Client) DescribeDomainTrafficData(request *DescribeDomainTrafficDataRequest) (response *DescribeDomainTrafficDataResponse, err error) {
	response = CreateDescribeDomainTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainTrafficDataWithChan invokes the cdn.DescribeDomainTrafficData API asynchronously
func (client *Client) DescribeDomainTrafficDataWithChan(request *DescribeDomainTrafficDataRequest) (<-chan *DescribeDomainTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainTrafficData(request)
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

// DescribeDomainTrafficDataWithCallback invokes the cdn.DescribeDomainTrafficData API asynchronously
func (client *Client) DescribeDomainTrafficDataWithCallback(request *DescribeDomainTrafficDataRequest, callback func(response *DescribeDomainTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainTrafficData(request)
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

// DescribeDomainTrafficDataRequest is the request struct for api DescribeDomainTrafficData
type DescribeDomainTrafficDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDomainTrafficDataResponse is the response struct for api DescribeDomainTrafficData
type DescribeDomainTrafficDataResponse struct {
	*responses.BaseResponse
	EndTime                string                 `json:"EndTime" xml:"EndTime"`
	StartTime              string                 `json:"StartTime" xml:"StartTime"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	DomainName             string                 `json:"DomainName" xml:"DomainName"`
	DataInterval           string                 `json:"DataInterval" xml:"DataInterval"`
	TrafficDataPerInterval TrafficDataPerInterval `json:"TrafficDataPerInterval" xml:"TrafficDataPerInterval"`
}

// CreateDescribeDomainTrafficDataRequest creates a request to invoke DescribeDomainTrafficData API
func CreateDescribeDomainTrafficDataRequest() (request *DescribeDomainTrafficDataRequest) {
	request = &DescribeDomainTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainTrafficData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainTrafficDataResponse creates a response to parse from DescribeDomainTrafficData response
func CreateDescribeDomainTrafficDataResponse() (response *DescribeDomainTrafficDataResponse) {
	response = &DescribeDomainTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
