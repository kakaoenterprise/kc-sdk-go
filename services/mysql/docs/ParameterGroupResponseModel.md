# ParameterGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ParameterGroupType**](ParameterGroupType.md) |  | 
**Id** | **string** | 적용된 MySQL 파라미터 그룹 ID | 
**ApplyStatus** | Pointer to [**NullableParameterGroupStatus**](ParameterGroupStatus.md) |  | [optional] 
**EngineVersion** | **string** | MySQL 엔진 버전 | 
**IsEngineVersionMismatch** | **bool** | 인스턴스 그룹의 엔진 버전과 파라미터 그룹 엔진 버전 불일치 여부 | 

## Methods

### NewParameterGroupResponseModel

`func NewParameterGroupResponseModel(type_ ParameterGroupType, id string, engineVersion string, isEngineVersionMismatch bool, ) *ParameterGroupResponseModel`

NewParameterGroupResponseModel instantiates a new ParameterGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupResponseModelWithDefaults

`func NewParameterGroupResponseModelWithDefaults() *ParameterGroupResponseModel`

NewParameterGroupResponseModelWithDefaults instantiates a new ParameterGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParameterGroupResponseModel) GetType() ParameterGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterGroupResponseModel) GetTypeOk() (*ParameterGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterGroupResponseModel) SetType(v ParameterGroupType)`

SetType sets Type field to given value.


### GetId

`func (o *ParameterGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetApplyStatus

`func (o *ParameterGroupResponseModel) GetApplyStatus() ParameterGroupStatus`

GetApplyStatus returns the ApplyStatus field if non-nil, zero value otherwise.

### GetApplyStatusOk

`func (o *ParameterGroupResponseModel) GetApplyStatusOk() (*ParameterGroupStatus, bool)`

GetApplyStatusOk returns a tuple with the ApplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyStatus

`func (o *ParameterGroupResponseModel) SetApplyStatus(v ParameterGroupStatus)`

SetApplyStatus sets ApplyStatus field to given value.

### HasApplyStatus

`func (o *ParameterGroupResponseModel) HasApplyStatus() bool`

HasApplyStatus returns a boolean if a field has been set.

### SetApplyStatusNil

`func (o *ParameterGroupResponseModel) SetApplyStatusNil(b bool)`

 SetApplyStatusNil sets the value for ApplyStatus to be an explicit nil

### UnsetApplyStatus
`func (o *ParameterGroupResponseModel) UnsetApplyStatus()`

UnsetApplyStatus ensures that no value is present for ApplyStatus, not even an explicit nil
### GetEngineVersion

`func (o *ParameterGroupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *ParameterGroupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *ParameterGroupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetIsEngineVersionMismatch

`func (o *ParameterGroupResponseModel) GetIsEngineVersionMismatch() bool`

GetIsEngineVersionMismatch returns the IsEngineVersionMismatch field if non-nil, zero value otherwise.

### GetIsEngineVersionMismatchOk

`func (o *ParameterGroupResponseModel) GetIsEngineVersionMismatchOk() (*bool, bool)`

GetIsEngineVersionMismatchOk returns a tuple with the IsEngineVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEngineVersionMismatch

`func (o *ParameterGroupResponseModel) SetIsEngineVersionMismatch(v bool)`

SetIsEngineVersionMismatch sets IsEngineVersionMismatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


