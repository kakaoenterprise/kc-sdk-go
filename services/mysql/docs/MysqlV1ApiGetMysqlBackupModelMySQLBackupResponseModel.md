# MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel

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
**ExtraInfo** | Pointer to [**NullableMysqlV1ApiGetMysqlBackupModelExtraInfoResponseModel**](MysqlV1ApiGetMysqlBackupModelExtraInfoResponseModel.md) |  | [optional] 
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

### NewMysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel

`func NewMysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel() *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel`

NewMysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel instantiates a new MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModelWithDefaults

`func NewMysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModelWithDefaults() *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel`

NewMysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModelWithDefaults instantiates a new MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetDescription

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiskSize

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### SetDiskSizeNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetDiskSizeNil(b bool)`

 SetDiskSizeNil sets the value for DiskSize to be an explicit nil

### UnsetDiskSize
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetDiskSize()`

UnsetDiskSize ensures that no value is present for DiskSize, not even an explicit nil
### GetExpireAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetExpireAt() string`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetExpireAtOk() (*string, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetExpireAt(v string)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetExpiryDuration

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetExtraInfo

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetExtraInfo() MysqlV1ApiGetMysqlBackupModelExtraInfoResponseModel`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetExtraInfoOk() (*MysqlV1ApiGetMysqlBackupModelExtraInfoResponseModel, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetExtraInfo(v MysqlV1ApiGetMysqlBackupModelExtraInfoResponseModel)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetInstanceGroupId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetInstanceGroupId() string`

GetInstanceGroupId returns the InstanceGroupId field if non-nil, zero value otherwise.

### GetInstanceGroupIdOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetInstanceGroupIdOk() (*string, bool)`

GetInstanceGroupIdOk returns a tuple with the InstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetInstanceGroupId(v string)`

SetInstanceGroupId sets InstanceGroupId field to given value.

### HasInstanceGroupId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasInstanceGroupId() bool`

HasInstanceGroupId returns a boolean if a field has been set.

### SetInstanceGroupIdNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetInstanceGroupIdNil(b bool)`

 SetInstanceGroupIdNil sets the value for InstanceGroupId to be an explicit nil

### UnsetInstanceGroupId
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetInstanceGroupId()`

UnsetInstanceGroupId ensures that no value is present for InstanceGroupId, not even an explicit nil
### GetInstanceGroupName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetInstanceGroupName() string`

GetInstanceGroupName returns the InstanceGroupName field if non-nil, zero value otherwise.

### GetInstanceGroupNameOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetInstanceGroupNameOk() (*string, bool)`

GetInstanceGroupNameOk returns a tuple with the InstanceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetInstanceGroupName(v string)`

SetInstanceGroupName sets InstanceGroupName field to given value.

### HasInstanceGroupName

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasInstanceGroupName() bool`

HasInstanceGroupName returns a boolean if a field has been set.

### SetInstanceGroupNameNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetInstanceGroupNameNil(b bool)`

 SetInstanceGroupNameNil sets the value for InstanceGroupName to be an explicit nil

### UnsetInstanceGroupName
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetInstanceGroupName()`

UnsetInstanceGroupName ensures that no value is present for InstanceGroupName, not even an explicit nil
### GetId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProjectId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSize

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetStatus() BackupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetStatusOk() (*BackupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetStatus(v BackupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetType() BackupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetTypeOk() (*BackupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetType(v BackupType)`

SetType sets Type field to given value.

### HasType

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetUpdatedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetEngineVersion

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


