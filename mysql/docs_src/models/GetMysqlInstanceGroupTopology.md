# GetMysqlInstanceGroupTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**NullableGetMysqlInstanceGroupTopologyInfo**](GetMysqlInstanceGroupTopologyInfo.md) | Primary 역할을 수행하는 MySQL 인스턴스 | [optional] 
**Standby** | Pointer to [**[]GetMysqlInstanceGroupTopologyInfo**](GetMysqlInstanceGroupTopologyInfo.md) | Standby 역할을 수행하는 MySQL 인스턴스 | [optional] 

## Methods

### NewGetMysqlInstanceGroupTopology

`func NewGetMysqlInstanceGroupTopology() *GetMysqlInstanceGroupTopology`

NewGetMysqlInstanceGroupTopology instantiates a new GetMysqlInstanceGroupTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlInstanceGroupTopologyWithDefaults

`func NewGetMysqlInstanceGroupTopologyWithDefaults() *GetMysqlInstanceGroupTopology`

NewGetMysqlInstanceGroupTopologyWithDefaults instantiates a new GetMysqlInstanceGroupTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *GetMysqlInstanceGroupTopology) GetPrimary() GetMysqlInstanceGroupTopologyInfo`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *GetMysqlInstanceGroupTopology) GetPrimaryOk() (*GetMysqlInstanceGroupTopologyInfo, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *GetMysqlInstanceGroupTopology) SetPrimary(v GetMysqlInstanceGroupTopologyInfo)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *GetMysqlInstanceGroupTopology) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *GetMysqlInstanceGroupTopology) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *GetMysqlInstanceGroupTopology) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetStandby

`func (o *GetMysqlInstanceGroupTopology) GetStandby() []GetMysqlInstanceGroupTopologyInfo`

GetStandby returns the Standby field if non-nil, zero value otherwise.

### GetStandbyOk

`func (o *GetMysqlInstanceGroupTopology) GetStandbyOk() (*[]GetMysqlInstanceGroupTopologyInfo, bool)`

GetStandbyOk returns a tuple with the Standby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandby

`func (o *GetMysqlInstanceGroupTopology) SetStandby(v []GetMysqlInstanceGroupTopologyInfo)`

SetStandby sets Standby field to given value.

### HasStandby

`func (o *GetMysqlInstanceGroupTopology) HasStandby() bool`

HasStandby returns a boolean if a field has been set.

### SetStandbyNil

`func (o *GetMysqlInstanceGroupTopology) SetStandbyNil(b bool)`

 SetStandbyNil sets the value for Standby to be an explicit nil

### UnsetStandby
`func (o *GetMysqlInstanceGroupTopology) UnsetStandby()`

UnsetStandby ensures that no value is present for Standby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


