# UpdatedAttachedVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 볼륨의 ID | 
**IsDeleteOnTermination** | **bool** | 인스턴스가 삭제될 때 해당 볼륨도 함께 삭제되는지 여부 | 

## Methods

### NewUpdatedAttachedVolume

`func NewUpdatedAttachedVolume(id string, isDeleteOnTermination bool, ) *UpdatedAttachedVolume`

NewUpdatedAttachedVolume instantiates a new UpdatedAttachedVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedAttachedVolumeWithDefaults

`func NewUpdatedAttachedVolumeWithDefaults() *UpdatedAttachedVolume`

NewUpdatedAttachedVolumeWithDefaults instantiates a new UpdatedAttachedVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatedAttachedVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedAttachedVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedAttachedVolume) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleteOnTermination

`func (o *UpdatedAttachedVolume) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *UpdatedAttachedVolume) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *UpdatedAttachedVolume) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


