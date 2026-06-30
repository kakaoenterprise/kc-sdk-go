# GetMySQLCustomParameterGroupsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomParameterGroups** | [**[]MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel**](MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel.md) | MySQL 커스텀 파라미터 그룹 목록 | 

## Methods

### NewGetMySQLCustomParameterGroupsResponseModel

`func NewGetMySQLCustomParameterGroupsResponseModel(customParameterGroups []MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel, ) *GetMySQLCustomParameterGroupsResponseModel`

NewGetMySQLCustomParameterGroupsResponseModel instantiates a new GetMySQLCustomParameterGroupsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMySQLCustomParameterGroupsResponseModelWithDefaults

`func NewGetMySQLCustomParameterGroupsResponseModelWithDefaults() *GetMySQLCustomParameterGroupsResponseModel`

NewGetMySQLCustomParameterGroupsResponseModelWithDefaults instantiates a new GetMySQLCustomParameterGroupsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomParameterGroups

`func (o *GetMySQLCustomParameterGroupsResponseModel) GetCustomParameterGroups() []MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel`

GetCustomParameterGroups returns the CustomParameterGroups field if non-nil, zero value otherwise.

### GetCustomParameterGroupsOk

`func (o *GetMySQLCustomParameterGroupsResponseModel) GetCustomParameterGroupsOk() (*[]MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel, bool)`

GetCustomParameterGroupsOk returns a tuple with the CustomParameterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameterGroups

`func (o *GetMySQLCustomParameterGroupsResponseModel) SetCustomParameterGroups(v []MysqlV1ApiListMysqlCustomParameterGroupsModelCustomParameterGroupResponseModel)`

SetCustomParameterGroups sets CustomParameterGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


