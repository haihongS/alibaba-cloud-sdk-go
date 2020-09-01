package cloudauth

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

// InitSmartVerify invokes the cloudauth.InitSmartVerify API synchronously
// api document: https://help.aliyun.com/api/cloudauth/initsmartverify.html
func (client *Client) InitSmartVerify(request *InitSmartVerifyRequest) (response *InitSmartVerifyResponse, err error) {
	response = CreateInitSmartVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// InitSmartVerifyWithChan invokes the cloudauth.InitSmartVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/initsmartverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InitSmartVerifyWithChan(request *InitSmartVerifyRequest) (<-chan *InitSmartVerifyResponse, <-chan error) {
	responseChan := make(chan *InitSmartVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitSmartVerify(request)
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

// InitSmartVerifyWithCallback invokes the cloudauth.InitSmartVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/initsmartverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InitSmartVerifyWithCallback(request *InitSmartVerifyRequest, callback func(response *InitSmartVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitSmartVerifyResponse
		var err error
		defer close(result)
		response, err = client.InitSmartVerify(request)
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

// InitSmartVerifyRequest is the request struct for api InitSmartVerify
type InitSmartVerifyRequest struct {
	*requests.RpcRequest
	Ip           string           `position:"Body" name:"Ip"`
	CertName     string           `position:"Body" name:"CertName"`
	Mobile       string           `position:"Body" name:"Mobile"`
	UserId       string           `position:"Body" name:"UserId"`
	Mode         string           `position:"Body" name:"Mode"`
	CertNo       string           `position:"Body" name:"CertNo"`
	OuterOrderNo string           `position:"Body" name:"OuterOrderNo"`
	CertType     string           `position:"Body" name:"CertType"`
	SceneId      requests.Integer `position:"Body" name:"SceneId"`
	MetaInfo     string           `position:"Body" name:"MetaInfo"`
	Ocr          string           `position:"Body" name:"Ocr"`
}

// InitSmartVerifyResponse is the response struct for api InitSmartVerify
type InitSmartVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateInitSmartVerifyRequest creates a request to invoke InitSmartVerify API
func CreateInitSmartVerifyRequest() (request *InitSmartVerifyRequest) {
	request = &InitSmartVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2020-06-18", "InitSmartVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInitSmartVerifyResponse creates a response to parse from InitSmartVerify response
func CreateInitSmartVerifyResponse() (response *InitSmartVerifyResponse) {
	response = &InitSmartVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
