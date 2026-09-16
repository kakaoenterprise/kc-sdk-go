# RestoreSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 볼륨에 대한 설명 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 신규 볼륨을 생성할 가용 영역 | 
**VolumeTypeId** | Pointer to **NullableString** | 생성할 볼륨 유형 ID - [List volume types](/openapi/bcs/list-volume-types)에서 확인 | [optional] 

## Methods

### NewRestoreSnapshot

`func NewRestoreSnapshot(name string, availabilityZone AvailabilityZone, ) *RestoreSnapshot`

NewRestoreSnapshot instantiates a new RestoreSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSnapshotWithDefaults

`func NewRestoreSnapshotWithDefaults() *RestoreSnapshot`

NewRestoreSnapshotWithDefaults instantiates a new RestoreSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestoreSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreSnapshot) SetName(v string)`

SetName sets Name field to given value.


### GetAvailabilityZone

`func (o *RestoreSnapshot) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RestoreSnapshot) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RestoreSnapshot) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetVolumeTypeId

`func (o *RestoreSnapshot) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *RestoreSnapshot) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *RestoreSnapshot) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.

### HasVolumeTypeId

`func (o *RestoreSnapshot) HasVolumeTypeId() bool`

HasVolumeTypeId returns a boolean if a field has been set.

### SetVolumeTypeIdNil

`func (o *RestoreSnapshot) SetVolumeTypeIdNil(b bool)`

 SetVolumeTypeIdNil sets the value for VolumeTypeId to be an explicit nil

### UnsetVolumeTypeId
`func (o *RestoreSnapshot) UnsetVolumeTypeId()`

UnsetVolumeTypeId ensures that no value is present for VolumeTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


