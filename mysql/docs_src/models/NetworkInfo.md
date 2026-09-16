# NetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupIds** | Pointer to **[]string** | 적용된 보안 그룹 ID 목록 | [optional] 
**PrimarySubnetInfo** | Pointer to [**NullableSubnetInfo**](SubnetInfo.md) | Primary 서브넷 정보 | [optional] 
**StandbySubnetInfo** | Pointer to [**[]SubnetInfo**](SubnetInfo.md) | Standby 서브넷 정보 | [optional] 

## Methods

### NewNetworkInfo

`func NewNetworkInfo() *NetworkInfo`

NewNetworkInfo instantiates a new NetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInfoWithDefaults

`func NewNetworkInfoWithDefaults() *NetworkInfo`

NewNetworkInfoWithDefaults instantiates a new NetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupIds

`func (o *NetworkInfo) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NetworkInfo) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *NetworkInfo) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *NetworkInfo) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIdsNil

`func (o *NetworkInfo) SetSecurityGroupIdsNil(b bool)`

 SetSecurityGroupIdsNil sets the value for SecurityGroupIds to be an explicit nil

### UnsetSecurityGroupIds
`func (o *NetworkInfo) UnsetSecurityGroupIds()`

UnsetSecurityGroupIds ensures that no value is present for SecurityGroupIds, not even an explicit nil
### GetPrimarySubnetInfo

`func (o *NetworkInfo) GetPrimarySubnetInfo() SubnetInfo`

GetPrimarySubnetInfo returns the PrimarySubnetInfo field if non-nil, zero value otherwise.

### GetPrimarySubnetInfoOk

`func (o *NetworkInfo) GetPrimarySubnetInfoOk() (*SubnetInfo, bool)`

GetPrimarySubnetInfoOk returns a tuple with the PrimarySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubnetInfo

`func (o *NetworkInfo) SetPrimarySubnetInfo(v SubnetInfo)`

SetPrimarySubnetInfo sets PrimarySubnetInfo field to given value.

### HasPrimarySubnetInfo

`func (o *NetworkInfo) HasPrimarySubnetInfo() bool`

HasPrimarySubnetInfo returns a boolean if a field has been set.

### SetPrimarySubnetInfoNil

`func (o *NetworkInfo) SetPrimarySubnetInfoNil(b bool)`

 SetPrimarySubnetInfoNil sets the value for PrimarySubnetInfo to be an explicit nil

### UnsetPrimarySubnetInfo
`func (o *NetworkInfo) UnsetPrimarySubnetInfo()`

UnsetPrimarySubnetInfo ensures that no value is present for PrimarySubnetInfo, not even an explicit nil
### GetStandbySubnetInfo

`func (o *NetworkInfo) GetStandbySubnetInfo() []SubnetInfo`

GetStandbySubnetInfo returns the StandbySubnetInfo field if non-nil, zero value otherwise.

### GetStandbySubnetInfoOk

`func (o *NetworkInfo) GetStandbySubnetInfoOk() (*[]SubnetInfo, bool)`

GetStandbySubnetInfoOk returns a tuple with the StandbySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbySubnetInfo

`func (o *NetworkInfo) SetStandbySubnetInfo(v []SubnetInfo)`

SetStandbySubnetInfo sets StandbySubnetInfo field to given value.

### HasStandbySubnetInfo

`func (o *NetworkInfo) HasStandbySubnetInfo() bool`

HasStandbySubnetInfo returns a boolean if a field has been set.

### SetStandbySubnetInfoNil

`func (o *NetworkInfo) SetStandbySubnetInfoNil(b bool)`

 SetStandbySubnetInfoNil sets the value for StandbySubnetInfo to be an explicit nil

### UnsetStandbySubnetInfo
`func (o *NetworkInfo) UnsetStandbySubnetInfo()`

UnsetStandbySubnetInfo ensures that no value is present for StandbySubnetInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


