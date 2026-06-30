# MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**NullableMysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel**](MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel.md) |  | [optional] 
**Standby** | Pointer to [**[]MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel**](MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel.md) |  | [optional] 

## Methods

### NewMysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel

`func NewMysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel() *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel`

NewMysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel instantiates a new MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModelWithDefaults

`func NewMysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModelWithDefaults() *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel`

NewMysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModelWithDefaults instantiates a new MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) GetPrimary() MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) GetPrimaryOk() (*MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) SetPrimary(v MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetStandby

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) GetStandby() []MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel`

GetStandby returns the Standby field if non-nil, zero value otherwise.

### GetStandbyOk

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) GetStandbyOk() (*[]MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel, bool)`

GetStandbyOk returns a tuple with the Standby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandby

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) SetStandby(v []MysqlV1ApiGetMysqlInstanceGroupModelTopologyInfoResponseModel)`

SetStandby sets Standby field to given value.

### HasStandby

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) HasStandby() bool`

HasStandby returns a boolean if a field has been set.

### SetStandbyNil

`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) SetStandbyNil(b bool)`

 SetStandbyNil sets the value for Standby to be an explicit nil

### UnsetStandby
`func (o *MysqlV1ApiGetMysqlInstanceGroupModelTopologyResponseModel) UnsetStandby()`

UnsetStandby ensures that no value is present for Standby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


