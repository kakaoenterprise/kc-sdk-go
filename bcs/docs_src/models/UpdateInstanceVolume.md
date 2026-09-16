# UpdateInstanceVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeleteOnTermination** | Pointer to **NullableBool** | 인스턴스 삭제 시 해당 볼륨을 자동으로 삭제할지 여부 | [optional] 
**Name** | Pointer to **NullableString** | 인스턴스에 연결된 볼륨의 새 이름 | [optional] 
**Description** | Pointer to **NullableString** | 볼륨에 대한 설명 - 필드를 사용하지 않거나 &#x60;null&#x60;이면 기존값 유지 | [optional] 

## Methods

### NewUpdateInstanceVolume

`func NewUpdateInstanceVolume() *UpdateInstanceVolume`

NewUpdateInstanceVolume instantiates a new UpdateInstanceVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceVolumeWithDefaults

`func NewUpdateInstanceVolumeWithDefaults() *UpdateInstanceVolume`

NewUpdateInstanceVolumeWithDefaults instantiates a new UpdateInstanceVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeleteOnTermination

`func (o *UpdateInstanceVolume) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *UpdateInstanceVolume) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *UpdateInstanceVolume) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.

### HasIsDeleteOnTermination

`func (o *UpdateInstanceVolume) HasIsDeleteOnTermination() bool`

HasIsDeleteOnTermination returns a boolean if a field has been set.

### SetIsDeleteOnTerminationNil

`func (o *UpdateInstanceVolume) SetIsDeleteOnTerminationNil(b bool)`

 SetIsDeleteOnTerminationNil sets the value for IsDeleteOnTermination to be an explicit nil

### UnsetIsDeleteOnTermination
`func (o *UpdateInstanceVolume) UnsetIsDeleteOnTermination()`

UnsetIsDeleteOnTermination ensures that no value is present for IsDeleteOnTermination, not even an explicit nil
### GetName

`func (o *UpdateInstanceVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstanceVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstanceVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInstanceVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateInstanceVolume) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateInstanceVolume) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateInstanceVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInstanceVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInstanceVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInstanceVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateInstanceVolume) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateInstanceVolume) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


