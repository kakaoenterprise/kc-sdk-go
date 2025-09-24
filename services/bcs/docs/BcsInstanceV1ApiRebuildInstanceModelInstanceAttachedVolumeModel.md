# BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeId** | **string** | 볼륨의 고유 ID | 
**IsDeleteOnTermination** | **bool** | 인스턴스가 삭제될 때 해당 볼륨이 자동으로 삭제되는지 여부 | 

## Methods

### NewBcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel

`func NewBcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel(volumeId string, isDeleteOnTermination bool, ) *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel`

NewBcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel instantiates a new BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModelWithDefaults

`func NewBcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModelWithDefaults() *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel`

NewBcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModelWithDefaults instantiates a new BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeId

`func (o *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### GetIsDeleteOnTermination

`func (o *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


