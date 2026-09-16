# ListMysqlInstanceGroupsTopologyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **NullableString** | 대상 MySQL 인스턴스 ID | [optional] 

## Methods

### NewListMysqlInstanceGroupsTopologyInfo

`func NewListMysqlInstanceGroupsTopologyInfo() *ListMysqlInstanceGroupsTopologyInfo`

NewListMysqlInstanceGroupsTopologyInfo instantiates a new ListMysqlInstanceGroupsTopologyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstanceGroupsTopologyInfoWithDefaults

`func NewListMysqlInstanceGroupsTopologyInfoWithDefaults() *ListMysqlInstanceGroupsTopologyInfo`

NewListMysqlInstanceGroupsTopologyInfoWithDefaults instantiates a new ListMysqlInstanceGroupsTopologyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ListMysqlInstanceGroupsTopologyInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListMysqlInstanceGroupsTopologyInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListMysqlInstanceGroupsTopologyInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListMysqlInstanceGroupsTopologyInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ListMysqlInstanceGroupsTopologyInfo) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ListMysqlInstanceGroupsTopologyInfo) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


