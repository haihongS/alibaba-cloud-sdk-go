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

// DescribeDomainRealTimeHttpCodeData invokes the cdn.DescribeDomainRealTimeHttpCodeData API synchronously
func (client *Client) DescribeDomainRealTimeHttpCodeData(request *DescribeDomainRealTimeHttpCodeDataRequest) (response *DescribeDomainRealTimeHttpCodeDataResponse, err error) {
	response = CreateDescribeDomainRealTimeHttpCodeDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeHttpCodeDataWithChan invokes the cdn.DescribeDomainRealTimeHttpCodeData API asynchronously
func (client *Client) DescribeDomainRealTimeHttpCodeDataWithChan(request *DescribeDomainRealTimeHttpCodeDataRequest) (<-chan *DescribeDomainRealTimeHttpCodeDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeHttpCodeDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeHttpCodeData(request)
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

// DescribeDomainRealTimeHttpCodeDataWithCallback invokes the cdn.DescribeDomainRealTimeHttpCodeData API asynchronously
func (client *Client) DescribeDomainRealTimeHttpCodeDataWithCallback(request *DescribeDomainRealTimeHttpCodeDataRequest, callback func(response *DescribeDomainRealTimeHttpCodeDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeHttpCodeDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeHttpCodeData(request)
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

// DescribeDomainRealTimeHttpCodeDataRequest is the request struct for api DescribeDomainRealTimeHttpCodeData
type DescribeDomainRealTimeHttpCodeDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainRealTimeHttpCodeDataResponse is the response struct for api DescribeDomainRealTimeHttpCodeData
type DescribeDomainRealTimeHttpCodeDataResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	DomainName           string               `json:"DomainName" xml:"DomainName"`
	StartTime            string               `json:"StartTime" xml:"StartTime"`
	EndTime              string               `json:"EndTime" xml:"EndTime"`
	DataInterval         string               `json:"DataInterval" xml:"DataInterval"`
	RealTimeHttpCodeData RealTimeHttpCodeData `json:"RealTimeHttpCodeData" xml:"RealTimeHttpCodeData"`
}

// CreateDescribeDomainRealTimeHttpCodeDataRequest creates a request to invoke DescribeDomainRealTimeHttpCodeData API
func CreateDescribeDomainRealTimeHttpCodeDataRequest() (request *DescribeDomainRealTimeHttpCodeDataRequest) {
	request = &DescribeDomainRealTimeHttpCodeDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeHttpCodeData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainRealTimeHttpCodeDataResponse creates a response to parse from DescribeDomainRealTimeHttpCodeData response
func CreateDescribeDomainRealTimeHttpCodeDataResponse() (response *DescribeDomainRealTimeHttpCodeDataResponse) {
	response = &DescribeDomainRealTimeHttpCodeDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
