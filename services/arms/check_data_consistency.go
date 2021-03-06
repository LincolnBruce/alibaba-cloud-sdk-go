package arms

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

// CheckDataConsistency invokes the arms.CheckDataConsistency API synchronously
// api document: https://help.aliyun.com/api/arms/checkdataconsistency.html
func (client *Client) CheckDataConsistency(request *CheckDataConsistencyRequest) (response *CheckDataConsistencyResponse, err error) {
	response = CreateCheckDataConsistencyResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDataConsistencyWithChan invokes the arms.CheckDataConsistency API asynchronously
// api document: https://help.aliyun.com/api/arms/checkdataconsistency.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDataConsistencyWithChan(request *CheckDataConsistencyRequest) (<-chan *CheckDataConsistencyResponse, <-chan error) {
	responseChan := make(chan *CheckDataConsistencyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDataConsistency(request)
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

// CheckDataConsistencyWithCallback invokes the arms.CheckDataConsistency API asynchronously
// api document: https://help.aliyun.com/api/arms/checkdataconsistency.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDataConsistencyWithCallback(request *CheckDataConsistencyRequest, callback func(response *CheckDataConsistencyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDataConsistencyResponse
		var err error
		defer close(result)
		response, err = client.CheckDataConsistency(request)
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

// CheckDataConsistencyRequest is the request struct for api CheckDataConsistency
type CheckDataConsistencyRequest struct {
	*requests.RpcRequest
	CurrentTimestamp requests.Integer `position:"Query" name:"CurrentTimestamp"`
	AppType          string           `position:"Query" name:"AppType"`
	Pid              string           `position:"Query" name:"Pid"`
	ProxyUserId      string           `position:"Query" name:"ProxyUserId"`
}

// CheckDataConsistencyResponse is the response struct for api CheckDataConsistency
type CheckDataConsistencyResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	IsDataConsistency bool   `json:"IsDataConsistency" xml:"IsDataConsistency"`
}

// CreateCheckDataConsistencyRequest creates a request to invoke CheckDataConsistency API
func CreateCheckDataConsistencyRequest() (request *CheckDataConsistencyRequest) {
	request = &CheckDataConsistencyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CheckDataConsistency", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckDataConsistencyResponse creates a response to parse from CheckDataConsistency response
func CreateCheckDataConsistencyResponse() (response *CheckDataConsistencyResponse) {
	response = &CheckDataConsistencyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
