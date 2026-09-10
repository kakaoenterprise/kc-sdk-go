# AllocatableResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **string** | 사용 가능한 CPU (코어 수) | 
**EphemeralStorage** | **string** | 사용 가능한 임시 스토리지 용량 | 
**Hugepages1Gi** | **string** | 1Gi HugePages 메모리 할당량 | 
**Hugepages2Mi** | **string** | 2Mi HugePages 메모리 할당량 | 
**Memory** | **string** | 사용 가능한 메모리 | 
**Pods** | **string** | 스케줄 가능한 파드 수 제한 | 

## Methods

### NewAllocatableResources

`func NewAllocatableResources(cpu string, ephemeralStorage string, hugepages1Gi string, hugepages2Mi string, memory string, pods string, ) *AllocatableResources`

NewAllocatableResources instantiates a new AllocatableResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocatableResourcesWithDefaults

`func NewAllocatableResourcesWithDefaults() *AllocatableResources`

NewAllocatableResourcesWithDefaults instantiates a new AllocatableResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *AllocatableResources) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AllocatableResources) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AllocatableResources) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### GetEphemeralStorage

`func (o *AllocatableResources) GetEphemeralStorage() string`

GetEphemeralStorage returns the EphemeralStorage field if non-nil, zero value otherwise.

### GetEphemeralStorageOk

`func (o *AllocatableResources) GetEphemeralStorageOk() (*string, bool)`

GetEphemeralStorageOk returns a tuple with the EphemeralStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorage

`func (o *AllocatableResources) SetEphemeralStorage(v string)`

SetEphemeralStorage sets EphemeralStorage field to given value.


### GetHugepages1Gi

`func (o *AllocatableResources) GetHugepages1Gi() string`

GetHugepages1Gi returns the Hugepages1Gi field if non-nil, zero value otherwise.

### GetHugepages1GiOk

`func (o *AllocatableResources) GetHugepages1GiOk() (*string, bool)`

GetHugepages1GiOk returns a tuple with the Hugepages1Gi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages1Gi

`func (o *AllocatableResources) SetHugepages1Gi(v string)`

SetHugepages1Gi sets Hugepages1Gi field to given value.


### GetHugepages2Mi

`func (o *AllocatableResources) GetHugepages2Mi() string`

GetHugepages2Mi returns the Hugepages2Mi field if non-nil, zero value otherwise.

### GetHugepages2MiOk

`func (o *AllocatableResources) GetHugepages2MiOk() (*string, bool)`

GetHugepages2MiOk returns a tuple with the Hugepages2Mi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages2Mi

`func (o *AllocatableResources) SetHugepages2Mi(v string)`

SetHugepages2Mi sets Hugepages2Mi field to given value.


### GetMemory

`func (o *AllocatableResources) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AllocatableResources) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AllocatableResources) SetMemory(v string)`

SetMemory sets Memory field to given value.


### GetPods

`func (o *AllocatableResources) GetPods() string`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *AllocatableResources) GetPodsOk() (*string, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *AllocatableResources) SetPods(v string)`

SetPods sets Pods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


