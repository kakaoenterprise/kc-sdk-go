# GetMysqlCustomParameterGroupCustomParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultParameterGroupId** | Pointer to **NullableString** | 기준이 되는 기본 MySQL 파라미터 그룹 ID | [optional] 
**Description** | Pointer to **NullableString** | MySQL 커스텀 파라미터 그룹 설명 | [optional] 
**EngineVersion** | Pointer to **NullableString** | MySQL 엔진 버전 | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** | 파라미터 적용 또는 동기화 과정에서 오류가 발생한 인스턴스 그룹 존재 여부 | [optional] 
**InstanceGroupCount** | **int32** | 해당 커스텀 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 수 | 
**IsRollbackPossible** | Pointer to **NullableBool** | 이전 파라미터 그룹으로 롤백 가능 여부 | [optional] 
**Name** | Pointer to **NullableString** | MySQL 커스텀 파라미터 그룹 이름 | [optional] 
**Id** | Pointer to **NullableString** | MySQL 커스텀 파라미터 그룹 ID | [optional] 
**Parameters** | Pointer to [**[]DataDetailParameters**](DataDetailParameters.md) | 커스텀 파라미터 목록 | [optional] 

## Methods

### NewGetMysqlCustomParameterGroupCustomParameterGroup

`func NewGetMysqlCustomParameterGroupCustomParameterGroup(instanceGroupCount int32, ) *GetMysqlCustomParameterGroupCustomParameterGroup`

NewGetMysqlCustomParameterGroupCustomParameterGroup instantiates a new GetMysqlCustomParameterGroupCustomParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlCustomParameterGroupCustomParameterGroupWithDefaults

`func NewGetMysqlCustomParameterGroupCustomParameterGroupWithDefaults() *GetMysqlCustomParameterGroupCustomParameterGroup`

NewGetMysqlCustomParameterGroupCustomParameterGroupWithDefaults instantiates a new GetMysqlCustomParameterGroupCustomParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultParameterGroupId

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetDefaultParameterGroupId() string`

GetDefaultParameterGroupId returns the DefaultParameterGroupId field if non-nil, zero value otherwise.

### GetDefaultParameterGroupIdOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetDefaultParameterGroupIdOk() (*string, bool)`

GetDefaultParameterGroupIdOk returns a tuple with the DefaultParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterGroupId

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetDefaultParameterGroupId(v string)`

SetDefaultParameterGroupId sets DefaultParameterGroupId field to given value.

### HasDefaultParameterGroupId

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasDefaultParameterGroupId() bool`

HasDefaultParameterGroupId returns a boolean if a field has been set.

### SetDefaultParameterGroupIdNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetDefaultParameterGroupIdNil(b bool)`

 SetDefaultParameterGroupIdNil sets the value for DefaultParameterGroupId to be an explicit nil

### UnsetDefaultParameterGroupId
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetDefaultParameterGroupId()`

UnsetDefaultParameterGroupId ensures that no value is present for DefaultParameterGroupId, not even an explicit nil
### GetDescription

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngineVersion

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil
### GetExistErrorSync

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil
### GetInstanceGroupCount

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.


### GetIsRollbackPossible

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetIsRollbackPossible() bool`

GetIsRollbackPossible returns the IsRollbackPossible field if non-nil, zero value otherwise.

### GetIsRollbackPossibleOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetIsRollbackPossibleOk() (*bool, bool)`

GetIsRollbackPossibleOk returns a tuple with the IsRollbackPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRollbackPossible

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetIsRollbackPossible(v bool)`

SetIsRollbackPossible sets IsRollbackPossible field to given value.

### HasIsRollbackPossible

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasIsRollbackPossible() bool`

HasIsRollbackPossible returns a boolean if a field has been set.

### SetIsRollbackPossibleNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetIsRollbackPossibleNil(b bool)`

 SetIsRollbackPossibleNil sets the value for IsRollbackPossible to be an explicit nil

### UnsetIsRollbackPossible
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetIsRollbackPossible()`

UnsetIsRollbackPossible ensures that no value is present for IsRollbackPossible, not even an explicit nil
### GetName

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetParameters

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetParameters() []DataDetailParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) GetParametersOk() (*[]DataDetailParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetParameters(v []DataDetailParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *GetMysqlCustomParameterGroupCustomParameterGroup) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


