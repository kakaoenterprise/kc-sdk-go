# GetMysqlDefaultParameterGroupDefaultParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 기본 MySQL 파라미터 그룹 ID | 
**EngineVersion** | **string** | 파라미터 그룹이 적용되는 MySQL 엔진 버전 | 
**Name** | **string** | 기본 파라미터 그룹 이름 | 
**Description** | **string** | 기본 파라미터 그룹 설명 | 
**InstanceGroupCount** | Pointer to **NullableInt32** | 해당 기본 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 수 | [optional] 
**ExistEngineVersionMismatch** | Pointer to **NullableBool** | 엔진 버전이 일치하지 않는 인스턴스 그룹 존재 여부 | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** | 파라미터 동기화 오류가 발생한 인스턴스 그룹 존재 여부 | [optional] 
**Parameters** | [**[]Parameter**](Parameter.md) | 기본 파라미터 그룹에 포함된 파라미터 목록 | 

## Methods

### NewGetMysqlDefaultParameterGroupDefaultParameterGroup

`func NewGetMysqlDefaultParameterGroupDefaultParameterGroup(id string, engineVersion string, name string, description string, parameters []Parameter, ) *GetMysqlDefaultParameterGroupDefaultParameterGroup`

NewGetMysqlDefaultParameterGroupDefaultParameterGroup instantiates a new GetMysqlDefaultParameterGroupDefaultParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlDefaultParameterGroupDefaultParameterGroupWithDefaults

`func NewGetMysqlDefaultParameterGroupDefaultParameterGroupWithDefaults() *GetMysqlDefaultParameterGroupDefaultParameterGroup`

NewGetMysqlDefaultParameterGroupDefaultParameterGroupWithDefaults instantiates a new GetMysqlDefaultParameterGroupDefaultParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetId(v string)`

SetId sets Id field to given value.


### GetEngineVersion

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetName

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstanceGroupCount

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.

### HasInstanceGroupCount

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) HasInstanceGroupCount() bool`

HasInstanceGroupCount returns a boolean if a field has been set.

### SetInstanceGroupCountNil

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetInstanceGroupCountNil(b bool)`

 SetInstanceGroupCountNil sets the value for InstanceGroupCount to be an explicit nil

### UnsetInstanceGroupCount
`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) UnsetInstanceGroupCount()`

UnsetInstanceGroupCount ensures that no value is present for InstanceGroupCount, not even an explicit nil
### GetExistEngineVersionMismatch

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetExistEngineVersionMismatch() bool`

GetExistEngineVersionMismatch returns the ExistEngineVersionMismatch field if non-nil, zero value otherwise.

### GetExistEngineVersionMismatchOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetExistEngineVersionMismatchOk() (*bool, bool)`

GetExistEngineVersionMismatchOk returns a tuple with the ExistEngineVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistEngineVersionMismatch

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetExistEngineVersionMismatch(v bool)`

SetExistEngineVersionMismatch sets ExistEngineVersionMismatch field to given value.

### HasExistEngineVersionMismatch

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) HasExistEngineVersionMismatch() bool`

HasExistEngineVersionMismatch returns a boolean if a field has been set.

### SetExistEngineVersionMismatchNil

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetExistEngineVersionMismatchNil(b bool)`

 SetExistEngineVersionMismatchNil sets the value for ExistEngineVersionMismatch to be an explicit nil

### UnsetExistEngineVersionMismatch
`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) UnsetExistEngineVersionMismatch()`

UnsetExistEngineVersionMismatch ensures that no value is present for ExistEngineVersionMismatch, not even an explicit nil
### GetExistErrorSync

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil
### GetParameters

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetParameters() []Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) GetParametersOk() (*[]Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetMysqlDefaultParameterGroupDefaultParameterGroup) SetParameters(v []Parameter)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


