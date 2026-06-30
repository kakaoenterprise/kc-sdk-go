# FlavorResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Flavor ID | 
**Name** | **string** | Flavor 이름 | 
**Type** | **string** |  | 
**Vcpus** | **int32** | vCPU 개수 | 
**Memory** | **int32** | 메모리 크기 (GB) | 
**MemoryMb** | **int32** | 메모리 크기 (MB) | 
**Group** | **string** | Flavor 그룹 | 
**Family** | **string** | Flavor 패밀리 | 
**AvailabilityZones** | **[]string** | 사용 가능한 가용 영역 목록 | 
**Deprecated** | **bool** | 사용 중단 여부 | 

## Methods

### NewFlavorResponseModel

`func NewFlavorResponseModel(id string, name string, type_ string, vcpus int32, memory int32, memoryMb int32, group string, family string, availabilityZones []string, deprecated bool, ) *FlavorResponseModel`

NewFlavorResponseModel instantiates a new FlavorResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorResponseModelWithDefaults

`func NewFlavorResponseModelWithDefaults() *FlavorResponseModel`

NewFlavorResponseModelWithDefaults instantiates a new FlavorResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlavorResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlavorResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlavorResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FlavorResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlavorResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlavorResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FlavorResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlavorResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlavorResponseModel) SetType(v string)`

SetType sets Type field to given value.


### GetVcpus

`func (o *FlavorResponseModel) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *FlavorResponseModel) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *FlavorResponseModel) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.


### GetMemory

`func (o *FlavorResponseModel) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *FlavorResponseModel) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *FlavorResponseModel) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetMemoryMb

`func (o *FlavorResponseModel) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *FlavorResponseModel) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *FlavorResponseModel) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.


### GetGroup

`func (o *FlavorResponseModel) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FlavorResponseModel) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FlavorResponseModel) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetFamily

`func (o *FlavorResponseModel) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *FlavorResponseModel) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *FlavorResponseModel) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetAvailabilityZones

`func (o *FlavorResponseModel) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *FlavorResponseModel) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *FlavorResponseModel) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetDeprecated

`func (o *FlavorResponseModel) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FlavorResponseModel) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FlavorResponseModel) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


