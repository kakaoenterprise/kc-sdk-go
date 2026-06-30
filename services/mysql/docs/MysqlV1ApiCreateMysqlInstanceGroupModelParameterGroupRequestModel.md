# MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ParameterGroupType**](ParameterGroupType.md) |  | 
**Id** | **string** | 적용할 파라미터 그룹의 ID | 

## Methods

### NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel(type_ ParameterGroupType, id string, ) *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModelWithDefaults

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModelWithDefaults() *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModelWithDefaults instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel) GetType() ParameterGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel) GetTypeOk() (*ParameterGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel) SetType(v ParameterGroupType)`

SetType sets Type field to given value.


### GetId

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


