package domain

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

// ScrollDomainList invokes the domain.ScrollDomainList API synchronously
func (client *Client) ScrollDomainList(request *ScrollDomainListRequest) (response *ScrollDomainListResponse, err error) {
	response = CreateScrollDomainListResponse()
	err = client.DoAction(request, response)
	return
}

// ScrollDomainListWithChan invokes the domain.ScrollDomainList API asynchronously
func (client *Client) ScrollDomainListWithChan(request *ScrollDomainListRequest) (<-chan *ScrollDomainListResponse, <-chan error) {
	responseChan := make(chan *ScrollDomainListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScrollDomainList(request)
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

// ScrollDomainListWithCallback invokes the domain.ScrollDomainList API asynchronously
func (client *Client) ScrollDomainListWithCallback(request *ScrollDomainListRequest, callback func(response *ScrollDomainListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScrollDomainListResponse
		var err error
		defer close(result)
		response, err = client.ScrollDomainList(request)
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

// ScrollDomainListRequest is the request struct for api ScrollDomainList
type ScrollDomainListRequest struct {
	*requests.RpcRequest
	ProductDomainType     string           `position:"Query" name:"ProductDomainType"`
	Excluded              string           `position:"Query" name:"Excluded"`
	StartLength           requests.Integer `position:"Query" name:"StartLength"`
	ExcludedSuffix        requests.Boolean `position:"Query" name:"ExcludedSuffix"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	Lang                  string           `position:"Query" name:"Lang"`
	ExcludedPrefix        requests.Boolean `position:"Query" name:"ExcludedPrefix"`
	KeyWord               string           `position:"Query" name:"KeyWord"`
	EndExpirationDate     requests.Integer `position:"Query" name:"EndExpirationDate"`
	Suffixs               string           `position:"Query" name:"Suffixs"`
	StartExpirationDate   requests.Integer `position:"Query" name:"StartExpirationDate"`
	DomainStatus          requests.Integer `position:"Query" name:"DomainStatus"`
	DomainGroupId         requests.Integer `position:"Query" name:"DomainGroupId"`
	KeyWordSuffix         requests.Boolean `position:"Query" name:"KeyWordSuffix"`
	ScrollId              string           `position:"Query" name:"ScrollId"`
	KeyWordPrefix         requests.Boolean `position:"Query" name:"KeyWordPrefix"`
	TradeType             requests.Integer `position:"Query" name:"TradeType"`
	EndRegistrationDate   requests.Integer `position:"Query" name:"EndRegistrationDate"`
	Form                  requests.Integer `position:"Query" name:"Form"`
	UserClientIp          string           `position:"Query" name:"UserClientIp"`
	StartRegistrationDate requests.Integer `position:"Query" name:"StartRegistrationDate"`
	EndLength             requests.Integer `position:"Query" name:"EndLength"`
}

// ScrollDomainListResponse is the response struct for api ScrollDomainList
type ScrollDomainListResponse struct {
	*responses.BaseResponse
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	PageSize     int                    `json:"PageSize" xml:"PageSize"`
	ScrollId     string                 `json:"ScrollId" xml:"ScrollId"`
	TotalItemNum int                    `json:"TotalItemNum" xml:"TotalItemNum"`
	Data         DataInScrollDomainList `json:"Data" xml:"Data"`
}

// CreateScrollDomainListRequest creates a request to invoke ScrollDomainList API
func CreateScrollDomainListRequest() (request *ScrollDomainListRequest) {
	request = &ScrollDomainListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "ScrollDomainList", "", "")
	request.Method = requests.POST
	return
}

// CreateScrollDomainListResponse creates a response to parse from ScrollDomainList response
func CreateScrollDomainListResponse() (response *ScrollDomainListResponse) {
	response = &ScrollDomainListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
