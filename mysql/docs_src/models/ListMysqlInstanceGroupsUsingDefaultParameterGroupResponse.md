# ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroups** | [**[]DefaultParameterGroupInstanceGroup**](DefaultParameterGroupInstanceGroup.md) | 기본 MySQL 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 목록 | 

## Methods

### NewListMysqlInstanceGroupsUsingDefaultParameterGroupResponse

`func NewListMysqlInstanceGroupsUsingDefaultParameterGroupResponse(instanceGroups []DefaultParameterGroupInstanceGroup, ) *ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse`

NewListMysqlInstanceGroupsUsingDefaultParameterGroupResponse instantiates a new ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstanceGroupsUsingDefaultParameterGroupResponseWithDefaults

`func NewListMysqlInstanceGroupsUsingDefaultParameterGroupResponseWithDefaults() *ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse`

NewListMysqlInstanceGroupsUsingDefaultParameterGroupResponseWithDefaults instantiates a new ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroups

`func (o *ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse) GetInstanceGroups() []DefaultParameterGroupInstanceGroup`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse) GetInstanceGroupsOk() (*[]DefaultParameterGroupInstanceGroup, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *ListMysqlInstanceGroupsUsingDefaultParameterGroupResponse) SetInstanceGroups(v []DefaultParameterGroupInstanceGroup)`

SetInstanceGroups sets InstanceGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


