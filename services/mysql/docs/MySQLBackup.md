# MySQLBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DiskSize** | Pointer to **NullableInt32** |  | [optional] 
**ExpireAt** | Pointer to **NullableString** |  | [optional] 
**ExpiryDuration** | Pointer to **NullableInt32** |  | [optional] 
**ExtraInfo** | Pointer to [**NullableExtraInfo**](ExtraInfo.md) |  | [optional] 
**InstanceGroupId** | Pointer to **NullableString** |  | [optional] 
**InstanceGroupName** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to [**NullableBackupStatus**](BackupStatus.md) |  | [optional] 
**Type** | Pointer to [**NullableBackupType**](BackupType.md) |  | [optional] 
**StartedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**EngineVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMySQLBackup

`func NewMySQLBackup() *MySQLBackup`

NewMySQLBackup instantiates a new MySQLBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySQLBackupWithDefaults

`func NewMySQLBackupWithDefaults() *MySQLBackup`

NewMySQLBackupWithDefaults instantiates a new MySQLBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MySQLBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MySQLBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MySQLBackup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MySQLBackup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MySQLBackup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MySQLBackup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *MySQLBackup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MySQLBackup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MySQLBackup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MySQLBackup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *MySQLBackup) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *MySQLBackup) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorName

`func (o *MySQLBackup) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *MySQLBackup) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *MySQLBackup) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *MySQLBackup) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *MySQLBackup) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *MySQLBackup) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetDescription

`func (o *MySQLBackup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MySQLBackup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MySQLBackup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MySQLBackup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MySQLBackup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MySQLBackup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiskSize

`func (o *MySQLBackup) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *MySQLBackup) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *MySQLBackup) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *MySQLBackup) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### SetDiskSizeNil

`func (o *MySQLBackup) SetDiskSizeNil(b bool)`

 SetDiskSizeNil sets the value for DiskSize to be an explicit nil

### UnsetDiskSize
`func (o *MySQLBackup) UnsetDiskSize()`

UnsetDiskSize ensures that no value is present for DiskSize, not even an explicit nil
### GetExpireAt

`func (o *MySQLBackup) GetExpireAt() string`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *MySQLBackup) GetExpireAtOk() (*string, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *MySQLBackup) SetExpireAt(v string)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *MySQLBackup) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *MySQLBackup) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *MySQLBackup) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetExpiryDuration

`func (o *MySQLBackup) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *MySQLBackup) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *MySQLBackup) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *MySQLBackup) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *MySQLBackup) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *MySQLBackup) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetExtraInfo

`func (o *MySQLBackup) GetExtraInfo() ExtraInfo`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *MySQLBackup) GetExtraInfoOk() (*ExtraInfo, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *MySQLBackup) SetExtraInfo(v ExtraInfo)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *MySQLBackup) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *MySQLBackup) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *MySQLBackup) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetInstanceGroupId

`func (o *MySQLBackup) GetInstanceGroupId() string`

GetInstanceGroupId returns the InstanceGroupId field if non-nil, zero value otherwise.

### GetInstanceGroupIdOk

`func (o *MySQLBackup) GetInstanceGroupIdOk() (*string, bool)`

GetInstanceGroupIdOk returns a tuple with the InstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupId

`func (o *MySQLBackup) SetInstanceGroupId(v string)`

SetInstanceGroupId sets InstanceGroupId field to given value.

### HasInstanceGroupId

`func (o *MySQLBackup) HasInstanceGroupId() bool`

HasInstanceGroupId returns a boolean if a field has been set.

### SetInstanceGroupIdNil

`func (o *MySQLBackup) SetInstanceGroupIdNil(b bool)`

 SetInstanceGroupIdNil sets the value for InstanceGroupId to be an explicit nil

### UnsetInstanceGroupId
`func (o *MySQLBackup) UnsetInstanceGroupId()`

UnsetInstanceGroupId ensures that no value is present for InstanceGroupId, not even an explicit nil
### GetInstanceGroupName

`func (o *MySQLBackup) GetInstanceGroupName() string`

GetInstanceGroupName returns the InstanceGroupName field if non-nil, zero value otherwise.

### GetInstanceGroupNameOk

`func (o *MySQLBackup) GetInstanceGroupNameOk() (*string, bool)`

GetInstanceGroupNameOk returns a tuple with the InstanceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupName

`func (o *MySQLBackup) SetInstanceGroupName(v string)`

SetInstanceGroupName sets InstanceGroupName field to given value.

### HasInstanceGroupName

`func (o *MySQLBackup) HasInstanceGroupName() bool`

HasInstanceGroupName returns a boolean if a field has been set.

### SetInstanceGroupNameNil

`func (o *MySQLBackup) SetInstanceGroupNameNil(b bool)`

 SetInstanceGroupNameNil sets the value for InstanceGroupName to be an explicit nil

### UnsetInstanceGroupName
`func (o *MySQLBackup) UnsetInstanceGroupName()`

UnsetInstanceGroupName ensures that no value is present for InstanceGroupName, not even an explicit nil
### GetId

`func (o *MySQLBackup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MySQLBackup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MySQLBackup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MySQLBackup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MySQLBackup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MySQLBackup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProjectId

`func (o *MySQLBackup) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MySQLBackup) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MySQLBackup) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MySQLBackup) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *MySQLBackup) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *MySQLBackup) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSize

`func (o *MySQLBackup) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MySQLBackup) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MySQLBackup) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MySQLBackup) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MySQLBackup) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MySQLBackup) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *MySQLBackup) GetStatus() BackupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MySQLBackup) GetStatusOk() (*BackupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MySQLBackup) SetStatus(v BackupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MySQLBackup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MySQLBackup) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MySQLBackup) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *MySQLBackup) GetType() BackupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MySQLBackup) GetTypeOk() (*BackupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MySQLBackup) SetType(v BackupType)`

SetType sets Type field to given value.

### HasType

`func (o *MySQLBackup) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MySQLBackup) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MySQLBackup) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartedAt

`func (o *MySQLBackup) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MySQLBackup) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MySQLBackup) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MySQLBackup) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *MySQLBackup) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *MySQLBackup) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetUpdatedAt

`func (o *MySQLBackup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MySQLBackup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MySQLBackup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MySQLBackup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *MySQLBackup) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *MySQLBackup) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetEngineVersion

`func (o *MySQLBackup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MySQLBackup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MySQLBackup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *MySQLBackup) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *MySQLBackup) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *MySQLBackup) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


