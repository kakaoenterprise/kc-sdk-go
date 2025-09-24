# RestoreVolumeSnapshotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **string** | 스냅샷의 고유 ID | 
**VolumeId** | **string** | 볼륨의 고유 ID | 
**VolumeName** | **string** | 복원 후 새로 생성된 볼륨의 이름 | 

## Methods

### NewRestoreVolumeSnapshotModel

`func NewRestoreVolumeSnapshotModel(snapshotId string, volumeId string, volumeName string, ) *RestoreVolumeSnapshotModel`

NewRestoreVolumeSnapshotModel instantiates a new RestoreVolumeSnapshotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreVolumeSnapshotModelWithDefaults

`func NewRestoreVolumeSnapshotModelWithDefaults() *RestoreVolumeSnapshotModel`

NewRestoreVolumeSnapshotModelWithDefaults instantiates a new RestoreVolumeSnapshotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *RestoreVolumeSnapshotModel) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RestoreVolumeSnapshotModel) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RestoreVolumeSnapshotModel) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetVolumeId

`func (o *RestoreVolumeSnapshotModel) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *RestoreVolumeSnapshotModel) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *RestoreVolumeSnapshotModel) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### GetVolumeName

`func (o *RestoreVolumeSnapshotModel) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *RestoreVolumeSnapshotModel) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *RestoreVolumeSnapshotModel) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


