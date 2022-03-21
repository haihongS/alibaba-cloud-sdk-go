package ecd

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

// Desktops is a nested struct in ecd response
type Desktops struct {
	GpuSpec            string      `json:"GpuSpec" xml:"GpuSpec"`
	Comments           string      `json:"Comments" xml:"Comments"`
	DesktopGroupName   string      `json:"DesktopGroupName" xml:"DesktopGroupName"`
	SystemDiskSize     int         `json:"SystemDiskSize" xml:"SystemDiskSize"`
	OfficeSiteName     string      `json:"OfficeSiteName" xml:"OfficeSiteName"`
	MaxDesktopsCount   int         `json:"MaxDesktopsCount" xml:"MaxDesktopsCount"`
	AllowAutoSetup     int         `json:"AllowAutoSetup" xml:"AllowAutoSetup"`
	Creator            string      `json:"Creator" xml:"Creator"`
	MinDesktopsCount   int         `json:"MinDesktopsCount" xml:"MinDesktopsCount"`
	BindAmount         int         `json:"BindAmount" xml:"BindAmount"`
	LoadPolicy         int         `json:"LoadPolicy" xml:"LoadPolicy"`
	DirectoryId        string      `json:"DirectoryId" xml:"DirectoryId"`
	Status             int         `json:"Status" xml:"Status"`
	ExpiredTime        string      `json:"ExpiredTime" xml:"ExpiredTime"`
	PayType            string      `json:"PayType" xml:"PayType"`
	OwnBundleName      string      `json:"OwnBundleName" xml:"OwnBundleName"`
	KeepDuration       int64       `json:"KeepDuration" xml:"KeepDuration"`
	CreationTime       string      `json:"CreationTime" xml:"CreationTime"`
	Memory             int64       `json:"Memory" xml:"Memory"`
	DesktopGroupId     string      `json:"DesktopGroupId" xml:"DesktopGroupId"`
	OfficeSiteId       string      `json:"OfficeSiteId" xml:"OfficeSiteId"`
	DataDiskSize       string      `json:"DataDiskSize" xml:"DataDiskSize"`
	ResType            int         `json:"ResType" xml:"ResType"`
	PolicyGroupName    string      `json:"PolicyGroupName" xml:"PolicyGroupName"`
	DirectoryType      string      `json:"DirectoryType" xml:"DirectoryType"`
	GpuCount           float64     `json:"GpuCount" xml:"GpuCount"`
	OwnBundleId        string      `json:"OwnBundleId" xml:"OwnBundleId"`
	AllowBufferCount   int         `json:"AllowBufferCount" xml:"AllowBufferCount"`
	OwnType            int         `json:"OwnType" xml:"OwnType"`
	OfficeSiteType     string      `json:"OfficeSiteType" xml:"OfficeSiteType"`
	ResetType          int         `json:"ResetType" xml:"ResetType"`
	PolicyGroupId      string      `json:"PolicyGroupId" xml:"PolicyGroupId"`
	SystemDiskCategory string      `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
	Cpu                int         `json:"Cpu" xml:"Cpu"`
	DataDiskCategory   string      `json:"DataDiskCategory" xml:"DataDiskCategory"`
	TimerInfos         []TimerInfo `json:"TimerInfos" xml:"TimerInfos"`
}
