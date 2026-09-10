# CreateMysqlInstanceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroup** | [**CreateMysqlInstanceGroup**](CreateMysqlInstanceGroup.md) | 생성할 MySQL 인스턴스 그룹 정보 | 

## Methods

### NewCreateMysqlInstanceGroupRequest

`func NewCreateMysqlInstanceGroupRequest(instanceGroup CreateMysqlInstanceGroup, ) *CreateMysqlInstanceGroupRequest`

NewCreateMysqlInstanceGroupRequest instantiates a new CreateMysqlInstanceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMysqlInstanceGroupRequestWithDefaults

`func NewCreateMysqlInstanceGroupRequestWithDefaults() *CreateMysqlInstanceGroupRequest`

NewCreateMysqlInstanceGroupRequestWithDefaults instantiates a new CreateMysqlInstanceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroup

`func (o *CreateMysqlInstanceGroupRequest) GetInstanceGroup() CreateMysqlInstanceGroup`

GetInstanceGroup returns the InstanceGroup field if non-nil, zero value otherwise.

### GetInstanceGroupOk

`func (o *CreateMysqlInstanceGroupRequest) GetInstanceGroupOk() (*CreateMysqlInstanceGroup, bool)`

GetInstanceGroupOk returns a tuple with the InstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroup

`func (o *CreateMysqlInstanceGroupRequest) SetInstanceGroup(v CreateMysqlInstanceGroup)`

SetInstanceGroup sets InstanceGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


