# MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultParameterGroupId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EngineVersion** | Pointer to **NullableString** |  | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** |  | [optional] 
**InstanceGroupCount** | Pointer to **NullableInt32** |  | [optional] 
**IsRollbackPossible** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel

`func NewMysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel() *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel`

NewMysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel instantiates a new MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModelWithDefaults

`func NewMysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModelWithDefaults() *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel`

NewMysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModelWithDefaults instantiates a new MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultParameterGroupId

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetDefaultParameterGroupId() string`

GetDefaultParameterGroupId returns the DefaultParameterGroupId field if non-nil, zero value otherwise.

### GetDefaultParameterGroupIdOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetDefaultParameterGroupIdOk() (*string, bool)`

GetDefaultParameterGroupIdOk returns a tuple with the DefaultParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterGroupId

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetDefaultParameterGroupId(v string)`

SetDefaultParameterGroupId sets DefaultParameterGroupId field to given value.

### HasDefaultParameterGroupId

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasDefaultParameterGroupId() bool`

HasDefaultParameterGroupId returns a boolean if a field has been set.

### SetDefaultParameterGroupIdNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetDefaultParameterGroupIdNil(b bool)`

 SetDefaultParameterGroupIdNil sets the value for DefaultParameterGroupId to be an explicit nil

### UnsetDefaultParameterGroupId
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetDefaultParameterGroupId()`

UnsetDefaultParameterGroupId ensures that no value is present for DefaultParameterGroupId, not even an explicit nil
### GetDescription

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngineVersion

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil
### GetExistErrorSync

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil
### GetInstanceGroupCount

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.

### HasInstanceGroupCount

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasInstanceGroupCount() bool`

HasInstanceGroupCount returns a boolean if a field has been set.

### SetInstanceGroupCountNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetInstanceGroupCountNil(b bool)`

 SetInstanceGroupCountNil sets the value for InstanceGroupCount to be an explicit nil

### UnsetInstanceGroupCount
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetInstanceGroupCount()`

UnsetInstanceGroupCount ensures that no value is present for InstanceGroupCount, not even an explicit nil
### GetIsRollbackPossible

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetIsRollbackPossible() bool`

GetIsRollbackPossible returns the IsRollbackPossible field if non-nil, zero value otherwise.

### GetIsRollbackPossibleOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetIsRollbackPossibleOk() (*bool, bool)`

GetIsRollbackPossibleOk returns a tuple with the IsRollbackPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRollbackPossible

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetIsRollbackPossible(v bool)`

SetIsRollbackPossible sets IsRollbackPossible field to given value.

### HasIsRollbackPossible

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasIsRollbackPossible() bool`

HasIsRollbackPossible returns a boolean if a field has been set.

### SetIsRollbackPossibleNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetIsRollbackPossibleNil(b bool)`

 SetIsRollbackPossibleNil sets the value for IsRollbackPossible to be an explicit nil

### UnsetIsRollbackPossible
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetIsRollbackPossible()`

UnsetIsRollbackPossible ensures that no value is present for IsRollbackPossible, not even an explicit nil
### GetName

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


