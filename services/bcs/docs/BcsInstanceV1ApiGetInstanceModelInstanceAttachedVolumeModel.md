# BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 연결된 볼륨의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**MountPoint** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Size** | **int32** | 볼륨 크기 (GB 단위) | 
**IsDeleteOnTermination** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**IsRoot** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel

`func NewBcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel(id string, size int32, ) *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel`

NewBcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel instantiates a new BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModelWithDefaults

`func NewBcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModelWithDefaults() *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel`

NewBcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModelWithDefaults instantiates a new BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMountPoint

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSize

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetSize(v int32)`

SetSize sets Size field to given value.


### GetIsDeleteOnTermination

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.

### HasIsDeleteOnTermination

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasIsDeleteOnTermination() bool`

HasIsDeleteOnTermination returns a boolean if a field has been set.

### SetIsDeleteOnTerminationNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetIsDeleteOnTerminationNil(b bool)`

 SetIsDeleteOnTerminationNil sets the value for IsDeleteOnTermination to be an explicit nil

### UnsetIsDeleteOnTermination
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetIsDeleteOnTermination()`

UnsetIsDeleteOnTermination ensures that no value is present for IsDeleteOnTermination, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsRoot

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### SetIsRootNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) SetIsRootNil(b bool)`

 SetIsRootNil sets the value for IsRoot to be an explicit nil

### UnsetIsRoot
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel) UnsetIsRoot()`

UnsetIsRoot ensures that no value is present for IsRoot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


