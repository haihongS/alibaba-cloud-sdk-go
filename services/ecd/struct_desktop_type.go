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

// DesktopType is a nested struct in ecd response
type DesktopType struct {
	SystemDiskSize     string              `json:"SystemDiskSize" xml:"SystemDiskSize"`
	DesktopTypeId      string              `json:"DesktopTypeId" xml:"DesktopTypeId"`
	DataDiskSize       string              `json:"DataDiskSize" xml:"DataDiskSize"`
	CpuCount           string              `json:"CpuCount" xml:"CpuCount"`
	GpuCount           float64             `json:"GpuCount" xml:"GpuCount"`
	GpuSpec            string              `json:"GpuSpec" xml:"GpuSpec"`
	InstanceTypeFamily string              `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
	MemorySize         string              `json:"MemorySize" xml:"MemorySize"`
	DesktopTypeStatus  string              `json:"DesktopTypeStatus" xml:"DesktopTypeStatus"`
	AllowDiskSize      []AllowDiskSizeItem `json:"AllowDiskSize" xml:"AllowDiskSize"`
}
