# ListMysqlDefaultParameterGroupsDefaultParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 기본 MySQL 파라미터 그룹 ID | 
**EngineVersion** | **string** | 파라미터 그룹이 적용되는 MySQL 엔진 버전 | 
**Name** | **string** | 기본 파라미터 그룹 이름 | 
**Description** | **string** | 기본 파라미터 그룹 설명 | 
**InstanceGroupCount** | Pointer to **NullableInt32** | 해당 기본 MySQL 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 수 | [optional] 
**ExistEngineVersionMismatch** | Pointer to **NullableBool** | 엔진 버전이 일치하지 않는 인스턴스 그룹 존재 여부 | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** | 파라미터 동기화 오류가 발생한 인스턴스 그룹 존재 여부 | [optional] 

## Methods

### NewListMysqlDefaultParameterGroupsDefaultParameterGroup

`func NewListMysqlDefaultParameterGroupsDefaultParameterGroup(id string, engineVersion string, name string, description string, ) *ListMysqlDefaultParameterGroupsDefaultParameterGroup`

NewListMysqlDefaultParameterGroupsDefaultParameterGroup instantiates a new ListMysqlDefaultParameterGroupsDefaultParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlDefaultParameterGroupsDefaultParameterGroupWithDefaults

`func NewListMysqlDefaultParameterGroupsDefaultParameterGroupWithDefaults() *ListMysqlDefaultParameterGroupsDefaultParameterGroup`

NewListMysqlDefaultParameterGroupsDefaultParameterGroupWithDefaults instantiates a new ListMysqlDefaultParameterGroupsDefaultParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetId(v string)`

SetId sets Id field to given value.


### GetEngineVersion

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetName

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstanceGroupCount

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.

### HasInstanceGroupCount

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) HasInstanceGroupCount() bool`

HasInstanceGroupCount returns a boolean if a field has been set.

### SetInstanceGroupCountNil

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetInstanceGroupCountNil(b bool)`

 SetInstanceGroupCountNil sets the value for InstanceGroupCount to be an explicit nil

### UnsetInstanceGroupCount
`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) UnsetInstanceGroupCount()`

UnsetInstanceGroupCount ensures that no value is present for InstanceGroupCount, not even an explicit nil
### GetExistEngineVersionMismatch

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetExistEngineVersionMismatch() bool`

GetExistEngineVersionMismatch returns the ExistEngineVersionMismatch field if non-nil, zero value otherwise.

### GetExistEngineVersionMismatchOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetExistEngineVersionMismatchOk() (*bool, bool)`

GetExistEngineVersionMismatchOk returns a tuple with the ExistEngineVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistEngineVersionMismatch

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetExistEngineVersionMismatch(v bool)`

SetExistEngineVersionMismatch sets ExistEngineVersionMismatch field to given value.

### HasExistEngineVersionMismatch

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) HasExistEngineVersionMismatch() bool`

HasExistEngineVersionMismatch returns a boolean if a field has been set.

### SetExistEngineVersionMismatchNil

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetExistEngineVersionMismatchNil(b bool)`

 SetExistEngineVersionMismatchNil sets the value for ExistEngineVersionMismatch to be an explicit nil

### UnsetExistEngineVersionMismatch
`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) UnsetExistEngineVersionMismatch()`

UnsetExistEngineVersionMismatch ensures that no value is present for ExistEngineVersionMismatch, not even an explicit nil
### GetExistErrorSync

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *ListMysqlDefaultParameterGroupsDefaultParameterGroup) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


