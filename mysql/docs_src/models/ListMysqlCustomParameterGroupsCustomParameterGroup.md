# ListMysqlCustomParameterGroupsCustomParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultParameterGroupId** | Pointer to **NullableString** | 기본 MySQL 파라미터 그룹 ID | [optional] 
**Description** | Pointer to **NullableString** | MySQL 커스텀 파라미터 그룹 설명 | [optional] 
**EngineVersion** | Pointer to **NullableString** | 파라미터 그룹이 적용되는 MySQL 엔진 버전 | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** | 파라미터 적용 또는 동기화 과정에서 오류가 발생한 인스턴스 그룹 존재 여부 | [optional] 
**InstanceGroupCount** | Pointer to **NullableInt32** | 해당 커스텀 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 수 | [optional] 
**IsRollbackPossible** | Pointer to **NullableBool** | 이전 파라미터 그룹으로 롤백 가능 여부 | [optional] 
**Name** | Pointer to **NullableString** | MySQL 커스텀 파라미터 그룹 이름 | [optional] 
**Id** | Pointer to **NullableString** | MySQL 커스텀 파라미터 그룹 ID | [optional] 

## Methods

### NewListMysqlCustomParameterGroupsCustomParameterGroup

`func NewListMysqlCustomParameterGroupsCustomParameterGroup() *ListMysqlCustomParameterGroupsCustomParameterGroup`

NewListMysqlCustomParameterGroupsCustomParameterGroup instantiates a new ListMysqlCustomParameterGroupsCustomParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlCustomParameterGroupsCustomParameterGroupWithDefaults

`func NewListMysqlCustomParameterGroupsCustomParameterGroupWithDefaults() *ListMysqlCustomParameterGroupsCustomParameterGroup`

NewListMysqlCustomParameterGroupsCustomParameterGroupWithDefaults instantiates a new ListMysqlCustomParameterGroupsCustomParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultParameterGroupId

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetDefaultParameterGroupId() string`

GetDefaultParameterGroupId returns the DefaultParameterGroupId field if non-nil, zero value otherwise.

### GetDefaultParameterGroupIdOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetDefaultParameterGroupIdOk() (*string, bool)`

GetDefaultParameterGroupIdOk returns a tuple with the DefaultParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterGroupId

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetDefaultParameterGroupId(v string)`

SetDefaultParameterGroupId sets DefaultParameterGroupId field to given value.

### HasDefaultParameterGroupId

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasDefaultParameterGroupId() bool`

HasDefaultParameterGroupId returns a boolean if a field has been set.

### SetDefaultParameterGroupIdNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetDefaultParameterGroupIdNil(b bool)`

 SetDefaultParameterGroupIdNil sets the value for DefaultParameterGroupId to be an explicit nil

### UnsetDefaultParameterGroupId
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetDefaultParameterGroupId()`

UnsetDefaultParameterGroupId ensures that no value is present for DefaultParameterGroupId, not even an explicit nil
### GetDescription

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngineVersion

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil
### GetExistErrorSync

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil
### GetInstanceGroupCount

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.

### HasInstanceGroupCount

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasInstanceGroupCount() bool`

HasInstanceGroupCount returns a boolean if a field has been set.

### SetInstanceGroupCountNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetInstanceGroupCountNil(b bool)`

 SetInstanceGroupCountNil sets the value for InstanceGroupCount to be an explicit nil

### UnsetInstanceGroupCount
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetInstanceGroupCount()`

UnsetInstanceGroupCount ensures that no value is present for InstanceGroupCount, not even an explicit nil
### GetIsRollbackPossible

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetIsRollbackPossible() bool`

GetIsRollbackPossible returns the IsRollbackPossible field if non-nil, zero value otherwise.

### GetIsRollbackPossibleOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetIsRollbackPossibleOk() (*bool, bool)`

GetIsRollbackPossibleOk returns a tuple with the IsRollbackPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRollbackPossible

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetIsRollbackPossible(v bool)`

SetIsRollbackPossible sets IsRollbackPossible field to given value.

### HasIsRollbackPossible

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasIsRollbackPossible() bool`

HasIsRollbackPossible returns a boolean if a field has been set.

### SetIsRollbackPossibleNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetIsRollbackPossibleNil(b bool)`

 SetIsRollbackPossibleNil sets the value for IsRollbackPossible to be an explicit nil

### UnsetIsRollbackPossible
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetIsRollbackPossible()`

UnsetIsRollbackPossible ensures that no value is present for IsRollbackPossible, not even an explicit nil
### GetName

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListMysqlCustomParameterGroupsCustomParameterGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


