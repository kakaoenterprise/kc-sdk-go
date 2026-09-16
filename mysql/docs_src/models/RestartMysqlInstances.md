# RestartMysqlInstances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceIds** | **[]string** | 재시작 대상 인스턴스의 ID 목록 | 

## Methods

### NewRestartMysqlInstances

`func NewRestartMysqlInstances(instanceIds []string, ) *RestartMysqlInstances`

NewRestartMysqlInstances instantiates a new RestartMysqlInstances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartMysqlInstancesWithDefaults

`func NewRestartMysqlInstancesWithDefaults() *RestartMysqlInstances`

NewRestartMysqlInstancesWithDefaults instantiates a new RestartMysqlInstances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceIds

`func (o *RestartMysqlInstances) GetInstanceIds() []string`

GetInstanceIds returns the InstanceIds field if non-nil, zero value otherwise.

### GetInstanceIdsOk

`func (o *RestartMysqlInstances) GetInstanceIdsOk() (*[]string, bool)`

GetInstanceIdsOk returns a tuple with the InstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIds

`func (o *RestartMysqlInstances) SetInstanceIds(v []string)`

SetInstanceIds sets InstanceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


