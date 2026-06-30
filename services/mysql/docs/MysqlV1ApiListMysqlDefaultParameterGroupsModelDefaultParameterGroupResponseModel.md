# MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 기본 MySQL 파라미터 그룹 ID | 
**EngineVersion** | **string** | 파라미터 그룹이 적용되는 MySQL 엔진 버전 | 
**Name** | **string** | 기본 파라미터 그룹 이름 | 
**Description** | **string** | 기본 파라미터 그룹 설명 | 
**InstanceGroupCount** | Pointer to **NullableInt32** |  | [optional] 
**ExistEngineVersionMismatch** | Pointer to **NullableBool** |  | [optional] 
**ExistErrorSync** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel

`func NewMysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel(id string, engineVersion string, name string, description string, ) *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel`

NewMysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel instantiates a new MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModelWithDefaults

`func NewMysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModelWithDefaults() *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel`

NewMysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModelWithDefaults instantiates a new MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetEngineVersion

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetName

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstanceGroupCount

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.

### HasInstanceGroupCount

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) HasInstanceGroupCount() bool`

HasInstanceGroupCount returns a boolean if a field has been set.

### SetInstanceGroupCountNil

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetInstanceGroupCountNil(b bool)`

 SetInstanceGroupCountNil sets the value for InstanceGroupCount to be an explicit nil

### UnsetInstanceGroupCount
`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) UnsetInstanceGroupCount()`

UnsetInstanceGroupCount ensures that no value is present for InstanceGroupCount, not even an explicit nil
### GetExistEngineVersionMismatch

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetExistEngineVersionMismatch() bool`

GetExistEngineVersionMismatch returns the ExistEngineVersionMismatch field if non-nil, zero value otherwise.

### GetExistEngineVersionMismatchOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetExistEngineVersionMismatchOk() (*bool, bool)`

GetExistEngineVersionMismatchOk returns a tuple with the ExistEngineVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistEngineVersionMismatch

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetExistEngineVersionMismatch(v bool)`

SetExistEngineVersionMismatch sets ExistEngineVersionMismatch field to given value.

### HasExistEngineVersionMismatch

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) HasExistEngineVersionMismatch() bool`

HasExistEngineVersionMismatch returns a boolean if a field has been set.

### SetExistEngineVersionMismatchNil

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetExistEngineVersionMismatchNil(b bool)`

 SetExistEngineVersionMismatchNil sets the value for ExistEngineVersionMismatch to be an explicit nil

### UnsetExistEngineVersionMismatch
`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) UnsetExistEngineVersionMismatch()`

UnsetExistEngineVersionMismatch ensures that no value is present for ExistEngineVersionMismatch, not even an explicit nil
### GetExistErrorSync

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *MysqlV1ApiListMysqlDefaultParameterGroupsModelDefaultParameterGroupResponseModel) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


