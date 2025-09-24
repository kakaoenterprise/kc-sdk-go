# BcsVolumeV1ApiGetVolumeModelVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 볼륨의 고유 ID | 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsBootable** | Pointer to **NullableBool** |  | [optional] 
**MountPoint** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**IsRoot** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**AttachStatus** | Pointer to **NullableString** |  | [optional] 
**LaunchedAt** | Pointer to **NullableTime** |  | [optional] 
**EncryptionKeyId** | Pointer to **NullableString** |  | [optional] 
**PreviousStatus** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**InstanceId** | Pointer to **NullableString** |  | [optional] 
**InstanceName** | Pointer to **NullableString** |  | [optional] 
**ImageMetadata** | Pointer to [**NullableBcsVolumeV1ApiGetVolumeModelImageMetaData**](BcsVolumeV1ApiGetVolumeModelImageMetaData.md) |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBcsVolumeV1ApiGetVolumeModelVolumeModel

`func NewBcsVolumeV1ApiGetVolumeModelVolumeModel(id string, ) *BcsVolumeV1ApiGetVolumeModelVolumeModel`

NewBcsVolumeV1ApiGetVolumeModelVolumeModel instantiates a new BcsVolumeV1ApiGetVolumeModelVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsVolumeV1ApiGetVolumeModelVolumeModelWithDefaults

`func NewBcsVolumeV1ApiGetVolumeModelVolumeModelWithDefaults() *BcsVolumeV1ApiGetVolumeModelVolumeModel`

NewBcsVolumeV1ApiGetVolumeModelVolumeModelWithDefaults instantiates a new BcsVolumeV1ApiGetVolumeModelVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetId(v string)`

SetId sets Id field to given value.


### GetSize

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetName

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsBootable

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.

### SetIsBootableNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetIsBootableNil(b bool)`

 SetIsBootableNil sets the value for IsBootable to be an explicit nil

### UnsetIsBootable
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetIsBootable()`

UnsetIsBootable ensures that no value is present for IsBootable, not even an explicit nil
### GetMountPoint

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetMetadata

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsEncrypted

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetIsRoot

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### SetIsRootNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetIsRootNil(b bool)`

 SetIsRootNil sets the value for IsRoot to be an explicit nil

### UnsetIsRoot
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetIsRoot()`

UnsetIsRoot ensures that no value is present for IsRoot, not even an explicit nil
### GetType

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetProjectId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAvailabilityZone

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAttachStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetAttachStatus() string`

GetAttachStatus returns the AttachStatus field if non-nil, zero value otherwise.

### GetAttachStatusOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetAttachStatusOk() (*string, bool)`

GetAttachStatusOk returns a tuple with the AttachStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetAttachStatus(v string)`

SetAttachStatus sets AttachStatus field to given value.

### HasAttachStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasAttachStatus() bool`

HasAttachStatus returns a boolean if a field has been set.

### SetAttachStatusNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetAttachStatusNil(b bool)`

 SetAttachStatusNil sets the value for AttachStatus to be an explicit nil

### UnsetAttachStatus
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetAttachStatus()`

UnsetAttachStatus ensures that no value is present for AttachStatus, not even an explicit nil
### GetLaunchedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetLaunchedAt() time.Time`

GetLaunchedAt returns the LaunchedAt field if non-nil, zero value otherwise.

### GetLaunchedAtOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetLaunchedAtOk() (*time.Time, bool)`

GetLaunchedAtOk returns a tuple with the LaunchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetLaunchedAt(v time.Time)`

SetLaunchedAt sets LaunchedAt field to given value.

### HasLaunchedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasLaunchedAt() bool`

HasLaunchedAt returns a boolean if a field has been set.

### SetLaunchedAtNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetLaunchedAtNil(b bool)`

 SetLaunchedAtNil sets the value for LaunchedAt to be an explicit nil

### UnsetLaunchedAt
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetLaunchedAt()`

UnsetLaunchedAt ensures that no value is present for LaunchedAt, not even an explicit nil
### GetEncryptionKeyId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.

### SetEncryptionKeyIdNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetEncryptionKeyIdNil(b bool)`

 SetEncryptionKeyIdNil sets the value for EncryptionKeyId to be an explicit nil

### UnsetEncryptionKeyId
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetEncryptionKeyId()`

UnsetEncryptionKeyId ensures that no value is present for EncryptionKeyId, not even an explicit nil
### GetPreviousStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.

### HasPreviousStatus

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasPreviousStatus() bool`

HasPreviousStatus returns a boolean if a field has been set.

### SetPreviousStatusNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetPreviousStatusNil(b bool)`

 SetPreviousStatusNil sets the value for PreviousStatus to be an explicit nil

### UnsetPreviousStatus
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetPreviousStatus()`

UnsetPreviousStatus ensures that no value is present for PreviousStatus, not even an explicit nil
### GetCreatedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetInstanceId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceName

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetImageMetadata

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetImageMetadata() BcsVolumeV1ApiGetVolumeModelImageMetaData`

GetImageMetadata returns the ImageMetadata field if non-nil, zero value otherwise.

### GetImageMetadataOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetImageMetadataOk() (*BcsVolumeV1ApiGetVolumeModelImageMetaData, bool)`

GetImageMetadataOk returns a tuple with the ImageMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMetadata

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetImageMetadata(v BcsVolumeV1ApiGetVolumeModelImageMetaData)`

SetImageMetadata sets ImageMetadata field to given value.

### HasImageMetadata

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasImageMetadata() bool`

HasImageMetadata returns a boolean if a field has been set.

### SetImageMetadataNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetImageMetadataNil(b bool)`

 SetImageMetadataNil sets the value for ImageMetadata to be an explicit nil

### UnsetImageMetadata
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetImageMetadata()`

UnsetImageMetadata ensures that no value is present for ImageMetadata, not even an explicit nil
### GetVolumeType

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *BcsVolumeV1ApiGetVolumeModelVolumeModel) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


