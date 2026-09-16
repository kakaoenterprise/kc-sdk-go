# ListMysqlDefaultParameterGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultParameterGroups** | [**[]ListMysqlDefaultParameterGroupsDefaultParameterGroup**](ListMysqlDefaultParameterGroupsDefaultParameterGroup.md) | 조회된 MySQL 기본 파라미터 그룹 목록 | 

## Methods

### NewListMysqlDefaultParameterGroupsResponse

`func NewListMysqlDefaultParameterGroupsResponse(defaultParameterGroups []ListMysqlDefaultParameterGroupsDefaultParameterGroup, ) *ListMysqlDefaultParameterGroupsResponse`

NewListMysqlDefaultParameterGroupsResponse instantiates a new ListMysqlDefaultParameterGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlDefaultParameterGroupsResponseWithDefaults

`func NewListMysqlDefaultParameterGroupsResponseWithDefaults() *ListMysqlDefaultParameterGroupsResponse`

NewListMysqlDefaultParameterGroupsResponseWithDefaults instantiates a new ListMysqlDefaultParameterGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultParameterGroups

`func (o *ListMysqlDefaultParameterGroupsResponse) GetDefaultParameterGroups() []ListMysqlDefaultParameterGroupsDefaultParameterGroup`

GetDefaultParameterGroups returns the DefaultParameterGroups field if non-nil, zero value otherwise.

### GetDefaultParameterGroupsOk

`func (o *ListMysqlDefaultParameterGroupsResponse) GetDefaultParameterGroupsOk() (*[]ListMysqlDefaultParameterGroupsDefaultParameterGroup, bool)`

GetDefaultParameterGroupsOk returns a tuple with the DefaultParameterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterGroups

`func (o *ListMysqlDefaultParameterGroupsResponse) SetDefaultParameterGroups(v []ListMysqlDefaultParameterGroupsDefaultParameterGroup)`

SetDefaultParameterGroups sets DefaultParameterGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


