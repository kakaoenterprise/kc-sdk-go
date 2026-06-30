# MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel

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
**Parameters** | [**[]ParameterResponseModel**](ParameterResponseModel.md) | 기본 파라미터 그룹에 포함된 파라미터 목록 | 

## Methods

### NewMysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel

`func NewMysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel(id string, engineVersion string, name string, description string, parameters []ParameterResponseModel, ) *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel`

NewMysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel instantiates a new MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModelWithDefaults

`func NewMysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModelWithDefaults() *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel`

NewMysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModelWithDefaults instantiates a new MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetEngineVersion

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetName

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstanceGroupCount

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetInstanceGroupCount() int32`

GetInstanceGroupCount returns the InstanceGroupCount field if non-nil, zero value otherwise.

### GetInstanceGroupCountOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetInstanceGroupCountOk() (*int32, bool)`

GetInstanceGroupCountOk returns a tuple with the InstanceGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupCount

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetInstanceGroupCount(v int32)`

SetInstanceGroupCount sets InstanceGroupCount field to given value.

### HasInstanceGroupCount

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) HasInstanceGroupCount() bool`

HasInstanceGroupCount returns a boolean if a field has been set.

### SetInstanceGroupCountNil

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetInstanceGroupCountNil(b bool)`

 SetInstanceGroupCountNil sets the value for InstanceGroupCount to be an explicit nil

### UnsetInstanceGroupCount
`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) UnsetInstanceGroupCount()`

UnsetInstanceGroupCount ensures that no value is present for InstanceGroupCount, not even an explicit nil
### GetExistEngineVersionMismatch

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetExistEngineVersionMismatch() bool`

GetExistEngineVersionMismatch returns the ExistEngineVersionMismatch field if non-nil, zero value otherwise.

### GetExistEngineVersionMismatchOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetExistEngineVersionMismatchOk() (*bool, bool)`

GetExistEngineVersionMismatchOk returns a tuple with the ExistEngineVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistEngineVersionMismatch

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetExistEngineVersionMismatch(v bool)`

SetExistEngineVersionMismatch sets ExistEngineVersionMismatch field to given value.

### HasExistEngineVersionMismatch

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) HasExistEngineVersionMismatch() bool`

HasExistEngineVersionMismatch returns a boolean if a field has been set.

### SetExistEngineVersionMismatchNil

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetExistEngineVersionMismatchNil(b bool)`

 SetExistEngineVersionMismatchNil sets the value for ExistEngineVersionMismatch to be an explicit nil

### UnsetExistEngineVersionMismatch
`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) UnsetExistEngineVersionMismatch()`

UnsetExistEngineVersionMismatch ensures that no value is present for ExistEngineVersionMismatch, not even an explicit nil
### GetExistErrorSync

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetExistErrorSync() bool`

GetExistErrorSync returns the ExistErrorSync field if non-nil, zero value otherwise.

### GetExistErrorSyncOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetExistErrorSyncOk() (*bool, bool)`

GetExistErrorSyncOk returns a tuple with the ExistErrorSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistErrorSync

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetExistErrorSync(v bool)`

SetExistErrorSync sets ExistErrorSync field to given value.

### HasExistErrorSync

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) HasExistErrorSync() bool`

HasExistErrorSync returns a boolean if a field has been set.

### SetExistErrorSyncNil

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetExistErrorSyncNil(b bool)`

 SetExistErrorSyncNil sets the value for ExistErrorSync to be an explicit nil

### UnsetExistErrorSync
`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) UnsetExistErrorSync()`

UnsetExistErrorSync ensures that no value is present for ExistErrorSync, not even an explicit nil
### GetParameters

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetParameters() []ParameterResponseModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) GetParametersOk() (*[]ParameterResponseModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MysqlV1ApiGetMysqlDefaultParameterGroupModelDefaultParameterGroupResponseModel) SetParameters(v []ParameterResponseModel)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


