# ApplyMysqlParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | MySQL 인스턴스 그룹에 적용할 파라미터 그룹 ID - &#x60;type&#x3D;DEFAULT&#x60;: [List MySQL default parameter groups](/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인 - &#x60;type&#x3D;CUSTOM&#x60;: [List MySQL custom parameter groups](/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 
**Type** | [**ParameterGroupType**](ParameterGroupType.md) | 파라미터 그룹 유형 | 

## Methods

### NewApplyMysqlParameterGroup

`func NewApplyMysqlParameterGroup(id string, type_ ParameterGroupType, ) *ApplyMysqlParameterGroup`

NewApplyMysqlParameterGroup instantiates a new ApplyMysqlParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyMysqlParameterGroupWithDefaults

`func NewApplyMysqlParameterGroupWithDefaults() *ApplyMysqlParameterGroup`

NewApplyMysqlParameterGroupWithDefaults instantiates a new ApplyMysqlParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplyMysqlParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplyMysqlParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplyMysqlParameterGroup) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ApplyMysqlParameterGroup) GetType() ParameterGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplyMysqlParameterGroup) GetTypeOk() (*ParameterGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplyMysqlParameterGroup) SetType(v ParameterGroupType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


