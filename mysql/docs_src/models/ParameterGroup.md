# ParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ParameterGroupType**](ParameterGroupType.md) | 파라미터 그룹 유형 | 
**Id** | **string** | 적용된 MySQL 파라미터 그룹 ID | 
**ApplyStatus** | Pointer to [**NullableParameterGroupStatus**](ParameterGroupStatus.md) | 파라미터 그룹 적용 상태 | [optional] 
**EngineVersion** | **string** | MySQL 엔진 버전 | 
**IsEngineVersionMismatch** | **bool** | 인스턴스 그룹의 엔진 버전과 파라미터 그룹 엔진 버전 불일치 여부 | 

## Methods

### NewParameterGroup

`func NewParameterGroup(type_ ParameterGroupType, id string, engineVersion string, isEngineVersionMismatch bool, ) *ParameterGroup`

NewParameterGroup instantiates a new ParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupWithDefaults

`func NewParameterGroupWithDefaults() *ParameterGroup`

NewParameterGroupWithDefaults instantiates a new ParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParameterGroup) GetType() ParameterGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterGroup) GetTypeOk() (*ParameterGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterGroup) SetType(v ParameterGroupType)`

SetType sets Type field to given value.


### GetId

`func (o *ParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterGroup) SetId(v string)`

SetId sets Id field to given value.


### GetApplyStatus

`func (o *ParameterGroup) GetApplyStatus() ParameterGroupStatus`

GetApplyStatus returns the ApplyStatus field if non-nil, zero value otherwise.

### GetApplyStatusOk

`func (o *ParameterGroup) GetApplyStatusOk() (*ParameterGroupStatus, bool)`

GetApplyStatusOk returns a tuple with the ApplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyStatus

`func (o *ParameterGroup) SetApplyStatus(v ParameterGroupStatus)`

SetApplyStatus sets ApplyStatus field to given value.

### HasApplyStatus

`func (o *ParameterGroup) HasApplyStatus() bool`

HasApplyStatus returns a boolean if a field has been set.

### SetApplyStatusNil

`func (o *ParameterGroup) SetApplyStatusNil(b bool)`

 SetApplyStatusNil sets the value for ApplyStatus to be an explicit nil

### UnsetApplyStatus
`func (o *ParameterGroup) UnsetApplyStatus()`

UnsetApplyStatus ensures that no value is present for ApplyStatus, not even an explicit nil
### GetEngineVersion

`func (o *ParameterGroup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *ParameterGroup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *ParameterGroup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetIsEngineVersionMismatch

`func (o *ParameterGroup) GetIsEngineVersionMismatch() bool`

GetIsEngineVersionMismatch returns the IsEngineVersionMismatch field if non-nil, zero value otherwise.

### GetIsEngineVersionMismatchOk

`func (o *ParameterGroup) GetIsEngineVersionMismatchOk() (*bool, bool)`

GetIsEngineVersionMismatchOk returns a tuple with the IsEngineVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEngineVersionMismatch

`func (o *ParameterGroup) SetIsEngineVersionMismatch(v bool)`

SetIsEngineVersionMismatch sets IsEngineVersionMismatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


