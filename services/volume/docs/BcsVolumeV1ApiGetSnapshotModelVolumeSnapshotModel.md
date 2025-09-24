# BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 스냅샷의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**VolumeId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**IsIncremental** | Pointer to **NullableBool** |  | [optional] 
**IsDependentSnapshot** | Pointer to **NullableBool** |  | [optional] 
**RealSize** | Pointer to **NullableInt64** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel

`func NewBcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel(id string, ) *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel`

NewBcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel instantiates a new BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModelWithDefaults

`func NewBcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModelWithDefaults() *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel`

NewBcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModelWithDefaults instantiates a new BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSize

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVolumeId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetUserId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetProjectId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetParentId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetCreatedAt

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetIsIncremental

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.

### HasIsIncremental

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasIsIncremental() bool`

HasIsIncremental returns a boolean if a field has been set.

### SetIsIncrementalNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetIsIncrementalNil(b bool)`

 SetIsIncrementalNil sets the value for IsIncremental to be an explicit nil

### UnsetIsIncremental
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetIsIncremental()`

UnsetIsIncremental ensures that no value is present for IsIncremental, not even an explicit nil
### GetIsDependentSnapshot

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetIsDependentSnapshot() bool`

GetIsDependentSnapshot returns the IsDependentSnapshot field if non-nil, zero value otherwise.

### GetIsDependentSnapshotOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetIsDependentSnapshotOk() (*bool, bool)`

GetIsDependentSnapshotOk returns a tuple with the IsDependentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDependentSnapshot

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetIsDependentSnapshot(v bool)`

SetIsDependentSnapshot sets IsDependentSnapshot field to given value.

### HasIsDependentSnapshot

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasIsDependentSnapshot() bool`

HasIsDependentSnapshot returns a boolean if a field has been set.

### SetIsDependentSnapshotNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetIsDependentSnapshotNil(b bool)`

 SetIsDependentSnapshotNil sets the value for IsDependentSnapshot to be an explicit nil

### UnsetIsDependentSnapshot
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetIsDependentSnapshot()`

UnsetIsDependentSnapshot ensures that no value is present for IsDependentSnapshot, not even an explicit nil
### GetRealSize

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetRealSize() int64`

GetRealSize returns the RealSize field if non-nil, zero value otherwise.

### GetRealSizeOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetRealSizeOk() (*int64, bool)`

GetRealSizeOk returns a tuple with the RealSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealSize

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetRealSize(v int64)`

SetRealSize sets RealSize field to given value.

### HasRealSize

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasRealSize() bool`

HasRealSize returns a boolean if a field has been set.

### SetRealSizeNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetRealSizeNil(b bool)`

 SetRealSizeNil sets the value for RealSize to be an explicit nil

### UnsetRealSize
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetRealSize()`

UnsetRealSize ensures that no value is present for RealSize, not even an explicit nil
### GetScheduleId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BcsVolumeV1ApiGetSnapshotModelVolumeSnapshotModel) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


