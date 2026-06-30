# NetworkInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupIds** | Pointer to **[]string** |  | [optional] 
**PrimarySubnetInfo** | Pointer to [**NullableSubnetInfoResponseModel**](SubnetInfoResponseModel.md) |  | [optional] 
**StandbySubnetInfo** | Pointer to [**[]SubnetInfoResponseModel**](SubnetInfoResponseModel.md) |  | [optional] 

## Methods

### NewNetworkInfoResponseModel

`func NewNetworkInfoResponseModel() *NetworkInfoResponseModel`

NewNetworkInfoResponseModel instantiates a new NetworkInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInfoResponseModelWithDefaults

`func NewNetworkInfoResponseModelWithDefaults() *NetworkInfoResponseModel`

NewNetworkInfoResponseModelWithDefaults instantiates a new NetworkInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupIds

`func (o *NetworkInfoResponseModel) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NetworkInfoResponseModel) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *NetworkInfoResponseModel) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *NetworkInfoResponseModel) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIdsNil

`func (o *NetworkInfoResponseModel) SetSecurityGroupIdsNil(b bool)`

 SetSecurityGroupIdsNil sets the value for SecurityGroupIds to be an explicit nil

### UnsetSecurityGroupIds
`func (o *NetworkInfoResponseModel) UnsetSecurityGroupIds()`

UnsetSecurityGroupIds ensures that no value is present for SecurityGroupIds, not even an explicit nil
### GetPrimarySubnetInfo

`func (o *NetworkInfoResponseModel) GetPrimarySubnetInfo() SubnetInfoResponseModel`

GetPrimarySubnetInfo returns the PrimarySubnetInfo field if non-nil, zero value otherwise.

### GetPrimarySubnetInfoOk

`func (o *NetworkInfoResponseModel) GetPrimarySubnetInfoOk() (*SubnetInfoResponseModel, bool)`

GetPrimarySubnetInfoOk returns a tuple with the PrimarySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubnetInfo

`func (o *NetworkInfoResponseModel) SetPrimarySubnetInfo(v SubnetInfoResponseModel)`

SetPrimarySubnetInfo sets PrimarySubnetInfo field to given value.

### HasPrimarySubnetInfo

`func (o *NetworkInfoResponseModel) HasPrimarySubnetInfo() bool`

HasPrimarySubnetInfo returns a boolean if a field has been set.

### SetPrimarySubnetInfoNil

`func (o *NetworkInfoResponseModel) SetPrimarySubnetInfoNil(b bool)`

 SetPrimarySubnetInfoNil sets the value for PrimarySubnetInfo to be an explicit nil

### UnsetPrimarySubnetInfo
`func (o *NetworkInfoResponseModel) UnsetPrimarySubnetInfo()`

UnsetPrimarySubnetInfo ensures that no value is present for PrimarySubnetInfo, not even an explicit nil
### GetStandbySubnetInfo

`func (o *NetworkInfoResponseModel) GetStandbySubnetInfo() []SubnetInfoResponseModel`

GetStandbySubnetInfo returns the StandbySubnetInfo field if non-nil, zero value otherwise.

### GetStandbySubnetInfoOk

`func (o *NetworkInfoResponseModel) GetStandbySubnetInfoOk() (*[]SubnetInfoResponseModel, bool)`

GetStandbySubnetInfoOk returns a tuple with the StandbySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbySubnetInfo

`func (o *NetworkInfoResponseModel) SetStandbySubnetInfo(v []SubnetInfoResponseModel)`

SetStandbySubnetInfo sets StandbySubnetInfo field to given value.

### HasStandbySubnetInfo

`func (o *NetworkInfoResponseModel) HasStandbySubnetInfo() bool`

HasStandbySubnetInfo returns a boolean if a field has been set.

### SetStandbySubnetInfoNil

`func (o *NetworkInfoResponseModel) SetStandbySubnetInfoNil(b bool)`

 SetStandbySubnetInfoNil sets the value for StandbySubnetInfo to be an explicit nil

### UnsetStandbySubnetInfo
`func (o *NetworkInfoResponseModel) UnsetStandbySubnetInfo()`

UnsetStandbySubnetInfo ensures that no value is present for StandbySubnetInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


