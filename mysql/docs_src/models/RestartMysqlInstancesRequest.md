# RestartMysqlInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroup** | [**RestartMysqlInstances**](RestartMysqlInstances.md) | 인스턴스 ID 목록을 포함하는 객체 | 

## Methods

### NewRestartMysqlInstancesRequest

`func NewRestartMysqlInstancesRequest(instanceGroup RestartMysqlInstances, ) *RestartMysqlInstancesRequest`

NewRestartMysqlInstancesRequest instantiates a new RestartMysqlInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartMysqlInstancesRequestWithDefaults

`func NewRestartMysqlInstancesRequestWithDefaults() *RestartMysqlInstancesRequest`

NewRestartMysqlInstancesRequestWithDefaults instantiates a new RestartMysqlInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroup

`func (o *RestartMysqlInstancesRequest) GetInstanceGroup() RestartMysqlInstances`

GetInstanceGroup returns the InstanceGroup field if non-nil, zero value otherwise.

### GetInstanceGroupOk

`func (o *RestartMysqlInstancesRequest) GetInstanceGroupOk() (*RestartMysqlInstances, bool)`

GetInstanceGroupOk returns a tuple with the InstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroup

`func (o *RestartMysqlInstancesRequest) SetInstanceGroup(v RestartMysqlInstances)`

SetInstanceGroup sets InstanceGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


