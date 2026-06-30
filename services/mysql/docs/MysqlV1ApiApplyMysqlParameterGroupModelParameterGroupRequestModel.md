# MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | MySQL 인스턴스 그룹에 적용할 파라미터 그룹 ID | 
**Type** | [**ParameterGroupType**](ParameterGroupType.md) |  | 

## Methods

### NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel

`func NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel(id string, type_ ParameterGroupType, ) *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel`

NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel instantiates a new MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModelWithDefaults

`func NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModelWithDefaults() *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel`

NewMysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModelWithDefaults instantiates a new MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel) GetType() ParameterGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel) GetTypeOk() (*ParameterGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MysqlV1ApiApplyMysqlParameterGroupModelParameterGroupRequestModel) SetType(v ParameterGroupType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


