# MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandbyPort** | Pointer to **NullableInt32** |  | [optional] 
**SubnetInfos** | [**[]MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel**](MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel.md) | 새로 생성할 인스턴스에 대한 서브넷 및 복제본(Replicas) 설정 목록 | 

## Methods

### NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel

`func NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel(subnetInfos []MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel, ) *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel`

NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel instantiates a new MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModelWithDefaults

`func NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModelWithDefaults() *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel`

NewMysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModelWithDefaults instantiates a new MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandbyPort

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) GetStandbyPort() int32`

GetStandbyPort returns the StandbyPort field if non-nil, zero value otherwise.

### GetStandbyPortOk

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) GetStandbyPortOk() (*int32, bool)`

GetStandbyPortOk returns a tuple with the StandbyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyPort

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) SetStandbyPort(v int32)`

SetStandbyPort sets StandbyPort field to given value.

### HasStandbyPort

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) HasStandbyPort() bool`

HasStandbyPort returns a boolean if a field has been set.

### SetStandbyPortNil

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) SetStandbyPortNil(b bool)`

 SetStandbyPortNil sets the value for StandbyPort to be an explicit nil

### UnsetStandbyPort
`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) UnsetStandbyPort()`

UnsetStandbyPort ensures that no value is present for StandbyPort, not even an explicit nil
### GetSubnetInfos

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) GetSubnetInfos() []MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel`

GetSubnetInfos returns the SubnetInfos field if non-nil, zero value otherwise.

### GetSubnetInfosOk

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) GetSubnetInfosOk() (*[]MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel, bool)`

GetSubnetInfosOk returns a tuple with the SubnetInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetInfos

`func (o *MysqlV1ApiScaleOutMysqlInstanceGroupModelInstanceGroupRequestModel) SetSubnetInfos(v []MysqlV1ApiScaleOutMysqlInstanceGroupModelSubnetInfoRequestModel)`

SetSubnetInfos sets SubnetInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


