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

// DtsJobStatus is a nested struct in dts response
type DtsJobStatus struct {
	Status                        string                        `json:"Status" xml:"Status"`
	DtsJobName                    string                        `json:"DtsJobName" xml:"DtsJobName"`
	Delay                         int64                         `json:"Delay" xml:"Delay"`
	ErrorMessage                  string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	ExpireTime                    string                        `json:"ExpireTime" xml:"ExpireTime"`
	DtsJobId                      string                        `json:"DtsJobId" xml:"DtsJobId"`
	CreateTime                    string                        `json:"CreateTime" xml:"CreateTime"`
	PayType                       string                        `json:"PayType" xml:"PayType"`
	Reserved                      string                        `json:"Reserved" xml:"Reserved"`
	ConsumptionClient             string                        `json:"ConsumptionClient" xml:"ConsumptionClient"`
	DbObject                      string                        `json:"DbObject" xml:"DbObject"`
	DtsJobClass                   string                        `json:"DtsJobClass" xml:"DtsJobClass"`
	ConsumptionCheckpoint         string                        `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	EndTimestamp                  string                        `json:"EndTimestamp" xml:"EndTimestamp"`
	AppName                       string                        `json:"AppName" xml:"AppName"`
	BeginTimestamp                string                        `json:"BeginTimestamp" xml:"BeginTimestamp"`
	DtsInstanceID                 string                        `json:"DtsInstanceID" xml:"DtsInstanceID"`
	DtsJobDirection               string                        `json:"DtsJobDirection" xml:"DtsJobDirection"`
	Checkpoint                    string                        `json:"Checkpoint" xml:"Checkpoint"`
	DataInitializationStatus      DataInitializationStatus      `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	DataSynchronizationStatus     DataSynchronizationStatus     `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	DataEtlStatus                 DataEtlStatus                 `json:"DataEtlStatus" xml:"DataEtlStatus"`
	DestinationEndpoint           DestinationEndpoint           `json:"DestinationEndpoint" xml:"DestinationEndpoint"`
	MigrationMode                 MigrationMode                 `json:"MigrationMode" xml:"MigrationMode"`
	Performance                   Performance                   `json:"Performance" xml:"Performance"`
	PrecheckStatus                PrecheckStatus                `json:"PrecheckStatus" xml:"PrecheckStatus"`
	ReverseJob                    ReverseJob                    `json:"ReverseJob" xml:"ReverseJob"`
	SourceEndpoint                SourceEndpoint                `json:"SourceEndpoint" xml:"SourceEndpoint"`
	StructureInitializationStatus StructureInitializationStatus `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
	RetryState                    RetryState                    `json:"RetryState" xml:"RetryState"`
	TagList                       []DtsTag                      `json:"TagList" xml:"TagList"`
}
