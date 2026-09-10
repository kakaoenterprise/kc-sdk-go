# ListMysqlInstanceGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroups** | [**[]ListMysqlInstanceGroupsInstanceGroup**](ListMysqlInstanceGroupsInstanceGroup.md) | MySQL 인스턴스 그룹 목록 | 

## Methods

### NewListMysqlInstanceGroupsResponse

`func NewListMysqlInstanceGroupsResponse(instanceGroups []ListMysqlInstanceGroupsInstanceGroup, ) *ListMysqlInstanceGroupsResponse`

NewListMysqlInstanceGroupsResponse instantiates a new ListMysqlInstanceGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstanceGroupsResponseWithDefaults

`func NewListMysqlInstanceGroupsResponseWithDefaults() *ListMysqlInstanceGroupsResponse`

NewListMysqlInstanceGroupsResponseWithDefaults instantiates a new ListMysqlInstanceGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroups

`func (o *ListMysqlInstanceGroupsResponse) GetInstanceGroups() []ListMysqlInstanceGroupsInstanceGroup`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *ListMysqlInstanceGroupsResponse) GetInstanceGroupsOk() (*[]ListMysqlInstanceGroupsInstanceGroup, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *ListMysqlInstanceGroupsResponse) SetInstanceGroups(v []ListMysqlInstanceGroupsInstanceGroup)`

SetInstanceGroups sets InstanceGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


