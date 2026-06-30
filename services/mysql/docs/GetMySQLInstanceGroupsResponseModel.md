# GetMySQLInstanceGroupsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroups** | [**[]MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel**](MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel.md) | MySQL 인스턴스 그룹 목록 | 

## Methods

### NewGetMySQLInstanceGroupsResponseModel

`func NewGetMySQLInstanceGroupsResponseModel(instanceGroups []MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel, ) *GetMySQLInstanceGroupsResponseModel`

NewGetMySQLInstanceGroupsResponseModel instantiates a new GetMySQLInstanceGroupsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMySQLInstanceGroupsResponseModelWithDefaults

`func NewGetMySQLInstanceGroupsResponseModelWithDefaults() *GetMySQLInstanceGroupsResponseModel`

NewGetMySQLInstanceGroupsResponseModelWithDefaults instantiates a new GetMySQLInstanceGroupsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroups

`func (o *GetMySQLInstanceGroupsResponseModel) GetInstanceGroups() []MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *GetMySQLInstanceGroupsResponseModel) GetInstanceGroupsOk() (*[]MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *GetMySQLInstanceGroupsResponseModel) SetInstanceGroups(v []MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel)`

SetInstanceGroups sets InstanceGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


