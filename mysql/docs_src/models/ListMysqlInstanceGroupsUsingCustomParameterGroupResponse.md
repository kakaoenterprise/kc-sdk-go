# ListMysqlInstanceGroupsUsingCustomParameterGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroups** | [**[]DefaultParameterGroupInstanceGroup**](DefaultParameterGroupInstanceGroup.md) | 커스텀 MySQL 파라미터 그룹을 사용하는 MySQL 인스턴스 그룹 목록 | 

## Methods

### NewListMysqlInstanceGroupsUsingCustomParameterGroupResponse

`func NewListMysqlInstanceGroupsUsingCustomParameterGroupResponse(instanceGroups []DefaultParameterGroupInstanceGroup, ) *ListMysqlInstanceGroupsUsingCustomParameterGroupResponse`

NewListMysqlInstanceGroupsUsingCustomParameterGroupResponse instantiates a new ListMysqlInstanceGroupsUsingCustomParameterGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstanceGroupsUsingCustomParameterGroupResponseWithDefaults

`func NewListMysqlInstanceGroupsUsingCustomParameterGroupResponseWithDefaults() *ListMysqlInstanceGroupsUsingCustomParameterGroupResponse`

NewListMysqlInstanceGroupsUsingCustomParameterGroupResponseWithDefaults instantiates a new ListMysqlInstanceGroupsUsingCustomParameterGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroups

`func (o *ListMysqlInstanceGroupsUsingCustomParameterGroupResponse) GetInstanceGroups() []DefaultParameterGroupInstanceGroup`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *ListMysqlInstanceGroupsUsingCustomParameterGroupResponse) GetInstanceGroupsOk() (*[]DefaultParameterGroupInstanceGroup, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *ListMysqlInstanceGroupsUsingCustomParameterGroupResponse) SetInstanceGroups(v []DefaultParameterGroupInstanceGroup)`

SetInstanceGroups sets InstanceGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


