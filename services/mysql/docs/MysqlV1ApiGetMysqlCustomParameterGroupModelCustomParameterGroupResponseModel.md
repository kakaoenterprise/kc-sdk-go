# MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultParameterGroupId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EngineVersion** | Pointer to **NullableString** |  | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** |  | [optional] 
**InstanceGroupCount** | **int32** | 해당 커스텀 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 수 | 
**IsRollbackPossible** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**[]DataDetailParametersResponseModel**](DataDetailParametersResponseModel.md) |  | [optional] 

## Methods

### NewMysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel

`func NewMysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel(instanceGroupCount int32, ) *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel`

NewMysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel instantiates a new MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModelWithDefaults

`func NewMysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModelWithDefaults() *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel`

NewMysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModelWithDefaults instantiates a new MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultParameterGroupId

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetDefaultParameterGroupId() string`

GetDefaultParameterGroupId returns the DefaultParameterGroupId field if non-nil, zero value otherwise.

### GetDefaultParameterGroupIdOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetDefaultParameterGroupIdOk() (*string, bool)`

GetDefaultParameterGroupIdOk returns a tuple with the DefaultParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterGroupId

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetDefaultParameterGroupId(v string)`

SetDefaultParameterGroupId sets DefaultParameterGroupId field to given value.

### HasDefaultParameterGroupId

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasDefaultParameterGroupId() bool`

HasDefaultParameterGroupId returns a boolean if a field has been set.

### SetDefaultParameterGroupIdNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetDefaultParameterGroupIdNil(b bool)`

 SetDefaultParameterGroupIdNil sets the value for DefaultParameterGroupId to be an explicit nil

### UnsetDefaultParameterGroupId
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetDefaultParameterGroupId()`

UnsetDefaultParameterGroupId ensures that no value is present for DefaultParameterGroupId, not even an explicit nil
### GetDescription

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngineVersion

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil
### GetExistErrorSync

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil
### GetInstanceGroupCount

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.


### GetIsRollbackPossible

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetIsRollbackPossible() bool`

GetIsRollbackPossible returns the IsRollbackPossible field if non-nil, zero value otherwise.

### GetIsRollbackPossibleOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetIsRollbackPossibleOk() (*bool, bool)`

GetIsRollbackPossibleOk returns a tuple with the IsRollbackPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRollbackPossible

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetIsRollbackPossible(v bool)`

SetIsRollbackPossible sets IsRollbackPossible field to given value.

### HasIsRollbackPossible

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasIsRollbackPossible() bool`

HasIsRollbackPossible returns a boolean if a field has been set.

### SetIsRollbackPossibleNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetIsRollbackPossibleNil(b bool)`

 SetIsRollbackPossibleNil sets the value for IsRollbackPossible to be an explicit nil

### UnsetIsRollbackPossible
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetIsRollbackPossible()`

UnsetIsRollbackPossible ensures that no value is present for IsRollbackPossible, not even an explicit nil
### GetName

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetParameters

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetParameters() []DataDetailParametersResponseModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) GetParametersOk() (*[]DataDetailParametersResponseModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetParameters(v []DataDetailParametersResponseModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *MysqlV1ApiGetMysqlCustomParameterGroupModelCustomParameterGroupResponseModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


