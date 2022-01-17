package companyreg

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

// GetPayrollInfo invokes the companyreg.GetPayrollInfo API synchronously
func (client *Client) GetPayrollInfo(request *GetPayrollInfoRequest) (response *GetPayrollInfoResponse, err error) {
	response = CreateGetPayrollInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetPayrollInfoWithChan invokes the companyreg.GetPayrollInfo API asynchronously
func (client *Client) GetPayrollInfoWithChan(request *GetPayrollInfoRequest) (<-chan *GetPayrollInfoResponse, <-chan error) {
	responseChan := make(chan *GetPayrollInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPayrollInfo(request)
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

// GetPayrollInfoWithCallback invokes the companyreg.GetPayrollInfo API asynchronously
func (client *Client) GetPayrollInfoWithCallback(request *GetPayrollInfoRequest, callback func(response *GetPayrollInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPayrollInfoResponse
		var err error
		defer close(result)
		response, err = client.GetPayrollInfo(request)
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

// GetPayrollInfoRequest is the request struct for api GetPayrollInfo
type GetPayrollInfoRequest struct {
	*requests.RpcRequest
	BizId string           `position:"Query" name:"BizId"`
	Id    requests.Integer `position:"Query" name:"Id"`
}

// GetPayrollInfoResponse is the response struct for api GetPayrollInfo
type GetPayrollInfoResponse struct {
	*responses.BaseResponse
	CorporateHousingAccumulationFund string `json:"CorporateHousingAccumulationFund" xml:"CorporateHousingAccumulationFund"`
	CorporateSocialInsurance         string `json:"CorporateSocialInsurance" xml:"CorporateSocialInsurance"`
	EmployeeTime                     string `json:"EmployeeTime" xml:"EmployeeTime"`
	Id                               int64  `json:"Id" xml:"Id"`
	IdNo                             string `json:"IdNo" xml:"IdNo"`
	Income                           string `json:"Income" xml:"Income"`
	Name                             string `json:"Name" xml:"Name"`
	Period                           string `json:"Period" xml:"Period"`
	PersonHousingAccumulationFund    string `json:"PersonHousingAccumulationFund" xml:"PersonHousingAccumulationFund"`
	PersonSocialInsurance            string `json:"PersonSocialInsurance" xml:"PersonSocialInsurance"`
	Phone                            string `json:"Phone" xml:"Phone"`
	CorpPension                      string `json:"CorpPension" xml:"CorpPension"`
	CorpMedicalInsurance             string `json:"CorpMedicalInsurance" xml:"CorpMedicalInsurance"`
	CorpUnemploymentInsurance        string `json:"CorpUnemploymentInsurance" xml:"CorpUnemploymentInsurance"`
	CorpInjuryInsurance              string `json:"CorpInjuryInsurance" xml:"CorpInjuryInsurance"`
	CorpMaternityInsurance           string `json:"CorpMaternityInsurance" xml:"CorpMaternityInsurance"`
	PersPension                      string `json:"PersPension" xml:"PersPension"`
	PersUnemploymentInsurance        string `json:"PersUnemploymentInsurance" xml:"PersUnemploymentInsurance"`
	PersMedicalInsurance             string `json:"PersMedicalInsurance" xml:"PersMedicalInsurance"`
	RequestId                        string `json:"RequestId" xml:"RequestId"`
}

// CreateGetPayrollInfoRequest creates a request to invoke GetPayrollInfo API
func CreateGetPayrollInfoRequest() (request *GetPayrollInfoRequest) {
	request = &GetPayrollInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "GetPayrollInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateGetPayrollInfoResponse creates a response to parse from GetPayrollInfo response
func CreateGetPayrollInfoResponse() (response *GetPayrollInfoResponse) {
	response = &GetPayrollInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
