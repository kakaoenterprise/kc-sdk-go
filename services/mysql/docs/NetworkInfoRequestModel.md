# NetworkInfoRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupIds** | **[]string** | 인스턴스 그룹에 적용할 보안 그룹 ID 목록 | 
**PrimarySubnetInfo** | [**MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel**](MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel.md) | Primary 서브넷 설정 | 
**StandbySubnetInfo** | Pointer to [**[]MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel**](MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel.md) |  | [optional] 

## Methods

### NewNetworkInfoRequestModel

`func NewNetworkInfoRequestModel(securityGroupIds []string, primarySubnetInfo MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel, ) *NetworkInfoRequestModel`

NewNetworkInfoRequestModel instantiates a new NetworkInfoRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInfoRequestModelWithDefaults

`func NewNetworkInfoRequestModelWithDefaults() *NetworkInfoRequestModel`

NewNetworkInfoRequestModelWithDefaults instantiates a new NetworkInfoRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupIds

`func (o *NetworkInfoRequestModel) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NetworkInfoRequestModel) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *NetworkInfoRequestModel) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.


### GetPrimarySubnetInfo

`func (o *NetworkInfoRequestModel) GetPrimarySubnetInfo() MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel`

GetPrimarySubnetInfo returns the PrimarySubnetInfo field if non-nil, zero value otherwise.

### GetPrimarySubnetInfoOk

`func (o *NetworkInfoRequestModel) GetPrimarySubnetInfoOk() (*MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel, bool)`

GetPrimarySubnetInfoOk returns a tuple with the PrimarySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubnetInfo

`func (o *NetworkInfoRequestModel) SetPrimarySubnetInfo(v MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel)`

SetPrimarySubnetInfo sets PrimarySubnetInfo field to given value.


### GetStandbySubnetInfo

`func (o *NetworkInfoRequestModel) GetStandbySubnetInfo() []MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel`

GetStandbySubnetInfo returns the StandbySubnetInfo field if non-nil, zero value otherwise.

### GetStandbySubnetInfoOk

`func (o *NetworkInfoRequestModel) GetStandbySubnetInfoOk() (*[]MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel, bool)`

GetStandbySubnetInfoOk returns a tuple with the StandbySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbySubnetInfo

`func (o *NetworkInfoRequestModel) SetStandbySubnetInfo(v []MysqlV1ApiCreateMysqlInstanceGroupModelSubnetInfoRequestModel)`

SetStandbySubnetInfo sets StandbySubnetInfo field to given value.

### HasStandbySubnetInfo

`func (o *NetworkInfoRequestModel) HasStandbySubnetInfo() bool`

HasStandbySubnetInfo returns a boolean if a field has been set.

### SetStandbySubnetInfoNil

`func (o *NetworkInfoRequestModel) SetStandbySubnetInfoNil(b bool)`

 SetStandbySubnetInfoNil sets the value for StandbySubnetInfo to be an explicit nil

### UnsetStandbySubnetInfo
`func (o *NetworkInfoRequestModel) UnsetStandbySubnetInfo()`

UnsetStandbySubnetInfo ensures that no value is present for StandbySubnetInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


