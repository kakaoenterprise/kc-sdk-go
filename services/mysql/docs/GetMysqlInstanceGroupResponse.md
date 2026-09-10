# GetMysqlInstanceGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroup** | [**GetMysqlInstanceGroupInstanceGroup**](GetMysqlInstanceGroupInstanceGroup.md) | MySQL 인스턴스 그룹 상세 정보 | 

## Methods

### NewGetMysqlInstanceGroupResponse

`func NewGetMysqlInstanceGroupResponse(instanceGroup GetMysqlInstanceGroupInstanceGroup, ) *GetMysqlInstanceGroupResponse`

NewGetMysqlInstanceGroupResponse instantiates a new GetMysqlInstanceGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlInstanceGroupResponseWithDefaults

`func NewGetMysqlInstanceGroupResponseWithDefaults() *GetMysqlInstanceGroupResponse`

NewGetMysqlInstanceGroupResponseWithDefaults instantiates a new GetMysqlInstanceGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroup

`func (o *GetMysqlInstanceGroupResponse) GetInstanceGroup() GetMysqlInstanceGroupInstanceGroup`

GetInstanceGroup returns the InstanceGroup field if non-nil, zero value otherwise.

### GetInstanceGroupOk

`func (o *GetMysqlInstanceGroupResponse) GetInstanceGroupOk() (*GetMysqlInstanceGroupInstanceGroup, bool)`

GetInstanceGroupOk returns a tuple with the InstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroup

`func (o *GetMysqlInstanceGroupResponse) SetInstanceGroup(v GetMysqlInstanceGroupInstanceGroup)`

SetInstanceGroup sets InstanceGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


