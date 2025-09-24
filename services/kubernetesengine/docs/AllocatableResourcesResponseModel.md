# AllocatableResourcesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **string** | 사용 가능한 CPU (코어 수) | 
**EphemeralStorage** | **string** | 사용 가능한 임시 스토리지 용량 | 
**Hugepages1Gi** | **string** | HugePages 메모리 할당량 | 
**Hugepages2Mi** | **string** | HugePages 메모리 할당량 | 
**Memory** | **string** | 사용 가능한 메모리 | 
**Pods** | **string** | 스케줄 가능한 파드 수 제한 | 

## Methods

### NewAllocatableResourcesResponseModel

`func NewAllocatableResourcesResponseModel(cpu string, ephemeralStorage string, hugepages1Gi string, hugepages2Mi string, memory string, pods string, ) *AllocatableResourcesResponseModel`

NewAllocatableResourcesResponseModel instantiates a new AllocatableResourcesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocatableResourcesResponseModelWithDefaults

`func NewAllocatableResourcesResponseModelWithDefaults() *AllocatableResourcesResponseModel`

NewAllocatableResourcesResponseModelWithDefaults instantiates a new AllocatableResourcesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *AllocatableResourcesResponseModel) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AllocatableResourcesResponseModel) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AllocatableResourcesResponseModel) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### GetEphemeralStorage

`func (o *AllocatableResourcesResponseModel) GetEphemeralStorage() string`

GetEphemeralStorage returns the EphemeralStorage field if non-nil, zero value otherwise.

### GetEphemeralStorageOk

`func (o *AllocatableResourcesResponseModel) GetEphemeralStorageOk() (*string, bool)`

GetEphemeralStorageOk returns a tuple with the EphemeralStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorage

`func (o *AllocatableResourcesResponseModel) SetEphemeralStorage(v string)`

SetEphemeralStorage sets EphemeralStorage field to given value.


### GetHugepages1Gi

`func (o *AllocatableResourcesResponseModel) GetHugepages1Gi() string`

GetHugepages1Gi returns the Hugepages1Gi field if non-nil, zero value otherwise.

### GetHugepages1GiOk

`func (o *AllocatableResourcesResponseModel) GetHugepages1GiOk() (*string, bool)`

GetHugepages1GiOk returns a tuple with the Hugepages1Gi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages1Gi

`func (o *AllocatableResourcesResponseModel) SetHugepages1Gi(v string)`

SetHugepages1Gi sets Hugepages1Gi field to given value.


### GetHugepages2Mi

`func (o *AllocatableResourcesResponseModel) GetHugepages2Mi() string`

GetHugepages2Mi returns the Hugepages2Mi field if non-nil, zero value otherwise.

### GetHugepages2MiOk

`func (o *AllocatableResourcesResponseModel) GetHugepages2MiOk() (*string, bool)`

GetHugepages2MiOk returns a tuple with the Hugepages2Mi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages2Mi

`func (o *AllocatableResourcesResponseModel) SetHugepages2Mi(v string)`

SetHugepages2Mi sets Hugepages2Mi field to given value.


### GetMemory

`func (o *AllocatableResourcesResponseModel) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AllocatableResourcesResponseModel) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AllocatableResourcesResponseModel) SetMemory(v string)`

SetMemory sets Memory field to given value.


### GetPods

`func (o *AllocatableResourcesResponseModel) GetPods() string`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *AllocatableResourcesResponseModel) GetPodsOk() (*string, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *AllocatableResourcesResponseModel) SetPods(v string)`

SetPods sets Pods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


