# ScaleOutMysqlInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandbyPort** | Pointer to **NullableInt32** | Standby 인스턴스에서 사용할 MySQL 포트 | [optional] 
**SubnetInfos** | [**[]SubnetInfoRequest**](SubnetInfoRequest.md) | 새로 생성할 인스턴스에 대한 서브넷 및 복제본(Replicas) 설정 목록 | 

## Methods

### NewScaleOutMysqlInstanceGroup

`func NewScaleOutMysqlInstanceGroup(subnetInfos []SubnetInfoRequest, ) *ScaleOutMysqlInstanceGroup`

NewScaleOutMysqlInstanceGroup instantiates a new ScaleOutMysqlInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleOutMysqlInstanceGroupWithDefaults

`func NewScaleOutMysqlInstanceGroupWithDefaults() *ScaleOutMysqlInstanceGroup`

NewScaleOutMysqlInstanceGroupWithDefaults instantiates a new ScaleOutMysqlInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandbyPort

`func (o *ScaleOutMysqlInstanceGroup) GetStandbyPort() int32`

GetStandbyPort returns the StandbyPort field if non-nil, zero value otherwise.

### GetStandbyPortOk

`func (o *ScaleOutMysqlInstanceGroup) GetStandbyPortOk() (*int32, bool)`

GetStandbyPortOk returns a tuple with the StandbyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyPort

`func (o *ScaleOutMysqlInstanceGroup) SetStandbyPort(v int32)`

SetStandbyPort sets StandbyPort field to given value.

### HasStandbyPort

`func (o *ScaleOutMysqlInstanceGroup) HasStandbyPort() bool`

HasStandbyPort returns a boolean if a field has been set.

### SetStandbyPortNil

`func (o *ScaleOutMysqlInstanceGroup) SetStandbyPortNil(b bool)`

 SetStandbyPortNil sets the value for StandbyPort to be an explicit nil

### UnsetStandbyPort
`func (o *ScaleOutMysqlInstanceGroup) UnsetStandbyPort()`

UnsetStandbyPort ensures that no value is present for StandbyPort, not even an explicit nil
### GetSubnetInfos

`func (o *ScaleOutMysqlInstanceGroup) GetSubnetInfos() []SubnetInfoRequest`

GetSubnetInfos returns the SubnetInfos field if non-nil, zero value otherwise.

### GetSubnetInfosOk

`func (o *ScaleOutMysqlInstanceGroup) GetSubnetInfosOk() (*[]SubnetInfoRequest, bool)`

GetSubnetInfosOk returns a tuple with the SubnetInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetInfos

`func (o *ScaleOutMysqlInstanceGroup) SetSubnetInfos(v []SubnetInfoRequest)`

SetSubnetInfos sets SubnetInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


