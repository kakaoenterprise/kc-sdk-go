# MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**NullableMysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel**](MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel.md) |  | [optional] 
**Standby** | Pointer to [**[]MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel**](MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel.md) |  | [optional] 

## Methods

### NewMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel

`func NewMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel() *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel`

NewMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel instantiates a new MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModelWithDefaults

`func NewMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModelWithDefaults() *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel`

NewMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModelWithDefaults instantiates a new MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) GetPrimary() MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) GetPrimaryOk() (*MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) SetPrimary(v MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetStandby

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) GetStandby() []MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel`

GetStandby returns the Standby field if non-nil, zero value otherwise.

### GetStandbyOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) GetStandbyOk() (*[]MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel, bool)`

GetStandbyOk returns a tuple with the Standby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandby

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) SetStandby(v []MysqlV1ApiListMysqlInstanceGroupsModelTopologyInfoResponseModel)`

SetStandby sets Standby field to given value.

### HasStandby

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) HasStandby() bool`

HasStandby returns a boolean if a field has been set.

### SetStandbyNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) SetStandbyNil(b bool)`

 SetStandbyNil sets the value for Standby to be an explicit nil

### UnsetStandby
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel) UnsetStandby()`

UnsetStandby ensures that no value is present for Standby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


