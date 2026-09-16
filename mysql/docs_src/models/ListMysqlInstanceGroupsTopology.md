# ListMysqlInstanceGroupsTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**NullableListMysqlInstanceGroupsTopologyInfo**](ListMysqlInstanceGroupsTopologyInfo.md) | Primary 인스턴스 정보 | [optional] 
**Standby** | Pointer to [**[]ListMysqlInstanceGroupsTopologyInfo**](ListMysqlInstanceGroupsTopologyInfo.md) | Standby 인스턴스 정보 목록 | [optional] 

## Methods

### NewListMysqlInstanceGroupsTopology

`func NewListMysqlInstanceGroupsTopology() *ListMysqlInstanceGroupsTopology`

NewListMysqlInstanceGroupsTopology instantiates a new ListMysqlInstanceGroupsTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstanceGroupsTopologyWithDefaults

`func NewListMysqlInstanceGroupsTopologyWithDefaults() *ListMysqlInstanceGroupsTopology`

NewListMysqlInstanceGroupsTopologyWithDefaults instantiates a new ListMysqlInstanceGroupsTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *ListMysqlInstanceGroupsTopology) GetPrimary() ListMysqlInstanceGroupsTopologyInfo`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *ListMysqlInstanceGroupsTopology) GetPrimaryOk() (*ListMysqlInstanceGroupsTopologyInfo, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *ListMysqlInstanceGroupsTopology) SetPrimary(v ListMysqlInstanceGroupsTopologyInfo)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *ListMysqlInstanceGroupsTopology) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *ListMysqlInstanceGroupsTopology) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *ListMysqlInstanceGroupsTopology) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetStandby

`func (o *ListMysqlInstanceGroupsTopology) GetStandby() []ListMysqlInstanceGroupsTopologyInfo`

GetStandby returns the Standby field if non-nil, zero value otherwise.

### GetStandbyOk

`func (o *ListMysqlInstanceGroupsTopology) GetStandbyOk() (*[]ListMysqlInstanceGroupsTopologyInfo, bool)`

GetStandbyOk returns a tuple with the Standby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandby

`func (o *ListMysqlInstanceGroupsTopology) SetStandby(v []ListMysqlInstanceGroupsTopologyInfo)`

SetStandby sets Standby field to given value.

### HasStandby

`func (o *ListMysqlInstanceGroupsTopology) HasStandby() bool`

HasStandby returns a boolean if a field has been set.

### SetStandbyNil

`func (o *ListMysqlInstanceGroupsTopology) SetStandbyNil(b bool)`

 SetStandbyNil sets the value for Standby to be an explicit nil

### UnsetStandby
`func (o *ListMysqlInstanceGroupsTopology) UnsetStandby()`

UnsetStandby ensures that no value is present for Standby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


