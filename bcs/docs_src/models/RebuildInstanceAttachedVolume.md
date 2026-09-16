# RebuildInstanceAttachedVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 볼륨의 고유 ID | 
**IsDeleteOnTermination** | **bool** | 인스턴스가 삭제될 때 해당 볼륨이 자동으로 삭제되는지 여부 | 

## Methods

### NewRebuildInstanceAttachedVolume

`func NewRebuildInstanceAttachedVolume(id string, isDeleteOnTermination bool, ) *RebuildInstanceAttachedVolume`

NewRebuildInstanceAttachedVolume instantiates a new RebuildInstanceAttachedVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuildInstanceAttachedVolumeWithDefaults

`func NewRebuildInstanceAttachedVolumeWithDefaults() *RebuildInstanceAttachedVolume`

NewRebuildInstanceAttachedVolumeWithDefaults instantiates a new RebuildInstanceAttachedVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RebuildInstanceAttachedVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebuildInstanceAttachedVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebuildInstanceAttachedVolume) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleteOnTermination

`func (o *RebuildInstanceAttachedVolume) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *RebuildInstanceAttachedVolume) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *RebuildInstanceAttachedVolume) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


