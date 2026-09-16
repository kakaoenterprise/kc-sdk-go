# UpdatedFlavor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vcpus** | **int32** | 가상 CPU 수 | 
**Ram** | **int32** | 메모리 크기 (MB 단위) | 
**Disk** | **int32** | 루트 디스크의 크기 (GB 단위) | 
**OriginalName** | **string** | Flavor의 원래 이름 (서버 생성 시 사용된 이름) | 

## Methods

### NewUpdatedFlavor

`func NewUpdatedFlavor(vcpus int32, ram int32, disk int32, originalName string, ) *UpdatedFlavor`

NewUpdatedFlavor instantiates a new UpdatedFlavor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedFlavorWithDefaults

`func NewUpdatedFlavorWithDefaults() *UpdatedFlavor`

NewUpdatedFlavorWithDefaults instantiates a new UpdatedFlavor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcpus

`func (o *UpdatedFlavor) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *UpdatedFlavor) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *UpdatedFlavor) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.


### GetRam

`func (o *UpdatedFlavor) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *UpdatedFlavor) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *UpdatedFlavor) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *UpdatedFlavor) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *UpdatedFlavor) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *UpdatedFlavor) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetOriginalName

`func (o *UpdatedFlavor) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *UpdatedFlavor) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *UpdatedFlavor) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


